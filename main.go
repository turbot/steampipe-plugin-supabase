package main

import (
	"github.com/turbot/steampipe-plugin-supabase/supabase"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: supabase.Plugin})
}
