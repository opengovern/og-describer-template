package maps
import (
	"github.com/opengovern/og-describer-azure/discovery/describers"
	"github.com/opengovern/og-describer-azure/discovery/provider"
	"github.com/opengovern/og-describer-azure/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-azure/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"Microsoft.App/containerApps": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.App/containerApps",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Container%20App.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AppContainerApps),
		GetDescriber:         nil,
	},

	"Microsoft.Blueprint/blueprints": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Blueprint/blueprints",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Blueprint.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.BlueprintBlueprint),
		GetDescriber:         nil,
	},

	"Microsoft.Cdn/profiles": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Cdn/profiles",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/CDN%20Profile.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.CdnProfiles),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/cloudServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/cloudServices",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Cloud%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeCloudServices),
		GetDescriber:         nil,
	},

	"Microsoft.ContainerInstance/containerGroups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.ContainerInstance/containerGroups",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Container%20Instance.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ContainerInstanceContainerGroups),
		GetDescriber:         nil,
	},

	"Microsoft.DataMigration/services": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataMigration/services",
		Tags:                 map[string][]string{
            "category": {"Migration"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Database%20Migration%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataMigrationServices),
		GetDescriber:         nil,
	},

	"Microsoft.DataProtection/backupVaults": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataProtection/backupVaults",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Backup%20vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataProtectionBackupVaults),
		GetDescriber:         nil,
	},

	"Microsoft.DataProtection/backupJobs": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataProtection/backupJobs",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataProtectionBackupJobs),
		GetDescriber:         nil,
	},

	"Microsoft.DataProtection/backupVaults/backupPolicies": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataProtection/backupVaults/backupPolicies",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Backup%20vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataProtectionBackupVaultsBackupPolicies),
		GetDescriber:         nil,
	},

	"Microsoft.Logic/integrationAccounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Logic/integrationAccounts",
		Tags:                 map[string][]string{
            "category": {"Integration"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Integration%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LogicIntegrationAccounts),
		GetDescriber:         nil,
	},

	"Microsoft.Network/bastionHosts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/bastionHosts",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Bastion.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkBastionHosts),
		GetDescriber:         nil,
	},

	"Microsoft.Network/connections": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/connections",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Hybrid%20Connection.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkConnections),
		GetDescriber:         nil,
	},

	"Microsoft.Network/firewallPolicies": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/firewallPolicies",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Firewall%20Policy.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.FirewallPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Network/localNetworkGateways": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/localNetworkGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Local%20Network%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LocalNetworkGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Network/privateLinkServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/privateLinkServices",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Private%20link%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PrivateLinkService),
		GetDescriber:         nil,
	},

	"Microsoft.Network/publicIPPrefixes": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/publicIPPrefixes",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Public%20IP%20Prefix.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PublicIPPrefix),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualHubs": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualHubs",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Virtual%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkVirtualHubs),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualWans": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualWans",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20WAN.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkVirtualWans),
		GetDescriber:         nil,
	},

	"Microsoft.Network/vpnGateways": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/vpnGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Network%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.VpnGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Network/vpnGateways/vpnConnections": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/vpnGateways/vpnConnections",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkVpnGatewaysVpnConnections),
		GetDescriber:         nil,
	},

	"Microsoft.Network/vpnSites": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/vpnSites",
		Tags:                 map[string][]string{
            "category": {"Networking"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkVpnGatewaysVpnSites),
		GetDescriber:         nil,
	},

	"Microsoft.OperationalInsights/workspaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.OperationalInsights/workspaces",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.OperationalInsightsWorkspaces),
		GetDescriber:         nil,
	},

	"Microsoft.StreamAnalytics/cluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.StreamAnalytics/cluster",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Stream%20Analytics%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StreamAnalyticsCluster),
		GetDescriber:         nil,
	},

	"Microsoft.TimeSeriesInsights/environments": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.TimeSeriesInsights/environments",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Time%20Series%20Insights%20Environment.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.TimeSeriesInsightsEnvironments),
		GetDescriber:         nil,
	},

	"Microsoft.VirtualMachineImages/imageTemplates": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.VirtualMachineImages/imageTemplates",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Image%20Template.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.VirtualMachineImagesImageTemplates),
		GetDescriber:         nil,
	},

	"Microsoft.Web/serverFarms": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Web/serverFarms",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Server%20Farm.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.WebServerFarms),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineScaleSets/virtualMachines": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineScaleSets/virtualMachines",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine%20Scale%20Set.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeVirtualMachineScaleSetVm),
		GetDescriber:         nil,
	},

	"Microsoft.Automation/automationAccounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Automation/automationAccounts",
		Tags:                 map[string][]string{
            "category": {"Management & Governance"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Automation%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AutomationAccounts),
		GetDescriber:         nil,
	},

	"Microsoft.Automation/automationAccounts/variables": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Automation/automationAccounts/variables",
		Tags:                 map[string][]string{
            "category": {"Management & Governance"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Automation%20Variable.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AutomationVariables),
		GetDescriber:         nil,
	},

	"Microsoft.Network/dnsZones": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/dnsZones",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DNS%20Zone%20(Public).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DNSZones),
		GetDescriber:         nil,
	},

	"Microsoft.Databricks/workspaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Databricks/workspaces",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Databricks.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DatabricksWorkspaces),
		GetDescriber:         nil,
	},

	"Microsoft.Network/privateDnsZones": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/privateDnsZones",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DNS%20Zone%20(Private).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PrivateDnsZones),
		GetDescriber:         nil,
	},

	"Microsoft.Network/privateEndpoints": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/privateEndpoints",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Private%20Endpoint.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PrivateEndpoints),
		GetDescriber:         nil,
	},

	"Microsoft.Network/networkWatchers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/networkWatchers",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Network%20Watcher.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkWatcher),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/subscriptions/resourceGroups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Resources/subscriptions/resourceGroups",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Resource%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ResourceGroup),
		GetDescriber:         nil,
	},

	"Microsoft.Web/staticSites": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Web/staticSites",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Static%20Web%20App.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AppServiceWebApp),
		GetDescriber:         nil,
	},

	"Microsoft.Web/sites/slots": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Web/sites/slots",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Static%20Web%20App.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AppServiceWebAppSlot),
		GetDescriber:         nil,
	},

	"Microsoft.CognitiveServices/accounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.CognitiveServices/accounts",
		Tags:                 map[string][]string{
            "category": {"AI + ML"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Cognitive%20Services.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.CognitiveAccount),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/managedInstances": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Sql/managedInstances",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Managed%20Instance.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MssqlManagedInstance),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/virtualclusters": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Sql/virtualclusters",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Database.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlVirtualClusters),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/managedInstances/databases": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Sql/managedInstances/databases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Managed%20Instance.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MssqlManagedInstanceDatabases),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/servers/databases": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Sql/servers/databases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Database.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlDatabase),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/largeFileSharesState": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/largeFileSharesState",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/File%20Share.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageFileShare),
		GetDescriber:         nil,
	},

	"Microsoft.DBforPostgreSQL/servers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DBforPostgreSQL/servers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20PostgreSQL.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PostgresqlServer),
		GetDescriber:         nil,
	},

	"Microsoft.DBforPostgreSQL/flexibleservers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DBforPostgreSQL/flexibleservers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20PostgreSQL.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PostgresqlFlexibleservers),
		GetDescriber:         nil,
	},

	"Microsoft.AnalysisServices/servers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.AnalysisServices/servers",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Analysis%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AnalysisService),
		GetDescriber:         nil,
	},

	"Microsoft.Security/pricings": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Security/pricings",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SecurityCenterSubscriptionPricing),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/guestDiagnosticSettings": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Insights/guestDiagnosticSettings",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Diagnostics%20Setting.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DiagnosticSetting),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/autoscaleSettings": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Insights/autoscaleSettings",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AutoscaleSetting),
		GetDescriber:         nil,
	},

	"Microsoft.Web/hostingEnvironments": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Web/hostingEnvironments",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Service%20Environment.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AppServiceEnvironment),
		GetDescriber:         nil,
	},

	"Microsoft.Cache/redis": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Cache/redis",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cache%20for%20Redis.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.RedisCache),
		GetDescriber:         nil,
	},

	"Microsoft.ContainerRegistry/registries": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.ContainerRegistry/registries",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Container%20Registry.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ContainerRegistry),
		GetDescriber:         nil,
	},

	"Microsoft.DataFactory/factories/pipelines": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataFactory/factories/pipelines",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Factory.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataFactoryPipeline),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/resourceSku": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/resourceSku",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeResourceSKU),
		GetDescriber:         nil,
	},

	"Microsoft.Network/expressRouteCircuits": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/expressRouteCircuits",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/ExpressRoute%20Circuit.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ExpressRouteCircuit),
		GetDescriber:         nil,
	},

	"Microsoft.Management/managementgroups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Management/managementgroups",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Management%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ManagementGroup),
		GetDescriber:         nil,
	},

	"microsoft.SqlVirtualMachine/SqlVirtualMachines": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "microsoft.SqlVirtualMachine/SqlVirtualMachines",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlServerVirtualMachine),
		GetDescriber:         nil,
	},

	"Microsoft.SqlVirtualMachine/SqlVirtualMachineGroups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.SqlVirtualMachine/SqlVirtualMachineGroups",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlServerVirtualMachineGroups),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/tableServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/tableServices",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Table.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageTableService),
		GetDescriber:         nil,
	},

	"Microsoft.Synapse/workspaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Synapse/workspaces",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Synapse%20Analytics.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SynapseWorkspace),
		GetDescriber:         nil,
	},

	"Microsoft.Synapse/workspaces/bigdatapools": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Synapse/workspaces/bigdatapools",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Synapse%20Analytics.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SynapseWorkspaceBigdataPools),
		GetDescriber:         nil,
	},

	"Microsoft.Synapse/workspaces/sqlpools": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Synapse/workspaces/sqlpools",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Synapse%20Analytics.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SynapseWorkspaceSqlpools),
		GetDescriber:         nil,
	},

	"Microsoft.StreamAnalytics/streamingJobs": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.StreamAnalytics/streamingJobs",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Stream%20Analytics%20job.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StreamAnalyticsJob),
		GetDescriber:         nil,
	},

	"Microsoft.CostManagement/CostBySubscription": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.CostManagement/CostBySubscription",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DailyCostBySubscription),
		GetDescriber:         nil,
	},

	"Microsoft.ContainerService/managedClusters": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.ContainerService/managedClusters",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/AKS%20Hybrid%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KubernetesCluster),
		GetDescriber:         nil,
	},

	"Microsoft.ContainerService/serviceVersions": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.ContainerService/serviceVersions",
		Tags:                 map[string][]string{
            "category": {"Container"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KubernetesServiceVersion),
		GetDescriber:         nil,
	},

	"Microsoft.DataFactory/factories": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataFactory/factories",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Factory.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataFactory),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/servers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Sql/servers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlServer),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/servers/jobagents": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Sql/servers/jobagents",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Elastic%20Job%20Agent.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlServerJobAgents),
		GetDescriber:         nil,
	},

	"Microsoft.Security/autoProvisioningSettings": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Security/autoProvisioningSettings",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SecurityCenterAutoProvisioning),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/logProfiles": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Insights/logProfiles",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LogProfile),
		GetDescriber:         nil,
	},

	"Microsoft.DataBoxEdge/dataBoxEdgeDevices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataBoxEdge/dataBoxEdgeDevices",
		Tags:                 map[string][]string{
            "category": {"IoT & Devices"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Box%20Edge.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataboxEdgeDevice),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LoadBalancer),
		GetDescriber:         nil,
	},

	"Microsoft.Network/azureFirewalls": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/azureFirewalls",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Firewall.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkAzureFirewall),
		GetDescriber:         nil,
	},

	"Microsoft.Management/locks": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Management/locks",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Resource%20Lock.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ManagementLock),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineScaleSets/networkInterfaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineScaleSets/networkInterfaces",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Network%20Interface.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeVirtualMachineScaleSetNetworkInterface),
		GetDescriber:         nil,
	},

	"Microsoft.Network/frontDoors": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/frontDoors",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Front%20Door.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.FrontDoor),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/policyAssignments": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Authorization/policyAssignments",
		Tags:                 map[string][]string{
            "category": {"Identify & Access"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Policy%20Assignment.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PolicyAssignment),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/userEffectiveAccess": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Authorization/userEffectiveAccess",
		Tags:                 map[string][]string{
            "category": {"Identify & Access"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Policy%20Assignment.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.UserEffectiveAccess),
		GetDescriber:         nil,
	},

	"Microsoft.Search/searchServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Search/searchServices",
		Tags:                 map[string][]string{
            "category": {"AI + ML"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Search%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SearchService),
		GetDescriber:         nil,
	},

	"Microsoft.Security/settings": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Security/settings",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SecurityCenterSetting),
		GetDescriber:         nil,
	},

	"Microsoft.RecoveryServices/vaults": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.RecoveryServices/vaults",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Recovery%20Services%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.RecoveryServicesVault),
		GetDescriber:         nil,
	},

	"Microsoft.RecoveryServices/vaults/backupJobs": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.RecoveryServices/vaults/backupJobs",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.RecoveryServicesBackupJobs),
		GetDescriber:         nil,
	},

	"Microsoft.RecoveryServices/vaults/backupPolicies": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.RecoveryServices/vaults/backupPolicies",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.RecoveryServicesBackupPolicies),
		GetDescriber:         nil,
	},

	"Microsoft.RecoveryServices/vaults/backupItems": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.RecoveryServices/vaults/backupItems",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.RecoveryServicesBackupItem),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/diskEncryptionSets": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/diskEncryptionSets",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Disk%20Encryption%20Set.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeDiskEncryptionSet),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/databaseAccounts/sqlDatabases": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/databaseAccounts/sqlDatabases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DocumentDBSQLDatabase),
		GetDescriber:         nil,
	},

	"Microsoft.EventGrid/topics": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.EventGrid/topics",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Event%20Grid%20Topic.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.EventGridTopic),
		GetDescriber:         nil,
	},

	"Microsoft.EventHub/namespaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.EventHub/namespaces",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Event%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.EventhubNamespace),
		GetDescriber:         nil,
	},

	"Microsoft.EventHub/namespaces/eventHubs": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.EventHub/namespaces/eventHubs",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Event%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.EventhubNamespaceEventhub),
		GetDescriber:         nil,
	},

	"Microsoft.MachineLearningServices/workspaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.MachineLearningServices/workspaces",
		Tags:                 map[string][]string{
            "category": {"AI + ML"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Machine%20Learning%20Studio%20Workspace%20(classic).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MachineLearningWorkspace),
		GetDescriber:         nil,
	},

	"Microsoft.Dashboard/grafana": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Dashboard/grafana",
		Tags:                 map[string][]string{
            "category": {"Managed Services"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Managed%20Grafana.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DashboardGrafana),
		GetDescriber:         nil,
	},

	"Microsoft.DesktopVirtualization/workspaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DesktopVirtualization/workspaces",
		Tags:                 map[string][]string{
            "category": {"End User"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Windows%20Virtual%20Desktop.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DesktopVirtualizationWorkspaces),
		GetDescriber:         nil,
	},

	"Microsoft.Network/trafficManagerProfiles": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/trafficManagerProfiles",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Traffic%20Manager%20profile.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.TrafficManagerProfile),
		GetDescriber:         nil,
	},

	"Microsoft.Network/dnsResolvers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/dnsResolvers",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DNS%20Private%20Resolver.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DNSResolvers),
		GetDescriber:         nil,
	},

	"Microsoft.CostManagement/CostByResourceType": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.CostManagement/CostByResourceType",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DailyCostByResourceType),
		GetDescriber:         nil,
	},

	"Microsoft.Network/networkInterfaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/networkInterfaces",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Network%20Interface.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkInterface),
		GetDescriber:         nil,
	},

	"Microsoft.Network/publicIPAddresses": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/publicIPAddresses",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Public%20IP%20Address.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PublicIPAddress),
		GetDescriber:         nil,
	},

	"Microsoft.HealthcareApis/services": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.HealthcareApis/services",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.HealthcareService),
		GetDescriber:         nil,
	},

	"Microsoft.ServiceBus/namespaces": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.ServiceBus/namespaces",
		Tags:                 map[string][]string{
            "category": {"Intergration"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Service%20Bus.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ServicebusNamespace),
		GetDescriber:         nil,
	},

	"Microsoft.Web/sites": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Web/sites",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AppServiceFunctionApp),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/availabilitySets": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/availabilitySets",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine%20Availability%20Set.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeAvailabilitySet),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualNetworks": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualNetworks",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Network.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.VirtualNetwork),
		GetDescriber:         nil,
	},

	"Microsoft.Security/securityContacts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Security/securityContacts",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SecurityCenterContact),
		GetDescriber:         nil,
	},

	"Microsoft.EventGrid/domains": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.EventGrid/domains",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Event%20Grid%20Domain.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.EventGridDomain),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/deletedVaults": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/deletedVaults",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DeletedVault),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/tableServices/tables": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/tableServices/tables",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Table.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageTable),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/snapshots": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/snapshots",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Managed%20Disk%20Snapshot.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeSnapshots),
		GetDescriber:         nil,
	},

	"Microsoft.Kusto/clusters": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Kusto/clusters",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_Grouped/Data/Azure%20Data%20Explorer%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KustoCluster),
		GetDescriber:         nil,
	},

	"Microsoft.StorageSync/storageSyncServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.StorageSync/storageSyncServices",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Sync%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageSync),
		GetDescriber:         nil,
	},

	"Microsoft.Security/locations/jitNetworkAccessPolicies": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Security/locations/jitNetworkAccessPolicies",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SecurityCenterJitNetworkAccessPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualNetworks/subnets": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualNetworks/subnets",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Subnet.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.Subnet),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/backendAddressPools": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/backendAddressPools",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer%20Backend%20pool.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LoadBalancerBackendAddressPool),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/loadBalancingRules": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/loadBalancingRules",
		Tags:                 map[string][]string{
            "category": {"Networking"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LoadBalancerRule),
		GetDescriber:         nil,
	},

	"Microsoft.DataLakeStore/accounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataLakeStore/accounts",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Lake.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataLakeStore),
		GetDescriber:         nil,
	},

	"Microsoft.StorageCache/caches": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.StorageCache/caches",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_Grouped/Data/HPC%20Cache.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.HpcCache),
		GetDescriber:         nil,
	},

	"Microsoft.Batch/batchAccounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Batch/batchAccounts",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Batch%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.BatchAccount),
		GetDescriber:         nil,
	},

	"Microsoft.Network/networkSecurityGroups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/networkSecurityGroups",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Network%20Security%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkSecurityGroup),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/roleDefinitions": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Authorization/roleDefinitions",
		Tags:                 map[string][]string{
            "category": {"Identify & Access"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Role%20(Custom).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.RoleDefinition),
		GetDescriber:         nil,
	},

	"Microsoft.Network/applicationSecurityGroups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/applicationSecurityGroups",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Security%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkApplicationSecurityGroups),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/roleAssignment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Authorization/roleAssignment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.RoleAssignment),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/databaseAccounts/mongodbDatabases": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DocumentDBMongoDatabase),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DocumentDBMongoCollection),
		GetDescriber:         nil,
	},

	"Microsoft.Network/networkWatchers/flowLogs": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/networkWatchers/flowLogs",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkWatcherFlowLog),
		GetDescriber:         nil,
	},

	"microsoft.Sql/servers/elasticpools": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "microsoft.Sql/servers/elasticpools",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Elastic%20Pool.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlServerElasticPool),
		GetDescriber:         nil,
	},

	"Microsoft.Security/subAssessments": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Security/subAssessments",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SecurityCenterSubAssessment),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/disks": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/disks",
		Tags:                 map[string][]string{
            "category": {"Storage"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeDisk),
		GetDescriber:         nil,
	},

	"Microsoft.Devices/ProvisioningServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Devices/ProvisioningServices",
		Tags:                 map[string][]string{
            "category": {"IoT & Devices"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/IoT%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.IOTHubDps),
		GetDescriber:         nil,
	},

	"Microsoft.HDInsight/clusters": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.HDInsight/clusters",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/HDInsight%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.HdInsightCluster),
		GetDescriber:         nil,
	},

	"Microsoft.ServiceFabric/clusters": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.ServiceFabric/clusters",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Service%20Fabric%20Managed%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ServiceFabricCluster),
		GetDescriber:         nil,
	},

	"Microsoft.SignalRService/signalR": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.SignalRService/signalR",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SignalR.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SignalrService),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/blob": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/blob",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Blob.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageBlob),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageaccounts/blobservices/containers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageaccounts/blobservices/containers",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Container.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageContainer),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/blobServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/blobServices",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Blob.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageBlobService),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/queueServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/queueServices",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Queue.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageQueue),
		GetDescriber:         nil,
	},

	"Microsoft.ApiManagement/service": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.ApiManagement/service",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/API%20Management%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.APIManagement),
		GetDescriber:         nil,
	},

	"Microsoft.ApiManagement/backend": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.ApiManagement/backend",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/API%20Management%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.APIManagementBackend),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineScaleSets": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineScaleSets",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine%20Scale%20Set.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeVirtualMachineScaleSet),
		GetDescriber:         nil,
	},

	"Microsoft.DataFactory/factories/datasets": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataFactory/factories/datasets",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Factory.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataFactoryDataset),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/policyDefinitions": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Authorization/policyDefinitions",
		Tags:                 map[string][]string{
            "category": {"Identify & Access"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Policy%20Definition.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PolicyDefinition),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/subscriptions/locations": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Resources/subscriptions/locations",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.Location),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/diskAccesses": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/diskAccesses",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Disk%20Access.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeDiskAccess),
		GetDescriber:         nil,
	},

	"Microsoft.DBforMySQL/servers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DBforMySQL/servers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20MySQL.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MysqlServer),
		GetDescriber:         nil,
	},

	"Microsoft.DBforMySQL/flexibleservers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DBforMySQL/flexibleservers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Wordpress%20and%20MySQL%20Flexible%20server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MysqlFlexibleservers),
		GetDescriber:         nil,
	},

	"Microsoft.Cache/redisenterprise": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Cache/redisenterprise",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cache%20for%20Redis.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.CacheRedisEnterprise),
		GetDescriber:         nil,
	},

	"Microsoft.DataLakeAnalytics/accounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DataLakeAnalytics/accounts",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Lake%20Analytics.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DataLakeAnalyticsAccount),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/activityLogAlerts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Insights/activityLogAlerts",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LogAlert),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/outboundRules": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/outboundRules",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer%20Backend%20Outbound%20Rule.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LoadBalancerOutboundRule),
		GetDescriber:         nil,
	},

	"Microsoft.HybridCompute/machines": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.HybridCompute/machines",
		Tags:                 map[string][]string{
            "category": {"Compute"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.HybridComputeMachine),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/inboundNatRules": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/inboundNatRules",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer%20Inbound%20NAT%20Rule.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LoadBalancerNatRule),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/providers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Resources/providers",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ResourceProvider),
		GetDescriber:         nil,
	},

	"Microsoft.Network/routeTables": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/routeTables",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Route%20Table.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.RouteTables),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/databaseAccounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/databaseAccounts",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.CosmosdbAccount),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/restorableDatabaseAccounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/restorableDatabaseAccounts",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.CosmosdbRestorableDatabaseAccount),
		GetDescriber:         nil,
	},

	"Microsoft.Network/applicationGateways": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/applicationGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ApplicationGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Security/automations": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Security/automations",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SecurityCenterAutomation),
		GetDescriber:         nil,
	},

	"Microsoft.Kubernetes/connectedClusters": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Kubernetes/connectedClusters",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Kubernetes%20Cluster%20(Operator%20Nexus).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.HybridKubernetesConnectedCluster),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults/keys": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults/keys",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KeyVaultKey),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults/certificates": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults/certificates",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KeyVaultCertificate),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults/keys/Versions": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults/keys/Versions",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KeyVaultKey),
		GetDescriber:         nil,
	},

	"Microsoft.DBforMariaDB/servers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DBforMariaDB/servers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20MariaDB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MariadbServer),
		GetDescriber:         nil,
	},

	"Microsoft.DBforMariaDB/servers/databases": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DBforMariaDB/servers/databases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20MariaDB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MariadbDatabases),
		GetDescriber:         nil,
	},

	"Microsoft.Web/plan": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Web/plan",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Service%20plan.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AppServicePlan),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/tenants": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Resources/tenants",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.Tenant),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualNetworkGateways": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualNetworkGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Network%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.VirtualNetworkGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Devices/iotHubs": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Devices/iotHubs",
		Tags:                 map[string][]string{
            "category": {"IoT & Devices"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/IoT%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.IOTHub),
		GetDescriber:         nil,
	},

	"Microsoft.Logic/workflows": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Logic/workflows",
		Tags:                 map[string][]string{
            "category": {"Integration"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LogicAppWorkflow),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/flexibleServers": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Sql/flexibleServers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlServerFlexibleServer),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/links": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Resources/links",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Resource%20Management%20Private%20Link.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ResourceLink),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/subscriptions": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Resources/subscriptions",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Subscription.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.Subscription),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/images": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/images",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine%20Image.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeImage),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachines": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachines",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeVirtualMachine),
		GetDescriber:         nil,
	},

	"Microsoft.Network/natGateways": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/natGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/NAT%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NatGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/probes": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/probes",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer%20Health%20Probe.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LoadBalancerProbe),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KeyVault),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/managedHsms": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/managedHsms",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault%20HSM.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KeyVaultManagedHardwareSecurityModule),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults/secrets": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults/secrets",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault%20Secret.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.KeyVaultSecret),
		GetDescriber:         nil,
	},

	"Microsoft.AppConfiguration/configurationStores": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.AppConfiguration/configurationStores",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/App%20Configuration.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.AppConfiguration),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.StorageAccount),
		GetDescriber:         nil,
	},

	"Microsoft.AppPlatform/Spring": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.AppPlatform/Spring",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Spring%20Cloud.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SpringCloudService),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/galleries": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/galleries",
		Tags:                 map[string][]string{
            "category": {"General"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Compute%20Gallery.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeGallery),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/hostGroups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/hostGroups",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Host%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeHostGroup),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/hostGroups/hosts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/hostGroups/hosts",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Host%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeHost),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/restorePointCollections": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/restorePointCollections",
		Tags:                 map[string][]string{
            "category": {"Backup"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeRestorePointCollection),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/sshPublicKeys": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Compute/sshPublicKeys",
		Tags:                 map[string][]string{
            "category": {"Management & Governance"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SSH%20key.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ComputeSSHPublicKey),
		GetDescriber:         nil,
	},

	"Microsoft.Cdn/profiles/endpoints": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Cdn/profiles/endpoints",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/CDN%20Profile.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.CdnEndpoint),
		GetDescriber:         nil,
	},

	"Microsoft.BotService/botServices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.BotService/botServices",
		Tags:                 map[string][]string{
            "category": {"AI + ML"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Bot%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.BotServiceBot),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/cassandraClusters": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/cassandraClusters",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Managed%20Instance%20for%20Apache%20Cassandra.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DocumentDBCassandraCluster),
		GetDescriber:         nil,
	},

	"Microsoft.Network/ddosProtectionPlans": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Network/ddosProtectionPlans",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DDoS%20Protection%20Plan.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetworkDDoSProtectionPlan),
		GetDescriber:         nil,
	},

	"microsoft.Sql/instancePools": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "microsoft.Sql/instancePools",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Instance%20Pool.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.SqlInstancePool),
		GetDescriber:         nil,
	},

	"microsoft.NetApp/netAppAccounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "microsoft.NetApp/netAppAccounts",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20NetApp%20Files.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetAppAccount),
		GetDescriber:         nil,
	},

	"Microsoft.NetApp/netAppAccounts/capacityPools": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.NetApp/netAppAccounts/capacityPools",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20NetApp%20Files.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.NetAppCapacityPool),
		GetDescriber:         nil,
	},

	"Microsoft.DesktopVirtualization/hostpools": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.DesktopVirtualization/hostpools",
		Tags:                 map[string][]string{
            "category": {"End User"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Windows%20Virtual%20Desktop.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DesktopVirtualizationHostPool),
		GetDescriber:         nil,
	},

	"Microsoft.Devtestlab/labs": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Devtestlab/labs",
		Tags:                 map[string][]string{
            "category": {"DevOps + Testing"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DevTest%20Lab.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.DevTestLabLab),
		GetDescriber:         nil,
	},

	"Microsoft.Purview/Accounts": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Purview/Accounts",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Purview%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PurviewAccount),
		GetDescriber:         nil,
	},

	"Microsoft.PowerBIDedicated/capacities": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.PowerBIDedicated/capacities",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Power%20BI%20Embedded.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.PowerBIDedicatedCapacity),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/components": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Insights/components",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Insights.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.ApplicationInsights),
		GetDescriber:         nil,
	},

	"Microsoft.Lighthouse/definition": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Lighthouse/definition",
		Tags:                 map[string][]string{
            "category": {},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LighthouseDefinition),
		GetDescriber:         nil,
	},

	"Microsoft.Lighthouse/assignment": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Lighthouse/assignment",
		Tags:                 map[string][]string{
            "category": {},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.LighthouseAssignments),
		GetDescriber:         nil,
	},

	"Microsoft.Maintenance/maintenanceConfigurations": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Maintenance/maintenanceConfigurations",
		Tags:                 map[string][]string{
            "category": {},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MaintenanceConfiguration),
		GetDescriber:         nil,
	},

	"Microsoft.Monitor/logProfiles": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Monitor/logProfiles",
		Tags:                 map[string][]string{
            "category": {},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.MonitorLogProfiles),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/subscriptions/resources": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Resources/subscriptions/resources",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeBySubscription(describers.Resources),
		GetDescriber:         nil,
	},
}


var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Microsoft.App/containerApps": {
		Name:         "Microsoft.App/containerApps",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Blueprint/blueprints": {
		Name:         "Microsoft.Blueprint/blueprints",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Cdn/profiles": {
		Name:         "Microsoft.Cdn/profiles",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/cloudServices": {
		Name:         "Microsoft.Compute/cloudServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.ContainerInstance/containerGroups": {
		Name:         "Microsoft.ContainerInstance/containerGroups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataMigration/services": {
		Name:         "Microsoft.DataMigration/services",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataProtection/backupVaults": {
		Name:         "Microsoft.DataProtection/backupVaults",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataProtection/backupJobs": {
		Name:         "Microsoft.DataProtection/backupJobs",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataProtection/backupVaults/backupPolicies": {
		Name:         "Microsoft.DataProtection/backupVaults/backupPolicies",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Logic/integrationAccounts": {
		Name:         "Microsoft.Logic/integrationAccounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/bastionHosts": {
		Name:         "Microsoft.Network/bastionHosts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/connections": {
		Name:         "Microsoft.Network/connections",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/firewallPolicies": {
		Name:         "Microsoft.Network/firewallPolicies",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/localNetworkGateways": {
		Name:         "Microsoft.Network/localNetworkGateways",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/privateLinkServices": {
		Name:         "Microsoft.Network/privateLinkServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/publicIPPrefixes": {
		Name:         "Microsoft.Network/publicIPPrefixes",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/virtualHubs": {
		Name:         "Microsoft.Network/virtualHubs",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/virtualWans": {
		Name:         "Microsoft.Network/virtualWans",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/vpnGateways": {
		Name:         "Microsoft.Network/vpnGateways",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/vpnGateways/vpnConnections": {
		Name:         "Microsoft.Network/vpnGateways/vpnConnections",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/vpnSites": {
		Name:         "Microsoft.Network/vpnSites",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.OperationalInsights/workspaces": {
		Name:         "Microsoft.OperationalInsights/workspaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.StreamAnalytics/cluster": {
		Name:         "Microsoft.StreamAnalytics/cluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.TimeSeriesInsights/environments": {
		Name:         "Microsoft.TimeSeriesInsights/environments",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.VirtualMachineImages/imageTemplates": {
		Name:         "Microsoft.VirtualMachineImages/imageTemplates",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Web/serverFarms": {
		Name:         "Microsoft.Web/serverFarms",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/virtualMachineScaleSets/virtualMachines": {
		Name:         "Microsoft.Compute/virtualMachineScaleSets/virtualMachines",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Automation/automationAccounts": {
		Name:         "Microsoft.Automation/automationAccounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Automation/automationAccounts/variables": {
		Name:         "Microsoft.Automation/automationAccounts/variables",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/dnsZones": {
		Name:         "Microsoft.Network/dnsZones",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Databricks/workspaces": {
		Name:         "Microsoft.Databricks/workspaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/privateDnsZones": {
		Name:         "Microsoft.Network/privateDnsZones",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/privateEndpoints": {
		Name:         "Microsoft.Network/privateEndpoints",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/networkWatchers": {
		Name:         "Microsoft.Network/networkWatchers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Resources/subscriptions/resourceGroups": {
		Name:         "Microsoft.Resources/subscriptions/resourceGroups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Web/staticSites": {
		Name:         "Microsoft.Web/staticSites",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Web/sites/slots": {
		Name:         "Microsoft.Web/sites/slots",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.CognitiveServices/accounts": {
		Name:         "Microsoft.CognitiveServices/accounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Sql/managedInstances": {
		Name:         "Microsoft.Sql/managedInstances",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Sql/virtualclusters": {
		Name:         "Microsoft.Sql/virtualclusters",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Sql/managedInstances/databases": {
		Name:         "Microsoft.Sql/managedInstances/databases",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Sql/servers/databases": {
		Name:         "Microsoft.Sql/servers/databases",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Storage/storageAccounts/largeFileSharesState": {
		Name:         "Microsoft.Storage/storageAccounts/largeFileSharesState",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DBforPostgreSQL/servers": {
		Name:         "Microsoft.DBforPostgreSQL/servers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DBforPostgreSQL/flexibleservers": {
		Name:         "Microsoft.DBforPostgreSQL/flexibleservers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.AnalysisServices/servers": {
		Name:         "Microsoft.AnalysisServices/servers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Security/pricings": {
		Name:         "Microsoft.Security/pricings",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Insights/guestDiagnosticSettings": {
		Name:         "Microsoft.Insights/guestDiagnosticSettings",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Insights/autoscaleSettings": {
		Name:         "Microsoft.Insights/autoscaleSettings",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Web/hostingEnvironments": {
		Name:         "Microsoft.Web/hostingEnvironments",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Cache/redis": {
		Name:         "Microsoft.Cache/redis",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.ContainerRegistry/registries": {
		Name:         "Microsoft.ContainerRegistry/registries",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataFactory/factories/pipelines": {
		Name:         "Microsoft.DataFactory/factories/pipelines",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/resourceSku": {
		Name:         "Microsoft.Compute/resourceSku",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/expressRouteCircuits": {
		Name:         "Microsoft.Network/expressRouteCircuits",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Management/managementgroups": {
		Name:         "Microsoft.Management/managementgroups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"microsoft.SqlVirtualMachine/SqlVirtualMachines": {
		Name:         "microsoft.SqlVirtualMachine/SqlVirtualMachines",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.SqlVirtualMachine/SqlVirtualMachineGroups": {
		Name:         "Microsoft.SqlVirtualMachine/SqlVirtualMachineGroups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Storage/storageAccounts/tableServices": {
		Name:         "Microsoft.Storage/storageAccounts/tableServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Synapse/workspaces": {
		Name:         "Microsoft.Synapse/workspaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Synapse/workspaces/bigdatapools": {
		Name:         "Microsoft.Synapse/workspaces/bigdatapools",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Synapse/workspaces/sqlpools": {
		Name:         "Microsoft.Synapse/workspaces/sqlpools",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.StreamAnalytics/streamingJobs": {
		Name:         "Microsoft.StreamAnalytics/streamingJobs",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.CostManagement/CostBySubscription": {
		Name:         "Microsoft.CostManagement/CostBySubscription",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.ContainerService/managedClusters": {
		Name:         "Microsoft.ContainerService/managedClusters",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.ContainerService/serviceVersions": {
		Name:         "Microsoft.ContainerService/serviceVersions",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataFactory/factories": {
		Name:         "Microsoft.DataFactory/factories",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Sql/servers": {
		Name:         "Microsoft.Sql/servers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Sql/servers/jobagents": {
		Name:         "Microsoft.Sql/servers/jobagents",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Security/autoProvisioningSettings": {
		Name:         "Microsoft.Security/autoProvisioningSettings",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Insights/logProfiles": {
		Name:         "Microsoft.Insights/logProfiles",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataBoxEdge/dataBoxEdgeDevices": {
		Name:         "Microsoft.DataBoxEdge/dataBoxEdgeDevices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/loadBalancers": {
		Name:         "Microsoft.Network/loadBalancers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/azureFirewalls": {
		Name:         "Microsoft.Network/azureFirewalls",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Management/locks": {
		Name:         "Microsoft.Management/locks",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/virtualMachineScaleSets/networkInterfaces": {
		Name:         "Microsoft.Compute/virtualMachineScaleSets/networkInterfaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/frontDoors": {
		Name:         "Microsoft.Network/frontDoors",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Authorization/policyAssignments": {
		Name:         "Microsoft.Authorization/policyAssignments",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Authorization/userEffectiveAccess": {
		Name:         "Microsoft.Authorization/userEffectiveAccess",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Search/searchServices": {
		Name:         "Microsoft.Search/searchServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Security/settings": {
		Name:         "Microsoft.Security/settings",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.RecoveryServices/vaults": {
		Name:         "Microsoft.RecoveryServices/vaults",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.RecoveryServices/vaults/backupJobs": {
		Name:         "Microsoft.RecoveryServices/vaults/backupJobs",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.RecoveryServices/vaults/backupPolicies": {
		Name:         "Microsoft.RecoveryServices/vaults/backupPolicies",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.RecoveryServices/vaults/backupItems": {
		Name:         "Microsoft.RecoveryServices/vaults/backupItems",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/diskEncryptionSets": {
		Name:         "Microsoft.Compute/diskEncryptionSets",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DocumentDB/databaseAccounts/sqlDatabases": {
		Name:         "Microsoft.DocumentDB/databaseAccounts/sqlDatabases",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.EventGrid/topics": {
		Name:         "Microsoft.EventGrid/topics",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.EventHub/namespaces": {
		Name:         "Microsoft.EventHub/namespaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.EventHub/namespaces/eventHubs": {
		Name:         "Microsoft.EventHub/namespaces/eventHubs",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.MachineLearningServices/workspaces": {
		Name:         "Microsoft.MachineLearningServices/workspaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Dashboard/grafana": {
		Name:         "Microsoft.Dashboard/grafana",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DesktopVirtualization/workspaces": {
		Name:         "Microsoft.DesktopVirtualization/workspaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/trafficManagerProfiles": {
		Name:         "Microsoft.Network/trafficManagerProfiles",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/dnsResolvers": {
		Name:         "Microsoft.Network/dnsResolvers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.CostManagement/CostByResourceType": {
		Name:         "Microsoft.CostManagement/CostByResourceType",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/networkInterfaces": {
		Name:         "Microsoft.Network/networkInterfaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/publicIPAddresses": {
		Name:         "Microsoft.Network/publicIPAddresses",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.HealthcareApis/services": {
		Name:         "Microsoft.HealthcareApis/services",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.ServiceBus/namespaces": {
		Name:         "Microsoft.ServiceBus/namespaces",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Web/sites": {
		Name:         "Microsoft.Web/sites",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/availabilitySets": {
		Name:         "Microsoft.Compute/availabilitySets",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/virtualNetworks": {
		Name:         "Microsoft.Network/virtualNetworks",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Security/securityContacts": {
		Name:         "Microsoft.Security/securityContacts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.EventGrid/domains": {
		Name:         "Microsoft.EventGrid/domains",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.KeyVault/deletedVaults": {
		Name:         "Microsoft.KeyVault/deletedVaults",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Storage/storageAccounts/tableServices/tables": {
		Name:         "Microsoft.Storage/storageAccounts/tableServices/tables",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/snapshots": {
		Name:         "Microsoft.Compute/snapshots",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Kusto/clusters": {
		Name:         "Microsoft.Kusto/clusters",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.StorageSync/storageSyncServices": {
		Name:         "Microsoft.StorageSync/storageSyncServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Security/locations/jitNetworkAccessPolicies": {
		Name:         "Microsoft.Security/locations/jitNetworkAccessPolicies",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/virtualNetworks/subnets": {
		Name:         "Microsoft.Network/virtualNetworks/subnets",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/loadBalancers/backendAddressPools": {
		Name:         "Microsoft.Network/loadBalancers/backendAddressPools",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/loadBalancers/loadBalancingRules": {
		Name:         "Microsoft.Network/loadBalancers/loadBalancingRules",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataLakeStore/accounts": {
		Name:         "Microsoft.DataLakeStore/accounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.StorageCache/caches": {
		Name:         "Microsoft.StorageCache/caches",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Batch/batchAccounts": {
		Name:         "Microsoft.Batch/batchAccounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/networkSecurityGroups": {
		Name:         "Microsoft.Network/networkSecurityGroups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Authorization/roleDefinitions": {
		Name:         "Microsoft.Authorization/roleDefinitions",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/applicationSecurityGroups": {
		Name:         "Microsoft.Network/applicationSecurityGroups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Authorization/roleAssignment": {
		Name:         "Microsoft.Authorization/roleAssignment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DocumentDB/databaseAccounts/mongodbDatabases": {
		Name:         "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections": {
		Name:         "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/networkWatchers/flowLogs": {
		Name:         "Microsoft.Network/networkWatchers/flowLogs",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"microsoft.Sql/servers/elasticpools": {
		Name:         "microsoft.Sql/servers/elasticpools",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Security/subAssessments": {
		Name:         "Microsoft.Security/subAssessments",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/disks": {
		Name:         "Microsoft.Compute/disks",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Devices/ProvisioningServices": {
		Name:         "Microsoft.Devices/ProvisioningServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.HDInsight/clusters": {
		Name:         "Microsoft.HDInsight/clusters",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.ServiceFabric/clusters": {
		Name:         "Microsoft.ServiceFabric/clusters",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.SignalRService/signalR": {
		Name:         "Microsoft.SignalRService/signalR",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Storage/storageAccounts/blob": {
		Name:         "Microsoft.Storage/storageAccounts/blob",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Storage/storageaccounts/blobservices/containers": {
		Name:         "Microsoft.Storage/storageaccounts/blobservices/containers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Storage/storageAccounts/blobServices": {
		Name:         "Microsoft.Storage/storageAccounts/blobServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Storage/storageAccounts/queueServices": {
		Name:         "Microsoft.Storage/storageAccounts/queueServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.ApiManagement/service": {
		Name:         "Microsoft.ApiManagement/service",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.ApiManagement/backend": {
		Name:         "Microsoft.ApiManagement/backend",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/virtualMachineScaleSets": {
		Name:         "Microsoft.Compute/virtualMachineScaleSets",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataFactory/factories/datasets": {
		Name:         "Microsoft.DataFactory/factories/datasets",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Authorization/policyDefinitions": {
		Name:         "Microsoft.Authorization/policyDefinitions",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Resources/subscriptions/locations": {
		Name:         "Microsoft.Resources/subscriptions/locations",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/diskAccesses": {
		Name:         "Microsoft.Compute/diskAccesses",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DBforMySQL/servers": {
		Name:         "Microsoft.DBforMySQL/servers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DBforMySQL/flexibleservers": {
		Name:         "Microsoft.DBforMySQL/flexibleservers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Cache/redisenterprise": {
		Name:         "Microsoft.Cache/redisenterprise",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DataLakeAnalytics/accounts": {
		Name:         "Microsoft.DataLakeAnalytics/accounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Insights/activityLogAlerts": {
		Name:         "Microsoft.Insights/activityLogAlerts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/loadBalancers/outboundRules": {
		Name:         "Microsoft.Network/loadBalancers/outboundRules",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.HybridCompute/machines": {
		Name:         "Microsoft.HybridCompute/machines",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/loadBalancers/inboundNatRules": {
		Name:         "Microsoft.Network/loadBalancers/inboundNatRules",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Resources/providers": {
		Name:         "Microsoft.Resources/providers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/routeTables": {
		Name:         "Microsoft.Network/routeTables",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DocumentDB/databaseAccounts": {
		Name:         "Microsoft.DocumentDB/databaseAccounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DocumentDB/restorableDatabaseAccounts": {
		Name:         "Microsoft.DocumentDB/restorableDatabaseAccounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/applicationGateways": {
		Name:         "Microsoft.Network/applicationGateways",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Security/automations": {
		Name:         "Microsoft.Security/automations",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Kubernetes/connectedClusters": {
		Name:         "Microsoft.Kubernetes/connectedClusters",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.KeyVault/vaults/keys": {
		Name:         "Microsoft.KeyVault/vaults/keys",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.KeyVault/vaults/certificates": {
		Name:         "Microsoft.KeyVault/vaults/certificates",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.KeyVault/vaults/keys/Versions": {
		Name:         "Microsoft.KeyVault/vaults/keys/Versions",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DBforMariaDB/servers": {
		Name:         "Microsoft.DBforMariaDB/servers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DBforMariaDB/servers/databases": {
		Name:         "Microsoft.DBforMariaDB/servers/databases",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Web/plan": {
		Name:         "Microsoft.Web/plan",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Resources/tenants": {
		Name:         "Microsoft.Resources/tenants",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/virtualNetworkGateways": {
		Name:         "Microsoft.Network/virtualNetworkGateways",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Devices/iotHubs": {
		Name:         "Microsoft.Devices/iotHubs",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Logic/workflows": {
		Name:         "Microsoft.Logic/workflows",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Sql/flexibleServers": {
		Name:         "Microsoft.Sql/flexibleServers",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Resources/links": {
		Name:         "Microsoft.Resources/links",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Resources/subscriptions": {
		Name:         "Microsoft.Resources/subscriptions",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/images": {
		Name:         "Microsoft.Compute/images",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/virtualMachines": {
		Name:         "Microsoft.Compute/virtualMachines",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/natGateways": {
		Name:         "Microsoft.Network/natGateways",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/loadBalancers/probes": {
		Name:         "Microsoft.Network/loadBalancers/probes",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.KeyVault/vaults": {
		Name:         "Microsoft.KeyVault/vaults",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.KeyVault/managedHsms": {
		Name:         "Microsoft.KeyVault/managedHsms",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.KeyVault/vaults/secrets": {
		Name:         "Microsoft.KeyVault/vaults/secrets",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.AppConfiguration/configurationStores": {
		Name:         "Microsoft.AppConfiguration/configurationStores",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Storage/storageAccounts": {
		Name:         "Microsoft.Storage/storageAccounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.AppPlatform/Spring": {
		Name:         "Microsoft.AppPlatform/Spring",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/galleries": {
		Name:         "Microsoft.Compute/galleries",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/hostGroups": {
		Name:         "Microsoft.Compute/hostGroups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/hostGroups/hosts": {
		Name:         "Microsoft.Compute/hostGroups/hosts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/restorePointCollections": {
		Name:         "Microsoft.Compute/restorePointCollections",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Compute/sshPublicKeys": {
		Name:         "Microsoft.Compute/sshPublicKeys",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Cdn/profiles/endpoints": {
		Name:         "Microsoft.Cdn/profiles/endpoints",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.BotService/botServices": {
		Name:         "Microsoft.BotService/botServices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DocumentDB/cassandraClusters": {
		Name:         "Microsoft.DocumentDB/cassandraClusters",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Network/ddosProtectionPlans": {
		Name:         "Microsoft.Network/ddosProtectionPlans",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"microsoft.Sql/instancePools": {
		Name:         "microsoft.Sql/instancePools",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"microsoft.NetApp/netAppAccounts": {
		Name:         "microsoft.NetApp/netAppAccounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.NetApp/netAppAccounts/capacityPools": {
		Name:         "Microsoft.NetApp/netAppAccounts/capacityPools",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.DesktopVirtualization/hostpools": {
		Name:         "Microsoft.DesktopVirtualization/hostpools",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Devtestlab/labs": {
		Name:         "Microsoft.Devtestlab/labs",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Purview/Accounts": {
		Name:         "Microsoft.Purview/Accounts",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.PowerBIDedicated/capacities": {
		Name:         "Microsoft.PowerBIDedicated/capacities",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Insights/components": {
		Name:         "Microsoft.Insights/components",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Lighthouse/definition": {
		Name:         "Microsoft.Lighthouse/definition",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Lighthouse/assignment": {
		Name:         "Microsoft.Lighthouse/assignment",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Maintenance/maintenanceConfigurations": {
		Name:         "Microsoft.Maintenance/maintenanceConfigurations",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Monitor/logProfiles": {
		Name:         "Microsoft.Monitor/logProfiles",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Resources/subscriptions/resources": {
		Name:         "Microsoft.Resources/subscriptions/resources",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},
}


var ResourceTypesList = []string{
  "Microsoft.App/containerApps",
  "Microsoft.Blueprint/blueprints",
  "Microsoft.Cdn/profiles",
  "Microsoft.Compute/cloudServices",
  "Microsoft.ContainerInstance/containerGroups",
  "Microsoft.DataMigration/services",
  "Microsoft.DataProtection/backupVaults",
  "Microsoft.DataProtection/backupJobs",
  "Microsoft.DataProtection/backupVaults/backupPolicies",
  "Microsoft.Logic/integrationAccounts",
  "Microsoft.Network/bastionHosts",
  "Microsoft.Network/connections",
  "Microsoft.Network/firewallPolicies",
  "Microsoft.Network/localNetworkGateways",
  "Microsoft.Network/privateLinkServices",
  "Microsoft.Network/publicIPPrefixes",
  "Microsoft.Network/virtualHubs",
  "Microsoft.Network/virtualWans",
  "Microsoft.Network/vpnGateways",
  "Microsoft.Network/vpnGateways/vpnConnections",
  "Microsoft.Network/vpnSites",
  "Microsoft.OperationalInsights/workspaces",
  "Microsoft.StreamAnalytics/cluster",
  "Microsoft.TimeSeriesInsights/environments",
  "Microsoft.VirtualMachineImages/imageTemplates",
  "Microsoft.Web/serverFarms",
  "Microsoft.Compute/virtualMachineScaleSets/virtualMachines",
  "Microsoft.Automation/automationAccounts",
  "Microsoft.Automation/automationAccounts/variables",
  "Microsoft.Network/dnsZones",
  "Microsoft.Databricks/workspaces",
  "Microsoft.Network/privateDnsZones",
  "Microsoft.Network/privateEndpoints",
  "Microsoft.Network/networkWatchers",
  "Microsoft.Resources/subscriptions/resourceGroups",
  "Microsoft.Web/staticSites",
  "Microsoft.Web/sites/slots",
  "Microsoft.CognitiveServices/accounts",
  "Microsoft.Sql/managedInstances",
  "Microsoft.Sql/virtualclusters",
  "Microsoft.Sql/managedInstances/databases",
  "Microsoft.Sql/servers/databases",
  "Microsoft.Storage/storageAccounts/largeFileSharesState",
  "Microsoft.DBforPostgreSQL/servers",
  "Microsoft.DBforPostgreSQL/flexibleservers",
  "Microsoft.AnalysisServices/servers",
  "Microsoft.Security/pricings",
  "Microsoft.Insights/guestDiagnosticSettings",
  "Microsoft.Insights/autoscaleSettings",
  "Microsoft.Web/hostingEnvironments",
  "Microsoft.Cache/redis",
  "Microsoft.ContainerRegistry/registries",
  "Microsoft.DataFactory/factories/pipelines",
  "Microsoft.Compute/resourceSku",
  "Microsoft.Network/expressRouteCircuits",
  "Microsoft.Management/managementgroups",
  "microsoft.SqlVirtualMachine/SqlVirtualMachines",
  "Microsoft.SqlVirtualMachine/SqlVirtualMachineGroups",
  "Microsoft.Storage/storageAccounts/tableServices",
  "Microsoft.Synapse/workspaces",
  "Microsoft.Synapse/workspaces/bigdatapools",
  "Microsoft.Synapse/workspaces/sqlpools",
  "Microsoft.StreamAnalytics/streamingJobs",
  "Microsoft.CostManagement/CostBySubscription",
  "Microsoft.ContainerService/managedClusters",
  "Microsoft.ContainerService/serviceVersions",
  "Microsoft.DataFactory/factories",
  "Microsoft.Sql/servers",
  "Microsoft.Sql/servers/jobagents",
  "Microsoft.Security/autoProvisioningSettings",
  "Microsoft.Insights/logProfiles",
  "Microsoft.DataBoxEdge/dataBoxEdgeDevices",
  "Microsoft.Network/loadBalancers",
  "Microsoft.Network/azureFirewalls",
  "Microsoft.Management/locks",
  "Microsoft.Compute/virtualMachineScaleSets/networkInterfaces",
  "Microsoft.Network/frontDoors",
  "Microsoft.Authorization/policyAssignments",
  "Microsoft.Authorization/userEffectiveAccess",
  "Microsoft.Search/searchServices",
  "Microsoft.Security/settings",
  "Microsoft.RecoveryServices/vaults",
  "Microsoft.RecoveryServices/vaults/backupJobs",
  "Microsoft.RecoveryServices/vaults/backupPolicies",
  "Microsoft.RecoveryServices/vaults/backupItems",
  "Microsoft.Compute/diskEncryptionSets",
  "Microsoft.DocumentDB/databaseAccounts/sqlDatabases",
  "Microsoft.EventGrid/topics",
  "Microsoft.EventHub/namespaces",
  "Microsoft.EventHub/namespaces/eventHubs",
  "Microsoft.MachineLearningServices/workspaces",
  "Microsoft.Dashboard/grafana",
  "Microsoft.DesktopVirtualization/workspaces",
  "Microsoft.Network/trafficManagerProfiles",
  "Microsoft.Network/dnsResolvers",
  "Microsoft.CostManagement/CostByResourceType",
  "Microsoft.Network/networkInterfaces",
  "Microsoft.Network/publicIPAddresses",
  "Microsoft.HealthcareApis/services",
  "Microsoft.ServiceBus/namespaces",
  "Microsoft.Web/sites",
  "Microsoft.Compute/availabilitySets",
  "Microsoft.Network/virtualNetworks",
  "Microsoft.Security/securityContacts",
  "Microsoft.EventGrid/domains",
  "Microsoft.KeyVault/deletedVaults",
  "Microsoft.Storage/storageAccounts/tableServices/tables",
  "Microsoft.Compute/snapshots",
  "Microsoft.Kusto/clusters",
  "Microsoft.StorageSync/storageSyncServices",
  "Microsoft.Security/locations/jitNetworkAccessPolicies",
  "Microsoft.Network/virtualNetworks/subnets",
  "Microsoft.Network/loadBalancers/backendAddressPools",
  "Microsoft.Network/loadBalancers/loadBalancingRules",
  "Microsoft.DataLakeStore/accounts",
  "Microsoft.StorageCache/caches",
  "Microsoft.Batch/batchAccounts",
  "Microsoft.Network/networkSecurityGroups",
  "Microsoft.Authorization/roleDefinitions",
  "Microsoft.Network/applicationSecurityGroups",
  "Microsoft.Authorization/roleAssignment",
  "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases",
  "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections",
  "Microsoft.Network/networkWatchers/flowLogs",
  "microsoft.Sql/servers/elasticpools",
  "Microsoft.Security/subAssessments",
  "Microsoft.Compute/disks",
  "Microsoft.Devices/ProvisioningServices",
  "Microsoft.HDInsight/clusters",
  "Microsoft.ServiceFabric/clusters",
  "Microsoft.SignalRService/signalR",
  "Microsoft.Storage/storageAccounts/blob",
  "Microsoft.Storage/storageaccounts/blobservices/containers",
  "Microsoft.Storage/storageAccounts/blobServices",
  "Microsoft.Storage/storageAccounts/queueServices",
  "Microsoft.ApiManagement/service",
  "Microsoft.ApiManagement/backend",
  "Microsoft.Compute/virtualMachineScaleSets",
  "Microsoft.DataFactory/factories/datasets",
  "Microsoft.Authorization/policyDefinitions",
  "Microsoft.Resources/subscriptions/locations",
  "Microsoft.Compute/diskAccesses",
  "Microsoft.DBforMySQL/servers",
  "Microsoft.DBforMySQL/flexibleservers",
  "Microsoft.Cache/redisenterprise",
  "Microsoft.DataLakeAnalytics/accounts",
  "Microsoft.Insights/activityLogAlerts",
  "Microsoft.Network/loadBalancers/outboundRules",
  "Microsoft.HybridCompute/machines",
  "Microsoft.Network/loadBalancers/inboundNatRules",
  "Microsoft.Resources/providers",
  "Microsoft.Network/routeTables",
  "Microsoft.DocumentDB/databaseAccounts",
  "Microsoft.DocumentDB/restorableDatabaseAccounts",
  "Microsoft.Network/applicationGateways",
  "Microsoft.Security/automations",
  "Microsoft.Kubernetes/connectedClusters",
  "Microsoft.KeyVault/vaults/keys",
  "Microsoft.KeyVault/vaults/certificates",
  "Microsoft.KeyVault/vaults/keys/Versions",
  "Microsoft.DBforMariaDB/servers",
  "Microsoft.DBforMariaDB/servers/databases",
  "Microsoft.Web/plan",
  "Microsoft.Resources/tenants",
  "Microsoft.Network/virtualNetworkGateways",
  "Microsoft.Devices/iotHubs",
  "Microsoft.Logic/workflows",
  "Microsoft.Sql/flexibleServers",
  "Microsoft.Resources/links",
  "Microsoft.Resources/subscriptions",
  "Microsoft.Compute/images",
  "Microsoft.Compute/virtualMachines",
  "Microsoft.Network/natGateways",
  "Microsoft.Network/loadBalancers/probes",
  "Microsoft.KeyVault/vaults",
  "Microsoft.KeyVault/managedHsms",
  "Microsoft.KeyVault/vaults/secrets",
  "Microsoft.AppConfiguration/configurationStores",
  "Microsoft.Storage/storageAccounts",
  "Microsoft.AppPlatform/Spring",
  "Microsoft.Compute/galleries",
  "Microsoft.Compute/hostGroups",
  "Microsoft.Compute/hostGroups/hosts",
  "Microsoft.Compute/restorePointCollections",
  "Microsoft.Compute/sshPublicKeys",
  "Microsoft.Cdn/profiles/endpoints",
  "Microsoft.BotService/botServices",
  "Microsoft.DocumentDB/cassandraClusters",
  "Microsoft.Network/ddosProtectionPlans",
  "microsoft.Sql/instancePools",
  "microsoft.NetApp/netAppAccounts",
  "Microsoft.NetApp/netAppAccounts/capacityPools",
  "Microsoft.DesktopVirtualization/hostpools",
  "Microsoft.Devtestlab/labs",
  "Microsoft.Purview/Accounts",
  "Microsoft.PowerBIDedicated/capacities",
  "Microsoft.Insights/components",
  "Microsoft.Lighthouse/definition",
  "Microsoft.Lighthouse/assignment",
  "Microsoft.Maintenance/maintenanceConfigurations",
  "Microsoft.Monitor/logProfiles",
  "Microsoft.Resources/subscriptions/resources",
}