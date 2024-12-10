package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderEnvGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_env_group",
		Description: "Information about environment group descriptions, including ID, name, owner, service links, and timestamps.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEnvGroup,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetEnvGroup,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the environment group."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the environment group."},
			{Name: "ownerId", Type: proto.ColumnType_STRING, Description: "The ID of the owner of the environment group."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the environment group was created."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the environment group."},
			{Name: "serviceLinks", Type: proto.ColumnType_JSON, Description: "A list of service links associated with the environment group."},
			{Name: "environmentId", Type: proto.ColumnType_STRING, Description: "The ID of the associated environment."},
		},
	}
}
