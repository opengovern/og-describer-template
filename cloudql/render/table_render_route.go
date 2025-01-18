package render

import (
	"context"
	"github.com/opengovern/og-describer-render/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderRoute(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_route",
		Description: "Information about route descriptions, including ID, type, source, destination, and priority.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoute,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetRoute,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the route.", Transform: transform.FromField("Description.ID")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the route.", Transform: transform.FromField("Description.Type")},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "The source of the route.", Transform: transform.FromField("Description.Source")},
			{Name: "destination", Type: proto.ColumnType_STRING, Description: "The destination of the route.", Transform: transform.FromField("Description.Destination")},
			{Name: "priority", Type: proto.ColumnType_INT, Description: "The priority of the route.", Transform: transform.FromField("Description.Priority")},
		}),
	}
}
