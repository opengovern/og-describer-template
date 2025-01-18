package describers

import (
	"context"
	"strings"

	

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func APIManagement(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armapimanagement.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServiceClient()

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
		for _, apiManagement := range page.Value {
			resource, err := getAPIMangement(ctx, diagnosticClient, apiManagement)
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

func getAPIMangement(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, apiManagement *armapimanagement.ServiceResource) (*models.Resource, error) {
	resourceGroup := strings.Split(*apiManagement.ID, "/")[4]
	accountListOpTemp := diagnosticClient.NewListPager(*apiManagement.ID, nil)
	var op []armmonitor.DiagnosticSettingsResource
	for accountListOpTemp.More() {
		accountOpPage, err := accountListOpTemp.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, accountOp := range accountOpPage.Value {
			op = append(op, *accountOp)
		}
	}
	resource := models.Resource{
		ID:       *apiManagement.ID,
		Name:     *apiManagement.Name,
		Location: *apiManagement.Location,
		Description: model.APIManagementDescription{
			APIManagement:               *apiManagement,
			DiagnosticSettingsResources: &op,
			ResourceGroup:               resourceGroup,
		},
	}
	return &resource, nil
}

func APIManagementBackend(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {

	clientFactory, err := armapimanagement.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewServiceClient()
	backendClient := clientFactory.NewBackendClient()

	pager := client.NewListPager(nil)
	var resources []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, apiManagementService := range page.Value {
			resources, err := listAPIMangementBackends(ctx, backendClient, apiManagementService)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					resources = append(resources, resource)
				}
			}
		}
	}
	return resources, nil

}

func listAPIMangementBackends(ctx context.Context, backendClient *armapimanagement.BackendClient, apiManagementService *armapimanagement.ServiceResource) ([]models.Resource, error) {

	resourceGroup := strings.Split(*apiManagementService.ID, "/")[4]
	pager := backendClient.NewListByServicePager(resourceGroup, *apiManagementService.Name, nil)

	var resources []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, backend := range page.Value {
			resource := GetAPIManagementBackend(ctx, apiManagementService, backend)
			resources = append(resources, *resource)
		}
	}
	return resources, nil

}

func GetAPIManagementBackend(ctx context.Context, service *armapimanagement.ServiceResource, backend *armapimanagement.BackendContract) *models.Resource {

	resourceGroup := strings.Split(*backend.ID, "/")[4]

	resource := models.Resource{
		ID:   *backend.ID,
		Name: *backend.Name,
		Description: model.APIManagementBackendDescription{
			APIManagementBackend: *backend,
			ServiceName:          *service.Name,
			ResourceGroup:        resourceGroup,
		},
	}
	return &resource
}
