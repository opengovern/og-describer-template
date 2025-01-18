package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func EventhubNamespace(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	clientFactory, err := armeventhub.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewNamespacesClient()
	eventhubClient := clientFactory.NewPrivateEndpointConnectionsClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, namespace := range page.Value {
			resource, err := getEventHubNamespace(ctx, diagnosticClient, client, eventhubClient, namespace)
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

func getEventHubNamespace(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, client *armeventhub.NamespacesClient, eventhubClient *armeventhub.PrivateEndpointConnectionsClient, namespace *armeventhub.EHNamespace) (*models.Resource, error) {
	resourceGroupName := strings.Split(string(*namespace.ID), "/")[4]
	var insightsListOp []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*namespace.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		insightsListOp = append(insightsListOp, page.Value...)
	}

	eventhubGetNetworkRuleSetOp, err := client.GetNetworkRuleSet(ctx, resourceGroupName, *namespace.Name, nil)
	if err != nil {
		return nil, err
	}

	pager2 := eventhubClient.NewListPager(resourceGroupName, *namespace.Name, nil)
	if err != nil {
		return nil, err
	}
	var eventhubListOp []*armeventhub.PrivateEndpointConnection
	for pager2.More() {
		page, err := pager2.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		eventhubListOp = append(eventhubListOp, page.Value...)
	}

	resource := models.Resource{
		ID:       *namespace.ID,
		Name:     *namespace.Name,
		Location: *namespace.Location,
		Description: model.EventhubNamespaceDescription{
			EHNamespace:                 *namespace,
			DiagnosticSettingsResources: insightsListOp,
			NetworkRuleSet:              eventhubGetNetworkRuleSetOp.NetworkRuleSet,
			PrivateEndpointConnection:   eventhubListOp,
			ResourceGroup:               resourceGroupName,
		},
	}
	return &resource, nil
}

func EventhubNamespaceEventhub(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armeventhub.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewNamespacesClient()
	eventhubClient := clientFactory.NewEventHubsClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, namespace := range page.Value {
			resourceGroupName := strings.Split(string(*namespace.ID), "/")[4]

			pager2 := eventhubClient.NewListByNamespacePager(resourceGroupName, *namespace.Name, nil)
			for pager2.More() {
				page, err := pager2.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, eh := range page.Value {
					resource := getEventhubNamespaceEventhub(ctx, namespace, eh)
					if stream != nil {
						if err := (*stream)(*resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, *resource)
					}
				}
			}
		}
	}
	return values, nil
}

func getEventhubNamespaceEventhub(ctx context.Context, namespace *armeventhub.EHNamespace, eh *armeventhub.Eventhub) *models.Resource {
	resourceGroupName := strings.Split(string(*namespace.ID), "/")[4]
	return &models.Resource{
		ID:       *namespace.ID,
		Name:     *namespace.Name,
		Location: *namespace.Location,
		Description: model.EventhubNamespaceEventhubDescription{
			EHNamespace:   *namespace,
			EventHub:      *eh,
			ResourceGroup: resourceGroupName,
		},
	}
}
