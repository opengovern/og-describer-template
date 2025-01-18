package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureDataBoxEdgeDevice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_databox_edge_device",
		Description: "Azure Data Box Edge Device",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetDataboxEdgeDevice,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDataboxEdgeDevice,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.ID")},
			{
				Name:        "friendly_name",
				Description: "The Data Box Edge/Gateway device name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.FriendlyName")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Type")},
			{
				Name:        "data_box_edge_device_status",
				Description: "The status of the Data Box Edge/Gateway device. Possible values include: 'ReadyToSetup', 'Online', 'Offline', 'NeedsAttention', 'Disconnected', 'PartiallyDisconnected', 'Maintenance'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.DataBoxEdgeDeviceStatus")},
			{
				Name:        "culture",
				Description: "The Data Box Edge/Gateway device culture.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.Culture")},
			{
				Name:        "description",
				Description: "he Description of the Data Box Edge/Gateway device.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.Description")},
			{
				Name:        "device_model",
				Description: "The Data Box Edge/Gateway device model.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.DeviceModel")},
			{
				Name:        "device_type",
				Description: "The type of the Data Box Edge/Gateway device. Possible values include: 'DataBoxEdgeDevice'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.DeviceType")},
			{
				Name:        "device_hcs_version",
				Description: "The device software version number of the device (eg: 1.2.18105.6).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.DeviceHcsVersion")},
			{
				Name:        "device_local_capacity",
				Description: "The Data Box Edge/Gateway device local capacity in MB.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Device.Properties.DeviceLocalCapacity")},
			{
				Name:        "device_software_version",
				Description: "The Data Box Edge/Gateway device software version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.DeviceSoftwareVersion")},
			{
				Name:        "etag",
				Description: "The etag for the devices.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Etag")},
			{
				Name:        "location",
				Description: "The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Location")},
			{
				Name:        "model_description",
				Description: "The description of the Data Box Edge/Gateway device model.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.ModelDescription")},
			{
				Name:        "node_count",
				Description: "The number of nodes in the cluster.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Device.Properties.NodeCount")},
			{
				Name:        "serial_number",
				Description: "The Serial Number of Data Box Edge/Gateway device.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.SerialNumber")},
			{
				Name:        "sku_name",
				Description: "SKU name of the resource. Possible values include: 'Gateway', 'Edge'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.SKU.Name")},
			{
				Name:        "sku_tier",
				Description: "The SKU tier. This is based on the SKU name. Possible values include: 'Standard'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.SKU.Tier")},
			{
				Name:        "time_zone",
				Description: "The Data Box Edge/Gateway device timezone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Properties.TimeZone")},
			{
				Name:        "configured_role_types",
				Description: "Type of compute roles configured.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Device.Properties.ConfiguredRoleTypes")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Device.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Device.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Device.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Device.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				//// HYDRATE FUNCTIONS
				Transform: transform.

					// Return nil, if no input provide
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Create session
