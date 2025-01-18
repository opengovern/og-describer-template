package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureRecoveryServicesVault(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_recovery_services_vault",
		Description: "Azure Recovery Services Vault",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetRecoveryServicesVault,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "Invalid input"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRecoveryServicesVault,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the recovery services vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Properties.ProvisioningState")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Etag")},
			{
				Name:        "private_endpoint_state_for_site_recovery",
				Description: "Private endpoint state for site recovery of the recovery services vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Properties.PrivateEndpointStateForSiteRecovery")},
			{
				Name:        "private_endpoint_state_for_backup",
				Description: "Private endpoint state for backup of the recovery services vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Properties.PrivateEndpointStateForBackup")},
			{
				Name:        "sku_name",
				Description: "The sku name of the recovery services vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.SKU.Name")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the recovery services vault.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.DiagnosticSettingsResource")},
			{
				Name:        "identity",
				Description: "Managed service identity of the recovery services vault.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vault.Identity")},
			{
				Name:        "private_endpoint_connections",
				Description: "List of private endpoint connections of the recovery services vault.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Vault.Properties.PrivateEndpointConnections")},
			{
				Name:        "upgrade_details",
				Description: "Upgrade details properties of the recovery services vault.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Vault.Properties.UpgradeDetails")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vault.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Vault.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

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

// Create session

// Return nil, if no input provide

// Create session

// If we return the API response directly, the output only gives
// the contents of DiagnosticSettings
