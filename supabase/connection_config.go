package supabase

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type supabaseConfig struct {
	AccessToken *string `cty:"access_token"`
	ApiKey      *string `cty:"api_key"`
	Url         *string `cty:"url"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"access_token": {
		Type: schema.TypeString,
	},
	"api_key": {
		Type: schema.TypeString,
	},
	"url": {
		Type: schema.TypeString,
	},
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
