package render

import (
	"context"
	"github.com/opengovern/og-describer-render/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_project",
		Description: "Information about project descriptions, including ID, name, owner, and associated environments.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProject,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetProject,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the project.", Transform: transform.FromField("Description.ID")},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the project was created.", Transform: transform.FromField("Description.CreatedAt")},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the project.", Transform: transform.FromField("Description.UpdatedAt")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the project.", Transform: transform.FromField("Description.Name")},
			{Name: "owner", Type: proto.ColumnType_JSON, Description: "Information about the owner of the project.", Transform: transform.FromField("Description.Owner")},
			{Name: "environment_ids", Type: proto.ColumnType_JSON, Description: "A list of environment IDs associated with the project.", Transform: transform.FromField("Description.EnvironmentIDs")},
		}),
	}
}
