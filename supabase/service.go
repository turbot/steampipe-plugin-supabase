package supabase

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-supabase/api"
)

// getClient:: returns Supabase client after authentication
func getClient(ctx context.Context, d *plugin.QueryData) (*api.Client, error) {
	conn, err := clientCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*api.Client), nil
}

// Get the cached version of the client
var clientCached = plugin.HydrateFunc(clientUncached).Memoize()

// clientUncached returns the Supabase client and cached the data
func clientUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	// Get the config
	supabaseConfig := GetConfig(d.Connection)

	// No creds
	if supabaseConfig.AccessToken == nil {
		return nil, fmt.Errorf("access_token must be configured")
	}

	// Start with an empty Supabase config
	config := api.ClientConfig{
		AccessToken: *supabaseConfig.AccessToken,
	}

	// Create the client
	client, err := api.CreateClient(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %s", err.Error())
	}

	return client, nil
}
