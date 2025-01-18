package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func HealthcareService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armhealthcareapis.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	privateEndpointClient := clientFactory.NewPrivateEndpointConnectionsClient()
	client := clientFactory.NewServicesClient()

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
		for _, v := range page.Value {
			resource, err := getHealthcareService(ctx, privateEndpointClient, diagnosticClient, v)
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

func getHealthcareService(ctx context.Context, privateEndpointClient *armhealthcareapis.PrivateEndpointConnectionsClient, diagnosticClient *armmonitor.DiagnosticSettingsClient, v *armhealthcareapis.ServicesDescription) (*models.Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	var opValue []*armmonitor.DiagnosticSettingsResource
	var opService []*armhealthcareapis.PrivateEndpointConnectionDescription
	if v.ID != nil {
		resourceId := v.ID

		pager := diagnosticClient.NewListPager(*resourceId, nil)
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			opValue = append(opValue, page.Value...)
		}

		if v.Name != nil {
			resourceGroup := strings.Split(*v.ID, "/")[4]
			resourceName := v.Name

			// SDK does not support pagination yet

			pager := privateEndpointClient.NewListByServicePager(resourceGroup, *resourceName, nil)
			page, err := pager.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			opService = append(opService, page.Value...)

		}
	} else {
		return nil, nil
	}

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.HealthcareServiceDescription{
			ServicesDescription:         *v,
			DiagnosticSettingsResources: opValue,
			PrivateEndpointConnections:  opService,
			ResourceGroup:               resourceGroup,
		},
	}

	return &resource, nil
}
