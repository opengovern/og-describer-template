package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func TimeSeriesInsightsEnvironments(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armtimeseriesinsights.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewEnvironmentsClient()

	list, err := client.ListBySubscription(ctx, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, record := range list.Value {
		resource := GetTimeSeriesInsightsEnvironments(ctx, record)
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

func GetTimeSeriesInsightsEnvironments(ctx context.Context, record armtimeseriesinsights.EnvironmentResourceClassification) *models.Resource {
	v := record.GetEnvironmentResource()
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.TimeSeriesInsightsEnvironmentsDescription{
			Environment:   v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
