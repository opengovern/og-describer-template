package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func DesktopVirtualizationWorkspaces(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armdesktopvirtualization.NewWorkspacesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewListBySubscriptionPager(&armdesktopvirtualization.WorkspacesClientListBySubscriptionOptions{})
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resourceGroupName := strings.Split(string(*v.ID), "/")[4]
			resource := &models.Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: *v.Location,
				Description: model.DesktopVirtualizationWorkspaceDescription{
					Workspace:     *v,
					ResourceGroup: resourceGroupName,
				},
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

func DesktopVirtualizationHostPool(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armdesktopvirtualization.NewHostPoolsClient(subscription, cred, nil)
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
			resource := getDesktopVirtualizationHostPool(ctx, v)
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

func getDesktopVirtualizationHostPool(ctx context.Context, v *armdesktopvirtualization.HostPool) *models.Resource {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.DesktopVirtualizationHostPoolDescription{
			HostPool:      *v,
			ResourceGroup: resourceGroupName,
		},
	}
	return &resource
}
