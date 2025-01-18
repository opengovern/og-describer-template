package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func SqlServer(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	virtualNetworkClient := clientFactory.NewVirtualNetworkRulesClient()
	privateEndpointClient := clientFactory.NewPrivateEndpointConnectionsClient()
	encryptionProtectorsClient := clientFactory.NewEncryptionProtectorsClient()
	firewallRulesClient := clientFactory.NewFirewallRulesClient()
	serverVulnerabilityClient := clientFactory.NewServerVulnerabilityAssessmentsClient()
	serverAzureClient := clientFactory.NewServerAzureADAdministratorsClient()
	serverSecurityClient := clientFactory.NewServerSecurityAlertPoliciesClient()
	serverBlobClient := clientFactory.NewServerBlobAuditingPoliciesClient()
	failoverClient := clientFactory.NewFailoverGroupsClient()
	automaticTuningClient := clientFactory.NewServerAutomaticTuningClient()
	client := clientFactory.NewServersClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resource, err := GetSqlServer(ctx, automaticTuningClient, failoverClient, virtualNetworkClient, privateEndpointClient, encryptionProtectorsClient, firewallRulesClient, serverVulnerabilityClient, serverAzureClient, serverSecurityClient, serverBlobClient, server)
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
	return values, err
}

func GetSqlServer(ctx context.Context, automaticTuningClient *armsql.ServerAutomaticTuningClient, failoverClient *armsql.FailoverGroupsClient, virtualNetworkClient *armsql.VirtualNetworkRulesClient, privateEndpointClient *armsql.PrivateEndpointConnectionsClient, encryptionProtectorsClient *armsql.EncryptionProtectorsClient, firewallRulesClient *armsql.FirewallRulesClient, serverVulnerabilityClient *armsql.ServerVulnerabilityAssessmentsClient, serverAzureClient *armsql.ServerAzureADAdministratorsClient, serverSecurityClient *armsql.ServerSecurityAlertPoliciesClient, serverBlobClient *armsql.ServerBlobAuditingPoliciesClient, server *armsql.Server) (*models.Resource, error) {
	resourceGroupName := strings.Split(string(*server.ID), "/")[4]

	pager1 := serverBlobClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var bop []*armsql.ServerBlobAuditingPolicy
	for pager1.More() {
		page, err := pager1.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		bop = append(bop, page.Value...)
	}

	pager2 := serverSecurityClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var sop []*armsql.ServerSecurityAlertPolicy
	for pager2.More() {
		page2, err := pager2.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		sop = append(sop, page2.Value...)
	}

	pager3 := serverAzureClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var adminOp []*armsql.ServerAzureADAdministrator
	for pager3.More() {
		page3, err := pager3.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		adminOp = append(adminOp, page3.Value...)
	}

	pager4 := serverVulnerabilityClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var vop []*armsql.ServerVulnerabilityAssessment
	for pager4.More() {
		page4, err := pager4.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		vop = append(vop, page4.Value...)
	}

	pager5 := firewallRulesClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var firewallOp []*armsql.FirewallRule
	for pager5.More() {
		page5, err := pager5.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		firewallOp = append(firewallOp, page5.Value...)
	}

	pager6 := encryptionProtectorsClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var eop []*armsql.EncryptionProtector
	for pager6.More() {
		page6, err := pager6.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		eop = append(eop, page6.Value...)
	}

	pager7 := privateEndpointClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var pop []*armsql.PrivateEndpointConnection
	for pager7.More() {
		page7, err := pager7.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		pop = append(pop, page7.Value...)
	}

	pager8 := virtualNetworkClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var nop []*armsql.VirtualNetworkRule
	for pager8.More() {
		page8, err := pager8.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		nop = append(nop, page8.Value...)
	}

	pager9 := failoverClient.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var fop []*armsql.FailoverGroup
	for pager9.More() {
		page9, err := pager9.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		fop = append(fop, page9.Value...)
	}

	automaticTuning, err := automaticTuningClient.Get(ctx, resourceGroupName, *server.Name, nil)
	if err != nil {
		return nil, err
	}

	resource := models.Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: model.SqlServerDescription{
			Server:                         *server,
			ServerBlobAuditingPolicies:     bop,
			ServerSecurityAlertPolicies:    sop,
			ServerAzureADAdministrators:    adminOp,
			ServerVulnerabilityAssessments: vop,
			FirewallRules:                  firewallOp,
			EncryptionProtectors:           eop,
			PrivateEndpointConnections:     pop,
			VirtualNetworkRules:            nop,
			FailoverGroups:                 fop,
			AutomaticTuning:                automaticTuning.ServerAutomaticTuning,
			ResourceGroup:                  resourceGroupName,
		},
	}
	return &resource, nil
}

func SqlServerJobAgents(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	serverClient := clientFactory.NewServersClient()
	client := clientFactory.NewJobAgentsClient()

	var values []models.Resource
	pager := serverClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resources, err := ListSqlServerJobAgents(ctx, client, server)
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
	return values, err
}

func ListSqlServerJobAgents(ctx context.Context, client *armsql.JobAgentsClient, server *armsql.Server) ([]models.Resource, error) {
	resourceGroupName := strings.Split(string(*server.ID), "/")[4]

	pager := client.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, job := range page.Value {
			resource := GetSqlServerJobAgent(ctx, server, job)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSqlServerJobAgent(ctx context.Context, server *armsql.Server, job *armsql.JobAgent) *models.Resource {
	jobResourceGroupName := strings.Split(string(*job.ID), "/")[4]

	resource := models.Resource{
		ID:       *job.ID,
		Name:     *job.Name,
		Location: *job.Location,
		Description: model.SqlServerJobAgentDescription{
			Server:        *server,
			JobAgent:      *job,
			ResourceGroup: jobResourceGroupName,
		},
	}
	return &resource
}

func SqlVirtualClusters(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewVirtualClustersClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSqlVirtualCluster(ctx, v)
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

func GetSqlVirtualCluster(ctx context.Context, v *armsql.VirtualCluster) *models.Resource {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.SqlVirtualClustersDescription{
			VirtualClusters: *v,
			ResourceGroup:   resourceGroupName,
		},
	}

	return &resource
}

func SqlServerElasticPool(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewServersClient()
	elasticPoolClient := clientFactory.NewElasticPoolsClient()
	activityClient := clientFactory.NewElasticPoolActivitiesClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resources, err := ListSqlServerElasticPools(ctx, elasticPoolClient, activityClient, server)
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

func ListSqlServerElasticPools(ctx context.Context, elasticPoolClient *armsql.ElasticPoolsClient, activityClient *armsql.ElasticPoolActivitiesClient, server *armsql.Server) ([]models.Resource, error) {
	if server == nil || server.ID == nil {
		return nil, nil
	}
	serverResourceGroup := strings.Split(string(*server.ID), "/")[4]

	var values []models.Resource
	name := *server.ID
	if server.Name != nil {
		name = *server.Name
	}
	pager := elasticPoolClient.NewListByServerPager(serverResourceGroup, name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := GetSqlServerElasticPool(ctx, server, activityClient, v)
			if err != nil {
				return nil, err
			}
			if resource == nil {
				continue
			}
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSqlServerElasticPool(ctx context.Context, server *armsql.Server, activityClient *armsql.ElasticPoolActivitiesClient, elasticPool *armsql.ElasticPool) (*models.Resource, error) {
	if elasticPool == nil || elasticPool.ID == nil {
		return nil, nil
	}
	resourceGroup := strings.Split(string(*elasticPool.ID), "/")[4]

	var totalDTU int32
	name := *server.ID
	if server.Name != nil {
		name = *server.Name
	}
	epName := *elasticPool.ID
	if elasticPool.Name != nil {
		epName = *elasticPool.Name
	}
	pager := activityClient.NewListByElasticPoolPager(resourceGroup, name, epName, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			if v.Properties == nil || v.Properties.RequestedDtu == nil {
				continue
			}
			totalDTU = totalDTU + *v.Properties.RequestedDtu
		}
	}
	resource := models.Resource{
		ID:       *elasticPool.ID,
		Name:     *elasticPool.Name,
		Location: *elasticPool.Location,
		Description: model.SqlServerElasticPoolDescription{
			TotalDTU:      totalDTU,
			Pool:          *elasticPool,
			ServerName:    name,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource, nil
}

func SqlServerVirtualMachine(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsqlvirtualmachine.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSQLVirtualMachinesClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSqlServerVirtualMachine(ctx, v)
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

func GetSqlServerVirtualMachine(ctx context.Context, vm *armsqlvirtualmachine.SQLVirtualMachine) *models.Resource {
	resourceGroup := strings.Split(string(*vm.ID), "/")[4]
	resource := models.Resource{
		ID:       *vm.ID,
		Name:     *vm.Name,
		Location: *vm.Location,
		Description: model.SqlServerVirtualMachineDescription{
			VirtualMachine: *vm,
			ResourceGroup:  resourceGroup,
		},
	}
	return &resource
}

func SqlServerVirtualMachineGroups(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsqlvirtualmachine.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewGroupsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vm := range page.Value {
			resource := GetSqlServerVirtualMachineGroups(ctx, vm)
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

func GetSqlServerVirtualMachineGroups(ctx context.Context, vm *armsqlvirtualmachine.Group) *models.Resource {
	resourceGroup := strings.Split(string(*vm.ID), "/")[4]

	resource := models.Resource{
		ID:       *vm.ID,
		Name:     *vm.Name,
		Location: *vm.Location,
		Description: model.SqlServerVirtualMachineGroupDescription{
			Group:         *vm,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func SqlServerFlexibleServer(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armmysqlflexibleservers.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServersClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, fs := range page.Value {
			resource := GetSqlServerFlexibleServer(ctx, fs)
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

func GetSqlServerFlexibleServer(ctx context.Context, fs *armmysqlflexibleservers.Server) *models.Resource {
	resourceGroup := strings.Split(string(*fs.ID), "/")[4]
	resource := models.Resource{
		ID:       *fs.ID,
		Name:     *fs.Name,
		Location: *fs.Location,
		Description: model.SqlServerFlexibleServerDescription{
			FlexibleServer: *fs,
			ResourceGroup:  resourceGroup,
		},
	}
	return &resource
}
