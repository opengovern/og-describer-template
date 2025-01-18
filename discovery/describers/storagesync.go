package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func StorageSync(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstoragesync.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServicesClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetStorageSync(ctx, v)
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

func GetStorageSync(ctx context.Context, storage *armstoragesync.Service) *models.Resource {
	resourceGroup := strings.Split(*storage.ID, "/")[4]

	resource := models.Resource{
		ID:       *storage.ID,
		Name:     *storage.Name,
		Location: *storage.Location,
		Description: model.StorageSyncDescription{
			Service:       *storage,
			ResourceGroup: resourceGroup,
		}}
	return &resource
}
