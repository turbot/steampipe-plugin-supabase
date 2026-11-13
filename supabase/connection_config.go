package supabase

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type supabaseConfig struct {
	AccessToken *string `hcl:"access_token"`
}

func ConfigInstance() interface{} {
	return &supabaseConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) supabaseConfig {
	if connection == nil || connection.GetConfig() == nil {
		return supabaseConfig{}
	}
	config, _ := connection.GetConfig().(supabaseConfig)
	return config
}
