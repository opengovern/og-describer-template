package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureMySQLServer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_mysql_server",
		Description: "Azure MySQL Server",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetMysqlServer,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404", "InvalidApiVersionParameter"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMysqlServer,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a server uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.ID")},
			{
				Name:        "type",
				Description: "The resource type of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Type")},
			{
				Name:        "state",
				Description: "The state of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.UserVisibleState")},
			{
				Name:        "version",
				Description: "Specifies the version of the server.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.Version"),
			},
			{
				Name:        "user_visible_state",
				Description: "A state of a server that is visible to user. Possible values include: 'Ready', 'Dropping', 'Disabled', 'Inaccessible'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.UserVisibleState"),
			},
			{
				Name:        "location",
				Description: "The resource location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Location")},
			{
				Name:        "administrator_login",
				Description: "Specifies the username of the administrator for this server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.AdministratorLogin")},
			{
				Name:        "backup_retention_days",
				Description: "Backup retention days for the server.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Server.Properties.StorageProfile.BackupRetentionDays")},
			{
				Name:        "byok_enforcement",
				Description: "Status showing whether the server data encryption is enabled with customer-managed keys.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.ByokEnforcement")},
			{
				Name:        "earliest_restore_date",
				Description: "Specifies the earliest restore point creation time.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Server.Properties.EarliestRestoreDate").Transform(convertDateToTime),
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.FullyQualifiedDomainName")},
			{
				Name:        "geo_redundant_backup",
				Description: "Indicates whether Geo-redundant is enabled, or not for server backup.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.StorageProfile.GeoRedundantBackup"),
			},
			{
				Name:        "infrastructure_encryption",
				Description: "Status showing whether the server enabled infrastructure encryption. Possible values include: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.InfrastructureEncryption"),
			},
			{
				Name:        "master_server_id",
				Description: "The master server id of a replica server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.MasterServerID")},
			{
				Name:        "minimal_tls_version",
				Description: "Enforce a minimal Tls version for the server. Possible values include: 'TLS10', 'TLS11', 'TLS12', 'TLSEnforcementDisabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.MinimalTLSVersion"),
			},
			{
				Name:        "public_network_access",
				Description: "Indicates whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'. Possible values include: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "replica_capacity",
				Description: "The maximum number of replicas that a master server can have.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Server.Properties.ReplicaCapacity")},
			{
				Name:        "replication_role",
				Description: "The replication role of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.ReplicationRole")},
			{
				Name:        "sku_capacity",
				Description: "The scale up/out capacity, representing server's compute units.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Server.SKU.Capacity")},
			{
				Name:        "sku_family",
				Description: "The family of hardware.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.SKU.Family")},
			{
				Name:        "sku_name",
				Description: "The name of the sku. For example: 'B_Gen4_1', 'GP_Gen5_8'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.SKU.Name")},
			{
				Name:        "sku_size",
				Description: "The size code, to be interpreted by resource as appropriate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.SKU.Size")},
			{
				Name:        "sku_tier",
				Description: "The tier of the particular SKU. Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.SKU.Tier"),
			},
			{
				Name:        "ssl_enforcement",
				Description: "Enable ssl enforcement or not when connect to server. Possible values include: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.SSLEnforcement"),
			},
			{
				Name:        "storage_auto_grow",
				Description: "Indicates whether storage auto grow is enabled, or not.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.StorageProfile.StorageAutogrow"),
			},
			{
				Name:        "storage_mb",
				Description: "Indicates max storage allowed for a server.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Server.Properties.StorageProfile.StorageMB")},
			{
				Name:        "private_endpoint_connections",
				Description: "A list of private endpoint connections on a server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractMySQLServerPrivateEndpointConnections),
			},
			{
				Name:        "server_configurations",
				Description: "The server configurations(parameters) details of the server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Configurations")},
			{
				Name:        "server_keys",
				Description: "The server keys of the server.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.ServerKeys")},
			{
				Name:        "server_security_alert_policy",
				Description: "Security alert policy associated with the MySQL Server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityAlertPolicies"),
			},
			{
				Name:        "vnet_rules",
				Description: "Rules represented by VNET.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VnetRules"),
			},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Server.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Server.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.

					// Currently the API does not support pagination
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// Error: mysql.ServersClient#Get: Invalid input: autorest/validation: validation failed: parameter=resourceGroupName
// constraint=MinLength value="" details: value length must be greater than or equal to 1

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

//// TRANSFORM FUNCTION

// If we return the API response directly, the output will not provide the properties of PrivateEndpointConnections
func extractMySQLServerPrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	server := d.HydrateItem.(opengovernance.MysqlServer).Description.Server
	var properties []map[string]interface{}

	if server.Properties.PrivateEndpointConnections != nil {
		for _, i := range server.Properties.PrivateEndpointConnections {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Properties != nil {
				if i.Properties.PrivateEndpoint != nil {
					objectMap["privateEndpointPropertyId"] = i.Properties.PrivateEndpoint.ID
				}
				if i.Properties.PrivateLinkServiceConnectionState != nil {
					if i.Properties.PrivateLinkServiceConnectionState.ActionsRequired != nil {
						if len(*i.Properties.PrivateLinkServiceConnectionState.ActionsRequired) > 0 {
							objectMap["privateLinkServiceConnectionStateActionsRequired"] = i.Properties.PrivateLinkServiceConnectionState.ActionsRequired
						}
					}
					if i.Properties.PrivateLinkServiceConnectionState.Status != nil {
						if len(*i.Properties.PrivateLinkServiceConnectionState.Status) > 0 {
							objectMap["privateLinkServiceConnectionStateStatus"] = i.Properties.PrivateLinkServiceConnectionState.Status
						}
					}
					if i.Properties.PrivateLinkServiceConnectionState.Description != nil {
						objectMap["privateLinkServiceConnectionStateDescription"] = i.Properties.PrivateLinkServiceConnectionState.Description
					}
				}
				if len(*i.Properties.ProvisioningState) > 0 {
					objectMap["provisioningState"] = i.Properties.ProvisioningState
				}
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}
