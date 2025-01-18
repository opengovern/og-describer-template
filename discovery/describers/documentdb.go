package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func DocumentDBSQLDatabase(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSQLResourcesClient()

	var values []models.Resource
	for _, rg := range rgs {
		accounts, err := documentDBDatabaseAccounts(ctx, cred, subscription, *rg.Name)
		if err != nil {
			return nil, err
		}

		for _, account := range accounts {

			pager := client.NewListSQLDatabasesPager(*rg.Name, *account.Name, nil)
			var it []*armcosmos.SQLDatabaseGetResults
			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, v := range page.Value {
					it = append(it, v)
				}
			}
			for _, v := range it {
				resource := getDocumentDBSQLDatabase(ctx, v, account, rg)
				if stream != nil {
					if err := (*stream)(*resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, *resource)
				}
			}
		}
	}

	return values, nil
}

func getDocumentDBSQLDatabase(ctx context.Context, v *armcosmos.SQLDatabaseGetResults, account *armcosmos.DatabaseAccountGetResults, rg armresources.ResourceGroup) *models.Resource {
	location := "global"
	if v.Location != nil {
		location = *v.Location
	}

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: model.CosmosdbSqlDatabaseDescription{
			Account:       *account,
			SqlDatabase:   *v,
			ResourceGroup: *rg.Name,
		},
	}
	return &resource
}

func DocumentDBMongoDatabase(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewMongoDBResourcesClient()

	var values []models.Resource
	for _, rg := range rgs {
		accounts, err := documentDBDatabaseAccounts(ctx, cred, subscription, *rg.Name)
		if err != nil {
			return nil, err
		}

		for _, account := range accounts {
			pager := client.NewListMongoDBDatabasesPager(*rg.Name, *account.Name, nil)
			var it []*armcosmos.MongoDBDatabaseGetResults
			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, v := range page.Value {
					it = append(it, v)
				}
			}
			for _, v := range it {
				resource := getDocumentDBMongoDatabase(ctx, v, account, rg)
				if stream != nil {
					if err := (*stream)(*resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, *resource)
				}
			}
		}
	}
	return values, nil
}

func getDocumentDBMongoDatabase(ctx context.Context, v *armcosmos.MongoDBDatabaseGetResults, account *armcosmos.DatabaseAccountGetResults, rg armresources.ResourceGroup) *models.Resource {
	location := ""
	if v.Location != nil {
		location = *v.Location
	}

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: model.CosmosdbMongoDatabaseDescription{
			Account:       *account,
			MongoDatabase: *v,
			ResourceGroup: *rg.Name,
		},
	}
	return &resource
}

func DocumentDBMongoCollection(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewMongoDBResourcesClient()

	var values []models.Resource
	for _, rg := range rgs {
		accounts, err := documentDBDatabaseAccounts(ctx, cred, subscription, *rg.Name)
		if err != nil {
			return nil, err
		}

		for _, account := range accounts {
			pager := client.NewListMongoDBDatabasesPager(*rg.Name, *account.Name, nil)
			var it []*armcosmos.MongoDBDatabaseGetResults
			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, v := range page.Value {
					it = append(it, v)
				}
			}
			for _, v := range it {
				resources, err := ListDocumentDBMongoDatabaseCollections(ctx, client, rg, account, v)
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
	}
	return values, nil
}

func ListDocumentDBMongoDatabaseCollections(ctx context.Context, client *armcosmos.MongoDBResourcesClient, rg armresources.ResourceGroup, account *armcosmos.DatabaseAccountGetResults, db *armcosmos.MongoDBDatabaseGetResults) ([]models.Resource, error) {
	pager := client.NewListMongoDBCollectionsPager(*rg.Name, *account.Name, *db.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getDocumentDBMongoCollection(ctx, client, v, account, rg, db)
			if err != nil {
				return nil, err
			}
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getDocumentDBMongoCollection(ctx context.Context, client *armcosmos.MongoDBResourcesClient, v *armcosmos.MongoDBCollectionGetResults, account *armcosmos.DatabaseAccountGetResults, rg armresources.ResourceGroup, db *armcosmos.MongoDBDatabaseGetResults) (*models.Resource, error) {
	tp, err := client.GetMongoDBCollectionThroughput(ctx, *rg.Name, *account.Name, *db.Name, *v.Name, nil)
	if err != nil {
		return nil, err
	}
	location := ""
	if v.Location != nil {
		location = *v.Location
	}

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: model.CosmosdbMongoCollectionDescription{
			Account:         *account,
			MongoDatabase:   *db,
			MongoCollection: *v,
			Throughput:      tp.ThroughputSettingsGetResults,
			ResourceGroup:   *rg.Name,
		},
	}
	return &resource, nil
}

func DocumentDBCassandraCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewCassandraClustersClient()

	var values []models.Resource
	pager := client.NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getDocumentDBCassandraCluster(ctx, v)
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

func getDocumentDBCassandraCluster(ctx context.Context, v *armcosmos.ClusterResource) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	location := "global"
	if v.Location != nil {
		location = *v.Location
	}
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: model.CosmosdbCassandraClusterDescription{
			CassandraCluster: *v,
			ResourceGroup:    resourceGroup,
		},
	}
	return &resource
}

func documentDBDatabaseAccounts(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, resourceGroup string) ([]*armcosmos.DatabaseAccountGetResults, error) {
	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDatabaseAccountsClient()

	var values []*armcosmos.DatabaseAccountGetResults
	pager := client.NewListByResourceGroupPager(resourceGroup, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		values = append(values, page.Value...)
	}

	return values, nil
}

func CosmosdbAccount(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDatabaseAccountsClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getCosmosdbAccount(ctx, v)
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

func getCosmosdbAccount(ctx context.Context, v *armcosmos.DatabaseAccountGetResults) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	location := ""
	if v.Location != nil {
		location = *v.Location
	}
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: model.CosmosdbAccountDescription{
			DatabaseAccountGetResults: *v,
			ResourceGroup:             resourceGroup,
		},
	}
	return &resource
}

func CosmosdbRestorableDatabaseAccount(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewRestorableDatabaseAccountsClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getRestorableDatabaseAccount(ctx, v)
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

func getRestorableDatabaseAccount(ctx context.Context, v *armcosmos.RestorableDatabaseAccountGetResult) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	location := ""
	if v.Location != nil {
		location = *v.Location
	}
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: model.CosmosdbRestorableDatabaseAccountDescription{
			Account:       *v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
