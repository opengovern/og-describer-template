package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureKustoCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_kusto_cluster",
		Description: "Azure Kusto Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetKustoCluster,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKustoCluster,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.Cluster.Name")},
			{
				Name:        "id",
				Description: "The resource Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ID")},
			{
				Name:        "provisioning_state",
				Description: "The provisioned state of the resource. Possible values include: 'Running', 'Creating', 'Deleting', 'Succeeded', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ProvisioningState")},
			{
				Name:        "state",
				Description: "The state of the resource. Possible values include: 'Creating', 'Deleted', 'Deleting', 'Running', 'Starting', 'Stopped', 'Stopping', 'Unavailable'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.State")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Type")},
			{
				Name:        "location",
				Description: "Specifies the name of the region, the resource is created at.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Location")},
			{
				Name:        "data_ingestion_uri",
				Description: "The cluster data ingestion URI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.DataIngestionURI")},
			{
				Name:        "etag",
				Description: "An ETag of the resource created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Etag")},
			{
				Name:        "enable_disk_encryption",
				Description: "A boolean value that indicates if the cluster's disks are encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.Properties.EnableDiskEncryption")},
			{
				Name:        "enable_double_encryption",
				Description: "A boolean value that indicates if double encryption is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.Properties.EnableDoubleEncryption")},
			{
				Name:        "enable_purge",
				Description: "A boolean value that indicates if the purge operations are enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.Properties.EnablePurge")},
			{
				Name:        "enable_streaming_ingest",
				Description: "A boolean value that indicates if the streaming ingest is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.Properties.EnableStreamingIngest")},
			{
				Name:        "engine_type",
				Description: "The engine type. Possible values include: 'EngineTypeV2', 'EngineTypeV3'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.EngineType")},
			{
				Name:        "sku_capacity",
				Description: "SKU capacity of the resource.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.SKU.Capacity")},
			{
				Name:        "sku_name",
				Description: "SKU name of the resource. Possible values include: 'KC8', 'KC16', 'KS8', 'KS16', 'D13V2', 'D14V2', 'L8', 'L16'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.SKU.Name")},
			{
				Name:        "sku_tier",
				Description: "SKU tier of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.SKU.Tier")},
			{
				Name:        "state_reason",
				Description: "SKU tier of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.StateReason")},
			{
				Name:        "uri",
				Description: "The cluster URI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.URI"),
			},
			{
				Name:        "identity",
				Description: "The identity of the cluster, if configured.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Identity"),
			},
			{
				Name:        "language_extensions",
				Description: "List of the cluster's language extensions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.LanguageExtensions")},
			{
				Name:        "key_vault_properties",
				Description: "KeyVault properties for the cluster encryption.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.KeyVaultProperties")},
			{
				Name:        "optimized_autoscale",
				Description: "Optimized auto scale definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.OptimizedAutoscale")},
			{
				Name:        "trusted_external_tenants",
				Description: "The cluster's external tenants.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.TrustedExternalTenants")},
			{
				Name:        "virtual_network_configuration",
				Description: "Virtual network definition of the resource.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Cluster.Properties.VirtualNetworkConfiguration")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Cluster.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Cluster.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				//Pagination does not support for kusto cluster list call till date

				//// HYDRATE FUNCTIONS

				// Return nil, if no input provide
				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
