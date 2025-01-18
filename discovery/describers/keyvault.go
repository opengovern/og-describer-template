package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azcertificates"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	

	"github.com/opengovern/og-util/pkg/concurrency"

	
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func KeyVaultKey(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armkeyvault.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultsClient := clientFactory.NewVaultsClient()
	keysClient := clientFactory.NewKeysClient()

	maxResults := int32(100)
	options := &armkeyvault.VaultsClientListOptions{
		Top: &maxResults,
	}
	pager := vaultsClient.NewListPager(options)

	wpe := concurrency.NewWorkPool(4)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			vault := v
			wpe.AddJob(func() (interface{}, error) {
				resourceGroup := strings.Split(*vault.ID, "/")[4]

				pager2 := keysClient.NewListPager(resourceGroup, *vault.Name, nil)
				var result []*armkeyvault.Key
				for pager2.More() {
					page2, err := pager2.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					result = append(result, page2.Value...)
				}
				wp := concurrency.NewWorkPool(8)
				for _, r := range result {
					resourceGroupCopy := resourceGroup
					vaultCopy := vault
					vCopy := r
					wp.AddJob(func() (interface{}, error) {
						op, err := keysClient.Get(ctx, resourceGroupCopy, *vaultCopy.Name, *vCopy.Name, nil)
						if err != nil {
							return nil, err
						}

						// In some cases resource does not give any notFound error
						// instead of notFound error, it returns empty data
						if op.ID == nil {
							return nil, nil
						}

						return models.Resource{
							ID:       *vCopy.ID,
							Name:     *vCopy.Name,
							Location: *vCopy.Location,
							Description: model.KeyVaultKeyDescription{
								Vault:         *vaultCopy,
								Key:           *vCopy,
								ResourceGroup: resourceGroupCopy,
							},
						}, nil
					})
				}

				results := wp.Run()
				var vvv []models.Resource
				for _, r := range results {
					if r.Error != nil {
						return nil, err
					}
					if r.Value == nil {
						continue
					}
					vvv = append(vvv, r.Value.(models.Resource))
				}
				return vvv, nil
			})
		}
	}

	results := wpe.Run()
	for _, result := range results {
		if result.Error != nil {
			return nil, err
		}
		if result.Value == nil {
			continue
		}
		values = append(values, result.Value.([]models.Resource)...)
	}

	if stream != nil {
		for _, resource := range values {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		}
		values = nil
	}
	return values, nil
}

func getKeyVaultKey(ctx context.Context, keysClient *armkeyvault.KeysClient, vCopy *armkeyvault.Key, resourceGroupCopy string, vaultCopy *armkeyvault.Resource) (*models.Resource, error) {
	op, err := keysClient.Get(ctx, resourceGroupCopy, *vaultCopy.Name, *vCopy.Name, nil)
	if err != nil {
		return nil, err
	}

	// In some cases resource does not give any notFound error
	// instead of notFound error, it returns empty data
	if op.ID == nil {
		return nil, nil
	}

	return &models.Resource{
		ID:       *vCopy.ID,
		Name:     *vCopy.Name,
		Location: *vCopy.Location,
		Description: model.KeyVaultKeyDescription{
			Vault:         *vaultCopy,
			Key:           *vCopy,
			ResourceGroup: resourceGroupCopy,
		},
	}, nil
}

func KeyVault(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armkeyvault.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultsClient := clientFactory.NewVaultsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	maxResults := int32(100)
	options := &armkeyvault.VaultsClientListOptions{
		Top: &maxResults,
	}
	var values []models.Resource
	pager := vaultsClient.NewListPager(options)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			resource, err := getKeyVault(ctx, vault, vaultsClient, diagnosticClient)
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

func getKeyVault(ctx context.Context, vault *armkeyvault.Resource, vaultsClient *armkeyvault.VaultsClient, diagnosticClient *armmonitor.DiagnosticSettingsClient) (*models.Resource, error) {
	name := *vault.Name
	resourceGroup := strings.Split(*vault.ID, "/")[4]

	keyVaultGetOp, err := vaultsClient.Get(ctx, resourceGroup, name, nil)
	if err != nil {
		return nil, err
	}

	var insightsListOp []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*vault.ID, nil)
	if err != nil {
		return nil, err
	}
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		insightsListOp = append(insightsListOp, page.Value...)
	}

	resource := models.Resource{
		ID:       *vault.ID,
		Name:     *vault.Name,
		Location: *vault.Location,
		Description: model.KeyVaultDescription{
			Resource:                    *vault,
			Vault:                       keyVaultGetOp.Vault,
			DiagnosticSettingsResources: insightsListOp,
			ResourceGroup:               resourceGroup,
		},
	}
	return &resource, nil
}

func DeletedVault(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armkeyvault.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultsClient := clientFactory.NewVaultsClient()

	var values []models.Resource
	pager := vaultsClient.NewListDeletedPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			resource := getDeletedVault(ctx, vault)
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

func getDeletedVault(ctx context.Context, vault *armkeyvault.DeletedVault) *models.Resource {
	resourceGroup := strings.Split(*vault.ID, "/")[4]

	resource := models.Resource{
		ID:       *vault.ID,
		Name:     *vault.Name,
		Location: *vault.Properties.Location,
		Description: model.KeyVaultDeletedVaultDescription{
			Vault:         *vault,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func KeyVaultManagedHardwareSecurityModule(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	maxResults := int32(100)

	clientFactory, err := armkeyvault.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewManagedHsmsClient()

	options := &armkeyvault.ManagedHsmsClientListBySubscriptionOptions{
		Top: &maxResults,
	}
	pager := client.NewListBySubscriptionPager(options)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			resource, err := getKeyVaultManagedHardwareSecurityModule(ctx, diagnosticClient, vault)
			for err != nil {
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

func getKeyVaultManagedHardwareSecurityModule(ctx context.Context, client *armmonitor.DiagnosticSettingsClient, vault *armkeyvault.ManagedHsm) (*models.Resource, error) {
	resourceGroup := strings.Split(*vault.ID, "/")[4]

	var keyvaultListOp []*armmonitor.DiagnosticSettingsResource
	pager := client.NewListPager(*vault.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		keyvaultListOp = append(keyvaultListOp, page.Value...)
	}

	resource := models.Resource{
		ID:       *vault.ID,
		Name:     *vault.Name,
		Location: *vault.Location,
		Description: model.KeyVaultManagedHardwareSecurityModuleDescription{
			ManagedHsm:                  *vault,
			DiagnosticSettingsResources: keyvaultListOp,
			ResourceGroup:               resourceGroup,
		},
	}
	return &resource, nil
}

func KeyVaultKeyVersion(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armkeyvault.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultsClient := clientFactory.NewVaultsClient()
	keysClient := clientFactory.NewKeysClient()

	maxResults := int32(100)
	options := &armkeyvault.VaultsClientListOptions{
		Top: &maxResults,
	}
	pager := vaultsClient.NewListPager(options)

	wpe := concurrency.NewWorkPool(4)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			vault := v
			wpe.AddJob(func() (interface{}, error) {
				resourceGroup := strings.Split(*vault.ID, "/")[4]

				pager2 := keysClient.NewListPager(resourceGroup, *vault.Name, nil)
				var result []*armkeyvault.Key
				for pager2.More() {
					page2, err := pager2.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					result = append(result, page2.Value...)
				}
				wp := concurrency.NewWorkPool(8)
				for _, r := range result {
					resourceGroupCopy := resourceGroup
					vaultCopy := vault
					vCopy := r
					wp.AddJob(func() (interface{}, error) {
						resources, err := ListKeyVaultKeyVersion(ctx, keysClient, vCopy, resourceGroupCopy, vaultCopy)
						if err != nil {
							return nil, err
						}
						return resources, nil
					})
				}

				results := wp.Run()
				var vvv []models.Resource
				for _, r := range results {
					if r.Error != nil {
						return nil, err
					}
					if r.Value == nil {
						continue
					}
					vvv = append(vvv, r.Value.(models.Resource))
				}
				return vvv, nil
			})
		}
	}

	results := wpe.Run()
	for _, result := range results {
		if result.Error != nil {
			return nil, err
		}
		if result.Value == nil {
			continue
		}
		values = append(values, result.Value.([]models.Resource)...)
	}

	if stream != nil {
		for _, resource := range values {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		}
		values = nil
	}
	return values, nil
}

func ListKeyVaultKeyVersion(ctx context.Context, keysClient *armkeyvault.KeysClient, vCopy *armkeyvault.Key, resourceGroupCopy string, vaultCopy *armkeyvault.Resource) ([]models.Resource, error) {
	op, err := keysClient.Get(ctx, resourceGroupCopy, *vaultCopy.Name, *vCopy.Name, nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := keysClient.NewListVersionsPager(resourceGroupCopy, *vaultCopy.Name, *vCopy.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetKeyVaultKeyVersion(ctx, resourceGroupCopy, vaultCopy, vCopy, v)
			values = append(values, *resource)
		}
	}
	// In some cases resource does not give any notFound error
	// instead of notFound error, it returns empty data
	if op.ID == nil {
		return nil, nil
	}

	return values, nil
}

func GetKeyVaultKeyVersion(ctx context.Context, resourceGroup string, vault *armkeyvault.Resource, key *armkeyvault.Key, version *armkeyvault.Key) *models.Resource {
	resource := models.Resource{
		ID:       *version.ID,
		Name:     *version.Name,
		Location: *version.Location,
		Description: model.KeyVaultKeyVersionDescription{
			Vault:         *vault,
			Key:           *key,
			Version:       *version,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func KeyVaultCertificate(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armkeyvault.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultsClient := clientFactory.NewVaultsClient()

	maxResults := int32(100)
	options := &armkeyvault.VaultsClientListOptions{
		Top: &maxResults,
	}
	var values []models.Resource
	pager := vaultsClient.NewListPager(options)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			resource, err := getKeyVaultCertificates(ctx, cred, vault, vaultsClient)
			if err != nil {
				return nil, err
			}
			for _, res := range resource {
				if stream != nil {
					if err := (*stream)(res); err != nil {
						return nil, err
					}
				} else {
					values = append(values, res)
				}
			}
		}
	}
	return values, nil
}

func getKeyVaultCertificates(ctx context.Context, cred *azidentity.ClientSecretCredential, vault *armkeyvault.Resource, vaultsClient *armkeyvault.VaultsClient) ([]models.Resource, error) {
	name := *vault.Name
	resourceGroup := strings.Split(*vault.ID, "/")[4]

	keyVaultGetOp, err := vaultsClient.Get(ctx, resourceGroup, name, nil)
	if err != nil {
		return nil, err
	}

	client, err := azcertificates.NewClient(*keyVaultGetOp.Vault.Properties.VaultURI, cred, nil)
	if err != nil {
		return nil, err
	}
	var resources []models.Resource
	pager2 := client.NewListCertificatesPager(nil)
	for pager2.More() {
		page, err := pager2.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, c := range page.Value {
			policy, err := client.GetCertificatePolicy(ctx, c.ID.Name(), nil)
			if err != nil {
				return nil, err
			}
			resource := models.Resource{
				ID:       *vault.ID,
				Name:     *vault.Name,
				Location: *vault.Location,
				Description: model.KeyVaultCertificateDescription{
					Policy:        policy.CertificatePolicy,
					Vault:         *vault,
					ResourceGroup: resourceGroup,
				},
			}
			resources = append(resources, resource)
		}
	}

	return resources, nil
}
