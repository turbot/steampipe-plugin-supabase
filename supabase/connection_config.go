package supabase

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type supabaseConfig struct {
	AccessToken *string `hcl:"access_token"`
}

func ConfigInstance() interface{} {
	return &supabaseConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) supabaseConfig {
	if connection == nil || connection.Config == nil {
		return supabaseConfig{}
	}
	config, _ := connection.Config.(supabaseConfig)
	return config
}
