package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderBlueprint(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_blueprint",
		Description: "Information about blueprint descriptions, including ID, name, status, and synchronization details.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBlueprint,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetBlueprint,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the blueprint."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the blueprint."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the blueprint (e.g., active, inactive)."},
			{Name: "autoSync", Type: proto.ColumnType_BOOL, Description: "Indicates whether auto-sync is enabled for the blueprint."},
			{Name: "repo", Type: proto.ColumnType_STRING, Description: "The repository associated with the blueprint."},
			{Name: "branch", Type: proto.ColumnType_STRING, Description: "The branch in the repository for the blueprint."},
			{Name: "lastSync", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last sync for the blueprint."},
		},
	}
}
