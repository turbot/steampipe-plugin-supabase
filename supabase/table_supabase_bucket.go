package supabase

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-supabase/api"
)

//// TABLE DEFINITION

func tableSupabaseBucket(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "supabase_bucket",
		Description: "Supabase Bucket",
		List: &plugin.ListConfig{
			Hydrate: listSupabaseBuckets,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSupabaseBucket,
			KeyColumns: plugin.AllColumns([]string{"project_id", "id"}),
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The display name of the bucket."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique identifier of the bucket."},
			{Name: "public", Type: proto.ColumnType_BOOL, Description: "True, if the bucket is public."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the bucket was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the bucket was last modified."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "The bucket owner."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "The ID of the project where the bucket is located."},
		},
	}
}

//// LIST FUNCTION

func listSupabaseBuckets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_bucket.listSupabaseBuckets", "connection_error", err)
		return nil, err
	}

	buckets, err := api.ListBuckets(ctx, client)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_bucket.listSupabaseBuckets", "query_error", err)
		return nil, err
	}

	for _, bucket := range buckets {
		// append project details
		bucket.ProjectId = extractProjectIdFromUrl(*client.Url)

		d.StreamListItem(ctx, bucket)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSupabaseBucket(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	projectID := d.EqualsQualString("project_id")
	id := d.EqualsQualString("id")

	// Return nil, if empty
	if projectID == "" || id == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_bucket.id", "connection_error", err)
		return nil, err
	}

	bucket, err := api.GetBucket(ctx, client, id)
	if err != nil {
		plugin.Logger(ctx).Error("supabase_bucket.id", "query_error", err)
		return nil, err
	}

	// Append project details
	bucket.ProjectId = extractProjectIdFromUrl(*client.Url)

	return bucket, nil
}
