package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func LoadBalancer(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewLoadBalancersClient(subscription, cred, nil)
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
		for _, loadBalancer := range page.Value {
			resource, err := getLoadBalancer(ctx, diagnosticClient, loadBalancer)
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

func getLoadBalancer(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, loadBalancer *armnetwork.LoadBalancer) (*models.Resource, error) {
	resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

	// Get diagnostic settings
	var diagnosticSettings []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*loadBalancer.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		diagnosticSettings = append(diagnosticSettings, page.Value...)
	}

	resource := models.Resource{
		ID:       *loadBalancer.ID,
		Name:     *loadBalancer.Name,
		Location: *loadBalancer.Location,
		Description: model.LoadBalancerDescription{
			ResourceGroup:     resourceGroup,
			DiagnosticSetting: diagnosticSettings,
			LoadBalancer:      *loadBalancer,
		},
	}
	return &resource, nil
}

func LoadBalancerBackendAddressPool(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewLoadBalancersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	addressClient, err := armnetwork.NewLoadBalancerBackendAddressPoolsClient(subscription, cred, nil)
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
		for _, loadBalancer := range page.Value {
			resources, err := listLoadBalancerBackendAddressPools(ctx, addressClient, loadBalancer)
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

func listLoadBalancerBackendAddressPools(ctx context.Context, addressClient *armnetwork.LoadBalancerBackendAddressPoolsClient, loadBalancer *armnetwork.LoadBalancer) ([]models.Resource, error) {
	resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

	pager := addressClient.NewListPager(resourceGroup, *loadBalancer.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, pool := range page.Value {
			resource := getLoadBalancerBackendAddressPools(ctx, loadBalancer, pool)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getLoadBalancerBackendAddressPools(ctx context.Context, loadBalancer *armnetwork.LoadBalancer, pool *armnetwork.BackendAddressPool) *models.Resource {
	location := "global"
	if pool.Properties.Location != nil {
		location = *pool.Properties.Location
	}
	resourceGroup := strings.Split(*pool.ID, "/")[4]
	resource := models.Resource{
		ID:       *pool.ID,
		Location: location,
		Description: model.LoadBalancerBackendAddressPoolDescription{
			ResourceGroup: resourceGroup,
			LoadBalancer:  *loadBalancer,
			Pool:          *pool,
		},
	}

	return &resource
}

func LoadBalancerNatRule(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewLoadBalancersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	natRulesClient, err := armnetwork.NewInboundNatRulesClient(subscription, cred, nil)
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
		for _, loadBalancer := range page.Value {
			resources, err := listLoadBalancerNatRules(ctx, natRulesClient, loadBalancer)
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

func listLoadBalancerNatRules(ctx context.Context, natRulesClient *armnetwork.InboundNatRulesClient, loadBalancer *armnetwork.LoadBalancer) ([]models.Resource, error) {
	resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

	pager := natRulesClient.NewListPager(resourceGroup, *loadBalancer.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, natRule := range page.Value {
			resource := getLoadBalancerNatRule(ctx, loadBalancer, natRule)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getLoadBalancerNatRule(ctx context.Context, loadBalancer *armnetwork.LoadBalancer, natRule *armnetwork.InboundNatRule) *models.Resource {
	resourceGroup := strings.Split(*natRule.ID, "/")[4]
	resource := models.Resource{
		ID:       *natRule.ID,
		Name:     *natRule.Name,
		Location: *loadBalancer.Location,
		Description: model.LoadBalancerNatRuleDescription{
			ResourceGroup:    resourceGroup,
			LoadBalancerName: *loadBalancer.Name,
			Rule:             *natRule,
		},
	}

	return &resource
}

func LoadBalancerOutboundRule(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewLoadBalancersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	outboundRulesClient, err := armnetwork.NewLoadBalancerOutboundRulesClient(subscription, cred, nil)
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
		for _, loadBalancer := range page.Value {
			resources, err := listLoadBalancerOutboundRules(ctx, outboundRulesClient, loadBalancer)
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

func listLoadBalancerOutboundRules(ctx context.Context, outboundRulesClient *armnetwork.LoadBalancerOutboundRulesClient, loadBalancer *armnetwork.LoadBalancer) ([]models.Resource, error) {
	resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

	pager := outboundRulesClient.NewListPager(resourceGroup, *loadBalancer.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, outboundRule := range page.Value {
			resource := getLoadBalancerOutboundRule(ctx, loadBalancer, outboundRule)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getLoadBalancerOutboundRule(ctx context.Context, loadBalancer *armnetwork.LoadBalancer, outboundRule *armnetwork.OutboundRule) *models.Resource {
	resourceGroup := strings.Split(*outboundRule.ID, "/")[4]
	resource := models.Resource{
		ID:       *outboundRule.ID,
		Name:     *outboundRule.Name,
		Location: *loadBalancer.Location,
		Description: model.LoadBalancerOutboundRuleDescription{
			ResourceGroup:    resourceGroup,
			LoadBalancerName: *loadBalancer.Name,
			Rule:             *outboundRule,
		},
	}

	return &resource
}

func LoadBalancerProbe(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewLoadBalancersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	probesClient, err := armnetwork.NewLoadBalancerProbesClient(subscription, cred, nil)
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
		for _, loadBalancer := range page.Value {
			resources, err := listLoadBalancerProbes(ctx, probesClient, loadBalancer)
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

func listLoadBalancerProbes(ctx context.Context, probesClient *armnetwork.LoadBalancerProbesClient, loadBalancer *armnetwork.LoadBalancer) ([]models.Resource, error) {
	resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

	pager := probesClient.NewListPager(resourceGroup, *loadBalancer.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, probe := range page.Value {
			resource := getLoadBalancerProbe(ctx, loadBalancer, probe)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getLoadBalancerProbe(ctx context.Context, loadBalancer *armnetwork.LoadBalancer, probe *armnetwork.Probe) *models.Resource {
	resourceGroup := strings.Split(*probe.ID, "/")[4]
	resource := models.Resource{
		ID:       *probe.ID,
		Name:     *probe.Name,
		Location: *loadBalancer.Location,
		Description: model.LoadBalancerProbeDescription{
			ResourceGroup:    resourceGroup,
			LoadBalancerName: *loadBalancer.Name,
			Probe:            *probe,
		},
	}

	return &resource
}

func LoadBalancerRule(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetwork.NewLoadBalancersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	rulesClient, err := armnetwork.NewLoadBalancerLoadBalancingRulesClient(subscription, cred, nil)
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
		for _, loadBalancer := range page.Value {
			resources, err := listLoadBalancerRules(ctx, rulesClient, loadBalancer)
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

func listLoadBalancerRules(ctx context.Context, rulesClient *armnetwork.LoadBalancerLoadBalancingRulesClient, loadBalancer *armnetwork.LoadBalancer) ([]models.Resource, error) {
	resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

	pager := rulesClient.NewListPager(resourceGroup, *loadBalancer.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, rule := range page.Value {
			resource := getLoadBalancerRule(ctx, loadBalancer, rule)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getLoadBalancerRule(ctx context.Context, loadBalancer *armnetwork.LoadBalancer, rule *armnetwork.LoadBalancingRule) *models.Resource {
	resourceGroup := strings.Split(*rule.ID, "/")[4]
	resource := models.Resource{
		ID:       *rule.ID,
		Name:     *rule.Name,
		Location: *loadBalancer.Location,
		Description: model.LoadBalancerRuleDescription{
			ResourceGroup:    resourceGroup,
			LoadBalancerName: *loadBalancer.Name,
			Rule:             *rule,
		},
	}

	return &resource
}
