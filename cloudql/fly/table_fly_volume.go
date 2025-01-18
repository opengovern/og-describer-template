package fly

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-fly/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableFlyVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "fly_volume",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListVolume,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetVolume,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "attached_alloc_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AttachedAllocID"), Description: "The ID of the attached allocation."},
			{Name: "attached_machine_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AttachedMachineID"), Description: "The ID of the attached machine."},
			{Name: "auto_backup_enabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.AutoBackupEnabled"), Description: "Indicates whether auto-backup is enabled."},
			{Name: "block_size", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.BlockSize"), Description: "The size of each block in the volume."},
			{Name: "blocks", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Blocks"), Description: "The total number of blocks in the volume."},
			{Name: "blocks_avail", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.BlocksAvail"), Description: "The number of available blocks in the volume."},
			{Name: "blocks_free", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.BlocksFree"), Description: "The number of free blocks in the volume."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The time when the volume was created."},
			{Name: "encrypted", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Encrypted"), Description: "Indicates whether the volume is encrypted."},
			{Name: "fstype", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.FSType"), Description: "The file system type of the volume."},
			{Name: "host_status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.HostStatus"), Description: "The host's status for the volume."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the volume."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the volume."},
			{Name: "region", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Region"), Description: "The region where the volume is located."},
			{Name: "size_gb", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.SizeGB"), Description: "The size of the volume in gigabytes."},
			{Name: "snapshot_retention", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.SnapshotRetention"), Description: "The number of snapshots retained for the volume."},
			{Name: "state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.State"), Description: "The current state of the volume."},
			{Name: "zone", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Zone"), Description: "The zone where the volume is located."},
		}),
	}
}
