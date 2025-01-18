package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureMSSQLElasticPool(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_mssql_elasticpool",
		Description: "Azure Microsoft SQL Elastic Pool",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group", "server_name"}),
			Hydrate:    opengovernance.GetSqlServerElasticPool,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404", "InvalidApiVersionParameter"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSqlServerElasticPool,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the elastic pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Name")},
			{
				Name:        "server_name",
				Description: "The name of the parent server of the elastic pool.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ServerName")},
			{
				Name:        "id",
				Description: "Contains ID to identify a elastic pool uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.ID")},
			{
				Name:        "type",
				Description: "The resource type of the elastic pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Type")},
			{
				Name:        "state",
				Description: "The state of the elastic pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Properties.State")},
			{
				Name:        "creation_date",
				Description: "The creation date of the elastic pool.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Pool.Properties.CreationDate"),
			},
			{
				Name:        "database_dtu_max",
				Description: "The maximum DTU any one database can consume.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Pool.Properties.PerDatabaseSettings.MaxCapacity")},
			{
				Name:        "database_dtu_min",
				Description: "The minimum DTU all databases are guaranteed.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Pool.Properties.PerDatabaseSettings.MinCapacity")},
			{
				Name:        "dtu",
				Description: "The total shared DTU for the database elastic pool.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.TotalDTU"),
			},
			{
				Name:        "edition",
				Description: "The edition of the elastic pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.SKU.Tier"),
			},
			{
				Name:        "kind",
				Description: "The kind of elastic pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Kind")},
			{
				Name:        "storage_mb",
				Description: "Storage limit for the database elastic pool in MB.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Pool.Properties.MaxSizeBytes")},
			{
				Name:        "zone_redundant",
				Description: "Whether or not this database elastic pool is zone redundant, which means the replicas of this database will be spread across multiple availability zones.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Pool.Properties.ZoneRedundant")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pool.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Pool.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Pool.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
