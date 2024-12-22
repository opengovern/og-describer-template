package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderHeader(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_header",
		Description: "Information about header descriptions, including ID, path, name, and value.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListHeader,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetHeader,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the header.", Transform: transform.FromField("Description.ID")},
			{Name: "path", Type: proto.ColumnType_STRING, Description: "The path of the header.", Transform: transform.FromField("Description.Path")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the header.", Transform: transform.FromField("Description.Name")},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "The value of the header.", Transform: transform.FromField("Description.Value")},
		}),
	}
}
