package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func DataLakeAnalyticsAccount(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armdatalakeanalytics.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAccountsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	if err != nil {
		return nil, err
	}
	accountsPages := client.NewListPager(nil)
	var values []models.Resource
	for accountsPages.More() {
		page, err := accountsPages.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			resource, err := getDataLakeAnalyticsAccount(ctx, account, client, diagnosticClient)
			if err != nil {
				return nil, err
			}
			if resource == nil {
				continue
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

func getDataLakeAnalyticsAccount(ctx context.Context, account *armdatalakeanalytics.AccountBasic, client *armdatalakeanalytics.AccountsClient, diagnosticClient *armmonitor.DiagnosticSettingsClient) (*models.Resource, error) {
	splitID := strings.Split(*account.ID, "/")
	name := *account.Name
	resourceGroup := splitID[4]
	accountGetOp, err := client.Get(ctx, resourceGroup, name, nil)
	if err != nil {
		return nil, err
	}
	if name == "" || resourceGroup == "" {
		return nil, nil
	}

	var accountListOp []armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*account.ID, nil)
	for pager.More() {
		accountOpPage, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, accountOp := range accountOpPage.Value {
			accountListOp = append(accountListOp, *accountOp)
		}
	}

	resource := models.Resource{
		ID:       *account.ID,
		Name:     *account.Name,
		Location: *account.Location,
		Description: model.DataLakeAnalyticsAccountDescription{
			DataLakeAnalyticsAccount:   accountGetOp.Account,
			DiagnosticSettingsResource: &accountListOp,
			ResourceGroup:              resourceGroup,
		},
	}
	return &resource, nil
}

func DataLakeStore(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armdatalakestore.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAccountsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			resource, err := getDataLakeStore(ctx, account, diagnosticClient, client)
			if err != nil {
				return nil, err
			}
			if resource == nil {
				continue
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

func getDataLakeStore(ctx context.Context, account *armdatalakestore.AccountBasic, diagnosticClient *armmonitor.DiagnosticSettingsClient, client *armdatalakestore.AccountsClient) (*models.Resource, error) {
	splitId := strings.Split(*account.ID, "/")
	name := *account.Name
	resourceGroup := splitId[4]

	if name == "" || resourceGroup == "" {
		return nil, nil
	}

	accountGetOp, err := client.Get(ctx, resourceGroup, name, nil)
	if err != nil {
		return nil, err
	}
	accountListOpTemp := diagnosticClient.NewListPager(*account.ID, nil)
	var accountListOp []armmonitor.DiagnosticSettingsResource
	for accountListOpTemp.More() {
		accountOpPage, err := accountListOpTemp.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, accountOp := range accountOpPage.Value {
			accountListOp = append(accountListOp, *accountOp)
		}
	}
	resource := models.Resource{
		ID:       *account.ID,
		Name:     name,
		Location: "",
		Description: model.DataLakeStoreDescription{
			DataLakeStoreAccount:       accountGetOp.Account,
			DiagnosticSettingsResource: &accountListOp,
			ResourceGroup:              resourceGroup,
		},
	}
	return &resource, nil
}
