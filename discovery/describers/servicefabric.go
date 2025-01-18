package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func ServiceFabricCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armservicefabric.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	clusterClient := clientFactory.NewClustersClient()

	var values []models.Resource
	list, err := clusterClient.List(ctx, nil)
	if err != nil {
		return nil, err
	}
	for _, cluster := range list.Value {
		resource := GetServiceFabricCluster(ctx, cluster)
		if stream != nil {
			if err := (*stream)(*resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetServiceFabricCluster(ctx context.Context, cluster *armservicefabric.Cluster) *models.Resource {
	resourceGroup := strings.Split(*cluster.ID, "/")[4]

	resource := models.Resource{
		ID:          *cluster.ID,
		Name:        *cluster.Name,
		Location:    *cluster.Location,
		Description: model.ServiceFabricClusterDescription{Cluster: *cluster, ResourceGroup: resourceGroup},
	}
	return &resource
}
