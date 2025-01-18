//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-azure/azure -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package provider

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azcertificates"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	appservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v2"
	analytics "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
	store "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/springappdiscovery/armspringappdiscovery"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder"
	"github.com/Azure/azure-sdk-for-go/services/preview/web/mgmt/2015-08-01-preview/web"
	azblobOld "github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/queue/queues"
	"github.com/tombuildsstuff/giovanni/storage/2019-12-12/blob/accounts"
)

type Metadata struct {
	ID               string
	Name             string
	SubscriptionID   string
	Location         string
	CloudEnvironment string
	ResourceType     string
	IntegrationID    string
}

//  ===================  APIManagement ==================

//index:microsoft_apimanagement_service
//getfilter:name=description.APIManagement.name
//getfilter:resource_group=description.ResourceGroup
type APIManagementDescription struct {
	APIManagement               armapimanagement.ServiceResource
	DiagnosticSettingsResources *[]armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

type APIManagementBackendDescription struct {
	APIManagementBackend armapimanagement.BackendContract
	ServiceName          string
	ResourceGroup        string
}

//  ===================  Automation ==================

//index:microsoft_automation_automationAccounts
type AutomationAccountsDescription struct {
	Automation    armautomation.Account
	ResourceGroup string
}

//index:microsoft_automation_automationVariables
type AutomationVariablesDescription struct {
	Automation    armautomation.Variable
	AccountName   string
	ResourceGroup string
}

//  ===================  App Configuration ==================

//index:microsoft_appconfiguration_configurationstores
//getfilter:name=description.ConfigurationStore.name
//getfilter:resource_group=description.ResourceGroup
type AppConfigurationDescription struct {
	ConfigurationStore          armappconfiguration.ConfigurationStore
	DiagnosticSettingsResources *[]armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== web ==================

//index:microsoft_web_hostingenvironments
//getfilter:name=description.AppServiceEnvironmentResource.name
//getfilter:resource_group=description.ResourceGroup
type AppServiceEnvironmentDescription struct {
	AppServiceEnvironmentResource appservice.EnvironmentResource
	ResourceGroup                 string
}

//index:microsoft_web_sites
//getfilter:name=description.Site.name
//getfilter:resource_group=description.ResourceGroup
type AppServiceFunctionAppDescription struct {
	Site               appservice.Site
	SiteAuthSettings   appservice.SiteAuthSettings
	SiteConfigResource appservice.SiteConfigResource
	ResourceGroup      string
}

//index:microsoft_web_staticsites
//getfilter:name=description.Site.name
//getfilter:resource_group=description.ResourceGroup
type AppServiceWebAppDescription struct {
	Site               appservice.Site
	SiteAuthSettings   appservice.SiteAuthSettings
	SiteConfigResource appservice.SiteConfigResource
	SiteLogConfig      appservice.SiteLogsConfig
	VnetInfo           appservice.VnetInfoResource
	StorageAccounts    map[string]*appservice.AzureStorageInfoValue
	ResourceGroup      string
}

type AppServiceWebAppSlotDescription struct {
	Site          appservice.Site
	AppName       string
	ResourceGroup string
}

//index:microsoft_web_plan
//getfilter:name=description.Site.name
//getfilter:resource_group=description.ResourceGroup
type AppServicePlanDescription struct {
	Plan          appservice.Plan
	Apps          []*appservice.Site
	ResourceGroup string
}

//index:microsoft_app_containerapp
type ContainerAppDescription struct {
	ResourceGroup string
	Server        appservice.ContainerApp
}

//index:microsoft_app_managedenvironment
type AppManagedEnvironmentDescription struct {
	ResourceGroup      string
	HostingEnvironment web.HostingEnvironment
}

//index:microsoft_web_serverfarm
type WebServerFarmsDescription struct {
	ResourceGroup string
	ServerFarm    appservice.Plan
}

//  =================== blueprint ==================

//index:microsoft_blueprint_blueprint
type BlueprintDescription struct {
	ResourceGroup string
	Blueprint     armblueprint.Blueprint
}

//  =================== compute ==================

//index:microsoft_compute_disks
//getfilter:name=description.Disk.name
//getfilter:resource_group=description.ResourceGroup
type ComputeDiskDescription struct {
	Disk          armcompute.Disk
	ResourceGroup string
}

//index:microsoft_compute_disksreadops
type ComputeDiskReadOpsDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_disksreadopsdaily
type ComputeDiskReadOpsDailyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_disksreadopshourly
type ComputeDiskReadOpsHourlyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_diskswriteops
type ComputeDiskWriteOpsDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_diskswriteopsdaily
type ComputeDiskWriteOpsDailyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_diskswriteopshourly
type ComputeDiskWriteOpsHourlyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_diskaccesses
//getfilter:name=description.DiskAccess.name
//getfilter:resource_group=description.ResourceGroup
type ComputeDiskAccessDescription struct {
	DiskAccess    armcompute.DiskAccess
	ResourceGroup string
}

//index:microsoft_compute_virtualmachinescalesets
//getfilter:name=description.VirtualMachineScaleSet.name
//getfilter:resource_group=description.ResourceGroup
type ComputeVirtualMachineScaleSetDescription struct {
	VirtualMachineScaleSet           armcompute.VirtualMachineScaleSet
	VirtualMachineScaleSetExtensions []armcompute.VirtualMachineScaleSetExtension
	ResourceGroup                    string
}

//index:microsoft_compute_virtualmachinescalesetnetworkinterface
type ComputeVirtualMachineScaleSetNetworkInterfaceDescription struct {
	VirtualMachineScaleSet armcompute.VirtualMachineScaleSet
	NetworkInterface       armnetwork.Interface
	ResourceGroup          string
}

//index:microsoft_compute_virtualmachinescalesetvm
//getfilter:scale_set_name=description.VirtualMachineScaleSet.name
//getfilter:instance_id=description.ScaleSetVM.InstanceID
//getfilter:resource_group=description.ResourceGroup
type ComputeVirtualMachineScaleSetVmDescription struct {
	VirtualMachineScaleSet armcompute.VirtualMachineScaleSet
	ScaleSetVM             armcompute.VirtualMachineScaleSetVM
	PowerState             string
	ResourceGroup          string
}

//index:microsoft_compute_snapshots
//getfilter:name=description.Snapshot.Name
//getfilter:resource_group=description.ResourceGroup
type ComputeSnapshotsDescription struct {
	Snapshot      armcompute.Snapshot
	ResourceGroup string
}

//index:microsoft_compute_availabilityset
//getfilter:name=description.AvailabilitySet.Name
//getfilter:resource_group=description.ResourceGroup
type ComputeAvailabilitySetDescription struct {
	AvailabilitySet armcompute.AvailabilitySet
	ResourceGroup   string
}

//index:microsoft_compute_diskencryptionset
//getfilter:name=description.DiskEncryptionSet.Name
//getfilter:resource_group=description.ResourceGroup
type ComputeDiskEncryptionSetDescription struct {
	DiskEncryptionSet armcompute.DiskEncryptionSet
	ResourceGroup     string
}

//index:microsoft_compute_gallery
//getfilter:name=description.ImageGallery.Name
//getfilter:resource_group=description.ResourceGroup
type ComputeImageGalleryDescription struct {
	ImageGallery  armcompute.Gallery
	ResourceGroup string
}

//index:microsoft_compute_image
//getfilter:name=Description.Image.Name
//getfilter:resource_group=Description.Image.ResourceGroup
type ComputeImageDescription struct {
	Image         armcompute.Image
	ResourceGroup string
}

type ComputeHostGroupDescription struct {
	HostGroup     armcompute.DedicatedHostGroup
	ResourceGroup string
}

type ComputeHostGroupHostDescription struct {
	Host          armcompute.DedicatedHost
	ResourceGroup string
}

type ComputeRestorePointCollectionDescription struct {
	RestorePointCollection armcompute.RestorePointCollection
	ResourceGroup          string
}

type ComputeSSHPublicKeyDescription struct {
	SSHPublicKey  armcompute.SSHPublicKeyResource
	ResourceGroup string
}

//  =================== databoxedge ==================

//index:microsoft_databoxedge_databoxedgedevices
//getfilter:name=description.Device.name
//getfilter:resource_group=description.ResourceGroup
type DataboxEdgeDeviceDescription struct {
	Device        armdataboxedge.Device
	ResourceGroup string
}

//  =================== healthcareapis ==================

//index:microsoft_healthcareapis_services
//getfilter:name=description.ServicesDescription.name
//getfilter:resource_group=description.ResourceGroup
type HealthcareServiceDescription struct {
	ServicesDescription         armhealthcareapis.ServicesDescription
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	PrivateEndpointConnections  []*armhealthcareapis.PrivateEndpointConnectionDescription
	ResourceGroup               string
}

//  =================== storagecache ==================

//index:microsoft_storagecache_caches
//getfilter:name=description.Cache.name
//getfilter:resource_group=description.ResourceGroup
type HpcCacheDescription struct {
	Cache         armstoragecache.Cache
	ResourceGroup string
}

//  =================== keyvault ==================

//index:microsoft_keyvault_vaults_keys
//getfilter:vault_name=description.Vault.name
//getfilter:name=description.Key.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultKeyDescription struct {
	Vault         armkeyvault.Resource
	Key           armkeyvault.Key
	ResourceGroup string
}

//index:microsoft_keyvault_vaults_keys_versions
//getfilter:vault_name=description.Vault.name
//getfilter:name=description.Version.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultKeyVersionDescription struct {
	Vault         armkeyvault.Resource
	Key           armkeyvault.Key
	Version       armkeyvault.Key
	ResourceGroup string
}

//  =================== containerservice ==================

//index:microsoft_containerservice_managedclusters
//getfilter:name=description.ManagedCluster.name
//getfilter:resource_group=description.ResourceGroup
type KubernetesClusterDescription struct {
	ManagedCluster armcontainerservice.ManagedCluster
	ResourceGroup  string
}

//index:microsoft_containerservice_serviceversion
//getfilter:name=description.Orchestrator.name
//getfilter:resource_group=description.ResourceGroup
type KubernetesServiceVersionDescription struct {
	Version armcontainerservice.KubernetesVersion
}

//  =================== containerinstance ==================

//index:microsoft_containerinstance_containergroup
type ContainerInstanceContainerGroupDescription struct {
	ResourceGroup  string
	ContainerGroup armcontainerinstance.ContainerGroup
}

//  =================== cdn ==================

//index:microsoft_cdn_profile
type CDNProfileDescription struct {
	ResourceGroup string
	Profile       armcdn.Profile
}

type CDNEndpointDescription struct {
	ResourceGroup string
	Endpoint      armcdn.Endpoint
}

//  =================== network ==================

//index:microsoft_network_networkinterfaces
//getfilter:name=description.Interface.name
//getfilter:resource_group=description.ResourceGroup
type NetworkInterfaceDescription struct {
	Interface     armnetwork.Interface
	ResourceGroup string
}

//index:microsoft_network_networkwatchers
//getfilter:network_watcher_name=description.NetworkWatcherName
//getfilter:name=description.ManagedCluster.name
//getfilter:resource_group=description.ResourceGroup
type NetworkWatcherFlowLogDescription struct {
	NetworkWatcherName string
	FlowLog            armnetwork.FlowLog
	ResourceGroup      string
}

//index:microsoft_network_routetables
//getfilter:name=description.RouteTable.Name
//getfilter:resource_group=description.ResourceGroup
type RouteTablesDescription struct {
	RouteTable    armnetwork.RouteTable
	ResourceGroup string
}

//index:microsoft_network_applicationsecuritygroups
//getfilter:name=description.ApplicationSecurityGroup.Name
//getfilter:resource_group=description.ResourceGroup
type NetworkApplicationSecurityGroupsDescription struct {
	ApplicationSecurityGroup armnetwork.ApplicationSecurityGroup
	ResourceGroup            string
}

//index:microsoft_network_azurefirewall
//getfilter:name=description.AzureFirewall.Name
//getfilter:resource_group=description.ResourceGroup
type NetworkAzureFirewallDescription struct {
	AzureFirewall armnetwork.AzureFirewall
	ResourceGroup string
}

//index:microsoft_network_expressroutecircuit
//getfilter:name=description.ExpressRouteCircuit.name
//getfilter:resource_group=description.ResourceGroup
type ExpressRouteCircuitDescription struct {
	ExpressRouteCircuit armnetwork.ExpressRouteCircuit
	ResourceGroup       string
}

//index:microsoft_network_virtualnetworkgateway
//getfilter:name=description.VirtualNetworkGateway.Name
//getfilter:resource_group=description.ResourceGroup
type VirtualNetworkGatewayDescription struct {
	VirtualNetwork                  string
	VirtualNetworkGateway           armnetwork.VirtualNetworkGateway
	ResourceGroup                   string
	VirtualNetworkGatewayConnection []*armnetwork.VirtualNetworkGatewayConnectionListEntity
}

//index:microsoft_network_dnszone
//getfilter:name=description.Zone.Name
//getfilter:resource_group=description.ResourceGroup
type DNSZoneDescription struct { // TODO: Implement describer func
	Zone          armdns.Zone
	ResourceGroup string
}

//index:microsoft_network_firewallpolicy
//getfilter:name=description.FirewallPolicy.Name
//getfilter:resource_group=description.ResourceGroup
type FirewallPolicyDescription struct {
	FirewallPolicy armnetwork.FirewallPolicy
	ResourceGroup  string
}

//index:microsoft_network_frontdoorwebapplicationfirewallpolicy
//getfilter:name=description.WebApplicationFirewallPolicy.Name
//getfilter:resource_group=description.ResourceGroup
type FrontdoorWebApplicationFirewallPolicyDescription struct { // TODO: Implement describer func
	WebApplicationFirewallPolicy armfrontdoor.WebApplicationFirewallPolicy
	ResourceGroup                string
}

//index:microsoft_network_localnetworkgateway
//getfilter:name=description.LocalNetworkGateway.Name
//getfilter:resource_group=description.ResourceGroup
type LocalNetworkGatewayDescription struct {
	LocalNetworkGateway armnetwork.LocalNetworkGateway
	ResourceGroup       string
}

//index:microsoft_network_natgateways
//getfilter:name=description.NatGateway.Name
//getfilter:resource_group=description.ResourceGroup
type NatGatewayDescription struct {
	NatGateway    armnetwork.NatGateway
	ResourceGroup string
}

//index:microsoft_network_privatelinkservice
//getfilter:name=description.PrivateLinkService.Name
//getfilter:resource_group=description.ResourceGroup
type PrivateLinkServiceDescription struct {
	PrivateLinkService armnetwork.PrivateLinkService
	ResourceGroup      string
}

//index:microsoft_network_routefilter
//getfilter:name=description.RouteFilter.Name
//getfilter:resource_group=description.ResourceGroup
type RouteFilterDescription struct {
	RouteFilter   armnetwork.RouteFilter
	ResourceGroup string
}

//index:microsoft_network_vpngateway
//getfilter:name=description.VpnGateway.Name
//getfilter:resource_group=description.ResourceGroup
type VpnGatewayDescription struct {
	VpnGateway    armnetwork.VPNGateway
	ResourceGroup string
}

//index:microsoft_network_vpngatewayvpnconnection
type VpnGatewayVpnConnectionDescription struct {
	ResourceGroup string
	VpnConnection armnetwork.VPNConnection
	VpnGateway    armnetwork.VPNGateway
}

//index:microsoft_network_vpnsite
type VpnSiteDescription struct {
	ResourceGroup string
	VpnSite       armnetwork.VPNSite
}

//index:microsoft_network_publicipaddresses
//getfilter:name=description.PublicIPAddress.Name
//getfilter:resource_group=description.ResourceGroup
type PublicIPAddressDescription struct {
	PublicIPAddress armnetwork.PublicIPAddress
	ResourceGroup   string
}

//index:microsoft_network_publicipprefix
type PublicIPPrefixDescription struct {
	ResourceGroup  string
	PublicIPPrefix armnetwork.PublicIPPrefix
}

//index:microsoft_network_dnszones
type DNSZonesDescription struct {
	ResourceGroup string
	DNSZone       armdns.Zone
}

//index:microsoft_network_bastianhosts
type BastionHostsDescription struct {
	ResourceGroup string
	BastianHost   armnetwork.BastionHost
}

//index:microsoft_network_connection
type ConnectionDescription struct {
	ResourceGroup string
	Connection    armnetwork.VirtualNetworkGatewayConnection
}

//index:microsoft_network_virtualhubs
type VirtualHubsDescription struct {
	ResourceGroup string
	VirtualHub    armnetwork.VirtualHub
}

//index:microsoft_network_virtualwans
type VirtualWansDescription struct {
	ResourceGroup string
	VirtualWan    armnetwork.VirtualWAN
}

//index:microsoft_network_dnsresolvers
type DNSResolverDescription struct {
	ResourceGroup string
	DNSResolver   armdnsresolver.DNSResolver
}

type TrafficManagerProfileDescription struct {
	ResourceGroup string
	Profile       armtrafficmanager.Profile
}

//index:microsoft_network_privatednszones
type PrivateDNSZonesDescription struct {
	ResourceGroup string
	PrivateZone   armprivatedns.PrivateZone
}

//index:microsoft_network_privateendpoint
type PrivateEndpointDescription struct {
	ResourceGroup   string
	PrivateEndpoint armnetwork.PrivateEndpoint
}

type NetworkDDoSProtectionPlanDescription struct {
	ResourceGroup      string
	DDoSProtectionPlan armnetwork.DdosProtectionPlan
}

//  =================== policy ==================

//index:microsoft_authorization_policyassignments
//getfilter:name=description.Assignment.name
type PolicyAssignmentDescription struct {
	Assignment armpolicy.Assignment
	Resource   armresources.GenericResource
}

//  =================== redis ==================

//index:microsoft_cache_redis
//getfilter:name=description.ResourceType.name
//getfilter:resource_group=description.ResourceGroup
type RedisCacheDescription struct {
	ResourceInfo  armredis.ResourceInfo
	ResourceGroup string
}

//index:microsoft_cache_redisenterprise
type RedisEnterpriseCacheDescription struct {
	ResourceGroup   string
	RedisEnterprise armredisenterprise.Cluster
}

//  =================== links ==================

//index:microsoft_resources_links
//getfilter:id=description.ResourceLink.id
type ResourceLinkDescription struct {
	ResourceLink armlinks.ResourceLink
}

//  =================== authorization ==================

//index:microsoft_authorization_elevateaccessroleassignment
//getfilter:id=description.RoleAssignment.id
type RoleAssignmentDescription struct {
	RoleAssignment armauthorization.RoleAssignment
}

//index:microsoft_authorization_roledefinitions
type RoleDefinitionDescription struct {
	RoleDefinition armauthorization.RoleDefinition
}

//index:microsoft_authorization_policydefinition
//getfilter:name=description.Definition.Name
type PolicyDefinitionDescription struct {
	Definition armpolicy.Definition
	TurboData  map[string]interface{}
}

//index:microsoft_authorization_usereffectiveaccess
type UserEffectiveAccessDescription struct {
	RoleAssignment    armauthorization.RoleAssignment
	PrincipalName     string
	PrincipalId       string
	PrincipalType     armauthorization.PrincipalType
	Scope             string
	ScopeType         string
	AssignmentType    string
	ParentPrincipalId *string
}

//  =================== security ==================

//index:microsoft_security_autoprovisioningsettings
//getfilter:name=description.AutoProvisioningSetting.name
type SecurityCenterAutoProvisioningDescription struct {
	AutoProvisioningSetting armsecurity.AutoProvisioningSetting
}

//index:microsoft_security_securitycontacts
//getfilter:name=description.Contact.name
type SecurityCenterContactDescription struct {
	Contact armsecurity.Contact
}

//index:microsoft_security_locations_jitnetworkaccesspolicies
type SecurityCenterJitNetworkAccessPolicyDescription struct {
	JitNetworkAccessPolicy armsecurity.JitNetworkAccessPolicy
}

//index:microsoft_security_settings
//getfilter:name=description.Setting.name
type SecurityCenterSettingDescription struct {
	Setting             armsecurity.Setting
	ExportSettingStatus bool
}

//index:microsoft_security_pricings
//getfilter:name=description.Pricing.Name
type SecurityCenterSubscriptionPricingDescription struct {
	Pricing armsecurity.Pricing
}

//index:microsoft_security_automations
//getfilter:name=description.Automation.name
//getfilter:resource_group=description.ResourceGroup
type SecurityCenterAutomationDescription struct {
	Automation    armsecurity.Automation
	ResourceGroup string
}

//index:microsoft_security_subassessments
type SecurityCenterSubAssessmentDescription struct {
	SubAssessment armsecurity.SubAssessment
	ResourceGroup string
}

//  =================== storage ==================

//index:microsoft_storage_storageaccounts_containers
//getfilter:name=description.ListContainerItem.name
//getfilter:resource_group=description.ResourceGroup
//getfilter:account_name=description.AccountName
type StorageContainerDescription struct {
	AccountName        string
	ListContainerItem  armstorage.ListContainerItem
	ImmutabilityPolicy armstorage.ImmutabilityPolicy
	ResourceGroup      string
}

//index:microsoft_storage_blobs
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageBlobDescription struct {
	Blob          azblobOld.BlobItemInternal
	AccountName   string
	IsSnapshot    bool
	ContainerName string
	ResourceGroup string
}

//index:microsoft_storage_blobservices
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageBlobServiceDescription struct {
	BlobService   armstorage.BlobServiceProperties
	AccountName   string
	Location      string
	ResourceGroup string
}

//index:microsoft_storage_queues
//listfilter:name=description.Queue.Name
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageQueueDescription struct {
	Queue         armstorage.ListQueue
	AccountName   string
	Location      string
	ResourceGroup string
}

//index:microsoft_storage_fileshares
//listfilter:name=description.FileShare.Name
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageFileShareDescription struct {
	FileShare     armstorage.FileShareItem
	AccountName   string
	Location      string
	ResourceGroup string
}

//index:microsoft_storage_tables
//listfilter:name=description.Table.Name
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageTableDescription struct {
	Table         armstorage.Table
	AccountName   string
	Location      string
	ResourceGroup string
}

//index:microsoft_storage_tableservices
//listfilter:name=description.TableService.Name
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageTableServiceDescription struct {
	TableService  armstorage.TableServiceProperties
	AccountName   string
	Location      string
	ResourceGroup string
}

//  =================== network ==================

//index:microsoft_network_virtualnetworks_subnets
//getfilter:name=description.Subnet.name
//getfilter:resource_group=description.ResourceGroup
//getfilter:virtual_network_name=description.VirtualNetworkName
type SubnetDescription struct {
	VirtualNetworkName string
	Subnet             armnetwork.Subnet
	ResourceGroup      string
}

//index:microsoft_network_virtualnetworks
//getfilter:name=description.VirtualNetwork.name
//getfilter:resource_group=description.ResourceGroup
type VirtualNetworkDescription struct {
	VirtualNetwork armnetwork.VirtualNetwork
	ResourceGroup  string
}

//  =================== subscriptions ==================

//index:microsoft_resources_tenants
type TenantDescription struct {
	TenantIDDescription armsubscription.TenantIDDescription
}

//index:microsoft_resources_subscriptions
type SubscriptionDescription struct {
	Subscription armsubscription.Subscription
	Tags         map[string][]string
}

//  =================== network ==================

//index:microsoft_network_applicationgateways
//getfilter:name=description.ApplicationGateway.name
//getfilter:resource_group=description.ResourceGroup
type ApplicationGatewayDescription struct {
	ApplicationGateway          armnetwork.ApplicationGateway
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== batch ==================

//index:microsoft_batch_batchaccounts
//getfilter:name=description.Account.name
//getfilter:resource_group=description.ResourceGroup
type BatchAccountDescription struct {
	Account                     armbatch.Account
	DiagnosticSettingsResources *[]armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== cognitiveservices ==================

//index:microsoft_cognitiveservices_accounts
//getfilter:name=description.Account.name
//getfilter:resource_group=description.ResourceGroup
type CognitiveAccountDescription struct {
	Account                     armcognitiveservices.Account
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== compute ==================

//index:microsoft_compute_virtualmachines
//getfilter:name=description.VirtualMachine.name
//getfilter:resource_group=description.ResourceGroup
type ComputeVirtualMachineDescription struct {
	VirtualMachine             armcompute.VirtualMachine
	VirtualMachineInstanceView armcompute.VirtualMachineInstanceView
	InterfaceIPConfigurations  []armnetwork.InterfaceIPConfiguration
	PublicIPs                  []string
	VirtualMachineExtension    []*armcompute.VirtualMachineExtension
	ExtensionsSettings         map[string]map[string]interface{}
	Assignments                *[]armguestconfiguration.Assignment
	ResourceGroup              string
}

//index:microsoft_compute_resourcesku
type ComputeResourceSKUDescription struct {
	ResourceSKU armcompute.ResourceSKU
}

//index:microsoft_compute_virtualmachinecpuutilization
type ComputeVirtualMachineCpuUtilizationDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_virtualmachinecpuutilizationdaily
type ComputeVirtualMachineCpuUtilizationDailyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_virtualmachinecpuutilizationhourly
type ComputeVirtualMachineCpuUtilizationHourlyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_cloudservice
type ComputeCloudServiceDescription struct {
	CloudService armcompute.CloudService
}

//  =================== containerregistry ==================

//index:microsoft_containerregistry_registries
//getfilter:name=description.Registry.name
//getfilter:resource_group=description.ResourceGroup
type ContainerRegistryDescription struct {
	Registry                      armcontainerregistry.Registry
	RegistryListCredentialsResult *armcontainerregistry.RegistryListCredentialsResult
	RegistryUsages                []*armcontainerregistry.RegistryUsage
	Webhooks                      []*armcontainerregistry.Webhook
	ResourceGroup                 string
}

//  =================== documentdb ==================

//index:microsoft_documentdb_databaseaccounts
//getfilter:name=description.DatabaseAccountGetResults.name
//getfilter:resource_group=description.ResourceGroup
type CosmosdbAccountDescription struct {
	DatabaseAccountGetResults armcosmos.DatabaseAccountGetResults
	ResourceGroup             string
}

//index:microsoft_documentdb_restorabledatabaseaccounts
//getfilter:name=description.Account.Name
//getfilter:resource_group=description.ResourceGroup
type CosmosdbRestorableDatabaseAccountDescription struct {
	Account       armcosmos.RestorableDatabaseAccountGetResult
	ResourceGroup string
}

//index:microsoft_documentdb_mongodatabases
//getfilter:account_name=description.Account.name
//getfilter:name=description.MongoDatabase.name
//getfilter:resource_group=description.ResourceGroup
type CosmosdbMongoDatabaseDescription struct {
	Account       armcosmos.DatabaseAccountGetResults
	MongoDatabase armcosmos.MongoDBDatabaseGetResults
	ResourceGroup string
}

//index:microsoft_documentdb_mongocollections
//getfilter:account_name=description.Account.name
//getfilter:name=description.MongoCollection.name
//getfilter:resource_group=description.ResourceGroup
type CosmosdbMongoCollectionDescription struct {
	Account         armcosmos.DatabaseAccountGetResults
	MongoDatabase   armcosmos.MongoDBDatabaseGetResults
	MongoCollection armcosmos.MongoDBCollectionGetResults
	Throughput      armcosmos.ThroughputSettingsGetResults
	ResourceGroup   string
}

//index:microsoft_documentdb_sqldatabases
//getfilter:account_name=description.Account.name
//getfilter:name=description.SqlDatabase.name
//getfilter:resource_group=description.ResourceGroup
type CosmosdbSqlDatabaseDescription struct {
	Account       armcosmos.DatabaseAccountGetResults
	SqlDatabase   armcosmos.SQLDatabaseGetResults
	ResourceGroup string
}

type CosmosdbCassandraClusterDescription struct {
	CassandraCluster armcosmos.ClusterResource
	ResourceGroup    string
}

//  =================== databricks ==================

//index:microsoft_databricks_workspace
type DatabricksWorkspaceDescription struct {
	Workspace     armdatabricks.Workspace
	ResourceGroup string
}

//  =================== datamigration ==================

//index:microsoft_datamigration_service
type DataMigrationServiceDescription struct {
	ResourceGroup string
	Service       armdatamigration.Service
}

//  =================== dataprotection ==================

//index:microsoft_dataprotection_backupvaults
type DataProtectionBackupVaultsDescription struct {
	ResourceGroup string
	BackupVaults  armdataprotection.BackupVaultResource
}

//index:microsoft_dataprotection_backupvaultsbackuppolicies
type DataProtectionBackupVaultsBackupPoliciesDescription struct {
	ResourceGroup  string
	BackupPolicies armdataprotection.BaseBackupPolicyResource
}

type DataProtectionJobDescription struct {
	DataProtectionJob armdataprotection.AzureBackupJobResource
	VaultName         string
	ResourceGroup     string
}

//  =================== datafactory ==================

//index:microsoft_datafactory_factories
//getfilter:name=description.Factory.name
//getfilter:resource_group=description.ResourceGroup
type DataFactoryDescription struct {
	Factory                    armdatafactory.Factory
	PrivateEndPointConnections []armdatafactory.PrivateEndpointConnectionResource
	ResourceGroup              string
}

//index:microsoft_datafactory_datafactorydatasets
//getfilter:factory_name=description.Factory.name
//getfilter:name=description.Dataset.name
//getfilter:resource_group=description.ResourceGroup
type DataFactoryDatasetDescription struct {
	Factory       armdatafactory.Factory
	Dataset       armdatafactory.DatasetResource
	ResourceGroup string
}

//index:microsoft_datafactory_datafactorypipelines
//getfilter:factory_name=description.Factory.name
//getfilter:name=description.Pipeline.name
//getfilter:resource_group=description.ResourceGroup
type DataFactoryPipelineDescription struct {
	Factory       armdatafactory.Factory
	Pipeline      armdatafactory.PipelineResource
	ResourceGroup string
}

//  =================== account ==================

//index:microsoft_datalakeanalytics_accounts
//getfilter:name=description.DataLakeAnalyticsAccount.name
//getfilter:resource_group=description.ResourceGroup
type DataLakeAnalyticsAccountDescription struct {
	DataLakeAnalyticsAccount   analytics.Account
	DiagnosticSettingsResource *[]armmonitor.DiagnosticSettingsResource
	ResourceGroup              string
}

//  =================== account ==================

//index:microsoft_datalakestore_accounts
//getfilter:name=description.DataLakeStoreAccount.name
//getfilter:resource_group=description.ResourceGroup
type DataLakeStoreDescription struct {
	DataLakeStoreAccount       store.Account
	DiagnosticSettingsResource *[]armmonitor.DiagnosticSettingsResource
	ResourceGroup              string
}

//  =================== insights ==================

type MonitoringMetric struct {
	// Resource Name
	DimensionValue string
	// MetadataValue represents a metric metadata value.
	MetaData *armmonitor.MetadataValue
	// Metric the result data of a query.
	Metric *armmonitor.Metric
	// The maximum metric value for the data point.
	Maximum *float64
	// The minimum metric value for the data point.
	Minimum *float64
	// The average of the metric values that correspond to the data point.
	Average *float64
	// The number of metric values that contributed to the aggregate value of this data point.
	SampleCount *float64
	// The sum of the metric values for the data point.
	Sum *float64
	// The time stamp used for the data point.
	TimeStamp string
	// The units in which the metric value is reported.
	Unit string
}

//index:microsoft_insights_guestdiagnosticsettings
//getfilter:name=description.DiagnosticSettingsResource.name
//getfilter:resource_group=description.ResourceGroup
type DiagnosticSettingDescription struct {
	DiagnosticSettingsResource armmonitor.DiagnosticSettingsResource
	ResourceGroup              string
}

//index:microsoft_insights_autoscalingsettings
//getfilter:name=description.AutoscaleSettingsResource.name
//getfilter:resource_group=description.ResourceGroup
type AutoscaleSettingDescription struct {
	AutoscaleSettingsResource armmonitor.AutoscaleSettingResource
	ResourceGroup             string
}

//  =================== eventgrid ==================

//index:microsoft_eventgrid_domains
//getfilter:name=description.Domain.name
//getfilter:resource_group=description.ResourceGroup
type EventGridDomainDescription struct {
	Domain                      armeventgrid.Domain
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== eventgrid ==================

//index:microsoft_eventgrid_topics
//getfilter:name=description.Topic.name
//getfilter:resource_group=description.ResourceGroup
type EventGridTopicDescription struct {
	Topic                       armeventgrid.Topic
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== eventhub ==================

//index:microsoft_eventhub_namespaces
//getfilter:name=description.EHNamespace.name
//getfilter:resource_group=description.ResourceGroup
type EventhubNamespaceDescription struct {
	EHNamespace                 armeventhub.EHNamespace
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	NetworkRuleSet              armeventhub.NetworkRuleSet
	PrivateEndpointConnection   []*armeventhub.PrivateEndpointConnection
	ResourceGroup               string
}

//index:microsoft_eventhub_namespaceseventhub
type EventhubNamespaceEventhubDescription struct {
	EHNamespace   armeventhub.EHNamespace
	EventHub      armeventhub.Eventhub
	ResourceGroup string
}

//  =================== frontdoor ==================

//index:microsoft_network_frontdoors
//getfilter:name=description.FrontDoor.name
//getfilter:resource_group=description.ResourceGroup
type FrontdoorDescription struct {
	FrontDoor                   armfrontdoor.FrontDoor
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== hdinsight ==================

//index:microsoft_hdinsight_clusterpools
//getfilter:name=description.Cluster.name
//getfilter:resource_group=description.ResourceGroup
type HdinsightClusterDescription struct {
	Cluster                     armhdinsight.Cluster
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== hybridcompute ==================

//index:microsoft_hybridcompute_machines
//getfilter:name=description.Machine.name
//getfilter:resource_group=description.ResourceGroup
type HybridComputeMachineDescription struct {
	Machine           armhybridcompute.Machine
	MachineExtensions []*armhybridcompute.MachineExtension
	ResourceGroup     string
}

//  =================== devices ==================

//index:microsoft_devices_iothubs
//getfilter:name=description.IotHubDescription.name
//getfilter:resource_group=description.ResourceGroup
type IOTHubDescription struct {
	IotHubDescription           armiothub.Description
	DiagnosticSettingsResources *[]armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//index:microsoft_devices_iothubdpses
//getfilter:name=description.IotHubDps.name
//getfilter:resource_group=description.ResourceGroup
type IOTHubDpsDescription struct {
	IotHubDps                   armdeviceprovisioningservices.ProvisioningServiceDescription
	DiagnosticSettingsResources *[]armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== keyvault ==================

//index:microsoft_keyvault_vaults
//getfilter:name=description.Resource.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultDescription struct {
	Resource                    armkeyvault.Resource
	Vault                       armkeyvault.Vault
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//index:microsoft_keyvault_vaults_certificates
//getfilter:name=description.Resource.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultCertificateDescription struct {
	Policy        azcertificates.CertificatePolicy
	Vault         armkeyvault.Resource
	ResourceGroup string
}

//index:microsoft_keyvault_deletedvaults
//getfilter:name=description.Vault.name
//getfilter:region=description.Vault.Properties.Location
type KeyVaultDeletedVaultDescription struct {
	Vault         armkeyvault.DeletedVault
	ResourceGroup string
}

//  =================== keyvault ==================

//index:microsoft_keyvault_managedhsms
//getfilter:name=description.ManagedHsm.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultManagedHardwareSecurityModuleDescription struct {
	ManagedHsm                  armkeyvault.ManagedHsm
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== secret ==================

//index:microsoft_keyvault_vaults_secrets
//getfilter:name=description.SecretItem.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultSecretDescription struct {
	SecretItem    armkeyvault.Secret
	Vault         armkeyvault.Vault
	TurboData     map[string]interface{}
	ResourceGroup string
}

//  =================== kusto ==================

//index:microsoft_kusto_clusters
//getfilter:name=description.Cluster.name
//getfilter:resource_group=description.ResourceGroup
type KustoClusterDescription struct {
	Cluster       armkusto.Cluster
	ResourceGroup string
}

//  =================== insights ==================

//index:microsoft_insights_activitylogalerts
//getfilter:name=description.ActivityLogAlertResource.name
//getfilter:resource_group=description.ResourceGroup
type LogAlertDescription struct {
	ActivityLogAlertResource armmonitor.ActivityLogAlertResource
	ResourceGroup            string
}

//  =================== insights ==================

//index:microsoft_insights_logprofiles
//getfilter:name=description.LogProfileResource.name
//getfilter:resource_group=description.ResourceGroup
type LogProfileDescription struct {
	LogProfileResource armmonitor.LogProfileResource
	ResourceGroup      string
}

//  =================== logic ==================

//index:microsoft_logic_workflows
//getfilter:name=description.Workflow.name
//getfilter:resource_group=description.ResourceGroup
type LogicAppWorkflowDescription struct {
	Workflow                    armlogic.Workflow
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//index:microsoft_logic_integrationaccounts
type LogicIntegrationAccountsDescription struct {
	ResourceGroup string
	Account       armlogic.IntegrationAccount
}

//  =================== machinelearningservices ==================

//index:microsoft_machinelearning_workspaces
//getfilter:name=description.Workspace.name
//getfilter:resource_group=description.ResourceGroup
type MachineLearningWorkspaceDescription struct {
	Workspace                   armmachinelearning.Workspace
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== mariadb ==================

//index:microsoft_dbformariadb_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type MariadbServerDescription struct {
	Server        armmariadb.Server
	ResourceGroup string
}

//index:microsoft_dbformariadb_databases
type MariadbDatabaseDescription struct {
	Server        armmariadb.Server
	Database      armmariadb.Database
	ResourceGroup string
}

//  =================== mysql ==================

//index:microsoft_dbformysql_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type MysqlServerDescription struct {
	Server                armmysql.Server
	Configurations        []*armmysql.Configuration
	ServerKeys            []*armmysql.ServerKey
	SecurityAlertPolicies []*armmysql.ServerSecurityAlertPolicy
	VnetRules             []*armmysql.VirtualNetworkRule
	ResourceGroup         string
}

//index:microsoft_dbformysql_flexibleservers
type MysqlFlexibleserverDescription struct {
	Server        armmysqlflexibleservers.Server
	ResourceGroup string
}

//  =================== network ==================

//index:microsoft_network_networksecuritygroups
//getfilter:name=description.SecurityGroup.name
//getfilter:resource_group=description.ResourceGroup
type NetworkSecurityGroupDescription struct {
	SecurityGroup               armnetwork.SecurityGroup
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//index:microsoft_network_networkwatchers
//getfilter:name=description.Watcher.name
//getfilter:resource_group=description.ResourceGroup
type NetworkWatcherDescription struct {
	Watcher       armnetwork.Watcher
	ResourceGroup string
}

//  =================== search ==================

//index:microsoft_search_searchservices
//getfilter:name=description.Service.name
//getfilter:resource_group=description.ResourceGroup
type SearchServiceDescription struct {
	Service                     armsearch.Service
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== servicefabric ==================

//index:microsoft_servicefabric_clusters
//getfilter:name=description.Cluster.name
//getfilter:resource_group=description.ResourceGroup
type ServiceFabricClusterDescription struct {
	Cluster       armservicefabric.Cluster
	ResourceGroup string
}

//  =================== servicebus ==================

//index:microsoft_servicebus_namespaces
//getfilter:name=description.SBNamespace.name
//getfilter:resource_group=description.ResourceGroup
type ServicebusNamespaceDescription struct {
	SBNamespace                 armservicebus.SBNamespace
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	NetworkRuleSet              []*armservicebus.NetworkRuleSet
	PrivateEndpointConnections  []*armservicebus.PrivateEndpointConnection
	AuthorizationRules          []*armservicebus.SBAuthorizationRule
	ResourceGroup               string
}

//  =================== signalr ==================

//index:microsoft_signalrservice_signalr
//getfilter:name=description.ResourceType.name
//getfilter:resource_group=description.ResourceGroup
type SignalrServiceDescription struct {
	ResourceInfo                armsignalr.ResourceInfo
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== appplatform ==================

//index:microsoft_appplatform_spring
//getfilter:name=description.ServiceResource.name
//getfilter:resource_group=description.ResourceGroup
type SpringCloudServiceDescription struct {
	DiagnosticSettingsResource *[]armmonitor.DiagnosticSettingsResource
	ResourceGroup              string
	Site                       *armspringappdiscovery.SpringbootsitesModel
	App                        armspringappdiscovery.SpringbootappsModel
}

//  =================== streamanalytics ==================

//index:microsoft_streamanalytics_streamingjobs
//getfilter:name=description.StreamingJob.name
//getfilter:resource_group=description.ResourceGroup
type StreamAnalyticsJobDescription struct {
	StreamingJob                armstreamanalytics.StreamingJob
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	ResourceGroup               string
}

//index:microsoft_streamanalytics_cluster
type StreamAnalyticsClusterDescription struct {
	ResourceGroup string
	StreamingJob  armstreamanalytics.Cluster
}

//index:microsoft_virtualmachineimages_imagetemplates
type VirtualMachineImagesImageTemplatesDescription struct {
	ResourceGroup string
	ImageTemplate armvirtualmachineimagebuilder.ImageTemplate
}

//  =================== operationalinsights ==================

//index:microsoft_operationalinsights_workspaces
type OperationalInsightsWorkspacesDescription struct {
	ResourceGroup string
	Workspace     armoperationalinsights.Workspace
}

//  =================== timeseriesinsight ==================

//index:microsoft_timeseriesinsight_environments
type TimeSeriesInsightsEnvironmentsDescription struct {
	ResourceGroup string
	Environment   *armtimeseriesinsights.EnvironmentResource
}

//  =================== synapse ==================

//index:microsoft_synapse_workspaces
//getfilter:name=description.Workspace.name
//getfilter:resource_group=description.ResourceGroup
type SynapseWorkspaceDescription struct {
	Workspace                      armsynapse.Workspace
	ServerVulnerabilityAssessments []*armsynapse.ServerVulnerabilityAssessment
	DiagnosticSettingsResources    []*armmonitor.DiagnosticSettingsResource
	ResourceGroup                  string
}

//index:microsoft_synapse_workspacesbigdatapools
type SynapseWorkspaceBigdatapoolsDescription struct {
	Workspace     armsynapse.Workspace
	BigDataPool   armsynapse.BigDataPoolResourceInfo
	ResourceGroup string
}

//index:microsoft_synapse_workspacessqlpools
type SynapseWorkspaceSqlpoolsDescription struct {
	Workspace     armsynapse.Workspace
	SqlPool       armsynapse.SQLPool
	ResourceGroup string
}

//  =================== sub ==================

//index:microsoft_resources_subscriptions_locations
//getfilter:name=description.Location.name
//getfilter:resource_group=description.ResourceGroup
type LocationDescription struct {
	Location      armsubscription.Location
	ResourceGroup string
}

//  =================== analysis ==================

//index:microsoft_analysisservice_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type AnalysisServiceServerDescription struct {
	ResourceGroup string
	Server        armanalysisservices.Server
}

//  =================== postgresql ==================

//index:microsoft_dbforpostgresql_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type PostgresqlServerDescription struct {
	Server                       armpostgresql.Server
	ServerAdministratorResources []*armpostgresql.ServerAdministratorResource
	Configurations               []*armpostgresql.Configuration
	ServerKeys                   []*armpostgresql.ServerKey
	FirewallRules                []*armpostgresql.FirewallRule
	ServerSecurityAlertPolicies  []*armpostgresql.ServerSecurityAlertPolicy
	ResourceGroup                string
}

//index:microsoft_dbforpostgresql_flexibleservers
type PostgresqlFlexibleServerDescription struct {
	ResourceGroup        string
	Server               armpostgresqlflexibleservers.Server
	ServerConfigurations []*armpostgresqlflexibleservers.Configuration
}

//  =================== storagesync ==================

//index:microsoft_storagesync_storagesyncservices
//getfilter:name=description.Service.name
//getfilter:resource_group=description.ResourceGroup
type StorageSyncDescription struct {
	Service       armstoragesync.Service
	ResourceGroup string
}

//  =================== sql ==================

//index:microsoft_sql_managedinstances
//getfilter:name=description.ManagedInstance.name
//getfilter:resource_group=description.ResourceGroup
type MssqlManagedInstanceDescription struct {
	ManagedInstance                         armsql.ManagedInstance
	ManagedInstanceVulnerabilityAssessments []*armsql.ManagedInstanceVulnerabilityAssessment
	ManagedDatabaseSecurityAlertPolicies    []*armsql.ManagedServerSecurityAlertPolicy
	ManagedInstanceEncryptionProtectors     []*armsql.ManagedInstanceEncryptionProtector
	ResourceGroup                           string
}

//index:microsoft_sql_managedinstancesdatabases
type MssqlManagedInstanceDatabasesDescription struct {
	ManagedInstance armsql.ManagedInstance
	Database        armsql.ManagedDatabase
	ResourceGroup   string
}

//index:microsoft_sql_servers_databases
//getfilter:name=description.Database.name
//getfilter:resource_group=description.ResourceGroup
type SqlDatabaseDescription struct {
	Database                           armsql.Database
	LongTermRetentionPolicy            armsql.LongTermRetentionPolicy
	TransparentDataEncryption          []*armsql.LogicalDatabaseTransparentDataEncryption
	DatabaseVulnerabilityAssessments   []*armsql.DatabaseVulnerabilityAssessment
	VulnerabilityAssessmentScanRecords []*armsql.VulnerabilityAssessmentScanRecord
	Advisors                           []*armsql.Advisor
	AuditPolicies                      []*armsql.DatabaseBlobAuditingPolicy
	ResourceGroup                      string
}

type SqlInstancePoolDescription struct {
	InstancePool  armsql.InstancePool
	ResourceGroup string
}

//  =================== sqlv3 ==================

//index:microsoft_sql_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type SqlServerDescription struct {
	Server                         armsql.Server
	ServerBlobAuditingPolicies     []*armsql.ServerBlobAuditingPolicy
	ServerSecurityAlertPolicies    []*armsql.ServerSecurityAlertPolicy
	ServerAzureADAdministrators    []*armsql.ServerAzureADAdministrator
	ServerVulnerabilityAssessments []*armsql.ServerVulnerabilityAssessment
	FirewallRules                  []*armsql.FirewallRule
	EncryptionProtectors           []*armsql.EncryptionProtector
	PrivateEndpointConnections     []*armsql.PrivateEndpointConnection
	VirtualNetworkRules            []*armsql.VirtualNetworkRule
	FailoverGroups                 []*armsql.FailoverGroup
	AutomaticTuning                armsql.ServerAutomaticTuning
	ResourceGroup                  string
}

//index:microsoft_sql_serversjobagent
type SqlServerJobAgentDescription struct {
	ResourceGroup string
	Server        armsql.Server
	JobAgent      armsql.JobAgent
}

//index:microsoft_sql_virtualclusters
type SqlVirtualClustersDescription struct {
	ResourceGroup   string
	VirtualClusters armsql.VirtualCluster
}

//index:microsoft_sql_elasticpools
//getfilter:name=description.Pool.Name
//getfilter:server_name=description.ServerName
//getfilter:resource_group=description.ResourceGroup
type SqlServerElasticPoolDescription struct {
	Pool          armsql.ElasticPool
	TotalDTU      int32
	ServerName    string
	ResourceGroup string
}

//index:microsoft_sql_virtualmachines
//getfilter:name=description.VirtualMachine.Name
//getfilter:resource_group=description.ResourceGroup
type SqlServerVirtualMachineDescription struct {
	VirtualMachine armsqlvirtualmachine.SQLVirtualMachine
	ResourceGroup  string
}

//index:microsoft_sql_virtualmachinegroups
type SqlServerVirtualMachineGroupDescription struct {
	Group         armsqlvirtualmachine.Group
	ResourceGroup string
}

//index:microsoft_sql_flexibleservers
//getfilter:name=description.FlexibleServer.Name
//getfilter:resource_group=description.ResourceGroup
type SqlServerFlexibleServerDescription struct {
	FlexibleServer armmysqlflexibleservers.Server
	ResourceGroup  string
}

//  =================== storage ==================

//index:microsoft_classicstorage_storageaccounts
//getfilter:name=description.Account.name
//getfilter:resource_group=description.ResourceGroup
type StorageAccountDescription struct {
	Account                     armstorage.Account
	ManagementPolicy            *armstorage.ManagementPolicy
	BlobServiceProperties       *armstorage.BlobServiceProperties
	Logging                     *accounts.Logging
	StorageServiceProperties    *queues.StorageServiceProperties
	FileServiceProperties       *armstorage.FileServiceProperties
	DiagnosticSettingsResources []*armmonitor.DiagnosticSettingsResource
	EncryptionScopes            []*armstorage.EncryptionScope
	TableProperties             aztables.ServiceProperties
	AccessKeys                  []map[string]interface{}
	ResourceGroup               string
}

//  =================== recoveryservice ==================

//index:microsoft_recoveryservices_vault
//getfilter:name=description.Vault.Name
//getfilter:resource_group=description.ResourceGroup
type RecoveryServicesVaultDescription struct {
	Vault                      armrecoveryservices.Vault
	DiagnosticSettingsResource []*armmonitor.DiagnosticSettingsResource
	ResourceGroup              string
}

//index:microsoft_recoveryservices_vault
//getfilter:name=description.Vault.Name
//getfilter:resource_group=description.ResourceGroup
type RecoveryServicesBackupJobDescription struct {
	Job struct {
		Name     *string
		ID       *string
		Type     *string
		ETag     *string
		Tags     map[string]*string
		Location *string
	}
	VaultName     string
	Properties    map[string]interface{}
	ResourceGroup string
}

//index:microsoft_recoveryservices_policy
//getfilter:name=description.Policy.Name
//getfilter:resource_group=description.ResourceGroup
type RecoveryServicesBackupPolicyDescription struct {
	Policy struct {
		Name     *string
		ID       *string
		Type     *string
		ETag     *string
		Tags     map[string]*string
		Location *string
	}
	Properties    map[string]interface{}
	VaultName     string
	ResourceGroup string
}

//index:microsoft_recoveryservices_item
//getfilter:name=description.Item.Name
//getfilter:resource_group=description.ResourceGroup
type RecoveryServicesBackupItemDescription struct {
	Item struct {
		Name     *string
		ID       *string
		Type     *string
		ETag     *string
		Tags     map[string]*string
		Location *string
	}
	Properties    map[string]interface{}
	VaultName     string
	ResourceGroup string
}

//  =================== kubernetes ==================

//index:microsoft_hybridkubernetes_connectedcluster
//getfilter:name=description.ConnectedCluster.Name
//getfilter:resource_group=description.ResourceGroup
type HybridKubernetesConnectedClusterDescription struct {
	ConnectedCluster           armhybridkubernetes.ConnectedCluster
	ConnectedClusterExtensions []*armkubernetesconfiguration.Extension
	ResourceGroup              string
}

//  =================== Cost ==================

type CostManagementQueryRow struct {
	UsageDate      int     `json:"UsageDate"`
	Cost           float64 `json:"Cost"`
	Currency       string  `json:"Currency"`
	ServiceName    *string `json:"ServiceName,omitempty"`
	PublisherType  *string `json:"PublisherType,omitempty"`
	SubscriptionID *string `json:"SubscriptionId,omitempty"`
}

//index:microsoft_costmanagement_costbyresourcetype
type CostManagementCostByResourceTypeDescription struct {
	CostManagementCostByResourceType CostManagementQueryRow
	CostDateMillis                   int64
}

//index:microsoft_costmanagement_costbysubscription
type CostManagementCostBySubscriptionDescription struct {
	CostManagementCostBySubscription CostManagementQueryRow
}

// =================== LB (loadbalancer) ==================

//index:microsoft_network_loadbalancers
//getfilter:name=description.LoadBalancer.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerDescription struct {
	LoadBalancer      armnetwork.LoadBalancer
	DiagnosticSetting []*armmonitor.DiagnosticSettingsResource
	ResourceGroup     string
}

//index:microsoft_lb_backendaddresspools
//getfilter:load_balancer_name=description.LoadBalancer.Name
//getfilter:name=description.Pool.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerBackendAddressPoolDescription struct {
	LoadBalancer  armnetwork.LoadBalancer
	Pool          armnetwork.BackendAddressPool
	ResourceGroup string
}

//index:microsoft_lb_natrules
//getfilter:load_balancer_name=description.LoadBalancerName
//getfilter:name=description.Rule.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerNatRuleDescription struct {
	Rule             armnetwork.InboundNatRule
	LoadBalancerName string
	ResourceGroup    string
}

//index:microsoft_lb_outboundrules
//getfilter:load_balancer_name=description.LoadBalancerName
//getfilter:name=description.Rule.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerOutboundRuleDescription struct {
	Rule             armnetwork.OutboundRule
	LoadBalancerName string
	ResourceGroup    string
}

//index:microsoft_lb_probes
//getfilter:load_balancer_name=description.LoadBalancerName
//getfilter:name=description.Probe.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerProbeDescription struct {
	Probe            armnetwork.Probe
	LoadBalancerName string
	ResourceGroup    string
}

//index:microsoft_lb_rules
//getfilter:load_balancer_name=description.LoadBalancerName
//getfilter:name=description.Rule.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerRuleDescription struct {
	Rule             armnetwork.LoadBalancingRule
	LoadBalancerName string
	ResourceGroup    string
}

// =================== Management ==================

//index:microsoft_management_groups
//getfilter:name=description.Group.Name
type ManagementGroupDescription struct {
	Group armmanagementgroups.ManagementGroup
}

//index:microsoft_management_locks
//getfilter:name=description.Lock.Name
//getfilter:resource_group=description.ResourceGroup
type ManagementLockDescription struct {
	Lock          armlocks.ManagementLockObject
	ResourceGroup string
}

// =================== Resources ==================

//index:microsoft_resources_providers
//getfilter:namespace=description.Provider.Namespace
type ResourceProviderDescription struct {
	Provider armresources.Provider
}

//index:microsoft_resources_resourcegroups
//getfilter:name=description.Group.Name
type ResourceGroupDescription struct {
	Group armresources.ResourceGroup
}

type GenericResourceDescription struct {
	GenericResource armresources.GenericResourceExpanded
	ResourceGroup   string
}

// =================== BotService ==================

type BotServiceBotDescription struct {
	Bot           armbotservice.Bot
	ResourceGroup string
}

// =================== NetApp ==================

type NetAppAccountDescription struct {
	Account       armnetapp.Account
	ResourceGroup string
}

type NetAppCapacityPoolDescription struct {
	CapacityPool  armnetapp.CapacityPool
	ResourceGroup string
}

// =================== Dashboard ==================

type DashboardGrafanaDescription struct {
	ResourceGroup string
	Grafana       armdashboard.ManagedGrafana
}

// =================== DesktopVirtualization ==================

type DesktopVirtualizationHostPoolDescription struct {
	HostPool      armdesktopvirtualization.HostPool
	ResourceGroup string
}

type DesktopVirtualizationWorkspaceDescription struct {
	ResourceGroup string
	Workspace     armdesktopvirtualization.Workspace
}

// =================== DevTestLab ==================

type DevTestLabLabDescription struct {
	Lab           armdevtestlabs.Lab
	ResourceGroup string
}

// =================== Purview ==================

type PurviewAccountDescription struct {
	Account       armpurview.Account
	ResourceGroup string
}

// =================== PowerBI ==================

type PowerBIDedicatedCapacityDescription struct {
	Capacity      armpowerbidedicated.DedicatedCapacity
	ResourceGroup string
}

// =================== applicationInsights =================

type ApplicationInsightsComponentDescription struct {
	Component     armapplicationinsights.Component
	ResourceGroup string
}

// =================== Alert Management =================

type AlertManagementDescription struct {
	Alert         armalertsmanagement.Alert
	ResourceGroup string
}

// =================== Lighthouse =================

type LighthouseDefinitionDescription struct {
	LighthouseDefinition armmanagedservices.RegistrationDefinition
	Scope                string
	ResourceGroup        string
}

type LighthouseAssignmentDescription struct {
	LighthouseAssignment armmanagedservices.RegistrationAssignment
	Scope                string
	ResourceGroup        string
}

// =================== Maintenance Configuration =================

type MaintenanceConfigurationDescription struct {
	MaintenanceConfiguration armmaintenance.Configuration
	ResourceGroup            string
}

// =================== Monitor Insights =================

type MonitorLogProfileDescription struct {
	LogProfile    armmonitor.LogProfileResource
	ResourceGroup string
}
