package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureHDInsightCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_hdinsight_cluster",
		Description: "Azure HDInsight Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetHdinsightCluster,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListHdinsightCluster,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name")},
			{
				Name:        "id",
				Description: "Fully qualified resource Id for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ID")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state, which only appears in the response. Possible values include: 'InProgress', 'Failed', 'Succeeded', 'Canceled', 'Deleting'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Cluster.Properties.ProvisioningState"),
			},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Type"),
			},
			{
				Name:        "cluster_hdp_version",
				Description: "The hdp version of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterHdpVersion")},
			{
				Name:        "cluster_id",
				Description: "The cluster id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterID")},
			{
				Name:        "cluster_state",
				Description: "The state of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterState")},
			{
				Name:        "cluster_version",
				Description: "The version of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterVersion")},
			{
				Name:        "created_date",
				Description: "The date on which the cluster was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.CreatedDate")},
			{
				Name:        "etag",
				Description: "The ETag for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Etag")},
			{
				Name:        "min_supported_tls_version",
				Description: "The minimal supported tls version of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.MinSupportedTLSVersion")},
			{
				Name:        "os_type",
				Description: "The type of operating system. Possible values include: 'Windows', 'Linux'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Cluster.Properties.OSType"),
			},
			{
				Name:        "tier",
				Description: "The cluster tier. Possible values include: 'Standard', 'Premium'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Cluster.Properties.Tier"),
			},
			{
				Name:        "cluster_definition",
				Description: "The cluster definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterDefinition")},
			{
				Name:        "compute_isolation_properties",
				Description: "The compute isolation properties of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ComputeIsolationProperties")},
			{
				Name:        "compute_profile",
				Description: "The complete profile of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ComputeProfile")},
			{
				Name:        "connectivity_endpoints",
				Description: "The list of connectivity endpoints.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ConnectivityEndpoints")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "disk_encryption_properties",
				Description: "The disk encryption properties of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.DiskEncryptionProperties")},
			{
				Name:        "encryption_in_transit_properties",
				Description: "The encryption-in-transit properties of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.EncryptionInTransitProperties")},
			{
				Name:        "errors",
				Description: "The list of errors.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.Errors")},
			{
				Name:        "excluded_services_config",
				Description: "The excluded services config of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ExcludedServicesConfig")},
			{
				Name:        "identity",
				Description: "The identity of the cluster, if configured.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Identity")},
			{
				Name:        "kafka_rest_properties",
				Description: "The cluster kafka rest proxy configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.KafkaRestProperties")},
			{
				Name:        "network_properties",
				Description: "The network properties of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.NetworkProperties")},
			{
				Name:        "quota_info",
				Description: "The quota information of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.QuotaInfo")},
			{
				Name:        "security_profile",
				Description: "The security profile of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.SecurityProfile")},
			{
				Name:        "storage_profile",
				Description: "The storage profile of the cluster.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Cluster.Properties.StorageProfile")},

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

				Transform: transform.FromField("Description.Cluster.Location").Transform(formatRegion).Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				//// HYDRATE FUNCTIONS
				Transform: transform.

					// Handle empty name or resourceGroup
					FromField("Description.ResourceGroup")},
		}),
	}
}

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

// Create session

// If we return the API response directly, the output does not provide all
// the contents of DiagnosticSettings
