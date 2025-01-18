package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func DatabricksWorkspaces(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armdatabricks.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewWorkspacesClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getDatabricksWorkspace(ctx, v)
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

func getDatabricksWorkspace(ctx context.Context, v *armdatabricks.Workspace) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.DatabricksWorkspaceDescription{
			Workspace:     *v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
