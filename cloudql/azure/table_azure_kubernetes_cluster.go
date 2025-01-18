package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureKubernetesCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_kubernetes_cluster",
		Description: "Azure Kubernetes Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetKubernetesCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListKubernetesCluster,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the cluster.",
				Transform:   transform.FromField("Description.ManagedCluster.Name")},
			{
				Name:        "id",
				Description: "The ID of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.ID")},
			{
				Name:        "type",
				Description: "The type of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Type")},
			{
				Name:        "location",
				Description: "The location where the cluster is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Location")},
			{
				Name:        "azure_portal_fqdn",
				Description: "FQDN for the master pool which used by proxy config.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.AzurePortalFQDN")},
			{
				Name:        "disk_encryption_set_id",
				Description: "ResourceId of the disk encryption set to use for enabling encryption at rest.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.DiskEncryptionSetID")},
			{
				Name:        "dns_prefix",
				Description: "DNS prefix specified when creating the managed cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.DNSPrefix")},
			{
				Name:        "enable_pod_security_policy",
				Description: "Whether to enable Kubernetes pod security policy (preview).",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.EnablePodSecurityPolicy")},
			{
				Name:        "enable_rbac",
				Description: "Whether to enable Kubernetes Role-Based Access Control.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.EnableRBAC")},
			{
				Name:        "fqdn",
				Description: "FQDN for the master pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.Fqdn")},
			{
				Name:        "fqdn_subdomain",
				Description: "FQDN subdomain specified when creating private cluster with custom private dns zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.FqdnSubdomain")},
			{
				Name:        "kubernetes_version",
				Description: "Version of Kubernetes specified when creating the managed cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.KubernetesVersion")},
			{
				Name:        "max_agent_pools",
				Description: "The max number of agent pools for the managed cluster.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.MaxAgentPools")},
			{
				Name:        "node_resource_group",
				Description: "Name of the resource group containing agent pool nodes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.NodeResourceGroup")},
			{
				Name:        "private_fqdn",
				Description: "FQDN of private cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.PrivateFQDN")},
			{
				Name:        "provisioning_state",
				Description: "The current deployment or provisioning state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.ProvisioningState")},
			{
				Name:        "aad_profile",
				Description: "Profile of Azure Active Directory configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.AADProfile")},
			{
				Name:        "addon_profiles",
				Description: "Profile of managed cluster add-on.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.AddonProfiles")},
			{
				Name:        "agent_pool_profiles",
				Description: "Properties of the agent pool.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.AgentPoolProfiles")},
			{
				Name:        "api_server_access_profile",
				Description: "Access profile for managed cluster API server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.APIServerAccessProfile")},
			{
				Name:        "auto_scaler_profile",
				Description: "Parameters to be applied to the cluster-autoscaler when enabled.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.AutoScalerProfile")},
			{
				Name:        "auto_upgrade_profile",
				Description: "Profile of auto upgrade configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.AutoUpgradeProfile")},
			{
				Name:        "identity",
				Description: "The identity of the managed cluster, if configured.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Identity")},
			{
				Name:        "identity_profile",
				Description: "Identities associated with the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.IdentityProfile")},
			{
				Name:        "linux_profile",
				Description: "Profile for Linux VMs in the container service cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.LinuxProfile")},
			{
				Name:        "network_profile",
				Description: "Profile of network configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.NetworkProfile")},
			{
				Name:        "pod_identity_profile",
				Description: "Profile of managed cluster pod identity.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.PodIdentityProfile")},
			{
				Name:        "power_state",
				Description: "Represents the Power State of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.PowerState")},
			{
				Name:        "service_principal_profile",
				Description: "Information about a service principal identity for the cluster to use for manipulating Azure APIs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Properties.ServicePrincipalProfile")},
			{
				Name:        "sku",
				Description: "The managed cluster SKU.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.SKU")},
			{
				Name:        "windows_profile",
				Description: "Profile for Windows VMs in the container service cluster.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.ManagedCluster.Properties.WindowsProfile")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedCluster.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedCluster.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ManagedCluster.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ManagedCluster.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit

				//// HYDRATE FUNCTIONS
				Transform: transform.

					// In some cases resource does not give any notFound error
					// instead of notFound error, it returns empty data
					FromField("Description.ResourceGroup")},
		}),
	}
}
