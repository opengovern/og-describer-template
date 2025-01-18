package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureMariaDBServer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_mariadb_server",
		Description: "Azure MariaDB Server",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetMariadbServer,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "400", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMariadbServer,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Name")},
			{
				Name:        "id",
				Description: "A fully qualified resource ID for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.ID")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Type")},
			{
				Name:        "version",
				Description: "Specifies the server version.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.Version"),
			},
			{
				Name:        "geo_redundant_backup_enabled",
				Description: "Indicates whether geo-redundant backup is enabled for server backup, or not.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.StorageProfile.GeoRedundantBackup"),
			},
			{
				Name:        "user_visible_state",
				Description: "A state of a server that is visible to user. Valid values are: 'Ready', 'Dropping', 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.UserVisibleState"),
			},
			{
				Name:        "administrator_login",
				Description: "The administrator's login name of a server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.AdministratorLogin")},
			{
				Name:        "auto_grow_enabled",
				Description: "Indicates whether storage auto grow is enabled for server, or not.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.StorageProfile.StorageAutogrow"),
			},
			{
				Name:        "backup_retention_days",
				Description: "Specifies the backup retention days for the server.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Server.Properties.StorageProfile.BackupRetentionDays")},
			{
				Name:        "earliest_restore_date",
				Description: "Specifies the earliest restore point creation time.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Server.Properties.EarliestRestoreDate").Transform(convertDateToTime),
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of a server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.FullyQualifiedDomainName")},
			{
				Name:        "master_service_id",
				Description: "The master server id of a replica server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.MasterServerID")},
			{
				Name:        "public_network_access",
				Description: "Indicates whether or not public network access is allowed for this server. Valid values are: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "replication_role",
				Description: "The replication role of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.ReplicationRole")},
			{
				Name:        "replica_capacity",
				Description: "The maximum number of replicas that a master server can have.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Server.Properties.ReplicaCapacity")},
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
				Description: "The name of the sku.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.SKU.Name")},
			{
				Name:        "sku_size",
				Description: "The size code, to be interpreted by resource as appropriate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.SKU.Size")},
			{
				Name:        "sku_tier",
				Description: "The tier of the particular SKU. Valid values are: 'Basic', 'GeneralPurpose', 'MemoryOptimized'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.SKU.Tier"),
			},
			{
				Name:        "ssl_enforcement",
				Description: "Indicates whether SSL enforcement is enabled, or not. Valid values are: 'Enabled', and 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Properties.SSLEnforcement"),
			},
			{
				Name:        "storage_mb",
				Description: "Specifies the max storage allowed for a server.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Server.Properties.StorageProfile.StorageMB")},
			{
				Name:        "private_endpoint_connections",
				Description: "A list of private endpoint connections on a server.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Server.Properties.PrivateEndpointConnections")},

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

					// Create session
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// Return nil, if no input provided

// Create session
