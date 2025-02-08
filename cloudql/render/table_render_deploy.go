package render

import (
	"context"
	"github.com/opengovern/og-describer-render/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderDeploy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_deploy",
		Description: "Information about deployment descriptions, including ID, commit details, image, status, and timestamps.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDeploy,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDeploy,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the deployment.", Transform: transform.FromField("Description.ID")},
			{Name: "commit", Type: proto.ColumnType_JSON, Description: "The commit details associated with the deployment.", Transform: transform.FromField("Description.Commit")},
			{Name: "image", Type: proto.ColumnType_JSON, Description: "The image details used in the deployment.", Transform: transform.FromField("Description.Image")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the deployment (e.g., pending, completed).", Transform: transform.FromField("Description.Status")},
			{Name: "trigger", Type: proto.ColumnType_STRING, Description: "The trigger that initiated the deployment.", Transform: transform.FromField("Description.Trigger")},
			{Name: "finished_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the deployment finished.", Transform: transform.FromField("Description.FinishedAt")},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the deployment was created.", Transform: transform.FromField("Description.CreatedAt")},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the deployment.", Transform: transform.FromField("Description.UpdatedAt")},
		}),
	}
}
