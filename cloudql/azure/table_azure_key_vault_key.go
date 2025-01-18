package azure

import (
	"context"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureKeyVaultKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_key_vault_key",
		Description: "Azure Key Vault Key",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"vault_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetKeyVaultKey,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKeyVaultKey,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a key uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.ID"),
			},
			{
				Name:        "vault_name",
				Description: "The friendly name that identifies the vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(extractVaultNameFromID),
			},
			{
				Name:        "enabled",
				Description: "Indicates whether the key is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Key.Properties.Attributes.Enabled")},
			{
				Name:        "key_type",
				Description: "The type of the key. Possible values are: 'EC', 'ECHSM', 'RSA', 'RSAHSM'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Properties.Kty"),
			},
			{
				Name:        "created_at",
				Description: "Specifies the time when the key is created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Key.Properties.Attributes.Created").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "curve_name",
				Description: "The elliptic curve name. Possible values are: 'P256', 'P384', 'P521', 'P256K'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Key.Properties.CurveName"),
			},
			{
				Name:        "expires_at",
				Description: "Specifies the time when the key wil expire.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Key.Properties.Attributes.Expires").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "key_size",
				Description: "The key size in bits.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Key.Properties.KeySize")},
			{
				Name:        "key_uri",
				Description: "The URI to retrieve the current version of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Properties.KeyURI")},
			{
				Name:        "key_uri_with_version",
				Description: "The URI to retrieve the specific version of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Properties.KeyURIWithVersion")},
			{
				Name:        "location",
				Description: "Azure location of the key vault resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Location")},
			{
				Name:        "not_before",
				Description: "Specifies the time before which the key is not usable.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Key.Properties.Attributes.NotBefore").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "recovery_level",
				Description: "The deletion recovery level currently in effect for the object. If it contains 'Purgeable', then the object can be permanently deleted by a privileged user; otherwise, only the system can purge the object at the end of the retention interval.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Key.Properties.Attributes.RecoveryLevel"),
			},
			{
				Name:        "type",
				Description: "Type of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Type")},
			{
				Name:        "updated_at",
				Description: "Specifies the time when the key was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Key.Properties.Attributes.Updated").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "key_ops",
				Description: "A list of key operations.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Key.Properties.KeyOps")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Key.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Key.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Key.Location").Transform(toLower),
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

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// Create session

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

//// TRANSFORM FUNCTIONS

func extractVaultNameFromID(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.KeyVaultKey).Description

	plugin.Logger(ctx).Error("getStorageTable", data)
	if data.Key.ID == nil {
		return data.Vault.Name, nil
	}

	vaultName := strings.Split(*data.Key.ID, "/")[8]
	return vaultName, nil
}
