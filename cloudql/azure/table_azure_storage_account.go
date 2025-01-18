package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureStorageAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_storage_account",
		Description: "Azure Storage Account",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetStorageAccount,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageAccount,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the storage account.",
				Transform:   transform.FromField("Description.Account.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a storage account uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.ID"),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Type")},
			{
				Name:        "access_tier",
				Description: "The access tier used for billing.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Properties.AccessTier"),
			},
			{
				Name:        "kind",
				Description: "The kind of the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Kind"),
			},
			{
				Name:        "sku_name",
				Description: "Contains sku name of the storage account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Contains sku tier of the storage account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.SKU.Tier"),
			},
			{
				Name:        "creation_time",
				Description: "Creation date and time of the storage account.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Account.Properties.CreationTime").Transform(convertDateToTime),
			},
			{
				Name:        "allow_blob_public_access",
				Description: "Specifies whether allow or disallow public access to all blobs or containers in the storage account.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Account.Properties.AllowBlobPublicAccess")},
			{
				Name:        "blob_change_feed_enabled",
				Description: "Specifies whether change feed event logging is enabled for the Blob service.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.BlobServiceProperties.BlobServiceProperties.ChangeFeed.Enabled")},
			{
				Name:        "blob_container_soft_delete_enabled",
				Description: "Specifies whether DeleteRetentionPolicy is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.BlobServiceProperties.BlobServiceProperties.ContainerDeleteRetentionPolicy.Enabled")},
			{
				Name:        "blob_container_soft_delete_retention_days",
				Description: "Indicates the number of days that the deleted item should be retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.BlobServiceProperties.BlobServiceProperties.ContainerDeleteRetentionPolicy.Days")},
			{
				Name:        "blob_restore_policy_days",
				Description: "Specifies how long the blob can be restored.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.BlobServiceProperties.BlobServiceProperties.RestorePolicy.Days")},
			{
				Name:        "blob_restore_policy_enabled",
				Description: "Specifies whether blob restore is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.BlobServiceProperties.BlobServiceProperties.RestorePolicy.Enabled")},
			{
				Name:        "blob_service_logging",
				Description: "Specifies the blob service properties for logging access.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Logging")},
			{
				Name:        "blob_soft_delete_enabled",
				Description: "Specifies whether DeleteRetentionPolicy is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.BlobServiceProperties.BlobServiceProperties.ContainerDeleteRetentionPolicy.Enabled")},
			{
				Name:        "blob_soft_delete_retention_days",
				Description: "Indicates the number of days that the deleted item should be retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.BlobServiceProperties.BlobServiceProperties.ContainerDeleteRetentionPolicy.Days")},
			{
				Name:        "blob_versioning_enabled",
				Description: "Specifies whether versioning is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.BlobServiceProperties.BlobServiceProperties.IsVersioningEnabled")},
			{
				Name:        "enable_https_traffic_only",
				Description: "Allows https traffic only to storage service if sets to true.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Account.Properties.EnableHTTPSTrafficOnly")},
			{
				Name:        "encryption_key_source",
				Description: "Contains the encryption keySource (provider).",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Properties.Encryption.KeySource"),
			},
			{
				Name:        "encryption_key_vault_properties_key_current_version_id",
				Description: "The object identifier of the current versioned Key Vault Key in use.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.Encryption.KeyVaultProperties.CurrentVersionedKeyIdentifier")},
			{
				Name:        "encryption_key_vault_properties_key_name",
				Description: "The name of KeyVault key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.Encryption.KeyVaultProperties.KeyName")},
			{
				Name:        "encryption_key_vault_properties_key_vault_uri",
				Description: "The Uri of KeyVault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.Encryption.KeyVaultProperties.KeyVaultURI")},
			{
				Name:        "encryption_key_vault_properties_key_version",
				Description: "The version of KeyVault key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.Encryption.KeyVaultProperties.KeyVersion")},
			{
				Name:        "encryption_key_vault_properties_last_rotation_time",
				Description: "Timestamp of last rotation of the Key Vault Key.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Account.Properties.Encryption.KeyVaultProperties.LastKeyRotationTimestamp").Transform(convertDateToTime),
			},
			{
				Name:        "failover_in_progress",
				Description: "Specifies whether the failover is in progress.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Account.Properties.FailoverInProgress")},
			{
				Name:        "file_soft_delete_enabled",
				Description: "Specifies whether DeleteRetentionPolicy is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.FileServiceProperties.FileServiceProperties.ShareDeleteRetentionPolicy.Enabled")},
			{
				Name:        "file_soft_delete_retention_days",
				Description: "Indicates the number of days that the deleted item should be retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FileServiceProperties.FileServiceProperties.ShareDeleteRetentionPolicy.Days")},
			{
				Name:        "is_hns_enabled",
				Description: "Specifies whether account HierarchicalNamespace is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Account.Properties.IsHnsEnabled")},
			{
				Name:        "queue_logging_delete",
				Description: "Specifies whether all delete requests should be logged.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.Delete")},
			{
				Name:        "queue_logging_read",
				Description: "Specifies whether all read requests should be logged.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.Read")},
			{
				Name:        "queue_logging_retention_days",
				Description: "Indicates the number of days that metrics or logging data should be retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Logging.RetentionPolicy.Days")},
			{
				Name:        "queue_logging_retention_enabled",
				Description: "Specifies whether a retention policy is enabled for the storage service.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Logging.RetentionPolicy.Enabled")},
			{
				Name:        "queue_logging_version",
				Description: "The version of Storage Analytics to configure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.Version")},
			{
				Name:        "queue_logging_write",
				Description: "Specifies whether all write requests should be logged.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.Write")},
			{
				Name:        "table_logging_read",
				Description: "Indicates whether all read requests should be logged.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.Read"),
			},
			{
				Name:        "table_logging_write",
				Description: "Indicates whether all write requests should be logged.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.Write"),
			},
			{
				Name:        "table_logging_delete",
				Description: "Indicates whether all delete requests should be logged.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.Delete"),
			},
			{
				Name:        "table_logging_version",
				Description: "The version of Analytics to configure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.Version"),
			},
			{
				Name:        "table_logging_retention_policy",
				Description: "The retention policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.StorageServiceProperties.Logging.RetentionPolicy"),
			},
			{
				Name:        "minimum_tls_version",
				Description: "Contains the minimum TLS version to be permitted on requests to storage.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Properties.MinimumTLSVersion"),
			},
			{
				Name:        "network_rule_bypass",
				Description: "Specifies whether traffic is bypassed for Logging/Metrics/AzureServices.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Properties.NetworkRuleSet.Bypass"),
			},
			{
				Name:        "network_rule_default_action",
				Description: "Specifies the default action of allow or deny when no other rules match.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Properties.NetworkRuleSet.DefaultAction"),
			},
			{
				Name:        "primary_blob_endpoint",
				Description: "Contains the blob endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PrimaryEndpoints.Blob")},
			{
				Name:        "primary_dfs_endpoint",
				Description: "Contains the dfs endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PrimaryEndpoints.Dfs")},
			{
				Name:        "primary_file_endpoint",
				Description: "Contains the file endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PrimaryEndpoints.File")},
			{
				Name:        "primary_location",
				Description: "Contains the location of the primary data center for the storage account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PrimaryLocation")},
			{
				Name:        "primary_queue_endpoint",
				Description: "Contains the queue endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PrimaryEndpoints.Queue")},
			{
				Name:        "primary_table_endpoint",
				Description: "Contains the table endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PrimaryEndpoints.Table")},
			{
				Name:        "primary_web_endpoint",
				Description: "Contains the web endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PrimaryEndpoints.Web")},
			{
				Name:        "public_network_access",
				Description: "Allow or disallow public network access to Storage Account. Value is optional but if passed in, must be Enabled or Disabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "status_of_primary",
				Description: "The status indicating whether the primary location of the storage account is available or unavailable. Possible values include: 'available', 'unavailable'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.StatusOfPrimary"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the storage account resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Properties.ProvisioningState"),
			},
			{
				Name:        "require_infrastructure_encryption",
				Description: "Specifies whether or not the service applies a secondary layer of encryption with platform managed keys for data at rest.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Account.Properties.Encryption.RequireInfrastructureEncryption")},
			{
				Name:        "secondary_location",
				Description: "Contains the location of the geo-replicated secondary for the storage account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.SecondaryLocation")},
			{
				Name:        "status_of_secondary",
				Description: "The status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS. Possible values include: 'available', 'unavailable'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.StatusOfSecondary"),
			},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the storage account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractStorageAccountDiagnosticSettings),
			},
			{
				Name:        "encryption_scope",
				Description: "Encryption scope details for the storage account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractStorageAccountEncryptionScope)},
			{
				Name:        "encryption_services",
				Description: "A list of services which support encryption.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.Encryption.Services")},
			{
				Name:        "lifecycle_management_policy",
				Description: "The managementpolicy associated with the specified storage account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractAzureStorageAccountLifecycleManagementPolicy),
			},
			{
				Name:        "network_ip_rules",
				Description: "A list of IP ACL rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.NetworkRuleSet.IPRules")},
			{
				Name:        "private_endpoint_connections",
				Description: "A list of private endpoint connection associated with the specified storage account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.PrivateEndpointConnections")},
			{
				Name:        "table_properties",
				Description: "Azure Analytics Logging settings of tables.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TableProperties")},
			{
				Name:        "access_keys",
				Description: "The list of access keys or Kerberos keys (if active directory enabled) for the specified storage account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AccessKeys")},
			{
				Name:        "virtual_network_rules",
				Description: "A list of virtual network rules.",
				Type:        proto.ColumnType_JSON,
				Transform:
				// Steampipe standard columns
				transform.FromField("Description.Account.Properties.NetworkRuleSet.VirtualNetworkRules")},
			{
				Name:        "sas_policy",
				Description: "A list of virtual network rules.",
				Type:        proto.ColumnType_JSON,
				Transform:
				// Steampipe standard columns
				transform.FromField("Description.Account.Properties.SasPolicy")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns
				Transform: transform.FromField("Description.Account.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.

					//// HYDRATE FUNCTIONS
					FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

func extractAzureStorageAccountLifecycleManagementPolicy(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	objectMap := make(map[string]interface{})
	if d.HydrateItem == nil {
		return objectMap, nil
	}
	op := d.HydrateItem.(opengovernance.StorageAccount).Description.ManagementPolicy
	if op == nil {
		return objectMap, nil
	}

	// Direct assignment returns ManagementPolicyProperties only
	if op.ID != nil {
		objectMap["id"] = op.ID
	}
	if op.Name != nil {
		objectMap["name"] = op.Name
	}
	if op.Type != nil {
		objectMap["type"] = op.Type
	}
	if op.Properties != nil {
		objectMap["properties"] = op.Properties
	}

	return objectMap, nil
}

func extractStorageAccountDiagnosticSettings(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	op := d.HydrateItem.(opengovernance.StorageAccount).Description.DiagnosticSettingsResources

	var diagnosticSettings []map[string]interface{}
	for _, i := range op {
		objectMap := make(map[string]interface{})
		if i.ID != nil {
			objectMap["ID"] = i.ID
		}
		if i.Name != nil {
			objectMap["Name"] = i.Name
		}
		if i.Type != nil {
			objectMap["Type"] = i.Type
		}
		if i.Properties != nil {
			objectMap["DiagnosticSettings"] = i.Properties
		}
		diagnosticSettings = append(diagnosticSettings, objectMap)
	}

	return diagnosticSettings, nil
}

func extractStorageAccountEncryptionScope(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	scopes := d.HydrateItem.(opengovernance.StorageAccount).Description.EncryptionScopes

	var res []interface{}

	for _, scope := range scopes {
		objMap := make(map[string]interface{})
		if scope.ID != nil {
			objMap["Id"] = scope.ID
		}
		if scope.Name != nil {
			objMap["Name"] = scope.Name
		}
		if scope.Type != nil {
			objMap["Type"] = scope.Type
		}
		if scope.EncryptionScopeProperties != nil {
			if *scope.EncryptionScopeProperties.Source != "" {
				objMap["Source"] = scope.EncryptionScopeProperties.Source
			}
			if *scope.EncryptionScopeProperties.State != "" {
				objMap["State"] = scope.EncryptionScopeProperties.State
			}
			if scope.EncryptionScopeProperties.CreationTime != nil {
				objMap["CreationTime"] = scope.EncryptionScopeProperties.CreationTime
			}
			if scope.EncryptionScopeProperties.LastModifiedTime != nil {
				objMap["LastModifiedTime"] = scope.EncryptionScopeProperties.LastModifiedTime
			}
			if scope.EncryptionScopeProperties.KeyVaultProperties != nil {
				if scope.EncryptionScopeProperties.KeyVaultProperties.KeyURI != nil {
					objMap["KeyURI"] = scope.EncryptionScopeProperties.KeyVaultProperties.KeyURI
				}
			}
		}

		res = append(res, objMap)
	}
	return res, nil
}
