package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func AnalysisService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armanalysisservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServersClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resource := getAnalysisService(ctx, server)
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

func getAnalysisService(ctx context.Context, server *armanalysisservices.Server) *models.Resource {
	resourceGroupName := strings.Split(*server.ID, "/")[4]

	resource := models.Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: model.AnalysisServiceServerDescription{
			Server:        *server,
			ResourceGroup: resourceGroupName,
		},
	}
	return &resource
}
