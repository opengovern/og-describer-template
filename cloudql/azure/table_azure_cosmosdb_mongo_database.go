package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureCosmosDBMongoDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_cosmosdb_mongo_database",
		Description: "Azure Cosmos DB Mongo Database",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"account_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetCosmosdbMongoDatabase,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "NotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCosmosdbMongoDatabase,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the Mongo DB database.",
				Transform:   transform.FromField("Description.MongoDatabase.Name")},
			{
				Name:        "account_name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the database account in which the database is created.",
				Transform:   transform.FromField("Description.Account.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a Mongo DB database uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoDatabase.ID"),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoDatabase.Type")},
			{
				Name:        "autoscale_settings_max_throughput",
				Description: "Contains maximum throughput, the resource can scale up to.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MongoDatabase.Properties.Options.AutoscaleSettings.MaxThroughput")},
			{
				Name:        "database_etag",
				Description: "A system generated property representing the resource etag required for optimistic concurrency control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoDatabase.Properties.Resource.Etag")},
			{
				Name:        "database_id",
				Description: "Name of the Cosmos DB MongoDB database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoDatabase.Properties.Resource.ID")},
			{
				Name:        "database_rid",
				Description: "A system generated unique identifier for database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoDatabase.Properties.Resource.Rid")},
			{
				Name:        "database_ts",
				Description: "A system generated property that denotes the last updated timestamp of the resource.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.MongoDatabase.Properties.Resource.Ts").Transform(transform.ToInt),
			},
			{
				Name:        "throughput",
				Description: "Contains the value of the Cosmos DB resource throughput or autoscaleSettings.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MongoDatabase.Properties.Options.Throughput")},
			{
				Name:        "throughput_settings",
				Description: "Contains the value of the Cosmos DB resource throughput or autoscaleSettings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MongoDatabase.Properties.Options.Throughput.ThroughputSettingsGetResults.Properties.Resource"),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoDatabase.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MongoDatabase.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.MongoDatabase.ID").Transform(idToAkas),
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
