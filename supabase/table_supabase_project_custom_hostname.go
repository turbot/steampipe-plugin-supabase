package supabase

import (
	"context"

	"github.com/supabase/cli/pkg/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableSupabaseProjectCustomHostname(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "supabase_project_custom_hostname",
		Description: "Supabase Project Custom Hostname",
		List: &plugin.ListConfig{
			ParentHydrate: listSupabaseProjects,
			Hydrate:       listSupabaseProjectCustomHostname,
		},
		Columns: []*plugin.Column{
			{Name: "custom_hostname", Type: proto.ColumnType_STRING, Description: "The custom hostname."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the hostname."},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "Specifies the custom hostname object for the specified project."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "The ID of the project."},
		},
	}
}

type CustomHostnamae struct {
	api.UpdateCustomHostnameResponse
	ProjectId string
}

//// LIST FUNCTION

func listSupabaseProjectCustomHostname(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the project data
	project := h.Item.(api.ProjectResponse)

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project_custom_hostname.listSupabaseProjectCustomHostname", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetCustomHostnameConfigWithResponse(ctx, project.Id)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project_custom_hostname.listSupabaseProjectCustomHostname", "query_error", err)
		return nil, err
	}

	// Return nil, if no custom hostname configured
	if resp.JSON200 == nil {
		return nil, nil
	}
	d.StreamListItem(ctx, CustomHostnamae{*resp.JSON200, project.Id})

	return nil, nil
}
