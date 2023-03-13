package supabase

import (
	"context"

	"github.com/supabase/cli/pkg/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableSupabaseFunction(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "supabase_function",
		Description: "Supabase Function",
		List: &plugin.ListConfig{
			ParentHydrate: listSupabaseProjects,
			Hydrate:       listSupabaseFunctions,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSupabaseFunction,
			KeyColumns: plugin.AllColumns([]string{"project_id", "slug"}),
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The display name of the function."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the function."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The function slug."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "The current version of the function."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the function."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the function was created.", Transform: transform.FromField("CreatedAt").Transform(transform.UnixMsToTimestamp)},
			{Name: "import_map", Type: proto.ColumnType_BOOL, Description: "If true, Supabase will automatically generate an import_map file based on the modules used within the function. It indicates whether or not Supabase should use an import_map file to load JavaScript modules within the function."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the function was last modified.", Transform: transform.FromField("UpdatedAt").Transform(transform.UnixMsToTimestamp)},
			{Name: "verify_jwt", Type: proto.ColumnType_BOOL, Description: "If true, it allows users to verify the authenticity of JSON Web Tokens (JWTs) issued by Supabase Authentication."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "The ID of the project where the function is located."},
		},
	}
}

type Function struct {
	api.FunctionResponse
	ProjectId string
}

//// LIST FUNCTION

func listSupabaseFunctions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the project data
	project := h.Item.(api.ProjectResponse)

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_function.listSupabaseFunctions", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetFunctionsWithResponse(ctx, project.Id)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_function.listSupabaseFunctions", "query_error", err)
		return nil, err
	}

	for _, function := range *resp.JSON200 {
		d.StreamListItem(ctx, Function{function, project.Id})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSupabaseFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	projectID := d.EqualsQualString("project_id")
	slug := d.EqualsQualString("slug")

	// Return nil, if empty
	if projectID == "" || slug == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_function.getSupabaseFunction", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetFunctionWithResponse(ctx, projectID, slug)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_function.getSupabaseFunction", "query_error", err)
		return nil, err
	}

	return resp.JSON200, nil
}
