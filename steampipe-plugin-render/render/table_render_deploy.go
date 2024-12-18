package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the deployment."},
			{Name: "commit", Type: proto.ColumnType_JSON, Description: "The commit details associated with the deployment."},
			{Name: "image", Type: proto.ColumnType_JSON, Description: "The image details used in the deployment."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the deployment (e.g., pending, completed)."},
			{Name: "trigger", Type: proto.ColumnType_STRING, Description: "The trigger that initiated the deployment."},
			{Name: "finishedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the deployment finished."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the deployment was created."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the deployment."},
		},
	}
}
