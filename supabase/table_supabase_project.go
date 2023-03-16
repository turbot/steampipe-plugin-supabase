package supabase

import (
	"context"

	"github.com/supabase/cli/pkg/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableSupabaseProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "supabase_project",
		Description: "Supabase Project",
		List: &plugin.ListConfig{
			Hydrate: listSupabaseProjects,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The display name of the project."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the project."},
			{Name: "organization_id", Type: proto.ColumnType_STRING, Description: "The organization ID."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the project was created."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "The project region."},
			{Name: "database", Type: proto.ColumnType_JSON, Description: "The database information."},
			{Name: "api_settings", Type: proto.ColumnType_JSON, Description: "The database information.", Hydrate: getProjectAPISettings, Transform: transform.FromValue()},
		},
	}
}

//// LIST FUNCTION

func listSupabaseProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project.listSupabaseProjects", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetProjectsWithResponse(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project.listSupabaseProjects", "query_error", err)
		return nil, err
	}

	for _, project := range *resp.JSON200 {
		d.StreamListItem(ctx, project)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getProjectAPISettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	projectId := h.Item.(api.ProjectResponse).Id

	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project.getProjectAPISettings", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetPostgRESTConfigWithResponse(ctx, projectId)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_project.getProjectAPISettings", "query_error", err)
		return nil, err
	}

	return resp.JSON200, nil
}
