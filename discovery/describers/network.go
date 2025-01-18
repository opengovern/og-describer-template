package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func NetworkInterface(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewInterfacesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getNetworkInterface(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getNetworkInterface(ctx context.Context, v *armnetwork.Interface) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.NetworkInterfaceDescription{
			Interface:     *v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func NetworkWatcherFlowLog(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	logsClient, err := armnetwork.NewFlowLogsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	watcherClient, err := armnetwork.NewWatchersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := watcherClient.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, watcher := range page.Value {
			resources, err := listWatcherFlowLogs(ctx, logsClient, watcher)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func listWatcherFlowLogs(ctx context.Context, logsClient *armnetwork.FlowLogsClient, watcher *armnetwork.Watcher) ([]models.Resource, error) {
	resourceGroupID := strings.Split(*watcher.ID, "/")[4]

	pager := logsClient.NewListPager(resourceGroupID, *watcher.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getWatcherFlowLog(ctx, watcher, v)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getWatcherFlowLog(ctx context.Context, watcher *armnetwork.Watcher, v *armnetwork.FlowLog) *models.Resource {
	resourceGroupID := strings.Split(*watcher.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.NetworkWatcherFlowLogDescription{
			NetworkWatcherName: *watcher.Name,
			FlowLog:            *v,
			ResourceGroup:      resourceGroupID,
		},
	}
	return &resource
}

func Subnet(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	subnetsClient, err := armnetwork.NewSubnetsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	virtualnetworkClient, err := armnetwork.NewVirtualNetworksClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := virtualnetworkClient.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, virtualnetwork := range page.Value {
			resources, err := listVirtualNetworkSubnets(ctx, subnetsClient, virtualnetwork)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func listVirtualNetworkSubnets(ctx context.Context, subnetsClient *armnetwork.SubnetsClient, virtualnetwork *armnetwork.VirtualNetwork) ([]models.Resource, error) {
	resourceGroupID := strings.Split(*virtualnetwork.ID, "/")[4]

	pager := subnetsClient.NewListPager(resourceGroupID, *virtualnetwork.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getVirtualNetworkSubnet(ctx, virtualnetwork, v)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getVirtualNetworkSubnet(ctx context.Context, virtualnetwork *armnetwork.VirtualNetwork, v *armnetwork.Subnet) *models.Resource {
	resourceGroupID := strings.Split(*virtualnetwork.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: model.SubnetDescription{
			VirtualNetworkName: *virtualnetwork.Name,
			Subnet:             *v,
			ResourceGroup:      resourceGroupID,
		},
	}
	return &resource
}

func VirtualNetwork(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewVirtualNetworksClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getVirtualNetwork(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getVirtualNetwork(ctx context.Context, v *armnetwork.VirtualNetwork) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.VirtualNetworkDescription{
			VirtualNetwork: *v,
			ResourceGroup:  resourceGroup,
		},
	}

	return &resource
}

func ApplicationGateway(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewApplicationGatewaysClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, gateway := range page.Value {
			resource, err := getApplicationGateway(ctx, diagnosticClient, gateway)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getApplicationGateway(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, gateway *armnetwork.ApplicationGateway) (*models.Resource, error) {
	resourceGroup := strings.Split(*gateway.ID, "/")[4]

	var networkListOp []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*gateway.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		networkListOp = append(networkListOp, page.Value...)
	}

	resource := models.Resource{
		ID:       *gateway.ID,
		Name:     *gateway.Name,
		Location: *gateway.Location,
		Description: model.ApplicationGatewayDescription{
			ApplicationGateway:          *gateway,
			DiagnosticSettingsResources: networkListOp,
			ResourceGroup:               resourceGroup,
		},
	}
	return &resource, nil
}

func NetworkSecurityGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewSecurityGroupsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, networkSecurityGroup := range page.Value {
			resource, err := getNetworkSecurityGroup(ctx, diagnosticClient, networkSecurityGroup)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getNetworkSecurityGroup(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, networkSecurityGroup *armnetwork.SecurityGroup) (*models.Resource, error) {
	resourceGroup := strings.Split(*networkSecurityGroup.ID, "/")[4]

	id := *networkSecurityGroup.ID
	pager := diagnosticClient.NewListPager(id, nil)
	var networkListOp []*armmonitor.DiagnosticSettingsResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "ResourceNotFound") || strings.Contains(err.Error(), "SubscriptionNotRegistered") {
				// ignore
			} else {
				return nil, err
			}
		}
		networkListOp = append(networkListOp, page.Value...)
	}

	resource := models.Resource{
		ID:       *networkSecurityGroup.ID,
		Name:     *networkSecurityGroup.Name,
		Location: *networkSecurityGroup.Location,
		Description: model.NetworkSecurityGroupDescription{
			SecurityGroup:               *networkSecurityGroup,
			DiagnosticSettingsResources: networkListOp,
			ResourceGroup:               resourceGroup,
		},
	}

	return &resource, nil
}

func NetworkWatcher(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewWatchersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, networkWatcher := range page.Value {
			resource := getNetworkWatcher(ctx, networkWatcher)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getNetworkWatcher(ctx context.Context, networkWatcher *armnetwork.Watcher) *models.Resource {
	resourceGroup := strings.Split(*networkWatcher.ID, "/")[4]

	resource := models.Resource{
		ID:       *networkWatcher.ID,
		Name:     *networkWatcher.Name,
		Location: *networkWatcher.Location,
		Description: model.NetworkWatcherDescription{
			Watcher:       *networkWatcher,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}

func RouteTables(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewRouteTablesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, routeTable := range page.Value {
			resource := getRouteTable(ctx, routeTable)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getRouteTable(ctx context.Context, routeTable *armnetwork.RouteTable) *models.Resource {
	resourceGroup := strings.Split(*routeTable.ID, "/")[4]

	resource := models.Resource{
		ID:       *routeTable.ID,
		Name:     *routeTable.Name,
		Location: *routeTable.Location,
		Description: model.RouteTablesDescription{
			ResourceGroup: resourceGroup,
			RouteTable:    *routeTable,
		},
	}
	return &resource
}

func NetworkApplicationSecurityGroups(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewApplicationSecurityGroupsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, applicationSecurityGroup := range page.Value {
			resource := getApplicationSecurityGroup(ctx, applicationSecurityGroup)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getApplicationSecurityGroup(ctx context.Context, applicationSecurityGroup *armnetwork.ApplicationSecurityGroup) *models.Resource {
	resourceGroup := strings.Split(*applicationSecurityGroup.ID, "/")[4]

	resource := models.Resource{
		ID:       *applicationSecurityGroup.ID,
		Name:     *applicationSecurityGroup.Name,
		Location: *applicationSecurityGroup.Location,
		Description: model.NetworkApplicationSecurityGroupsDescription{
			ApplicationSecurityGroup: *applicationSecurityGroup,
			ResourceGroup:            resourceGroup,
		},
	}

	return &resource
}

func NetworkAzureFirewall(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewAzureFirewallsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, azureFirewall := range page.Value {
			resource := getAzureFirewall(ctx, azureFirewall)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getAzureFirewall(ctx context.Context, azureFirewall *armnetwork.AzureFirewall) *models.Resource {
	resourceGroup := strings.Split(*azureFirewall.ID, "/")[4]

	resource := models.Resource{
		ID:       *azureFirewall.ID,
		Name:     *azureFirewall.Name,
		Location: *azureFirewall.Location,
		Description: model.NetworkAzureFirewallDescription{
			AzureFirewall: *azureFirewall,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}

func ExpressRouteCircuit(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewExpressRouteCircuitsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, expressRouteCircuit := range page.Value {
			resource := getExpressRouteCircuit(ctx, expressRouteCircuit)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getExpressRouteCircuit(ctx context.Context, expressRouteCircuit *armnetwork.ExpressRouteCircuit) *models.Resource {
	resourceGroup := strings.Split(*expressRouteCircuit.ID, "/")[4]

	resource := models.Resource{
		ID:       *expressRouteCircuit.ID,
		Name:     *expressRouteCircuit.Name,
		Location: *expressRouteCircuit.Location,
		Description: model.ExpressRouteCircuitDescription{
			ExpressRouteCircuit: *expressRouteCircuit,
			ResourceGroup:       resourceGroup,
		},
	}
	return &resource
}

func VirtualNetworkGateway(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewVirtualNetworkGatewaysClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, rg := range rgs {
		resources, err := getResourceGroupVirtualNetworkGateway(ctx, client, rg)
		if err != nil {
			return nil, err
		}
		for _, resource := range resources {
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func getResourceGroupVirtualNetworkGateway(ctx context.Context, client *armnetwork.VirtualNetworkGatewaysClient, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	pager := client.NewListPager(*resourceGroup.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, virtualNetworkGateway := range page.Value {
			resource, err := getVirtualNetworkGateway(ctx, client, virtualNetworkGateway)
			if err != nil {
				return nil, err
			}
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getVirtualNetworkGateway(ctx context.Context, client *armnetwork.VirtualNetworkGatewaysClient, virtualNetworkGateway *armnetwork.VirtualNetworkGateway) (*models.Resource, error) {
	resourceGroup := strings.Split(*virtualNetworkGateway.ID, "/")[4]

	var gatewayConnections []*armnetwork.VirtualNetworkGatewayConnectionListEntity
	pager := client.NewListConnectionsPager(resourceGroup, *virtualNetworkGateway.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		gatewayConnections = append(gatewayConnections, page.Value...)
	}
	var virtualNetwork string
	if virtualNetworkGateway.Properties != nil {
		if len(virtualNetworkGateway.Properties.IPConfigurations) > 0 {
			for _, config := range virtualNetworkGateway.Properties.IPConfigurations {
				if config != nil && config.Properties != nil && config.Properties.Subnet != nil && config.Properties.Subnet.ID != nil {
					split := strings.Split(*config.Properties.Subnet.ID, "/subnets")
					if len(split) > 0 {
						virtualNetwork = split[0]
					}
				}
			}
		}
	}

	resource := models.Resource{
		ID:       *virtualNetworkGateway.ID,
		Name:     *virtualNetworkGateway.Name,
		Location: *virtualNetworkGateway.Location,
		Description: model.VirtualNetworkGatewayDescription{
			ResourceGroup:                   resourceGroup,
			VirtualNetworkGateway:           *virtualNetworkGateway,
			VirtualNetworkGatewayConnection: gatewayConnections,
			VirtualNetwork:                  virtualNetwork,
		},
	}

	return &resource, nil
}

func FirewallPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewFirewallPoliciesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, firewallPolicy := range page.Value {
			resource := getFirewallPolicy(ctx, firewallPolicy)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getFirewallPolicy(ctx context.Context, firewallPolicy *armnetwork.FirewallPolicy) *models.Resource {
	resourceGroup := strings.Split(*firewallPolicy.ID, "/")[4]

	resource := models.Resource{
		ID:       *firewallPolicy.ID,
		Name:     *firewallPolicy.Name,
		Location: *firewallPolicy.Location,
		Description: model.FirewallPolicyDescription{
			ResourceGroup:  resourceGroup,
			FirewallPolicy: *firewallPolicy,
		},
	}

	return &resource
}

func LocalNetworkGateway(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewLocalNetworkGatewaysClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, rg := range rgs {
		resources, err := ListResourceGroupLocalNetworkGateways(ctx, client, rg)
		if err != nil {
			return nil, err
		}
		for _, resource := range resources {
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ListResourceGroupLocalNetworkGateways(ctx context.Context, client *armnetwork.LocalNetworkGatewaysClient, rg armresources.ResourceGroup) ([]models.Resource, error) {
	var values []models.Resource
	pager := client.NewListPager(*rg.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, localNetworkGateway := range page.Value {
			resource := getLocalNetworkGateway(ctx, localNetworkGateway)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getLocalNetworkGateway(ctx context.Context, localNetworkGateway *armnetwork.LocalNetworkGateway) *models.Resource {
	resourceGroup := strings.Split(*localNetworkGateway.ID, "/")[4]

	resource := models.Resource{
		ID:       *localNetworkGateway.ID,
		Name:     *localNetworkGateway.Name,
		Location: *localNetworkGateway.Location,
		Description: model.LocalNetworkGatewayDescription{
			ResourceGroup:       resourceGroup,
			LocalNetworkGateway: *localNetworkGateway,
		},
	}

	return &resource
}

func NatGateway(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewNatGatewaysClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, rg := range rgs {
		resources, err := ListResourceGroupNatGateways(ctx, client, rg)
		if err != nil {
			return nil, err
		}
		for _, resource := range resources {
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ListResourceGroupNatGateways(ctx context.Context, client *armnetwork.NatGatewaysClient, rg armresources.ResourceGroup) ([]models.Resource, error) {
	var values []models.Resource
	pager := client.NewListPager(*rg.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, natGateway := range page.Value {
			resource := getNatGateway(ctx, natGateway)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getNatGateway(ctx context.Context, natGateway *armnetwork.NatGateway) *models.Resource {
	resourceGroup := strings.Split(*natGateway.ID, "/")[4]

	resource := models.Resource{
		ID:       *natGateway.ID,
		Name:     *natGateway.Name,
		Location: *natGateway.Location,
		Description: model.NatGatewayDescription{
			ResourceGroup: resourceGroup,
			NatGateway:    *natGateway,
		},
	}

	return &resource
}

func PrivateLinkService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewPrivateLinkServicesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, rg := range rgs {
		resources, err := ListResourceGroupPrivateLinkServices(ctx, client, rg)
		if err != nil {
			return nil, err
		}
		for _, resource := range resources {
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ListResourceGroupPrivateLinkServices(ctx context.Context, client *armnetwork.PrivateLinkServicesClient, rg armresources.ResourceGroup) ([]models.Resource, error) {
	var values []models.Resource
	pager := client.NewListPager(*rg.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, privateLinkService := range page.Value {
			resource := getPrivateLinkService(ctx, privateLinkService)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getPrivateLinkService(ctx context.Context, privateLinkService *armnetwork.PrivateLinkService) *models.Resource {
	resourceGroup := strings.Split(*privateLinkService.ID, "/")[4]

	resource := models.Resource{
		ID:       *privateLinkService.ID,
		Name:     *privateLinkService.Name,
		Location: *privateLinkService.Location,
		Description: model.PrivateLinkServiceDescription{
			ResourceGroup:      resourceGroup,
			PrivateLinkService: *privateLinkService,
		},
	}

	return &resource
}

func RouteFilter(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewRouteFiltersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, routeFilter := range page.Value {
			resource := getRouteFilter(ctx, routeFilter)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getRouteFilter(ctx context.Context, routeFilter *armnetwork.RouteFilter) *models.Resource {
	resourceGroup := strings.Split(*routeFilter.ID, "/")[4]

	resource := models.Resource{
		ID:       *routeFilter.ID,
		Name:     *routeFilter.Name,
		Location: *routeFilter.Location,
		Description: model.RouteFilterDescription{
			ResourceGroup: resourceGroup,
			RouteFilter:   *routeFilter,
		},
	}

	return &resource
}

func VpnGateway(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewVPNGatewaysClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vpnGateway := range page.Value {
			resource := getVpnGateway(ctx, vpnGateway)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getVpnGateway(ctx context.Context, vpnGateway *armnetwork.VPNGateway) *models.Resource {
	resourceGroup := strings.Split(*vpnGateway.ID, "/")[4]

	resource := models.Resource{
		ID:       *vpnGateway.ID,
		Name:     *vpnGateway.Name,
		Location: *vpnGateway.Location,
		Description: model.VpnGatewayDescription{
			ResourceGroup: resourceGroup,
			VpnGateway:    *vpnGateway,
		},
	}

	return &resource
}

func NetworkVpnGatewaysVpnConnections(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewVPNGatewaysClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	connClient, err := armnetwork.NewVPNConnectionsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vpnGateway := range page.Value {
			resources, err := ListNetworkVpnGatewayVpnConnections(ctx, connClient, vpnGateway)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func ListNetworkVpnGatewayVpnConnections(ctx context.Context, connClient *armnetwork.VPNConnectionsClient, vpnGateway *armnetwork.VPNGateway) ([]models.Resource, error) {
	resourceGroup := strings.Split(*vpnGateway.ID, "/")[4]

	var values []models.Resource
	pager := connClient.NewListByVPNGatewayPager(resourceGroup, *vpnGateway.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vpnConn := range page.Value {
			resource := getNetworkVpnGatewaysVpnConnections(ctx, vpnGateway, vpnConn)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getNetworkVpnGatewaysVpnConnections(ctx context.Context, vpnGateway *armnetwork.VPNGateway, vpnConn *armnetwork.VPNConnection) *models.Resource {
	resourceGroup := strings.Split(*vpnConn.ID, "/")[4]

	resource := models.Resource{
		ID:       *vpnConn.ID,
		Name:     *vpnConn.Name,
		Location: *vpnGateway.Location,
		Description: model.VpnGatewayVpnConnectionDescription{
			VpnConnection: *vpnConn,
			VpnGateway:    *vpnGateway,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}

func NetworkVpnGatewaysVpnSites(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewVPNSitesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vpnSite := range page.Value {
			resource := getNetworkVpnGatewaysVpnSites(ctx, vpnSite)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getNetworkVpnGatewaysVpnSites(ctx context.Context, v *armnetwork.VPNSite) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.VpnSiteDescription{
			ResourceGroup: resourceGroup,
			VpnSite:       *v,
		},
	}

	return &resource
}

func PublicIPAddress(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewPublicIPAddressesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, resourceGroup := range resourceGroups {
		resources, err := ListResourceGroupPublicIPAddresses(ctx, client, resourceGroup)
		if err != nil {
			return nil, err
		}
		for _, resource := range resources {
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ListResourceGroupPublicIPAddresses(ctx context.Context, client *armnetwork.PublicIPAddressesClient, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	pager := client.NewListPager(*resourceGroup.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, publicIPAddress := range page.Value {
			resource := getPublicIPAddress(ctx, resourceGroup, publicIPAddress)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getPublicIPAddress(ctx context.Context, resourceGroup armresources.ResourceGroup, publicIPAddress *armnetwork.PublicIPAddress) *models.Resource {
	resource := models.Resource{
		ID:       *publicIPAddress.ID,
		Name:     *publicIPAddress.Name,
		Location: *publicIPAddress.Location,
		Description: model.PublicIPAddressDescription{
			ResourceGroup:   *resourceGroup.Name,
			PublicIPAddress: *publicIPAddress,
		},
	}
	return &resource
}

func PublicIPPrefix(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewPublicIPPrefixesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, resourceGroup := range resourceGroups {
		resources, err := ListResourceGroupPublicIPPrefixes(ctx, client, resourceGroup)
		if err != nil {
			return nil, err
		}
		for _, resource := range resources {
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ListResourceGroupPublicIPPrefixes(ctx context.Context, client *armnetwork.PublicIPPrefixesClient, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	pager := client.NewListPager(*resourceGroup.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, publicIPPrefix := range page.Value {
			resource := getPublicIPPrefix(ctx, resourceGroup, publicIPPrefix)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getPublicIPPrefix(ctx context.Context, resourceGroup armresources.ResourceGroup, publicIPPrefix *armnetwork.PublicIPPrefix) *models.Resource {
	resource := models.Resource{
		ID:       *publicIPPrefix.ID,
		Name:     *publicIPPrefix.Name,
		Location: *publicIPPrefix.Location,
		Description: model.PublicIPPrefixDescription{
			ResourceGroup:  *resourceGroup.Name,
			PublicIPPrefix: *publicIPPrefix,
		},
	}

	return &resource
}

func DNSZones(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armdns.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewZonesClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, dnsZone := range page.Value {
			resource := GetDNSZone(ctx, dnsZone)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetDNSZone(ctx context.Context, dnsZone *armdns.Zone) *models.Resource {
	resourceGroup := strings.Split(*dnsZone.ID, "/")[4]
	resource := models.Resource{
		ID:       *dnsZone.ID,
		Name:     *dnsZone.Name,
		Location: *dnsZone.Location,
		Description: model.DNSZonesDescription{
			DNSZone:       *dnsZone,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}

func DNSResolvers(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armdnsresolver.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDNSResolversClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, dnsResolver := range page.Value {
			resource := GetDNSResolver(ctx, dnsResolver)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}
func GetDNSResolver(ctx context.Context, dnsResolver *armdnsresolver.DNSResolver) *models.Resource {
	resourceGroup := strings.Split(*dnsResolver.ID, "/")[4]
	resource := models.Resource{
		ID:       *dnsResolver.ID,
		Name:     *dnsResolver.Name,
		Location: *dnsResolver.Location,
		Description: model.DNSResolverDescription{
			DNSResolver:   *dnsResolver,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}

func TrafficManagerProfile(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armtrafficmanager.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewProfilesClient()

	pager := client.NewListBySubscriptionPager(&armtrafficmanager.ProfilesClientListBySubscriptionOptions{})
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, profile := range page.Value {
			resource := GetTrafficManagerProfile(ctx, profile)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetTrafficManagerProfile(ctx context.Context, profile *armtrafficmanager.Profile) *models.Resource {
	resourceGroup := strings.Split(*profile.ID, "/")[4]
	resource := models.Resource{
		ID:       *profile.ID,
		Name:     *profile.Name,
		Location: *profile.Location,
		Description: model.TrafficManagerProfileDescription{
			Profile:       *profile,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}

func PrivateDnsZones(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armprivatedns.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewPrivateZonesClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, privateZone := range page.Value {
			resource := GetPrivateDnsZone(ctx, privateZone)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetPrivateDnsZone(ctx context.Context, privateZone *armprivatedns.PrivateZone) *models.Resource {
	resourceGroup := strings.Split(*privateZone.ID, "/")[4]
	resource := models.Resource{
		ID:       *privateZone.ID,
		Name:     *privateZone.Name,
		Location: *privateZone.Location,
		Description: model.PrivateDNSZonesDescription{
			PrivateZone:   *privateZone,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}

func PrivateEndpoints(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewPrivateEndpointsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, resourceGroup := range resourceGroups {
		resources, err := ListResourceGroupPrivateEndpoints(ctx, client, resourceGroup)
		if err != nil {
			return nil, err
		}
		for _, resource := range resources {
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ListResourceGroupPrivateEndpoints(ctx context.Context, client *armnetwork.PrivateEndpointsClient, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	pager := client.NewListPager(*resourceGroup.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, privateEndpoint := range page.Value {
			resource := GetPrivateEndpoint(ctx, resourceGroup, privateEndpoint)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetPrivateEndpoint(ctx context.Context, resourceGroup armresources.ResourceGroup, v *armnetwork.PrivateEndpoint) *models.Resource {
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.PrivateEndpointDescription{
			PrivateEndpoint: *v,
			ResourceGroup:   *resourceGroup.Name,
		},
	}

	return &resource
}

func NetworkBastionHosts(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewBastionHostsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, bastionHost := range page.Value {
			resource := GetBastionHost(ctx, bastionHost)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetBastionHost(ctx context.Context, v *armnetwork.BastionHost) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.BastionHostsDescription{
			BastianHost:   *v,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}

func NetworkConnections(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewVirtualNetworkGatewayConnectionsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, resourceGroup := range resourceGroups {
		resources, err := ListResourceGroupNetworkCOnnections(ctx, client, resourceGroup)
		if err != nil {
			return nil, err
		}
		for _, resource := range resources {
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func ListResourceGroupNetworkCOnnections(ctx context.Context, client *armnetwork.VirtualNetworkGatewayConnectionsClient, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	pager := client.NewListPager(*resourceGroup.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, connection := range page.Value {
			resource := GetNetworkConnection(ctx, resourceGroup, connection)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetNetworkConnection(ctx context.Context, resourceGroup armresources.ResourceGroup, v *armnetwork.VirtualNetworkGatewayConnection) *models.Resource {
	resourceGroupName := strings.Split(*v.ID, "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ConnectionDescription{
			Connection:    *v,
			ResourceGroup: resourceGroupName,
		},
	}

	return &resource
}

func NetworkVirtualHubs(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewVirtualHubsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetNetworkVirtualHub(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetNetworkVirtualHub(ctx context.Context, v *armnetwork.VirtualHub) *models.Resource {
	resourceGroupName := strings.Split(*v.ID, "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.VirtualHubsDescription{
			VirtualHub:    *v,
			ResourceGroup: resourceGroupName,
		},
	}

	return &resource
}

func NetworkVirtualWans(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewVirtualWansClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetNetworkVirtualWan(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetNetworkVirtualWan(ctx context.Context, v *armnetwork.VirtualWAN) *models.Resource {
	resourceGroupName := strings.Split(*v.ID, "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.VirtualWansDescription{
			VirtualWan:    *v,
			ResourceGroup: resourceGroupName,
		},
	}
	return &resource
}

func NetworkDDoSProtectionPlan(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewDdosProtectionPlansClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetNetworkDDoSProtectionPlan(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetNetworkDDoSProtectionPlan(ctx context.Context, v *armnetwork.DdosProtectionPlan) *models.Resource {
	resourceGroupName := strings.Split(*v.ID, "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.NetworkDDoSProtectionPlanDescription{
			DDoSProtectionPlan: *v,
			ResourceGroup:      resourceGroupName,
		},
	}
	return &resource
}
