package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func DataProtectionBackupVaults(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armdataprotection.NewBackupVaultsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewGetInSubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getDataProtectionBackupVaults(ctx, v)
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

func getDataProtectionBackupVaults(ctx context.Context, v *armdataprotection.BackupVaultResource) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.DataProtectionBackupVaultsDescription{
			BackupVaults:  *v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func DataProtectionBackupVaultsBackupPolicies(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armdataprotection.NewBackupVaultsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	policiesClient, err := armdataprotection.NewBackupPoliciesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewGetInSubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := getDataProtectionBackupVaultsBackupPolicies(ctx, policiesClient, v)
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

func getDataProtectionBackupVaultsBackupPolicies(ctx context.Context, client *armdataprotection.BackupPoliciesClient, v *armdataprotection.BackupVaultResource) ([]models.Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	pager := client.NewListPager(resourceGroup, *v.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, p := range page.Value {
			resourceGroup := strings.Split(*v.ID, "/")[4]

			resource := models.Resource{
				ID:       *p.ID,
				Name:     *p.Name,
				Location: *v.Location,
				Description: model.DataProtectionBackupVaultsBackupPoliciesDescription{
					BackupPolicies: *p,
					ResourceGroup:  resourceGroup,
				},
			}
			values = append(values, resource)
		}
	}
	return values, nil
}

func DataProtectionBackupJobs(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {

	client, err := armdataprotection.NewBackupVaultsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	jobsClient, err := armdataprotection.NewJobsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewGetInSubscriptionPager(nil)
	var resources []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			jobs, err := listDataProtectionBackupJobs(ctx, jobsClient, vault)
			if err != nil {
				return nil, err
			}
			for _, job := range jobs {
				if stream != nil {
					if err := (*stream)(job); err != nil {
						return nil, err
					}
				} else {
					resources = append(resources, job)
				}
			}
		}
	}
	return resources, nil
}

func listDataProtectionBackupJobs(ctx context.Context, jobsClient *armdataprotection.JobsClient, vault *armdataprotection.BackupVaultResource) ([]models.Resource, error) {

	resourceGroup := strings.Split(*vault.ID, "/")[4]

	pager := jobsClient.NewListPager(resourceGroup, *vault.Name, nil)

	var resources []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, job := range page.Value {
			resource := GetDataPotectionJob(ctx, vault, job)
			resources = append(resources, *resource)
		}
	}
	return resources, nil
}

func GetDataPotectionJob(ctx context.Context, vault *armdataprotection.BackupVaultResource, job *armdataprotection.AzureBackupJobResource) *models.Resource {

	resourceGroup := strings.Split(*job.ID, "/")[4]

	resource := models.Resource{
		ID:   *job.ID,
		Name: *job.Name,
		Description: model.DataProtectionJobDescription{
			DataProtectionJob: *job,
			VaultName:         *vault.Name,
			ResourceGroup:     resourceGroup,
		},
	}
	return &resource
}
