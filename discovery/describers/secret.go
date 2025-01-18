package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func KeyVaultSecret(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armkeyvault.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	vaultsClient := clientFactory.NewVaultsClient()
	secretsClient := clientFactory.NewSecretsClient()

	maxResults := int32(100)
	options := armkeyvault.VaultsClientListOptions{
		Top: &maxResults,
	}
	pager := vaultsClient.NewListPager(&options)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			//vaultURI := "https://" + *vault.Name + ".vault.azure.net/"
			maxResults := int32(25)
			rgs, err := listResourceGroups(ctx, cred, subscription)
			if err != nil {
				return nil, err
			}
			for _, rg := range rgs {
				options := armkeyvault.SecretsClientListOptions{
					Top: &maxResults,
				}
				var vaultName string
				splited := strings.Split(*vault.Name, "/")
				if len(splited) > 1 {
					vaultName = splited[8]
				} else {
					vaultName = *vault.Name
				}
				pager := secretsClient.NewListPager(*rg.Name, vaultName, &options)
				for pager.More() {
					page, err := pager.NextPage(ctx)
					if err != nil {
						if strings.Contains(err.Error(), "could not be found") {
							break
						}
						return nil, err
					}
					for _, sc := range page.Value {
						splitID := strings.Split(*sc.ID, "/")
						splitVaultID := strings.Split(*vault.ID, "/")
						akas := []string{"azure:///subscriptions/" + subscription + "/resourceGroups/" + splitVaultID[4] +
							"/providers/Microsoft.KeyVault/vaults/" + *vault.Name + "/secrets/" + splitID[4],
							"azure:///subscriptions/" + subscription + "/resourcegroups/" + splitVaultID[4] +
								"/providers/microsoft.keyvault/vaults/" + *vault.Name + "/secrets/" + splitID[4]}

						turbotData := map[string]interface{}{
							"SubscriptionId": subscription,
							"ResourceGroup":  splitVaultID[4],
							"Location":       vault.Location,
							"Akas":           akas,
						}

						name := *vault.Name
						resourceGroup := strings.Split(*vault.ID, "/")[4]

						keyVaultGetOp, err := vaultsClient.Get(ctx, resourceGroup, name, nil)
						if err != nil {
							return nil, err
						}

						resource := models.Resource{
							ID:       *sc.ID,
							Name:     *sc.ID,
							Location: "global",
							Description: model.KeyVaultSecretDescription{
								SecretItem:    *sc,
								Vault:         keyVaultGetOp.Vault,
								TurboData:     turbotData,
								ResourceGroup: *rg.Name,
							},
						}
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
	}
	return values, nil
}
