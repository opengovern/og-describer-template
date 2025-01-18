package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureKeyVaultCertificate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_key_vault_certificate",
		Description: "Azure Key Vault Certificate",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"vault_name", "resource_group"}),
			Hydrate:    opengovernance.GetKeyVaultCertificate,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKeyVaultCertificate,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Contains ID to identify a key uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.ID"),
			},
			{
				Name:        "vault_name",
				Description: "The friendly name that identifies the vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Name"),
			},
			{
				Name:        "attributes",
				Description: "Certificate attributes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Attributes"),
			},
			{
				Name:        "issuer_parameters",
				Description: "Issuer parameters.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.IssuerParameters"),
			},
			{
				Name:        "key_properties",
				Description: "Key properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.KeyProperties"),
			},
			{
				Name:        "lifetime_actions",
				Description: "Lifetime actions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.LifetimeActions"),
			},
			{
				Name:        "secret_properties",
				Description: "Secret properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.SecretProperties"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Policy.ID").Transform(idToAkas),
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

				//// LIST FUNCTION
				// Get the details of key vault
				Transform: transform.

					// Create session
					FromField("Description.ResourceGroup")},
		}),
	}
}
