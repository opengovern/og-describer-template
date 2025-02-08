package render

import (
	"context"
	"github.com/opengovern/og-describer-render/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderBlueprint(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_blueprint",
		Description: "Information about blueprint descriptions, including ID, name, status, and synchronization details.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBlueprint,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetBlueprint,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the blueprint.", Transform: transform.FromField("Description.ID")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the blueprint.", Transform: transform.FromField("Description.Name")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the blueprint (e.g., active, inactive).", Transform: transform.FromField("Description.Status")},
			{Name: "auto_sync", Type: proto.ColumnType_BOOL, Description: "Indicates whether auto-sync is enabled for the blueprint.", Transform: transform.FromField("Description.AutoSync")},
			{Name: "repo", Type: proto.ColumnType_STRING, Description: "The repository associated with the blueprint.", Transform: transform.FromField("Description.Repo")},
			{Name: "branch", Type: proto.ColumnType_STRING, Description: "The branch in the repository for the blueprint.", Transform: transform.FromField("Description.Branch")},
			{Name: "last_sync", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last sync for the blueprint.", Transform: transform.FromField("Description.LastSync")},
		}),
	}
}
