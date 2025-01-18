package fly

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-fly/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableFlyMachine(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "fly_machine",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMachine,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetMachine,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "checks", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Checks"), Description: "The health checks associated with the machine."},
			{Name: "config", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Config"), Description: "The configuration details of the machine."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The time when the machine was created."},
			{Name: "events", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Events"), Description: "The events associated with the machine."},
			{Name: "host_status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.HostStatus"), Description: "The status of the machine's host."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the machine."},
			{Name: "image_ref", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.ImageRef"), Description: "The reference to the image used by the machine."},
			{Name: "incomplete_config", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.IncompleteConfig"), Description: "The incomplete configuration of the machine."},
			{Name: "instance_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.InstanceID"), Description: "The unique instance ID of the machine."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the machine."},
			{Name: "nonce", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Nonce"), Description: "The nonce associated with the machine."},
			{Name: "private_ip", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.PrivateIP"), Description: "The private IP address of the machine."},
			{Name: "region", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Region"), Description: "The region where the machine is located."},
			{Name: "state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.State"), Description: "The current state of the machine."},
			{Name: "updated_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.UpdatedAt"), Description: "The time when the machine was last updated."},
		}),
	}
}
