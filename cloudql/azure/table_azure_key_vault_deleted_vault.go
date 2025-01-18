package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureKeyVaultDeletedVault(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_key_vault_deleted_vault",
		Description: "Azure Key Vault Deleted Vault",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "region"}),
			Hydrate:    opengovernance.GetKeyVaultDeletedVault,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKeyVaultDeletedVault,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the deleted vault.",
				Transform:   transform.FromField("Description.Vault.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a deleted vault uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.ID")},
			{
				Name:        "vault_id",
				Description: "The resource id of the original vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Properties.VaultID")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Type")},
			{
				Name:        "deletion_date",
				Description: "The deleted date of the vault.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Vault.Properties.DeletionDate").Transform(convertDateToTime),
			},
			{
				Name:        "scheduled_purge_date",
				Description: "The scheduled purged date of the vault.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Vault.Properties.ScheduledPurgeDate").Transform(convertDateToTime),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vault.Properties.Tags")},
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

				Transform: transform.FromField("Description.Vault.Properties.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Properties.VaultID").Transform(extractResourceGroupFromID)},
		}),
	}
}
