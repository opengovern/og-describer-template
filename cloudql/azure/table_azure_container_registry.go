package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureContainerRegistry(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_container_registry",
		Description: "Azure Container Registry",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetContainerRegistry,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "Invalid input", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListContainerRegistry,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Registry.Name")},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Registry.ID")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Registry.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the container registry at the time the operation was called. Valid values are: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Failed', 'Canceled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Registry.Properties.ProvisioningState"),
			},
			{
				Name:        "creation_date",
				Description: "The creation date of the container registry.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Registry.Properties.CreationDate").Transform(convertDateToTime),
			},
			{
				Name:        "admin_user_enabled",
				Description: "Indicates whether the admin user is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Registry.Properties.AdminUserEnabled")},
			{
				Name:        "data_endpoint_enabled",
				Description: "Enable a single data endpoint per region for serving data.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Registry.Properties.DataEndpointEnabled")},
			{
				Name:        "login_server",
				Description: "The URL that can be used to log into the container registry.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Registry.Properties.LoginServer")},
			{
				Name:        "network_rule_bypass_options",
				Description: "Indicates whether to allow trusted Azure services to access a network restricted registry. Valid values are: 'AzureServices', 'None'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Registry.Properties.NetworkRuleBypassOptions"),
			},
			{
				Name:        "public_network_access",
				Description: "Indicates whether or not public network access is allowed for the container registry. Valid values are: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Registry.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "sku_name",
				Description: "The SKU name of the container registry. Required for registry creation. Valid values are: 'Classic', 'Basic', 'Standard', 'Premium'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Registry.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The SKU tier based on the SKU name. Valid values are: 'Classic', 'Basic', 'Standard', 'Premium'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Registry.SKU.Tier"),
			},
			{
				Name:        "status",
				Description: "The current status of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Registry.Properties.Status.DisplayStatus")},
			{
				Name:        "status_message",
				Description: "The detailed message for the status, including alerts and error messages.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Registry.Properties.Status.Message")},
			{
				Name:        "status_timestamp",
				Description: "The timestamp when the status was changed to the current value.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Registry.Properties.Status.Timestamp").Transform(convertDateToTime),
			},
			{
				Name:        "storage_account_id",
				Description: "The resource ID of the storage account. Only applicable to Classic SKU.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegistryProperties.StorageAccount.ID"),
			},
			{
				Name:        "zone_redundancy",
				Description: "Indicates whether or not zone redundancy is enabled for this container registry. Valid values are: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Registry.Properties.ZoneRedundancy"),
			},
			{
				Name:        "data_endpoint_host_names",
				Description: "A list of host names that will serve data when dataEndpointEnabled is true.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Registry.Properties.DataEndpointHostNames")},
			{
				Name:        "encryption",
				Description: "The encryption settings of container registry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Registry.Properties.Encryption")},
			{
				Name:        "identity",
				Description: "The identity of the container registry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Registry.Identity")},
			{
				Name:        "login_credentials",
				Description: "The login credentials for the specified container registry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RegistryListCredentialsResult")},
			{
				Name:        "network_rule_set",
				Description: "The network rule set for a container registry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Registry.Properties.NetworkRuleSet")},
			{
				Name:        "policies",
				Description: "The policies for a container registry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Registry.Properties.Policies")},
			{
				Name:        "private_endpoint_connections",
				Description: "A list of private endpoint connections for a container registry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Registry.Properties.PrivateEndpointConnections")},
			{
				Name:        "system_data",
				Description: "Metadata pertaining to creation and last modification of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Registry.SystemData")},
			{
				Name:        "usages",
				Description: "Specifies the quota usages for the specified container registry.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.RegistryUsages")},

			{
				Name:        "webhooks",
				Description: "Webhooks in Azure Container Registry provide a way to trigger custom actions in response to events happening within the registry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Webhooks")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Registry.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Registry.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Registry.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Registry.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				// Create session
				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// Return nil, if no input provided

// Create session

// Create session

// Create session
