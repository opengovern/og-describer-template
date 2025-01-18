package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureMySQLFlexibleServer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_mysql_flexible_server",
		Description: "Azure MySQL Flexible Server",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetSqlServerFlexibleServer,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSqlServerFlexibleServer,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a server uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.ID")},
			{
				Name:        "type",
				Description: "The resource type of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Type")},
			{
				Name:        "state",
				Description: "The state of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.State")},
			{
				Name:        "version",
				Description: "Specifies the version of the server.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FlexibleServer.Properties.Version"),
			},
			{
				Name:        "administrator_login",
				Description: "The administrator's login name of a server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.AdministratorLogin")},
			{
				Name:        "availability_zone",
				Description: "Availability Zone information of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.AvailabilityZone")},
			{
				Name:        "backup_retention_days",
				Description: "Backup retention days for the server.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.Backup.BackupRetentionDays")},
			{
				Name:        "create_mode",
				Description: "The mode to create a new server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.CreateMode")},
			{
				Name:        "earliest_restore_date",
				Description: "Specifies the earliest restore point creation time.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.FlexibleServer.Properties.Backup.EarliestRestoreDate").Transform(convertDateToTime),
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.FullyQualifiedDomainName")},
			{
				Name:        "geo_redundant_backup",
				Description: "Indicates whether Geo-redundant is enabled, or not for server backup.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FlexibleServer.Properties.Backup.GeoRedundantBackup"),
			},
			{
				Name:        "location",
				Description: "The server location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Location")},
			{
				Name:        "public_network_access",
				Description: "Whether or not public network access is allowed for this server.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FlexibleServer.Properties.Network.PublicNetworkAccess"),
			},
			{
				Name:        "replica_capacity",
				Description: "The maximum number of replicas that a primary server can have.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.ReplicaCapacity")},
			{
				Name:        "replication_role",
				Description: "The replication role of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.ReplicationRole")},
			{
				Name:        "restore_point_in_time",
				Description: "Restore point creation time (ISO8601 format), specifying the time to restore from.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.RestorePointInTime")},
			{
				Name:        "sku_name",
				Description: "The name of the sku.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.SKU.Name")},
			{
				Name:        "sku_tier",
				Description: "The tier of the particular SKU.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FlexibleServer.SKU.Tier"),
			},
			{
				Name:        "source_server_resource_id",
				Description: "The source MySQL server id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.SourceServerResourceID")},
			{
				Name:        "storage_auto_grow",
				Description: "Indicates whether storage auto grow is enabled, or not.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FlexibleServer.Properties.Storage.AutoGrow"),
			},
			{
				Name:        "storage_iops",
				Description: "Storage IOPS for a server.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.Storage.Iops")},
			{
				Name:        "storage_size_gb",
				Description: "Indicates max storage allowed for a server.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.Storage.StorageSizeGB")},
			{
				Name:        "storage_sku",
				Description: "The sku name of the server storage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.Storage.StorageSKU")},
			{
				Name:        "flexible_server_configurations",
				Description: "The server configurations(parameters) details of the server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FlexibleServer.Location")},
			{
				Name:        "high_availability",
				Description: "High availability related properties of a server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.HighAvailability")},
			{
				Name:        "network",
				Description: "Network related properties of a server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.Network")},
			{
				Name:        "maintenance_window",
				Description: "Maintenance window of a server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FlexibleServer.Properties.MaintenanceWindow")},
			{
				Name:        "system_data",
				Description: "The system metadata relating to this server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FlexibleServer.SystemData")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlexibleServer.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FlexibleServer.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.FlexibleServer.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FlexibleServer.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
