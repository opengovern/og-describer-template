package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureIotHubDps(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_iothub_dps",
		Description: "Azure Iot Hub Dps",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetIOTHubDps,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIOTHubDps,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.ID")},
			{
				Name:        "state",
				Description: "Current state of the provisioning service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Properties.State")},
			{
				Name:        "provisioning_state",
				Description: "The ARM provisioning state of the provisioning service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Type")},
			{
				Name:        "allocation_policy",
				Description: "Allocation policy to be used by this provisioning service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Properties.AllocationPolicy")},
			{
				Name:        "device_provisioning_host_name",
				Description: "Device endpoint for this provisioning service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Properties.DeviceProvisioningHostName")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Etag")},
			{
				Name:        "id_scope",
				Description: "Unique identifier of this provisioning service..",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Properties.IDScope")},
			{
				Name:        "service_operations_host_name",
				Description: "Service endpoint for provisioning service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Properties.ServiceOperationsHostName")},
			{
				Name:        "sku_capacity",
				Description: "Iot dps SKU capacity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.SKU.Capacity")},
			{
				Name:        "sku_name",
				Description: "Iot dps SKU name.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.IotHubDps.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Iot dps SKU tier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.SKU.Tier")},
			{
				Name:        "authorization_policies",
				Description: "List of authorization keys for a provisioning service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDps.Properties.AuthorizationPolicies")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the iot dps.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "iot_hubs",
				Description: "List of IoT hubs associated with this provisioning service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDps.Properties.IotHubs")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDps.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDps.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.IotHubDps.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.IotHubDps.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
