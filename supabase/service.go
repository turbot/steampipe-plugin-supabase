package supabase

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	supabase "github.com/supabase/cli/pkg/api"
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

	// No creds
	if supabaseConfig.AccessToken == nil {
		return nil, fmt.Errorf("access_token must be configured")
	}

	provider, err := securityprovider.NewSecurityProviderBearerToken(*supabaseConfig.AccessToken)
	if err != nil {
		return nil, err
	}

	if t, ok := http.DefaultTransport.(*http.Transport); ok {
		dialContext := t.DialContext
		t.DialContext = func(ctx context.Context, network, address string) (net.Conn, error) {
			conn, err := dialContext(ctx, network, address)
			// Workaround when pure Go DNS resolver fails https://github.com/golang/go/issues/12524
			if err, ok := err.(*net.OpError); ok && err.Op == "dial" {
				if ip := fallbackLookupIP(ctx, address); ip != "" {
					return dialContext(ctx, network, ip)
				}
			}
			return conn, err
		}
	}

	apiClient, err := supabase.NewClientWithResponses(
		"https://api.supabase.com",
		supabase.WithRequestEditorFn(provider.Intercept),
		supabase.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("User-Agent", "Steampipe")
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return apiClient, nil
}
