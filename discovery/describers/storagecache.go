package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v2"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func HpcCache(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armstoragecache.NewCachesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := client.NewListPager(nil)
	if pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetHpcCache(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, err
}

func GetHpcCache(ctx context.Context, v *armstoragecache.Cache) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.HpcCacheDescription{
			Cache:         *v,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}
