package render

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderRoute(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_route",
		Description: "Information about route descriptions, including ID, type, source, destination, and priority.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
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
