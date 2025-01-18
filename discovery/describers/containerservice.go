package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func KubernetesCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armcontainerservice.NewManagedClustersClient(subscription, cred, nil)
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
			resource := getKubernatesCluster(ctx, v)
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

func getKubernatesCluster(ctx context.Context, v *armcontainerservice.ManagedCluster) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.KubernetesClusterDescription{
			ManagedCluster: *v,
			ResourceGroup:  resourceGroup,
		},
	}
	return &resource
}

func KubernetesServiceVersion(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	subClient, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		return nil, err
	}

	client, err := armcontainerservice.NewManagedClustersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := subClient.NewListLocationsPager(subscription, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, location := range page.Value {
			services, err := listLocationKubernatesServices(ctx, client, location)
			if err != nil {
				return nil, err
			}
			values = append(values, services...)
		}
	}
	return values, nil
}

func listLocationKubernatesServices(ctx context.Context, client *armcontainerservice.ManagedClustersClient, location *armsubscriptions.Location) ([]models.Resource, error) {
	kubernetesVersions, err := client.ListKubernetesVersions(ctx, *location.ID, nil)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, v := range kubernetesVersions.Values {
		resource := getKubernatesService(ctx, location, v)
		values = append(values, *resource)
	}
	return values, nil
}

func getKubernatesService(ctx context.Context, location *armsubscriptions.Location, v *armcontainerservice.KubernetesVersion) *models.Resource {
	resource := models.Resource{
		ID:       *v.Version,
		Name:     *v.Version,
		Type:     *v.Version,
		Location: *location.ID,
		Description: model.KubernetesServiceVersionDescription{
			Version: *v,
		},
	}
	return &resource
}
