package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func LogicAppWorkflow(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armlogic.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewWorkflowsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, workflow := range page.Value {
			resource, err := getLogicAppWorkflow(ctx, diagnosticClient, workflow)
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

func getLogicAppWorkflow(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, workflow *armlogic.Workflow) (*models.Resource, error) {
	resourceGroup := strings.Split(*workflow.ID, "/")[4]

	var logicListOp []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*workflow.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		logicListOp = append(logicListOp, page.Value...)
	}

	resource := models.Resource{
		ID:       *workflow.ID,
		Name:     *workflow.Name,
		Location: *workflow.Location,
		Description: model.LogicAppWorkflowDescription{
			Workflow:                    *workflow,
			DiagnosticSettingsResources: logicListOp,
			ResourceGroup:               resourceGroup,
		},
	}
	return &resource, nil
}

func LogicIntegrationAccounts(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armlogic.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewIntegrationAccountsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			resource := getLogicIntegrationAccounts(ctx, account)
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

func getLogicIntegrationAccounts(ctx context.Context, account *armlogic.IntegrationAccount) *models.Resource {
	resourceGroup := strings.Split(*account.ID, "/")[4]

	resource := models.Resource{
		ID:       *account.ID,
		Name:     *account.Name,
		Location: *account.Location,
		Description: model.LogicIntegrationAccountsDescription{
			Account:       *account,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
