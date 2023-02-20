package supabase

import (
	"context"

	"github.com/supabase/cli/pkg/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableSupabaseSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "supabase_secret",
		Description: "Supabase Secret",
		List: &plugin.ListConfig{
			ParentHydrate: listSupabaseProjects,
			Hydrate:       listSupabaseSecrets,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the secret."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "The secret value."},
		},
	}
}

//// LIST FUNCTION

func listSupabaseSecrets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the project data
	project := h.Item.(api.ProjectResponse)

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_secret.listSupabaseSecrets", "connection_error", err)
		return nil, err
	}

	resp, err := client.GetSecretsWithResponse(ctx, project.Id)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_secret.listSupabaseSecrets", "query_error", err)
		return nil, err
	}

	for _, secret := range *resp.JSON200 {
		// // append project details
		// function.ProjectId = project.Id

		d.StreamListItem(ctx, secret)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
