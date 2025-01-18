package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureCosmosDBMongoCollection(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_cosmosdb_mongo_collection",
		Description: "Azure Cosmos DB Mongo Collection",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"account_name", "name", "resource_group", "database_name"}),
			Hydrate:    opengovernance.GetCosmosdbMongoCollection,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "NotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCosmosdbMongoCollection,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name: "database_name", Require: plugin.Required,
				},
				{
					Name: "account_name", Require: plugin.Optional,
				},
			},
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the Mongo DB collection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoCollection.Name"),
			},
			{
				Name:        "account_name",
				Description: "The friendly name that identifies the cosmosdb account in which the collection is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name"),
			},
			{
				Name:        "database_name",
				Description: "The friendly name that identifies the database in which the collection is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoDatabase.Name"),
			},
			{
				Name:        "id",
				Description: "Contains ID to identify a Mongo DB collection uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoCollection.ID"),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoCollection.Type"),
			},
			{
				Name:        "analytical_storage_ttl",
				Description: "Analytical TTL.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MongoCollection.Properties.Resource.AnalyticalStorageTTL"),
			},
			{
				Name:        "autoscale_settings_max_throughput",
				Description: "Contains maximum throughput, the resource can scale up to.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MongoCollection.Properties.AutoscaleSettings.MaxThroughput"),
			},
			{
				Name:        "collection_etag",
				Description: "A system generated property representing the resource etag required for optimistic concurrency control.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoCollection.Properties.Resource.Etag"),
			},
			{
				Name:        "collection_id",
				Description: "Name of the Cosmos DB MongoDB collection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoCollection.Properties.Resource.ID"),
			},
			{
				Name:        "collection_rid",
				Description: "A system generated unique identifier for collection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoCollection.Properties.Resource.Rid"),
			},
			{
				Name:        "collection_ts",
				Description: "A system generated property that denotes the last updated timestamp of the resource.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MongoCollection.Properties.Resource.Ts").Transform(transform.ToInt),
			},
			{
				Name:        "shard_key",
				Description: "A key-value pair of shard keys to be applied for the request.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MongoCollection.Properties.Resource.ShardKey"),
			},
			{
				Name:        "indexes",
				Description: "List of index keys.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MongoCollection.Properties.Resource.Indexes"),
			},
			{
				Name:        "throughput",
				Description: "Contains the value of the Cosmos DB resource throughput.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.MongoCollection.Properties.Options.Throughput"),
			},
			{
				Name:        "throughput_settings",
				Description: "Contains the Cosmos DB resource throughput or autoscaleSettings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Throughput"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoCollection.Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MongoCollection.Tags"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MongoCollection.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MongoCollection.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
