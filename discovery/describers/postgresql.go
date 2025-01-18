package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func PostgresqlServer(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armpostgresql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	firewallClient := clientFactory.NewFirewallRulesClient()
	keysClient := clientFactory.NewServerKeysClient()
	confClient := clientFactory.NewConfigurationsClient()
	adminClient := clientFactory.NewServerAdministratorsClient()
	client := clientFactory.NewServersClient()
	alertPolicyClient := clientFactory.NewServerSecurityAlertPoliciesClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resource, err := GetPostgresqlServer(ctx, firewallClient, keysClient, confClient, adminClient, alertPolicyClient, server)
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

func GetPostgresqlServer(ctx context.Context, firewallClient *armpostgresql.FirewallRulesClient, keysClient *armpostgresql.ServerKeysClient, confClient *armpostgresql.ConfigurationsClient, adminClient *armpostgresql.ServerAdministratorsClient, alertPolicyClient *armpostgresql.ServerSecurityAlertPoliciesClient, server *armpostgresql.Server) (*models.Resource, error) {
	resourceGroupName := strings.Split(string(*server.ID), "/")[4]

	pager := adminClient.NewListPager(resourceGroupName, *server.Name, nil)
	var adminListOp []*armpostgresql.ServerAdministratorResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			break
		}
		adminListOp = append(adminListOp, page.Value...)
	}

	pager2 := confClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var confListByServerOp []*armpostgresql.Configuration
	for pager2.More() {
		page, err := pager2.NextPage(ctx)
		if err != nil {
			break
		}
		confListByServerOp = append(confListByServerOp, page.Value...)
	}

	pager3 := keysClient.NewListPager(resourceGroupName, *server.Name, nil)
	var kop []*armpostgresql.ServerKey
	for pager3.More() {
		page, err := pager3.NextPage(ctx)
		if err != nil {
			break
		}
		kop = append(kop, page.Value...)
	}

	pager4 := firewallClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var firewallListByServerOp []*armpostgresql.FirewallRule
	for pager4.More() {
		page, err := pager4.NextPage(ctx)
		if err != nil {
			break
		}
		firewallListByServerOp = append(firewallListByServerOp, page.Value...)
	}

	pager5 := alertPolicyClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var serverSecurityAlertPolicies []*armpostgresql.ServerSecurityAlertPolicy
	for pager5.More() {
		page, err := pager5.NextPage(ctx)
		if err != nil {
			break
		}
		serverSecurityAlertPolicies = append(serverSecurityAlertPolicies, page.Value...)
	}

	resource := models.Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: model.PostgresqlServerDescription{
			Server:                       *server,
			ServerAdministratorResources: adminListOp,
			Configurations:               confListByServerOp,
			ServerKeys:                   kop,
			FirewallRules:                firewallListByServerOp,
			ServerSecurityAlertPolicies:  serverSecurityAlertPolicies,
			ResourceGroup:                resourceGroupName,
		},
	}

	return &resource, nil
}

func PostgresqlFlexibleservers(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {

	client, err := armpostgresqlflexibleservers.NewServersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	configurationsClient, err := armpostgresqlflexibleservers.NewConfigurationsClient(subscription, cred, nil)

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resource := GetPostgresqlFlexibleserver(ctx, configurationsClient, server)
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

func GetPostgresqlFlexibleserver(ctx context.Context, configurationsClient *armpostgresqlflexibleservers.ConfigurationsClient, server *armpostgresqlflexibleservers.Server) *models.Resource {
	resourceGroupName := strings.Split(string(*server.ID), "/")[4]

	pager := configurationsClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var serverConfigurations []*armpostgresqlflexibleservers.Configuration
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			break
		}
		serverConfigurations = append(serverConfigurations, page.Value...)
	}

	resource := models.Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: model.PostgresqlFlexibleServerDescription{
			Server:               *server,
			ServerConfigurations: serverConfigurations,
			ResourceGroup:        resourceGroupName,
		},
	}
	return &resource
}
