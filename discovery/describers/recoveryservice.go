package describers

import (
	"context"
	"reflect"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func RecoveryServicesVault(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armrecoveryservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewVaultsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	var values []models.Resource
	pager := client.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			resource, err := GetRecoveryServicesVault(ctx, diagnosticClient, vault)
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

func GetRecoveryServicesVault(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, vault *armrecoveryservices.Vault) (*models.Resource, error) {
	resourceGroup := strings.Split(*vault.ID, "/")[4]

	var diagnostic []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*vault.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		diagnostic = append(diagnostic, page.Value...)
	}

	resource := models.Resource{
		ID:       *vault.ID,
		Name:     *vault.Name,
		Location: *vault.Location,
		Description: model.RecoveryServicesVaultDescription{
			Vault:                      *vault,
			DiagnosticSettingsResource: diagnostic,
			ResourceGroup:              resourceGroup,
		},
	}
	return &resource, nil
}

func RecoveryServicesBackupJobs(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	vaultClientFactory, err := armrecoveryservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultClient := vaultClientFactory.NewVaultsClient()

	clientFactory, err := armrecoveryservicesbackup.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewBackupJobsClient()

	var values []models.Resource
	pager := vaultClient.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			if vault.ID == nil || vault.Name == nil {
				continue
			}
			resourceGroup := strings.Split(*vault.ID, "/")[4]
			vaultBackupJobs, err := ListRecoveryServicesVaultBackupJobs(ctx, client, *vault.Name, resourceGroup)
			if err != nil {
				return nil, err
			}
			for _, resource := range vaultBackupJobs {
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

func ListRecoveryServicesVaultBackupJobs(ctx context.Context, client *armrecoveryservicesbackup.BackupJobsClient, vaultName, resourceGroup string) ([]models.Resource, error) {
	pager := client.NewListPager(vaultName, resourceGroup, nil)
	var resources []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, job := range page.Value {
			resource, err := GetRecoveryServicesBackupJob(resourceGroup, vaultName, job)
			if err != nil {
				return nil, err
			}
			resources = append(resources, *resource)
		}
	}
	return resources, nil
}

func GetRecoveryServicesBackupJob(resourceGroup, vaultName string, job *armrecoveryservicesbackup.JobResource) (*models.Resource, error) {
	properties, err := backupJobProperties(job)
	if err != nil {
		return nil, err
	}
	resource := models.Resource{
		Description: model.RecoveryServicesBackupJobDescription{
			Job: struct {
				Name     *string
				ID       *string
				Type     *string
				ETag     *string
				Tags     map[string]*string
				Location *string
			}{
				Name:     job.Name,
				ID:       job.ID,
				Location: job.Location,
				Type:     job.Type,
				Tags:     job.Tags,
				ETag:     job.ETag,
			},
			VaultName:     vaultName,
			Properties:    properties,
			ResourceGroup: resourceGroup,
		},
	}
	if job.ID != nil {
		resource.ID = *job.ID
	}
	if job.Name != nil {
		resource.Name = *job.Name
	}
	if job.Location != nil {
		resource.Location = *job.Location
	}
	return &resource, nil
}

func backupJobProperties(data *armrecoveryservicesbackup.JobResource) (map[string]interface{}, error) {
	output := make(map[string]interface{})

	if data.Properties != nil {
		if data.Properties.GetJob() != nil {
			if data.Properties.GetJob().ActivityID != nil {
				output["ActivityID"] = data.Properties.GetJob().ActivityID
			}
			if data.Properties.GetJob().BackupManagementType != nil {
				output["BackupManagementType"] = data.Properties.GetJob().BackupManagementType
			}
			if data.Properties.GetJob().JobType != nil {
				output["JobType"] = data.Properties.GetJob().JobType
			}
			if data.Properties.GetJob().EndTime != nil {
				output["EndTime"] = data.Properties.GetJob().EndTime
			}
			if data.Properties.GetJob().EntityFriendlyName != nil {
				output["EntityFriendlyName"] = data.Properties.GetJob().EntityFriendlyName
			}
			if data.Properties.GetJob().Operation != nil {
				output["Operation"] = data.Properties.GetJob().Operation
			}
			if data.Properties.GetJob().StartTime != nil {
				output["StartTime"] = data.Properties.GetJob().StartTime
			}
			if data.Properties.GetJob().Status != nil {
				output["Status"] = data.Properties.GetJob().Status
			}
		}
	}
	return output, nil
}

func RecoveryServicesBackupPolicies(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	vaultClientFactory, err := armrecoveryservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultClient := vaultClientFactory.NewVaultsClient()

	clientFactory, err := armrecoveryservicesbackup.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewBackupPoliciesClient()

	var values []models.Resource
	pager := vaultClient.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			if vault.ID == nil || vault.Name == nil {
				continue
			}
			resourceGroup := strings.Split(*vault.ID, "/")[4]
			vaultBackupJobs, err := ListRecoveryServicesVaultBackupPolicies(ctx, client, *vault.Name, resourceGroup)
			if err != nil {
				return nil, err
			}
			for _, resource := range vaultBackupJobs {
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

func ListRecoveryServicesVaultBackupPolicies(ctx context.Context, client *armrecoveryservicesbackup.BackupPoliciesClient, vaultName, resourceGroup string) ([]models.Resource, error) {
	pager := client.NewListPager(vaultName, resourceGroup, nil)
	var resources []models.Resource

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, policy := range page.Value {
			resource := GetRecoveryServicesBackupPolicy(policy, vaultName, resourceGroup)
			if err != nil {
				return nil, err
			}
			resources = append(resources, resource)
		}
	}
	return resources, nil
}

func GetRecoveryServicesBackupPolicy(policy *armrecoveryservicesbackup.ProtectionPolicyResource, vaultName, resourceGroup string) models.Resource {
	return models.Resource{
		Description: model.RecoveryServicesBackupPolicyDescription{
			ResourceGroup: resourceGroup,
			VaultName:     vaultName,
			Policy: struct {
				Name     *string
				ID       *string
				Type     *string
				ETag     *string
				Tags     map[string]*string
				Location *string
			}{
				Name:     policy.Name,
				ID:       policy.ID,
				Location: policy.Location,
				Type:     policy.Type,
				Tags:     policy.Tags,
				ETag:     policy.ETag,
			},
			Properties: extractData(policy.Properties),
		},
	}
}

func RecoveryServicesBackupItem(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	vaultClientFactory, err := armrecoveryservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultClient := vaultClientFactory.NewVaultsClient()

	clientFactory, err := armrecoveryservicesbackup.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewBackupProtectedItemsClient()

	var values []models.Resource
	pager := vaultClient.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			if vault.ID == nil || vault.Name == nil {
				continue
			}
			resourceGroup := strings.Split(*vault.ID, "/")[4]
			vaultBackupJobs, err := ListRecoveryServicesVaultBackupItems(ctx, client, *vault.Name, resourceGroup)
			if err != nil {
				return nil, err
			}
			for _, resource := range vaultBackupJobs {
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

func ListRecoveryServicesVaultBackupItems(ctx context.Context, client *armrecoveryservicesbackup.BackupProtectedItemsClient, vaultName, resourceGroup string) ([]models.Resource, error) {
	pager := client.NewListPager(vaultName, resourceGroup, nil)
	var resources []models.Resource

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.Value {
			resource := GetRecoveryServicesBackupItem(item, vaultName, resourceGroup)
			resources = append(resources, resource)
		}
	}
	return resources, nil
}

func GetRecoveryServicesBackupItem(item *armrecoveryservicesbackup.ProtectedItemResource, vaultName, resourceGroup string) models.Resource {
	return models.Resource{
		Description: model.RecoveryServicesBackupItemDescription{
			ResourceGroup: resourceGroup,
			VaultName:     vaultName,
			Item: struct {
				Name     *string
				ID       *string
				Type     *string
				ETag     *string
				Tags     map[string]*string
				Location *string
			}{
				Name:     item.Name,
				ID:       item.ID,
				Location: item.Location,
				Type:     item.Type,
				Tags:     item.Tags,
				ETag:     item.ETag,
			},
			Properties: extractData(item.Properties),
		},
	}
}

func extractData(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(obj)

	// If the object is a pointer, we need to dereference it
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Check if v is a struct
	if v.Kind() != reflect.Struct {
		return result
	}

	// Iterate over the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		// Only deal with exported fields
		if fieldType.PkgPath != "" {
			continue
		}

		// If the field is a pointer, dereference it
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				result[fieldType.Name] = nil
				continue
			}
			field = field.Elem()
		}

		// If the field is a struct, recursively extract its data
		if field.Kind() == reflect.Struct {
			result[fieldType.Name] = extractData(field.Interface())
		} else {
			result[fieldType.Name] = field.Interface()
		}
	}

	return result
}
