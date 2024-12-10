package render

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderDisk(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_disk",
		Description: "Information about disk descriptions, including ID, name, size, mount path, and timestamps.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the disk."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the disk."},
			{Name: "sizeGB", Type: proto.ColumnType_INT, Description: "The size of the disk in gigabytes."},
			{Name: "mountPath", Type: proto.ColumnType_STRING, Description: "The mount path of the disk."},
			{Name: "serviceId", Type: proto.ColumnType_STRING, Description: "The ID of the service associated with the disk."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the disk was created."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the disk."},
		},
	}
}
