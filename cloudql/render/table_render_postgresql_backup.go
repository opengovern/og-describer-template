package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderPostgresqlBackup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_postgresql_backup",
		Description: "Information about postgres backups",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPostgresqlBackup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetPostgresqlBackup,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the backup", Transform: transform.FromField("Description.ID")},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromField("Description.URL")},
			{Name: "created_at", Type: proto.ColumnType_STRING, Description: "", Transform: transform.FromField("Description.CreatedAt")},
		}),
	}
}
