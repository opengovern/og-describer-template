package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureServiceFabricCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_service_fabric_cluster",
		Description: "Azure Service Fabric Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetServiceFabricCluster,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListServiceFabricCluster,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Azure resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name")},
			{
				Name:        "id",
				Description: "Azure resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ID")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the cluster resource. Possible values include: 'Updating', 'Succeeded', 'Failed', 'Canceled'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "Azure resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Type")},
			{
				Name:        "cluster_code_version",
				Description: "The service fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterCodeVersion")},
			{
				Name:        "cluster_endpoint",
				Description: "The azure resource provider endpoint. A system service in the cluster connects to this  endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterEndpoint")},
			{
				Name:        "cluster_id",
				Description: "A service generated unique identifier for the cluster resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterID")},
			{
				Name:        "cluster_state",
				Description: "The current state of the cluster. Possible values include: 'WaitingForNodes', 'Deploying', 'BaselineUpgrade', 'UpdatingUserConfiguration', 'UpdatingUserCertificate', 'UpdatingInfrastructure', 'EnforcingClusterVersion', 'UpgradeServiceUnreachable', 'AutoScale', 'Ready'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ClusterState")},
			{
				Name:        "event_store_service_enabled",
				Description: "Indicates if the event store service is enabled.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Cluster.Properties.EventStoreServiceEnabled"), Default: false,
			},
			{
				Name:        "etag",
				Description: "Azure resource etag.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Etag")},
			{
				Name:        "management_endpoint",
				Description: "The http management endpoint of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ManagementEndpoint")},
			{
				Name:        "reliability_level",
				Description: "The reliability level sets the replica set size of system services. Possible values include: 'None', 'Bronze', 'Silver', 'Gold', 'Platinum'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.ReliabilityLevel")},
			{
				Name:        "upgrade_mode",
				Description: "The upgrade mode of the cluster when new service fabric runtime version is available. Possible values include: 'Automatic', 'Manual'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.UpgradeMode")},
			{
				Name:        "vm_image",
				Description: "The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Properties.VMImage")},
			{
				Name:        "add_on_features",
				Description: "The list of add-on features to enable in the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.AddOnFeatures")},
			{
				Name:        "available_cluster_versions",
				Description: "The service fabric runtime versions available for this cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.AvailableClusterVersions")},
			{
				Name:        "azure_active_directory",
				Description: "The azure active directory authentication settings of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.AzureActiveDirectory")},
			{
				Name:        "certificate",
				Description: "The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.Certificate")},
			{
				Name:        "certificate_common_names",
				Description: "Describes a list of server certificates referenced by common name that are used to secure the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.CertificateCommonNames")},
			{
				Name:        "client_certificate_common_names",
				Description: "The list of client certificates referenced by common name that are allowed to manage the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ClientCertificateCommonNames")},
			{
				Name:        "client_certificate_thumbprints",
				Description: "The list of client certificates referenced by thumbprint that are allowed to manage the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ClientCertificateThumbprints")},
			{
				Name:        "diagnostics_storage_account_config",
				Description: "The storage account information for storing service fabric diagnostic logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.DiagnosticsStorageAccountConfig")},
			{
				Name:        "fabric_settings",
				Description: "The list of custom fabric settings to configure the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.FabricSettings")},
			{
				Name:        "node_types",
				Description: "The list of node types in the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.NodeTypes")},
			{
				Name:        "reverse_proxy_certificate",
				Description: "The server certificate used by reverse proxy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ReverseProxyCertificate")},
			{
				Name:        "reverse_proxy_certificate_common_names",
				Description: "Describes a list of server certificates referenced by common name that are used to secure the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Properties.ReverseProxyCertificateCommonNames")},
			{
				Name:        "upgrade_description",
				Description: "The policy to use when upgrading the cluster.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Cluster.Properties.UpgradeDescription")},

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

				Transform: transform.

					// The API does not support pagination
					FromField("Description.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTIONS

// Handle empty name or resourceGroup

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data
