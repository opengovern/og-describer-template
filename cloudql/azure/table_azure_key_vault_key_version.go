package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureKeyVaultKeyVersion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_key_vault_key_version",
		Description: "Azure Key Vault Key Version",
		List: &plugin.ListConfig{
			Hydrate:       opengovernance.ListKeyVaultKeyVersion,
			ParentHydrate: opengovernance.ListKeyVault,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name: "key_name", Require: plugin.Optional,
				},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the key version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Name"),
			},
			{
				Name:        "key_name",
				Description: "The friendly name that identifies the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Name"),
			},
			{
				Name:        "key_id",
				Description: "Contains ID to identify a key uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.ID"),
			},
			{
				Name:        "id",
				Description: "Contains ID to identify a key version uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.ID"),
			},
			{
				Name:        "vault_name",
				Description: "The friendly name that identifies the vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Name"),
			},
			{
				Name:        "enabled",
				Description: "Indicates whether the key version is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Version.Properties.Attributes.Enabled"),
			},
			{
				Name:        "key_type",
				Description: "The type of the key. Possible values are: 'EC', 'ECHSM', 'RSA', 'RSAHSM'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Properties.Kty"),
			},
			{
				Name:        "created_at",
				Description: "Specifies the time when the key version is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Version.Properties.Attributes.Created").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "curve_name",
				Description: "The elliptic curve name. Possible values are: 'P256', 'P384', 'P521', 'P256K'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Properties.CurveName"),
			},
			{
				Name:        "expires_at",
				Description: "Specifies the time when the key version wil expire.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Version.Properties.Attributes.Expires").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "key_size",
				Description: "The key size in bits.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Version.Properties.KeySize"),
			},
			{
				Name:        "key_uri",
				Description: "The URI to retrieve the current version of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Properties.KeyURI"),
			},
			{
				Name:        "key_uri_with_version",
				Description: "The URI to retrieve the specific version of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Properties.KeyURIWithVersion"),
			},
			{
				Name:        "location",
				Description: "Azure location of the key vault resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Location"),
			},
			{
				Name:        "not_before",
				Description: "Specifies the time before which the key version is not usable.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Version.Properties.Attributes.NotBefore").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "recovery_level",
				Description: "The deletion recovery level currently in effect for the object. If it contains 'Purgeable', then the object can be permanently deleted by a privileged user; otherwise, only the system can purge the object at the end of the retention interval.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Properties.Attributes.RecoveryLevel"),
			},
			{
				Name:        "type",
				Description: "Type of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Type"),
			},
			{
				Name:        "updated_at",
				Description: "Specifies the time when the key was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Version.Properties.Attributes.Updated").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "key_ops",
				Description: "A list of key operations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Version.Properties.JSONWebKeyOperation"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Version.Tags"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Version.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Version.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
