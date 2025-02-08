package render

import (
	"context"
	"github.com/opengovern/og-describer-render/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderDisk(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_disk",
		Description: "Information about disk descriptions, including ID, name, size, mount path, and timestamps.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDisk,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDisk,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the disk.", Transform: transform.FromField("Description.ID")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the disk.", Transform: transform.FromField("Description.Name")},
			{Name: "size_gb", Type: proto.ColumnType_INT, Description: "The size of the disk in gigabytes.", Transform: transform.FromField("Description.SizeGB")},
			{Name: "mount_path", Type: proto.ColumnType_STRING, Description: "The mount path of the disk.", Transform: transform.FromField("Description.MountPath")},
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "The ID of the service associated with the disk.", Transform: transform.FromField("Description.ServiceID")},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the disk was created.", Transform: transform.FromField("Description.CreatedAt")},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the disk.", Transform: transform.FromField("Description.UpdatedAt")},
		}),
	}
}
