package describers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func Tenant(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewTenantsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetTenand(ctx, v)
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

func GetTenand(ctx context.Context, v *armsubscription.TenantIDDescription) *models.Resource {
	name := ""

	resource := models.Resource{
		ID:       *v.ID,
		Name:     name,
		Location: "global",
		Description: model.TenantDescription{
			TenantIDDescription: *v, // TODO has much less values
		},
	}

	return &resource
}

func Subscription(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}
	resourceClientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewSubscriptionsClient()
	tagsClient := resourceClientFactory.NewTagsClient()

	op, err := client.Get(ctx, subscription, nil)
	if err != nil {
		return nil, err
	}

	tags := map[string][]string{}
	pager := tagsClient.NewListPager(&armresources.TagsClientListOptions{})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, tag := range page.Value {
			if tag.TagName == nil {
				continue
			}

			var values []string
			for _, v := range tag.Values {
				if v.TagValue != nil {
					values = append(values, *v.TagValue)
				}
			}

			tags[*tag.TagName] = values
		}
	}

	var values []models.Resource
	resource := models.Resource{
		ID:       *op.ID,
		Name:     *op.DisplayName,
		Location: "global",
		Description: model.SubscriptionDescription{
			Subscription: op.Subscription,
			Tags:         tags,
		},
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}
