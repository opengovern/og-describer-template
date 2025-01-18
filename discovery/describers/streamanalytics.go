package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func StreamAnalyticsJob(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstreamanalytics.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	streamingJobsClient := clientFactory.NewStreamingJobsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	var values []models.Resource
	pager := streamingJobsClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := GetStreamAnalyticsJob(ctx, diagnosticClient, v)
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

func GetStreamAnalyticsJob(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, streamingJob *armstreamanalytics.StreamingJob) (*models.Resource, error) {
	resourceGroup := strings.Split(*streamingJob.ID, "/")[4]

	pager := diagnosticClient.NewListPager(*streamingJob.ID, nil)
	var streamanalyticsListOp []*armmonitor.DiagnosticSettingsResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		streamanalyticsListOp = append(streamanalyticsListOp, page.Value...)
	}

	resource := models.Resource{
		ID:       *streamingJob.ID,
		Name:     *streamingJob.Name,
		Location: *streamingJob.Location,
		Description: model.StreamAnalyticsJobDescription{
			StreamingJob:                *streamingJob,
			DiagnosticSettingsResources: streamanalyticsListOp,
			ResourceGroup:               resourceGroup,
		},
	}
	return &resource, nil
}

func StreamAnalyticsCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstreamanalytics.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	streamingJobsClient := clientFactory.NewClustersClient()

	var values []models.Resource
	pager := streamingJobsClient.NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetStreamAnalyticsCluster(ctx, v)
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

func GetStreamAnalyticsCluster(ctx context.Context, streamingJob *armstreamanalytics.Cluster) *models.Resource {
	resourceGroup := strings.Split(*streamingJob.ID, "/")[4]

	resource := models.Resource{
		ID:       *streamingJob.ID,
		Name:     *streamingJob.Name,
		Location: *streamingJob.Location,
		Description: model.StreamAnalyticsClusterDescription{
			StreamingJob:  *streamingJob,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
