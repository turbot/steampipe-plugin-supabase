package supabase

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	supabase "github.com/supabase/cli/pkg/api"
)

const (
	BaseURL = "https://api.supabase.com"
)

// getClient:: returns Supabase client after authentication
func getClient(ctx context.Context, d *plugin.QueryData) (*supabase.ClientWithResponses, error) {
	conn, err := clientCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*supabase.ClientWithResponses), nil
}

// Get the cached version of the client
var clientCached = plugin.HydrateFunc(clientUncached).Memoize()

// clientUncached is the actual implementation of getClient, which should
// be run only once per connection. Do not call this directly, use
// getClient instead.
func clientUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	// Get the config
	supabaseConfig := GetConfig(d.Connection)

	/*
		Returns credentials by precedence.
		Precedence of credentials:
		  - Credentials set in config
		  - Value set using SUPABASE_ACCESS_TOKEN env var
	*/
	accessToken := os.Getenv("SUPABASE_ACCESS_TOKEN")
	if supabaseConfig.AccessToken != nil {
		accessToken = *supabaseConfig.AccessToken
	}

	// No creds
	if accessToken == "" {
		return nil, fmt.Errorf("access_token must be configured")
	}

	provider, err := securityprovider.NewSecurityProviderBearerToken(*supabaseConfig.AccessToken)
	if err != nil {
		return nil, err
	}

	apiClient, err := supabase.NewClientWithResponses(
		BaseURL,
		supabase.WithRequestEditorFn(provider.Intercept),
		supabase.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("User-Agent", "Steampipe")
			return nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	return apiClient, nil
}
