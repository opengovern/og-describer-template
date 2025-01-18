package azure

import (
	"context"
	"strings"

	secret "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureKeyVaultSecret(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_key_vault_secret",
		Description: "Azure Key Vault Secret",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"vault_name", "name"}),
			Hydrate:    opengovernance.GetKeyVaultSecret,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "404", "SecretDisabled"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKeyVaultSecret,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the secret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecretItem.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a secret uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecretItem.ID")},
			{
				Name:        "vault_name",
				Description: "The friendly name that identifies the vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(extractVaultNameFromSecretID, "VaultName"),
			},
			{
				Name:        "enabled",
				Description: "Indicates whether the secret is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.SecretItem.Properties.Attributes.Enabled")},
			{
				Name:        "managed",
				Description: "Indicates whether the secret's lifetime is managed by key vault, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "content_type",
				Description: "Specifies the type of the secret value such as a password.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecretItem.Properties.ContentType")},
			{
				Name:        "created_at",
				Description: "Specifies the time when the secret is created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.SecretItem.Properties.Attributes.Created").Transform(convertDateUnixToTime),
			},
			{
				Name:        "expires_at",
				Description: "Specifies the time when the secret will expire.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.SecretItem.Properties.Attributes.Expires").Transform(convertDateUnixToTime).Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "kid",
				Description: "If this is a secret backing a KV certificate, then this field specifies the corresponding key backing the KV certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecretBundle.Kid")},
			{
				Name:        "not_before",
				Description: "Specifies the time before which the secret is not usable.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.SecretItem.Properties.Attributes.NotBefore").Transform(convertDateUnixToTime),
			},
			{
				Name:        "recoverable_days",
				Description: "Specifies the soft delete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Vault.Properties.SoftDeleteRetentionInDays"),
			},
			{
				Name:        "recovery_level",
				Description: "The deletion recovery level currently in effect for the object. If it contains 'Purgeable', then the object can be permanently deleted by a privileged user; otherwise, only the system can purge the object at the end of the retention interval.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(extractRecoveryLevel),
			},
			{
				Name:        "updated_at",
				Description: "Specifies the time when the secret was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.SecretItem.Properties.Attributes.Updated").Transform(convertDateUnixToTime),
			},
			{
				Name:        "value",
				Description: "Specifies the secret value.",
				Type:        proto.ColumnType_STRING,

				// Steampipe standard columns
				Transform: transform.FromField("Description.SecretItem.Properties.Value")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(extractVaultNameFromSecretID, "Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecretItem.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns
				Transform: transform.FromField("Description.TurboData.Akas")},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.TurboData.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				// Get the details of key vault

				// Create session

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit
				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.TurboData.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTIONS

func getKeyVaultSecret(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getKeyVaultSecret")

	var vaultName, name string
	if h.Item != nil {
		data := h.Item.(secret.SecretItem)
		splitID := strings.Split(*data.ID, "/")
		vaultName = strings.Split(splitID[2], ".")[0]
		name = splitID[4]

		// Operation get is not allowed on a disabled secret
		if !*data.Attributes.Enabled {
			return nil, nil
		}
	} else {
		vaultName = d.EqualsQuals["vault_name"].GetStringValue()
		name = d.EqualsQuals["name"].GetStringValue()
	}

	// Create session
	session, err := GetNewSession(ctx, d, "VAULT")
	if err != nil {
		return nil, err
	}

	client := secret.New()
	client.Authorizer = session.Authorizer

	vaultURI := "https://" + vaultName + ".vault.azure.net/"

	op, err := client.GetSecret(ctx, vaultURI, name, "")
	if err != nil {
		return nil, err
	}

	// In some cases resource does not give any notFound error
	// instead of notFound error, it returns empty data
	if op.ID != nil {
		return op, nil
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func extractVaultNameFromSecretID(ctx context.Context, d *transform.TransformData) (any, error) {
	secretID := keyVaultSecretData(d.HydrateItem)
	param := d.Param.(string)

	splitID := strings.Split(secretID, "/")
	if len(splitID) < 5 {
		return nil, nil
	}

	result := map[string]string{
		"VaultName": strings.Split(splitID[2], ".")[0],
		"Name":      splitID[4],
	}

	return result[param], nil
}

func extractRecoveryLevel(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	purge := d.HydrateItem.(opengovernance.KeyVaultSecret).Description.Vault.Properties.EnablePurgeProtection

	if purge != nil && *purge {
		return "Purgeable", nil
	} else {
		return "Other", nil
	}
}

func keyVaultSecretData(item interface{}) string {
	switch item := item.(type) {
	case secret.SecretItem:
		if item.ID != nil {
			return *item.ID
		}
	case secret.SecretBundle:
		if item.ID != nil {
			return *item.ID
		}
	}
	return ""
}
