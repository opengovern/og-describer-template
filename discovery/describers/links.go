package describers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func ResourceLink(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armlinks.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewResourceLinksClient()

	pager := client.NewListAtSubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getResourceLink(ctx, v)
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

func getResourceLink(ctx context.Context, v *armlinks.ResourceLink) *models.Resource {
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: model.ResourceLinkDescription{
			ResourceLink: *v,
		},
	}
	return &resource
}
