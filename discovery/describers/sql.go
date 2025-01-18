package describers

import (
	"context"

	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"

	"strings"

	
)

func MssqlManagedInstance(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewManagedInstancesClient()
	managedInstanceClient := clientFactory.NewManagedInstanceVulnerabilityAssessmentsClient()
	managedServerClient := clientFactory.NewManagedServerSecurityAlertPoliciesClient()
	managedInstanceEncClient := clientFactory.NewManagedInstanceEncryptionProtectorsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, managedInstance := range page.Value {
			resource, err := GetMssqlManagedInstance(ctx, managedInstanceClient, managedServerClient, managedInstanceEncClient, managedInstance)
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

func GetMssqlManagedInstance(ctx context.Context, managedInstanceClient *armsql.ManagedInstanceVulnerabilityAssessmentsClient, managedServerClient *armsql.ManagedServerSecurityAlertPoliciesClient, managedInstanceEncClient *armsql.ManagedInstanceEncryptionProtectorsClient, managedInstance *armsql.ManagedInstance) (*models.Resource, error) {
	resourceGroup := strings.Split(string(*managedInstance.ID), "/")[4]
	managedInstanceName := *managedInstance.Name

	var viop []*armsql.ManagedInstanceVulnerabilityAssessment
	pager1 := managedInstanceClient.NewListByInstancePager(resourceGroup, managedInstanceName, nil)
	for pager1.More() {
		page1, err := pager1.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		viop = append(viop, page1.Value...)
	}

	var vsop []*armsql.ManagedServerSecurityAlertPolicy
	pager2 := managedServerClient.NewListByInstancePager(resourceGroup, managedInstanceName, nil)
	for pager2.More() {
		page2, err := pager2.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		vsop = append(vsop, page2.Value...)
	}

	var veop []*armsql.ManagedInstanceEncryptionProtector
	pager3 := managedInstanceEncClient.NewListByInstancePager(resourceGroup, managedInstanceName, nil)
	for pager3.More() {
		page3, err := pager3.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		veop = append(veop, page3.Value...)
	}

	resource := models.Resource{
		ID:       *managedInstance.ID,
		Name:     *managedInstance.Name,
		Location: *managedInstance.Location,
		Description: model.MssqlManagedInstanceDescription{
				ManagedInstance:                         *managedInstance,
				ManagedInstanceVulnerabilityAssessments: viop,
				ManagedDatabaseSecurityAlertPolicies:    vsop,
				ManagedInstanceEncryptionProtectors:     veop,
				ResourceGroup:                           resourceGroup,
			},
	}

	return &resource, nil
}

func MssqlManagedInstanceDatabases(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewManagedInstancesClient()
	dbClient := clientFactory.NewManagedDatabasesClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, managedInstance := range page.Value {
			resources, err := ListManagedInstanceDatabases(ctx, dbClient, managedInstance)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				for _, resource := range resources {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				}
			} else {
				values = append(values, resources...)
			}
		}
	}
	return values, nil
}

func ListManagedInstanceDatabases(ctx context.Context, dbClient *armsql.ManagedDatabasesClient, managedInstance *armsql.ManagedInstance) ([]models.Resource, error) {
	resourceGroup := strings.Split(*managedInstance.ID, "/")[4]

	var values []models.Resource
	pager := dbClient.NewListByInstancePager(resourceGroup, *managedInstance.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, db := range page.Value {
			resource := GetManagedInstanceDatabases(ctx, managedInstance, db)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetManagedInstanceDatabases(ctx context.Context, managedInstance *armsql.ManagedInstance, db *armsql.ManagedDatabase) *models.Resource {
	resourceGroup := strings.Split(string(*managedInstance.ID), "/")[4]

	resource := models.Resource{
		ID:       *db.ID,
		Name:     *db.Name,
		Location: *db.Location,
		Description: model.MssqlManagedInstanceDatabasesDescription{
				ManagedInstance: *managedInstance,
				Database:        *db,
				ResourceGroup:   resourceGroup,
			},
	}
	return &resource
}

func SqlDatabase(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	parentClient := clientFactory.NewServersClient()
	databaseVulnerabilityScanClient := clientFactory.NewDatabaseVulnerabilityAssessmentScansClient()
	databaseVulnerabilityClient := clientFactory.NewDatabaseVulnerabilityAssessmentsClient()
	transparentDataClient := clientFactory.NewTransparentDataEncryptionsClient()
	longTermClient := clientFactory.NewLongTermRetentionPoliciesClient()
	databasesClientClient := clientFactory.NewDatabasesClient()
	client := clientFactory.NewDatabasesClient()
	advisorsClient := clientFactory.NewDatabaseAdvisorsClient()
	recoverableClient := clientFactory.NewRecoverableDatabasesClient()
	auditingPolicyClient := clientFactory.NewDatabaseBlobAuditingPoliciesClient()

	pager := parentClient.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resources, err := ListServerSqlDatabases(ctx, recoverableClient, advisorsClient, databaseVulnerabilityScanClient, databaseVulnerabilityClient, transparentDataClient, longTermClient, databasesClientClient, auditingPolicyClient, client, server)
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

func ListServerSqlDatabases(ctx context.Context, recoverableClient *armsql.RecoverableDatabasesClient, advisorsClient *armsql.DatabaseAdvisorsClient, databaseVulnerabilityScanClient *armsql.DatabaseVulnerabilityAssessmentScansClient, databaseVulnerabilityClient *armsql.DatabaseVulnerabilityAssessmentsClient, transparentDataClient *armsql.TransparentDataEncryptionsClient, longTermClient *armsql.LongTermRetentionPoliciesClient, databasesClientClient *armsql.DatabasesClient, auditingPoliciesClient *armsql.DatabaseBlobAuditingPoliciesClient, client *armsql.DatabasesClient, server *armsql.Server) ([]models.Resource, error) {
	resourceGroupName := strings.Split(string(*server.ID), "/")[4]
	pager := client.NewListByServerPager(resourceGroupName, *server.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, database := range page.Value {
			resource, err := GetSqlDatabase(ctx, recoverableClient, advisorsClient, databaseVulnerabilityScanClient, databaseVulnerabilityClient, transparentDataClient, longTermClient, databasesClientClient, auditingPoliciesClient, server, database)
			if err != nil {
				return nil, err
			}
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSqlDatabase(ctx context.Context, recoverableClient *armsql.RecoverableDatabasesClient, advisorsClient *armsql.DatabaseAdvisorsClient, databaseVulnerabilityScanClient *armsql.DatabaseVulnerabilityAssessmentScansClient, databaseVulnerabilityClient *armsql.DatabaseVulnerabilityAssessmentsClient, transparentDataClient *armsql.TransparentDataEncryptionsClient, longTermClient *armsql.LongTermRetentionPoliciesClient, databasesClientClient *armsql.DatabasesClient, auditingPoliciesClient *armsql.DatabaseBlobAuditingPoliciesClient, server *armsql.Server, database *armsql.Database) (*models.Resource, error) {
	serverName := strings.Split(*database.ID, "/")[8]
	databaseName := *database.Name
	resourceGroupName := strings.Split(string(*database.ID), "/")[4]

	var longTermRetentionPolicies []*armsql.LongTermRetentionPolicy
	pager1 := longTermClient.NewListByDatabasePager(resourceGroupName, serverName, databaseName, nil)
	for pager1.More() {
		page1, err := pager1.NextPage(ctx)
		if err != nil {
			break
		}
		longTermRetentionPolicies = append(longTermRetentionPolicies, page1.Value...)
	}
	var longTermRetentionPolicy armsql.LongTermRetentionPolicy
	if len(longTermRetentionPolicies) > 0 {
		longTermRetentionPolicy = *longTermRetentionPolicies[0]
	}

	pager2 := transparentDataClient.NewListByDatabasePager(resourceGroupName, serverName, databaseName, nil)
	var transparentDataOp []*armsql.LogicalDatabaseTransparentDataEncryption
	for pager2.More() {
		page2, err := pager2.NextPage(ctx)
		if err != nil {
			break
		}
		transparentDataOp = append(transparentDataOp, page2.Value...)
	}

	var c []*armsql.DatabaseVulnerabilityAssessment
	var v []*armsql.VulnerabilityAssessmentScanRecord
	pager3 := databaseVulnerabilityClient.NewListByDatabasePager(resourceGroupName, serverName, databaseName, nil)
	for pager3.More() {
		page3, err := pager3.NextPage(ctx)
		if err != nil {
			if !strings.Contains(err.Error(), "VulnerabilityAssessmentInvalidPolicy") {
				return nil, err
			}
		} else {
			for _, assessment := range c {
				pager4 := databaseVulnerabilityScanClient.NewListByDatabasePager(resourceGroupName, serverName, databaseName, armsql.VulnerabilityAssessmentName(*assessment.Name), nil)
				for pager4.More() {
					page4, err := pager4.NextPage(ctx)
					if err != nil {
						break
					}
					v = append(v, page4.Value...)
				}
			}
		}
		c = append(c, page3.Value...)
	}

	var auditPolicies []*armsql.DatabaseBlobAuditingPolicy
	pager4 := auditingPoliciesClient.NewListByDatabasePager(resourceGroupName, serverName, databaseName, nil)
	for pager4.More() {
		page4, err := pager4.NextPage(ctx)
		if err != nil {
			break
		}
		auditPolicies = append(auditPolicies, page4.Value...)
	}

	advisors, err := advisorsClient.ListByDatabase(ctx, resourceGroupName, serverName, databaseName, nil)
	if err != nil {
		//IGNORE ERROR
	}

	getOp, err := databasesClientClient.Get(ctx, resourceGroupName, serverName, databaseName, nil)
	if err != nil {
		return nil, err
	}

	resource := models.Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: model.SqlDatabaseDescription{
				Database:                           getOp.Database,
				LongTermRetentionPolicy:            longTermRetentionPolicy,
				TransparentDataEncryption:          transparentDataOp,
				DatabaseVulnerabilityAssessments:   c,
				VulnerabilityAssessmentScanRecords: v,
				Advisors:                           advisors.AdvisorArray,
				AuditPolicies:                      auditPolicies,
				ResourceGroup:                      resourceGroupName,
			},
	}
	return &resource, nil
}

func SqlInstancePool(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsql.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	parentClient := clientFactory.NewInstancePoolsClient()

	pager := parentClient.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, instancePool := range page.Value {
			resource, err := GetSqlInstancePool(ctx, clientFactory, instancePool)
			if err != nil {
				return nil, err
			}
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSqlInstancePool(ctx context.Context, clientFactory *armsql.ClientFactory, v *armsql.InstancePool) (*models.Resource, error) {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.SqlInstancePoolDescription{
				InstancePool:  *v,
				ResourceGroup: resourceGroupName,
			},
	}
	return &resource, nil
}
