package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureDataLakeStore(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_data_lake_store",
		Description: "Azure Data Lake Store",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetDataLakeStore,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDataLakeStore,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Type")},
			{
				Name:        "account_id",
				Description: "The unique identifier associated with this data lake store account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.AccountID")},
			{
				Name:        "creation_time",
				Description: "The account creation time.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.DataLakeStoreAccount.Properties.CreationTime").Transform(convertDateToTime),
			},
			{
				Name:        "current_tier",
				Description: "The commitment tier in use for current month.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.CurrentTier")},
			{
				Name:        "default_group",
				Description: "The default owner group for all new folders and files created in the data lake store account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.DefaultGroup")},
			{
				Name:        "encryption_state",
				Description: "The current state of encryption for this data lake store account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.EncryptionState")},
			{
				Name:        "encryption_provisioning_state",
				Description: "The current state of encryption provisioning for this data lake store account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.EncryptionProvisioningState")},
			{
				Name:        "endpoint",
				Description: "The full cname endpoint for this account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.Endpoint")},
			{
				Name:        "firewall_state",
				Description: "The current state of the IP address firewall for this data lake store account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.FirewallState")},
			{
				Name:        "last_modified_time",
				Description: "The account last modified time.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.DataLakeStoreAccount.Properties.LastModifiedTime").Transform(convertDateToTime),
			},
			{
				Name:        "new_tier",
				Description: "The commitment tier to use for next month.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.NewTier")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning status of the data lake store account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.ProvisioningState")},
			{
				Name:        "state",
				Description: "The state of the data lake store account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.State")},
			{
				Name:        "trusted_id_provider_state",
				Description: "The current state of the trusted identity provider feature for this data lake store account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.TrustedIDProviderState")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the data lake store.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource")},
			{
				Name:        "encryption_config",
				Description: "The key vault encryption configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.EncryptionConfig")},
			{
				Name:        "firewall_allow_azure_ips",
				Description: "The current state of allowing or disallowing IPs originating within azure through the firewall. If the firewall is disabled, this is not enforced.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.FirewallAllowAzureIPs")},
			{
				Name:        "firewall_rules",
				Description: "The list of firewall rules associated with this data lake store account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.FirewallRules")},
			{
				Name:        "identity",
				Description: "The key vault encryption identity, if any.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Identity")},
			{
				Name:        "trusted_id_providers",
				Description: "The list of trusted identity providers associated with this data lake store account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Properties.TrustedIDProviders")},
			{
				Name:        "virtual_network_rules",
				Description: "The list of virtual network rules associated with this data lake store account.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.DataLakeStoreAccount.Properties.VirtualNetworkRules")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeStoreAccount.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.DataLakeStoreAccount.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DataLakeStoreAccount.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
