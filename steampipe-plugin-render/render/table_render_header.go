package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the header."},
			{Name: "path", Type: proto.ColumnType_STRING, Description: "The path of the header."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the header."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "The value of the header."},
		},
	}
}
