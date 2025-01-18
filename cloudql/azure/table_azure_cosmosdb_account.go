package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureCosmosDBAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_cosmosdb_account",
		Description: "Azure Cosmos DB Account",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetCosmosdbAccount,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCosmosdbAccount,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the database account.",
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a database account uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.ID"),
			},
			{
				Name:        "kind",
				Description: "Indicates the type of database account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Kind"),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Type")},
			{
				Name:        "connector_offer",
				Description: "The cassandra connector offer type for the Cosmos DB database C* account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.ConnectorOffer"),
			},
			{
				Name:        "consistency_policy_max_interval",
				Description: "The time amount of staleness (in seconds) tolerated, when used with the Bounded Staleness consistency level.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.ConsistencyPolicy.MaxIntervalInSeconds")},
			{
				Name:        "consistency_policy_max_staleness_prefix",
				Description: "The number of stale requests tolerated, when used with the Bounded Staleness consistency level.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.ConsistencyPolicy.MaxStalenessPrefix")},
			{
				Name:        "database_account_offer_type",
				Description: "The offer type for the Cosmos DB database account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.DatabaseAccountOfferType"),
			},
			{
				Name:        "default_consistency_level",
				Description: "The default consistency level and configuration settings of the Cosmos DB account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.ConsistencyPolicy.DefaultConsistencyLevel"),
			},
			{
				Name:        "disable_key_based_metadata_write_access",
				Description: "Disable write operations on metadata resources (databases, containers, throughput) via account keys.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.DisableKeyBasedMetadataWriteAccess"), Default: false,
			},
			{
				Name:        "document_endpoint",
				Description: "The connection endpoint for the Cosmos DB database account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.DocumentEndpoint")},
			{
				Name:        "enable_analytical_storage",
				Description: "Specifies whether to enable storage analytics, or not.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.EnableAnalyticalStorage"), Default: false,
			},
			{
				Name:        "enable_automatic_failover",
				Description: "Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.EnableAutomaticFailover"), Default: false,
			},
			{
				Name:        "enable_cassandra_connector",
				Description: "Enables the cassandra connector on the Cosmos DB C* account.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.EnableCassandraConnector"), Default: false,
			},
			{
				Name:        "enable_free_tier",
				Description: "Specifies whether free Tier is enabled for Cosmos DB database account, or not.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.EnableFreeTier"), Default: false,
			},
			{
				Name:        "enable_multiple_write_locations",
				Description: "Enables the account to write in multiple locations.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.EnableMultipleWriteLocations"), Default: false,
			},
			{
				Name:        "is_virtual_network_filter_enabled",
				Description: "Specifies whether to enable/disable Virtual Network ACL rules.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.IsVirtualNetworkFilterEnabled"), Default: false,
			},
			{
				Name:        "disable_local_auth",
				Description: "Disable local authentication and ensure only MSI and AAD can be used exclusively for authentication. Defaults to false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.DisableLocalAuth"),
				Default:     false,
			},
			{
				Name:        "key_vault_key_uri",
				Description: "The URI of the key vault, used to encrypt the Cosmos DB database account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.KeyVaultKeyURI")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the database account resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.ProvisioningState")},
			{
				Name:        "public_network_access",
				Description: "Indicates whether requests from Public Network are allowed.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "server_version",
				Description: "Describes the ServerVersion of an a MongoDB account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.APIProperties.ServerVersion"),
			},
			{
				Name:        "backup_policy",
				Description: "The object representing the policy for taking backups on an account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.BackupPolicy"),
			},
			{
				Name:        "capabilities",
				Description: "A list of Cosmos DB capabilities for the account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.Capabilities")},
			{
				Name:        "cors",
				Description: "A list of CORS policy for the Cosmos DB database account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.Cors")},
			{
				Name:        "failover_policies",
				Description: "A list of regions ordered by their failover priorities.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.FailoverPolicies")},
			{
				Name:        "ip_rules",
				Description: "A list of IP rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.IPRules")},
			{
				Name:        "locations",
				Description: "A list of all locations that are enabled for the Cosmos DB account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.Locations")},
			{
				Name:        "private_endpoint_connections",
				Description: "A list of Private Endpoint Connections configured for the Cosmos DB account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.PrivateEndpointConnections")},
			{
				Name:        "read_locations",
				Description: "A list of read locations enabled for the Cosmos DB account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.ReadLocations")},
			{
				Name:        "restore_parameters",
				Description: "Parameters to indicate the information about the restore.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.RestoreParameters"),
			},
			{
				Name:        "virtual_network_rules",
				Description: "A list of Virtual Network ACL rules configured for the Cosmos DB account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Properties.VirtualNetworkRules")},
			{
				Name:        "write_locations",
				Description: "A list of write locations enabled for the Cosmos DB account.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.DatabaseAccountGetResults.Properties.WriteLocations")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseAccountGetResults.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.DatabaseAccountGetResults.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DatabaseAccountGetResults.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}
