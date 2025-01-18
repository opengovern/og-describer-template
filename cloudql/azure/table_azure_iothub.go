package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureIotHub(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_iothub",
		Description: "Azure Iot Hub",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetIOTHub,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIOTHub,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.ID")},
			{
				Name:        "state",
				Description: "The iot hub state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.State")},
			{
				Name:        "provisioning_state",
				Description: "The iot hub provisioning state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Type")},
			{
				Name:        "comments",
				Description: "Iot hub comments.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.Comments")},
			{
				Name:        "enable_file_upload_notifications",
				Description: "Indicates if file upload notifications are enabled for the iot hub.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.EnableFileUploadNotifications")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Etag")},
			{
				Name:        "features",
				Description: "The capabilities and features enabled for the iot hub. Possible values include: 'None', 'DeviceManagement'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.Features")},
			{
				Name:        "host_name",
				Description: "The name of the host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.HostName")},
			{
				Name:        "min_tls_version",
				Description: "Specifies the minimum TLS version to support for this iot hub.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.MinTLSVersion")},
			{
				Name:        "public_network_access",
				Description: "Indicates whether requests from public network are allowed.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.IotHubDescription.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "sku_capacity",
				Description: "Iot hub SKU capacity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.SKU.Capacity")},
			{
				Name:        "sku_name",
				Description: "Iot hub SKU name.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.IotHubDescription.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Iot hub SKU tier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.SKU.Tier")},
			{
				Name:        "authorization_policies",
				Description: "The shared access policies you can use to secure a connection to the iot hub.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.AuthorizationPolicies")},
			{
				Name:        "cloud_to_device",
				Description: "CloudToDevice properties of the iot hub.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.CloudToDevice")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the iot hub.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "event_hub_endpoints",
				Description: "The event hub-compatible endpoint properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.EventHubEndpoints")},
			{
				Name:        "ip_filter_rules",
				Description: "The IP filter rules of the iot hub.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.IPFilterRules")},
			{
				Name:        "locations",
				Description: "Primary and secondary location for iot hub.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.Locations")},
			{
				Name:        "messaging_endpoints",
				Description: "The messaging endpoint properties for the file upload notification queue.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.MessagingEndpoints")},
			{
				Name:        "private_endpoint_connections",
				Description: "Private endpoint connections created on this iot hub.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.PrivateEndpointConnections")},
			{
				Name:        "routing",
				Description: "Routing properties of the iot hub.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Properties.Routing")},
			{
				Name:        "storage_endpoints",
				Description: "The list of azure storage endpoints where you can upload files.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.IotHubDescription.Properties.StorageEndpoints")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IotHubDescription.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IotHubDescription.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.IotHubDescription.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.IotHubDescription.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTIONS

// Create session

// Return nil, if no input provide

// Create session

// If we return the API response directly, the output only gives
// the contents of DiagnosticSettings
