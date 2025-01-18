package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureHybridKubernetesConnectedCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_hybrid_kubernetes_connected_cluster",
		Description: "Azure Hybrid Kubernetes Connected Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetHybridKubernetesConnectedCluster,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListHybridKubernetesConnectedCluster,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConnectedCluster.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "connectivity_status",
				Description: "Represents the connectivity status of the connected cluster.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.ConnectivityStatus")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the connected cluster resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConnectedCluster.Type")},
			{
				Name:        "agent_public_key_certificate",
				Description: "Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.AgentPublicKeyCertificate")},
			{
				Name:        "agent_version",
				Description: "Version of the agent running on the connected cluster resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.AgentVersion")},
			{
				Name:        "created_at",
				Description: "The timestamp of resource creation (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ConnectedCluster.SystemData.CreatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "created_by",
				Description: "The identity that created the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.SystemData.CreatedBy")},
			{
				Name:        "created_by_type",
				Description: "The type of identity that created the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.SystemData.CreatedByType")},
			{
				Name:        "distribution",
				Description: "The Kubernetes distribution running on this connected cluster.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.Distribution")},
			{
				Name:        "infrastructure",
				Description: "The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.Infrastructure")},
			{
				Name:        "kubernetes_version",
				Description: "The Kubernetes version of the connected cluster resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.KubernetesVersion")},
			{
				Name:        "last_connectivity_time",
				Description: "Time representing the last instance when heart beat was received from the cluster.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.LastConnectivityTime").Transform(convertDateToTime),
			},
			{
				Name:        "last_modified_at",
				Description: "The timestamp of resource last modification (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ConnectedCluster.SystemData.LastModifiedAt").Transform(convertDateToTime),
			},
			{
				Name:        "last_modified_by",
				Description: "The identity that last modified the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.SystemData.LastModifiedBy")},
			{
				Name:        "last_modified_by_type",
				Description: "The type of identity that last modified the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.SystemData.LastModifiedByType")},
			{
				Name:        "location",
				Description: "Location of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConnectedCluster.Location")},
			{
				Name:        "managed_identity_certificate_expiration_time",
				Description: "Expiration time of the managed identity certificate.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.ManagedIdentityCertificateExpirationTime").Transform(convertDateToTime),
			},
			{
				Name:        "offering",
				Description: "Connected cluster offering.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.Offering")},
			{
				Name:        "total_core_count",
				Description: "Number of CPU cores present in the connected cluster resource.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.TotalCoreCount")},
			{
				Name:        "total_node_count",
				Description: "Number of nodes present in the connected cluster resource.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.ConnectedCluster.Properties.TotalNodeCount")},
			{
				Name:        "extensions",
				Description: "The extensions of the connected cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConnectedClusterExtensions"),
			},
			{
				Name:        "identity",
				Description: "The identity of the connected cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConnectedCluster.Identity")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConnectedCluster.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConnectedCluster.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ConnectedCluster.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConnectedCluster.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.

					//// HYDRATE FUNCTIONS
					FromField("Description.ResourceGroup")},
		}),
	}
}
