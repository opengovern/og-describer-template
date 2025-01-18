package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func KustoCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armkusto.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClustersClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, kusto := range page.Value {
			resource := getKustoCluster(ctx, kusto)
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

func getKustoCluster(ctx context.Context, kusto *armkusto.Cluster) *models.Resource {
	resourceGroup := strings.Split(*kusto.ID, "/")[4]

	resource := models.Resource{
		ID:       *kusto.ID,
		Name:     *kusto.Name,
		Location: *kusto.Location,
		Description: model.KustoClusterDescription{
			Cluster:       *kusto,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
