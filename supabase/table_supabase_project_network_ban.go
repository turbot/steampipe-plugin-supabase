package supabase

import (
	"context"

	"github.com/supabase/cli/pkg/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableSupabaseProjectNetworkBans(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "supabase_project_network_ban",
		Description: "Supabase Project Network Bans",
		List: &plugin.ListConfig{
			ParentHydrate: listSupabaseProjects,
			Hydrate:       listSupabaseProjectNetworkBans,
		},
		Columns: []*plugin.Column{
			{Name: "address", Type: proto.ColumnType_IPADDR, Description: "The name of the secret."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "The ID of the project."},
		},
	}
}

type NetworkBans struct {
	Address   string
	ProjectId string
}

//// LIST FUNCTION

func listSupabaseProjectNetworkBans(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the project data
	project := h.Item.(api.ProjectResponse)

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project_network_ban.listSupabaseProjectNetworkBans", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetNetworkBansWithResponse(ctx, project.Id)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project_network_ban.listSupabaseProjectNetworkBans", "query_error", err)
		return nil, err
	}

	if resp.JSON201.BannedIpv4Addresses != nil {
		for _, address := range resp.JSON201.BannedIpv4Addresses {
			d.StreamListItem(ctx, NetworkBans{address, project.Id})

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}
