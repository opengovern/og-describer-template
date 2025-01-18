package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v2"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func DataFactory(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armdatafactory.NewFactoriesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	connClient, err := armdatafactory.NewPrivateEndPointConnectionsClient(subscription, cred, nil)
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
			resource, err := getDataFactory(ctx, connClient, v)
			if err != nil {
				return nil, err
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

func getDataFactory(ctx context.Context, connClient *armdatafactory.PrivateEndPointConnectionsClient, factory *armdatafactory.Factory) (*models.Resource, error) {
	resourceGroup := strings.Split(*factory.ID, "/")[4]

	pager := connClient.NewListByFactoryPager(resourceGroup, *factory.Name, nil)
	var datafactoryListByFactoryOp []armdatafactory.PrivateEndpointConnectionResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			datafactoryListByFactoryOp = append(datafactoryListByFactoryOp, *v)
		}
	}

	resource := models.Resource{
		ID:       *factory.ID,
		Name:     *factory.Name,
		Location: *factory.Location,
		Description: model.DataFactoryDescription{
			Factory:                    *factory,
			PrivateEndPointConnections: datafactoryListByFactoryOp,
			ResourceGroup:              resourceGroup,
		},
	}
	return &resource, nil
}

func DataFactoryDataset(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armdatafactory.NewFactoriesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	datasetsClient, err := armdatafactory.NewDatasetsClient(subscription, cred, nil)
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
			resources, err := getDataFactoryDataset(ctx, datasetsClient, v)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func getDataFactoryDataset(ctx context.Context, client *armdatafactory.DatasetsClient, factory *armdatafactory.Factory) ([]models.Resource, error) {
	factoryName := *factory.Name
	factoryResourceGroup := strings.Split(*factory.ID, "/")[4]

	pager := client.NewListByFactoryPager(factoryResourceGroup, factoryName, nil)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, dataset := range page.Value {
			resource := models.Resource{
				ID:       *dataset.ID,
				Name:     *dataset.Name,
				Location: *factory.Location,
				Description: model.DataFactoryDatasetDescription{
					Factory:       *factory,
					Dataset:       *dataset,
					ResourceGroup: factoryResourceGroup,
				},
			}
			values = append(values, resource)
		}
	}

	return values, nil
}

func DataFactoryPipeline(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armdatafactory.NewFactoriesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pipelineClient, err := armdatafactory.NewPipelinesClient(subscription, cred, nil)
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
			resources, err := getDataFactoryPipeline(ctx, pipelineClient, v)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func getDataFactoryPipeline(ctx context.Context, client *armdatafactory.PipelinesClient, factory *armdatafactory.Factory) ([]models.Resource, error) {
	factoryName := *factory.Name
	factoryResourceGroup := strings.Split(*factory.ID, "/")[4]

	pager := client.NewListByFactoryPager(factoryResourceGroup, factoryName, nil)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, pipeline := range page.Value {
			resource := models.Resource{
				ID:       *pipeline.ID,
				Name:     *pipeline.Name,
				Location: *factory.Location,
				Description: model.DataFactoryPipelineDescription{
					Factory:       *factory,
					Pipeline:      *pipeline,
					ResourceGroup: factoryResourceGroup,
				},
			}
			values = append(values, resource)
		}
	}

	return values, nil
}
