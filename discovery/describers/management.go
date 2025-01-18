package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func ManagementGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, group := range page.Value {
			resource, err := getManagementGroup(ctx, client, group)
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

func getManagementGroup(ctx context.Context, client *armmanagementgroups.Client, group *armmanagementgroups.ManagementGroupInfo) (*models.Resource, error) {
	info, err := client.Get(ctx, *group.Name, nil)
	if err != nil {
		return nil, err
	}

	resource := &models.Resource{
		ID:   *info.ID,
		Name: *info.Name,
		Description: model.ManagementGroupDescription{
			Group: info.ManagementGroup,
		},
	}
	return resource, nil
}

func ManagementLock(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armlocks.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewManagementLocksClient()

	pager := client.NewListAtSubscriptionLevelPager(nil)
	var values []models.Resource
	for pager.More() {
		pagem, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, lockObject := range pagem.Value {
			resource := getManagementLock(ctx, lockObject)
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

func getManagementLock(ctx context.Context, lockObject *armlocks.ManagementLockObject) *models.Resource {
	resourceGroup := strings.Split(*lockObject.ID, "/")[4]
	resource := models.Resource{
		ID:       *lockObject.ID,
		Name:     *lockObject.Name,
		Location: "global",
		Description: model.ManagementLockDescription{
			Lock:          *lockObject,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
