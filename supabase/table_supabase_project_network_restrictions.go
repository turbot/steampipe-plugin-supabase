package supabase

import (
	"context"

	"github.com/supabase/cli/pkg/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableSupabaseProjectNetworkRestrictions(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "supabase_project_network_restrictions",
		Description: "Supabase Project Network Restrictions",
		List: &plugin.ListConfig{
			ParentHydrate: listSupabaseProjects,
			Hydrate:       listSupabaseProjectNetworkRestrictions,
		},
		Columns: []*plugin.Column{
			{Name: "entitlement", Type: proto.ColumnType_STRING, Description: "Indicates whether the Supabase project has access to network restrictions or not. Possible values are: 'allowed', 'disallowed'."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the network restrictions for the Supabase project. Possible values are: 'stored', 'applied'."},
			{Name: "config", Type: proto.ColumnType_JSON, Description: "Specifies the current network restrictions configuration for the Supabase project."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "The ID of the project."},
		},
	}
}

type NetworkRestrictions struct {
	api.GetNetworkRestrictionsResponse
	ProjectId string
}

//// LIST FUNCTION

func listSupabaseProjectNetworkRestrictions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the project data
	project := h.Item.(api.ProjectResponse)

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project_network_restrictions.listSupabaseProjectNetworkRestrictions", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetNetworkRestrictionsWithResponse(ctx, project.Id)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project_network_restrictions.listSupabaseProjectNetworkRestrictions", "query_error", err)
		return nil, err
	}
	if resp.JSON200 != nil {
		d.StreamListItem(ctx, NetworkRestrictions{*resp, project.Id})

	}

	return nil, nil
}
