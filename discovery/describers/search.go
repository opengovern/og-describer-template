package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func SearchService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsearch.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServicesClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	var values []models.Resource
	pager := client.NewListBySubscriptionPager(nil, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := GetSearchService(ctx, diagnosticClient, v)
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

func GetSearchService(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, v *armsearch.Service) (*models.Resource, error) {
	id := v.ID

	var searchListOp []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*id, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		searchListOp = append(searchListOp, page.Value...)
	}

	resourceGroupName := strings.Split(string(*v.ID), "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.SearchServiceDescription{
			Service:                     *v,
			DiagnosticSettingsResources: searchListOp,
			ResourceGroup:               resourceGroupName,
		},
	}
	return &resource, nil
}
