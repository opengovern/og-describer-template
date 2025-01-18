package describers

import (
	"context"
	"strings"

	

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
	
)

func ServiceBusQueue(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armservicebus.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewQueuesClient()

	var values []models.Resource
	for _, rg := range rgs {
		resources, err := ListResourceGroupServiceBusQueue(ctx, cred, subscription, client, rg)
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

func ListResourceGroupServiceBusQueue(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, client *armservicebus.QueuesClient, rg armresources.ResourceGroup) ([]models.Resource, error) {
	ns, err := serviceBusNamespace(ctx, cred, subscription, *rg.Name)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, n := range ns {
		resources, err := ListNamespaceServiceBusQueues(ctx, client, rg, n)
		if err != nil {
			return nil, err
		}
		values = append(values, resources...)
	}
	return values, nil
}

func ListNamespaceServiceBusQueues(ctx context.Context, client *armservicebus.QueuesClient, rg armresources.ResourceGroup, n *armservicebus.SBNamespace) ([]models.Resource, error) {
	pager := client.NewListByNamespacePager(*rg.Name, *n.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetServiceBusQueue(ctx, v)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetServiceBusQueue(ctx context.Context, v *armservicebus.SBQueue) *models.Resource {
	resource := models.Resource{
		ID:          *v.ID,
		Name:        *v.Name,
		Location:    "global",
		Description: v,
	}
	return &resource
}

func ServiceBusTopic(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armservicebus.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewTopicsClient()

	var values []models.Resource
	for _, rg := range rgs {
		resources, err := ListResourceGroupServiceBusTopic(ctx, cred, subscription, client, rg)
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

func ListResourceGroupServiceBusTopic(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, client *armservicebus.TopicsClient, rg armresources.ResourceGroup) ([]models.Resource, error) {
	ns, err := serviceBusNamespace(ctx, cred, subscription, *rg.Name)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, n := range ns {
		resources, err := ListNamespaceServiceBusTopics(ctx, client, rg, n)
		if err != nil {
			return nil, err
		}
		values = append(values, resources...)
	}
	return values, nil
}

func ListNamespaceServiceBusTopics(ctx context.Context, client *armservicebus.TopicsClient, rg armresources.ResourceGroup, n *armservicebus.SBNamespace) ([]models.Resource, error) {
	pager := client.NewListByNamespacePager(*rg.Name, *n.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetServiceBusTopic(ctx, v)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetServiceBusTopic(ctx context.Context, v *armservicebus.SBTopic) *models.Resource {
	resource := models.Resource{
		ID:          *v.ID,
		Name:        *v.Name,
		Location:    "global",
		Description: v,
	}
	return &resource
}

func serviceBusNamespace(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, resourceGroup string) ([]*armservicebus.SBNamespace, error) {
	clientFactory, err := armservicebus.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewNamespacesClient()

	var values []*armservicebus.SBNamespace
	pager := client.NewListByResourceGroupPager(resourceGroup, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		values = append(values, page.Value...)
	}
	return values, nil
}

func ServicebusNamespace(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armservicebus.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	servicebusClient := clientFactory.NewPrivateEndpointConnectionsClient()
	namespaceClient := clientFactory.NewNamespacesClient()
	client := clientFactory.NewNamespacesClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, namespace := range page.Value {
			resource, err := GetServicebusNamespace(ctx, namespaceClient, servicebusClient, diagnosticClient, namespace)
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

func GetServicebusNamespace(ctx context.Context, namespaceClient *armservicebus.NamespacesClient, servicebusClient *armservicebus.PrivateEndpointConnectionsClient, diagnosticClient *armmonitor.DiagnosticSettingsClient, namespace *armservicebus.SBNamespace) (*models.Resource, error) {
	resourceGroup := strings.Split(*namespace.ID, "/")[4]

	var insightsListOp []*armmonitor.DiagnosticSettingsResource
	pager1 := diagnosticClient.NewListPager(*namespace.ID, nil)
	for pager1.More() {
		page1, err := pager1.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		insightsListOp = append(insightsListOp, page1.Value...)
	}

	var servicebusGetNetworkRuleSetOp []*armservicebus.NetworkRuleSet
	pager2 := namespaceClient.NewListNetworkRuleSetsPager(resourceGroup, *namespace.Name, nil)
	for pager2.More() {
		page2, err := pager2.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		servicebusGetNetworkRuleSetOp = append(servicebusGetNetworkRuleSetOp, page2.Value...)
	}

	var servicebusListOp []*armservicebus.PrivateEndpointConnection
	pager3 := servicebusClient.NewListPager(resourceGroup, *namespace.Name, nil)
	for pager3.More() {
		page, err := pager3.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		servicebusListOp = append(servicebusListOp, page.Value...)
	}

	var servicebusAuthorizationRules []*armservicebus.SBAuthorizationRule
	pager4 := namespaceClient.NewListAuthorizationRulesPager(resourceGroup, *namespace.Name, nil)
	for pager4.More() {
		page, err := pager4.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		servicebusAuthorizationRules = append(servicebusAuthorizationRules, page.Value...)
	}

	resource := models.Resource{
		ID:       *namespace.ID,
		Name:     *namespace.Name,
		Location: *namespace.Location,
		Description: model.ServicebusNamespaceDescription{
			SBNamespace:                 *namespace,
			DiagnosticSettingsResources: insightsListOp,
			NetworkRuleSet:              servicebusGetNetworkRuleSetOp,
			PrivateEndpointConnections:  servicebusListOp,
			AuthorizationRules:          servicebusAuthorizationRules,
			ResourceGroup:               resourceGroup,
		},
	}
	return &resource, nil
}
