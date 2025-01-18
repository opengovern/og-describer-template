package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v2"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func NetAppAccount(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetapp.NewAccountsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getNetAppAccount(ctx, v)
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

func getNetAppAccount(ctx context.Context, v *armnetapp.Account) *models.Resource {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.NetAppAccountDescription{
			Account:       *v,
			ResourceGroup: resourceGroupName,
		},
	}

	return &resource
}

func NetAppCapacityPool(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armnetapp.NewAccountsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	poolsClient, err := armnetapp.NewPoolsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := listNetAppAccountPools(ctx, poolsClient, v)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				for _, resource := range resources {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				}
			} else {
				values = append(values, resources...)
			}
		}
	}
	return values, nil
}

func listNetAppAccountPools(ctx context.Context, poolsClient *armnetapp.PoolsClient, v *armnetapp.Account) ([]models.Resource, error) {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]

	pager := poolsClient.NewListPager(resourceGroupName, *v.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, pool := range page.Value {
			resource := getNetAppCapacityPool(ctx, v, pool)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getNetAppCapacityPool(ctx context.Context, v *armnetapp.Account, pool *armnetapp.CapacityPool) *models.Resource {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.NetAppCapacityPoolDescription{
			CapacityPool:  *pool,
			ResourceGroup: resourceGroupName,
		},
	}
	return &resource
}
