package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureComputeDiskEncryptionSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_disk_encryption_set",
		Description: "Azure Compute Disk Encryption Set",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetComputeDiskEncryptionSet,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeDiskEncryptionSet,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the disk encryption set",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskEncryptionSet.Name")},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskEncryptionSet.ID")},
			{
				Name:        "type",
				Description: "The type of the resource in Azure",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskEncryptionSet.Type")},
			{
				Name:        "provisioning_state",
				Description: "The disk encryption set provisioning state",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Properties.ProvisioningState")},
			{
				Name:        "active_key_source_vault_id",
				Description: "Resource id of the KeyVault containing the key or secret",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Properties.ActiveKey.SourceVault.ID")},
			{
				Name:        "active_key_url",
				Description: "Url pointing to a key or secret in KeyVault",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Properties.ActiveKey.KeyURL")},
			{
				Name:        "encryption_type",
				Description: "Contains the type of the encryption",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Properties.EncryptionType"),
			},
			{
				Name:        "identity_principal_id",
				Description: "The object id of the Managed Identity Resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Identity.PrincipalID")},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant id of the Managed Identity Resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Identity.TenantID")},
			{
				Name:        "identity_type",
				Description: "The type of Managed Identity used by the DiskEncryptionSet",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Identity.Type"),
			},
			{
				Name:        "previous_keys",
				Description: "A list of key vault keys previously used by this disk encryption set while a key rotation is in progress",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiskEncryptionSet.Properties.PreviousKeys")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiskEncryptionSet.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.DiskEncryptionSet.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskEncryptionSet.Location").Transform(toLower),
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

//// HYDRATE FUNCTION ////

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data
