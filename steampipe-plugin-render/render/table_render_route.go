package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderRoute(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_route",
		Description: "Information about route descriptions, including ID, type, source, destination, and priority.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoute,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetRoute,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the route."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the route."},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "The source of the route."},
			{Name: "destination", Type: proto.ColumnType_STRING, Description: "The destination of the route."},
			{Name: "priority", Type: proto.ColumnType_INT, Description: "The priority of the route."},
		},
	}
}
