package describers

import (
	"context"
	"strings"

	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"

)

func MysqlServer(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armmysql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	serversClient := clientFactory.NewServersClient()
	keysClient := clientFactory.NewServerKeysClient()
	configClient := clientFactory.NewConfigurationsClient()
	securityAlertPolicyClient := clientFactory.NewServerSecurityAlertPoliciesClient()
	vnetRulesClient := clientFactory.NewVirtualNetworkRulesClient()

	pager := serversClient.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, r := range page.Value {
			resource, err := getMysqlServer(ctx, keysClient, configClient, securityAlertPolicyClient, vnetRulesClient, r)
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

func getMysqlServer(ctx context.Context, keysClient *armmysql.ServerKeysClient, configClient *armmysql.ConfigurationsClient, securityAlertPolicyClient *armmysql.ServerSecurityAlertPoliciesClient, vnetRulesClient *armmysql.VirtualNetworkRulesClient, server *armmysql.Server) (*models.Resource, error) {
	resourceGroup := strings.Split(string(*server.ID), "/")[4]
	serverName := *server.Name

	pager1 := configClient.NewListByServerPager(resourceGroup, serverName, nil)
	var configurations []*armmysql.Configuration
	for pager1.More() {
		page, err := pager1.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		configurations = append(configurations, page.Value...)
	}

	pager2 := keysClient.NewListPager(resourceGroup, serverName, nil)
	var keys []*armmysql.ServerKey
	for pager2.More() {
		page, err := pager2.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		keys = append(keys, page.Value...)
	}

	pager3 := securityAlertPolicyClient.NewListByServerPager(resourceGroup, serverName, nil)
	var alertPolicies []*armmysql.ServerSecurityAlertPolicy
	for pager3.More() {
		page, err := pager3.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		alertPolicies = append(alertPolicies, page.Value...)
	}

	pager4 := vnetRulesClient.NewListByServerPager(resourceGroup, serverName, nil)
	var vnetRules []*armmysql.VirtualNetworkRule
	for pager4.More() {
		page, err := pager4.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		vnetRules = append(vnetRules, page.Value...)
	}

	resource := models.Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: model.MysqlServerDescription{
			Server:                *server,
			Configurations:        configurations,
			ServerKeys:            keys,
			SecurityAlertPolicies: alertPolicies,
			VnetRules:             vnetRules,
			ResourceGroup:         resourceGroup,
		},
	}
	return &resource, nil
}

func MysqlFlexibleservers(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armmysqlflexibleservers.NewClientFactory(subscription, cred, nil)
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
			resource := getMysqlFlexibleservers(ctx, server)
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

func getMysqlFlexibleservers(ctx context.Context, server *armmysqlflexibleservers.Server) *models.Resource {
	resourceGroup := strings.Split(string(*server.ID), "/")[4]

	resource := models.Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: model.MysqlFlexibleserverDescription{
			Server:        *server,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
