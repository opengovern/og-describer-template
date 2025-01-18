package describers

import (
	"context"
	"strings"

	

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func listResourceGroups(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string) ([]armresources.ResourceGroup, error) {
	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewResourceGroupsClient()
	pager := client.NewListPager(nil)
	var values []armresources.ResourceGroup
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			values = append(values, *v)
		}
	}
	return values, nil
}

func ResourceProvider(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewProvidersClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, provider := range page.Value {
			resource := GetResourceProvider(ctx, provider)
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

func GetResourceProvider(ctx context.Context, provider *armresources.Provider) *models.Resource {
	resource := models.Resource{
		ID:       *provider.ID,
		Location: "global",
		Description: model.ResourceProviderDescription{
			Provider: *provider,
		},
	}

	return &resource
}

func ResourceGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewResourceGroupsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, group := range page.Value {
			resource := GetResourceGroup(ctx, group)
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

func GetResourceGroup(ctx context.Context, group *armresources.ResourceGroup) *models.Resource {
	resource := models.Resource{
		ID:       *group.ID,
		Name:     *group.Name,
		Location: *group.Location,
		Description: model.ResourceGroupDescription{
			Group: *group,
		},
	}

	return &resource
}

func Resources(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {

	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClient()

	var resources []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, genericResource := range page.Value {
			resource := GetResource(ctx, genericResource)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, *resource)
			}
		}
	}
	return resources, nil

}

func GetResource(ctx context.Context, genericResource *armresources.GenericResourceExpanded) *models.Resource {

	resourceGroupName := strings.Split(string(*genericResource.ID), "/")[4]

	resource := models.Resource{
		ID:       *genericResource.ID,
		Name:     *genericResource.Name,
		Location: *genericResource.Location,
		Description: model.GenericResourceDescription{
			GenericResource: *genericResource,
			ResourceGroup:   resourceGroupName,
		},
	}
	return &resource
}
