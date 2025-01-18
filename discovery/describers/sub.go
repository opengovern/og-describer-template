package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func Location(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSubscriptionsClient()

	var values []models.Resource
	pager := client.NewListLocationsPager(subscription, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetLocation(ctx, v)
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

func GetLocation(ctx context.Context, location *armsubscription.Location) *models.Resource {
	resourceGroup := strings.Split(*location.ID, "/")[4]

	resource := models.Resource{
		ID:       *location.ID,
		Name:     *location.Name,
		Location: "global",
		Description: model.LocationDescription{
			Location:      *location,
			ResourceGroup: resourceGroup,
		},
	}

	return &resource
}
