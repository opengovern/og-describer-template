package describers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func PolicyAssignment(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armpolicy.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAssignmentsClient()

	resourceClient, err := armresources.NewClient(subscription, cred, nil)

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetPolicyAssignment(ctx, resourceClient, v)
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

func GetPolicyAssignment(ctx context.Context, resourceClient *armresources.Client, v *armpolicy.Assignment) *models.Resource {
	location := "global"
	if v.Location != nil {
		location = *v.Location
	}

	res, err := resourceClient.GetByID(ctx, *v.ID, "2021-04-01", nil)

	if err == nil {
		location = *res.Location
	}

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: model.PolicyAssignmentDescription{
			Assignment: *v,
			Resource:   res.GenericResource,
		},
	}

	return &resource
}
