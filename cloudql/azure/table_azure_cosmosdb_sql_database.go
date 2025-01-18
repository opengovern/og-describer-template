package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureCosmosDBSQLDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_cosmosdb_sql_database",
		Description: "Azure Cosmos DB SQL Database",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"account_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetCosmosdbSqlDatabase,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "NotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCosmosdbSqlDatabase,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the sql database.",
				Transform:   transform.FromField("Description.SqlDatabase.Name")},
			{
				Name:        "account_name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the database account in which the database is created.",
				Transform:   transform.FromField("Description.Account.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a sql database uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SqlDatabase.ID"),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SqlDatabase.Type")},
			{
				Name:        "autoscale_settings_max_throughput",
				Description: "Contains maximum throughput, the resource can scale up to.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.SqlDatabase.Properties.Options.AutoscaleSettings.MaxThroughput")},
			{
				Name:        "database_colls",
				Description: "A system generated property that specified the addressable path of the collections resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SqlDatabase.Properties.Resource.Colls")},
			{
				Name:        "database_etag",
				Description: "A system generated property representing the resource etag required for optimistic concurrency control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SqlDatabase.Properties.Resource.Etag")},
			{
				Name:        "database_id",
				Description: "Name of the Cosmos DB SQL database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SqlDatabase.Properties.Resource.ID")},
			{
				Name:        "database_rid",
				Description: "A system generated unique identifier for database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SqlDatabase.Properties.Resource.Rid")},
			{
				Name:        "database_ts",
				Description: "A system generated property that denotes the last updated timestamp of the resource.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.SqlDatabase.Properties.Resource.Ts").Transform(transform.ToInt),
			},
			{
				Name:        "database_users",
				Description: "A system generated property that specifies the addressable path of the users resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SqlDatabase.Properties.Resource.Users")},
			{
				Name:        "throughput",
				Description: "Contains the value of the Cosmos DB resource throughput or autoscaleSettings.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.SqlDatabase.Properties.Options.Throughput")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SqlDatabase.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SqlDatabase.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.SqlDatabase.ID").Transform(idToAkas),
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

				Transform: transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}
