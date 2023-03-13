package supabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableSupabaseOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "supabase_organization",
		Description: "Supabase Organization",
		List: &plugin.ListConfig{
			Hydrate: listSupabaseOrganizations,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The display name of the organization."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the organization."},
		},
	}
}

//// LIST FUNCTION

func listSupabaseOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_organization.listSupabaseOrganizations", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetOrganizationsWithResponse(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_organization.listSupabaseOrganizations", "query_error", err)
		return nil, err
	}

	for _, organization := range *resp.JSON200 {
		d.StreamListItem(ctx, organization)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
