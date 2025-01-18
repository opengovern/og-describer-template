package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureKeyVault(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_key_vault",
		Description: "Azure Key Vault",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetKeyVault,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKeyVault,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the vault.",
				Transform:   transform.FromField("Description.Vault.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a vault uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.ID")},
			{
				Name:        "vault_uri",
				Description: "Contains URI of the vault for performing operations on keys and secrets.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Properties.VaultURI")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Type")},
			{
				Name:        "create_mode",
				Description: "The vault's create mode to indicate whether the vault need to be recovered or not. Possible values include: 'default', 'recover'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Properties.CreateMode")},
			{
				Name:        "enabled_for_deployment",
				Description: "Indicates whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Vault.Properties.EnabledForDeployment"), Default: false,
			},
			{
				Name:        "enabled_for_disk_encryption",
				Description: "Indicates whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Vault.Properties.EnabledForDiskEncryption"), Default: false,
			},
			{
				Name:        "enabled_for_template_deployment",
				Description: "Indicates whether Azure Resource Manager is permitted to retrieve secrets from the key vault.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Vault.Properties.EnabledForTemplateDeployment"), Default: false,
			},
			{
				Name:        "enable_rbac_authorization",
				Description: "Property that controls how data actions are authorized.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Vault.Properties.EnableRbacAuthorization"), Default: false,
			},
			{
				Name:        "purge_protection_enabled",
				Description: "Indicates whether protection against purge is enabled for this vault.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Vault.Properties.EnablePurgeProtection"), Default: false,
			},
			{
				Name:        "soft_delete_enabled",
				Description: "Indicates whether the 'soft delete' functionality is enabled for this key vault.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Vault.Properties.EnableSoftDelete")},
			{
				Name:        "soft_delete_retention_in_days",
				Description: "Contains softDelete data retention days.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Vault.Properties.SoftDeleteRetentionInDays")},
			{
				Name:        "sku_family",
				Description: "Contains SKU family name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Vault.Properties.SKU.Family")},
			{
				Name:        "sku_name",
				Description: "SKU name to specify whether the key vault is a standard vault or a premium vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Properties.SKU.Name"),
			},
			{
				Name:        "tenant_id",
				Description: "The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Vault.Properties.TenantID"),
			},
			{
				Name:        "access_policies",
				Description: "A list of 0 to 1024 identities that have access to the key vault.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractKeyVaultAccessPolicies),
			},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the vault.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "network_acls",
				Description: "Rules governing the accessibility of the key vault from specific network locations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vault.Properties.NetworkACLs")},
			{
				Name:        "private_endpoint_connections",
				Description: "List of private endpoint connections associated with the key vault.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractKeyVaultPrivateEndpointConnections),
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
				Transform:   transform.FromField("Description.Vault.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Vault.ID").Transform(idToAkas),
			},

			// Azure standard columns
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

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type PrivateEndpointConnectionInfo struct {
	PrivateEndpointId                               string
	PrivateLinkServiceConnectionStateStatus         string
	PrivateLinkServiceConnectionStateDescription    string
	PrivateLinkServiceConnectionStateActionRequired string
	ProvisioningState                               string
}

//// TRANSFORM FUNCTIONS

func extractKeyVaultPrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	vault := d.HydrateItem.(opengovernance.KeyVault).Description.Vault
	plugin.Logger(ctx).Trace("extractKeyVaultPrivateEndpointConnections")
	var privateEndpointDetails []PrivateEndpointConnectionInfo
	var privateEndpoint PrivateEndpointConnectionInfo
	if vault.Properties.PrivateEndpointConnections != nil {
		for _, connection := range vault.Properties.PrivateEndpointConnections {
			// Below checks are required for handling invalid memory address or nil pointer dereference error
			if connection.Properties != nil {
				if connection.Properties.PrivateEndpoint != nil {
					privateEndpoint.PrivateEndpointId = *connection.Properties.PrivateEndpoint.ID
				}
				if connection.Properties.PrivateLinkServiceConnectionState != nil {
					if connection.Properties.PrivateLinkServiceConnectionState.ActionsRequired != nil {
						privateEndpoint.PrivateLinkServiceConnectionStateActionRequired = string(*connection.Properties.PrivateLinkServiceConnectionState.ActionsRequired)
					}
					if connection.Properties.PrivateLinkServiceConnectionState.Description != nil {
						privateEndpoint.PrivateLinkServiceConnectionStateDescription = *connection.Properties.PrivateLinkServiceConnectionState.Description
					}
					if connection.Properties.PrivateLinkServiceConnectionState.Status != nil {
						if *connection.Properties.PrivateLinkServiceConnectionState.Status != "" {
							privateEndpoint.PrivateLinkServiceConnectionStateStatus = string(*connection.Properties.PrivateLinkServiceConnectionState.Status)
						}
					}
				}
				if connection.Properties.ProvisioningState != nil {
					if *connection.Properties.ProvisioningState != "" {
						privateEndpoint.ProvisioningState = string(*connection.Properties.ProvisioningState)
					}
				}
			}
			privateEndpointDetails = append(privateEndpointDetails, privateEndpoint)
		}
	}

	return privateEndpointDetails, nil
}

// If we return the API response directly, the output will not provide the properties of AccessPolicies
func extractKeyVaultAccessPolicies(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	vault := d.HydrateItem.(opengovernance.KeyVault).Description.Vault
	var policies []map[string]interface{}

	if vault.Properties.AccessPolicies != nil {
		for _, i := range vault.Properties.AccessPolicies {
			objectMap := make(map[string]interface{})
			if i.TenantID != nil {
				objectMap["tenantId"] = i.TenantID
			}
			if i.ObjectID != nil {
				objectMap["objectId"] = i.ObjectID
			}
			if i.ApplicationID != nil {
				objectMap["applicationId"] = i.ApplicationID
			}
			if i.Permissions != nil {
				if i.Permissions.Keys != nil {
					objectMap["permissionsKeys"] = i.Permissions.Keys
				}
				if i.Permissions.Secrets != nil {
					objectMap["permissionsSecrets"] = i.Permissions.Secrets
				}
				if i.Permissions.Keys != nil {
					objectMap["permissionsCertificates"] = i.Permissions.Certificates
				}
				if i.Permissions.Keys != nil {
					objectMap["permissionsStorage"] = i.Permissions.Storage
				}
			}
			policies = append(policies, objectMap)
		}
	}

	return policies, nil
}
