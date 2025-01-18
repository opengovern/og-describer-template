package describers

import (
	"context"
	"fmt"
	"strings"


	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	azblobOld "github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/opengovern/og-util/pkg/concurrency"

	"github.com/Azure/go-autorest/autorest"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/queue/queues"
	"github.com/tombuildsstuff/giovanni/storage/2019-12-12/blob/accounts"
)

func StorageContainer(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstorage.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewBlobContainersClient()
	storageClient := clientFactory.NewAccountsClient()

	wpe := concurrency.NewWorkPool(4)
	var values []models.Resource
	pager := storageClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, ac := range page.Value {
			account := ac
			wpe.AddJob(func() (interface{}, error) {
				results, err := ListAccountStorageContainers(ctx, client, account)
				if err != nil {
					return nil, err
				}
				return results, nil
			})
		}
	}

	results := wpe.Run()
	for _, r := range results {
		if r.Error != nil {
			return nil, err
		}
		if r.Value == nil {
			continue
		}
		values = append(values, r.Value.([]models.Resource)...)
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

func ListAccountStorageContainers(ctx context.Context, client *armstorage.BlobContainersClient, account *armstorage.Account) ([]models.Resource, error) {
	resourceGroup := &strings.Split(string(*account.ID), "/")[4]
	var resources []models.Resource
	pager := client.NewListPager(*resourceGroup, *account.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vl := range page.Value {
			resource, err := GetAccountStorageContainter(ctx, client, vl, account)
			if err != nil {
				return nil, err
			}
			resources = append(resources, *resource)
		}
	}
	return resources, nil
}

func GetAccountStorageContainter(ctx context.Context, client *armstorage.BlobContainersClient, v *armstorage.ListContainerItem, acc *armstorage.Account) (*models.Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	accountName := strings.Split(*v.ID, "/")[8]

	op, err := client.GetImmutabilityPolicy(ctx, resourceGroup, accountName, *v.Name, nil)
	if err != nil {
		return nil, err
	}

	return &models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: model.StorageContainerDescription{
			AccountName:        *acc.Name,
			ListContainerItem:  *v,
			ImmutabilityPolicy: op.ImmutabilityPolicy,
			ResourceGroup:      resourceGroup,
		},
	}, nil
}

func StorageAccount(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {

	clientFactory, err := armstorage.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	encryptionScopesStorageClient := clientFactory.NewEncryptionScopesClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	fileServicesStorageClient := clientFactory.NewFileServicesClient()

	blobServicesStorageClient := clientFactory.NewBlobServicesClient()

	managementPoliciesStorageClient := clientFactory.NewManagementPoliciesClient()

	storageClient := clientFactory.NewAccountsClient()

	pager := storageClient.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			resource, err := GetStorageAccount(ctx, storageClient, encryptionScopesStorageClient, diagnosticClient, fileServicesStorageClient, blobServicesStorageClient, managementPoliciesStorageClient, account)
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

func GetStorageAccount(ctx context.Context, storageClient *armstorage.AccountsClient, encryptionScopesStorageClient *armstorage.EncryptionScopesClient, diagnosticClient *armmonitor.DiagnosticSettingsClient, fileServicesStorageClient *armstorage.FileServicesClient, blobServicesStorageClient *armstorage.BlobServicesClient, managementPoliciesStorageClient *armstorage.ManagementPoliciesClient, account *armstorage.Account) (*models.Resource, error) {
	resourceGroup := &strings.Split(*account.ID, "/")[4]

	var managementPolicy *armstorage.ManagementPolicy
	storageGetOp, err := managementPoliciesStorageClient.Get(ctx, *resourceGroup, *account.Name, armstorage.ManagementPolicyNameDefault, nil)
	if err != nil {
		if !strings.Contains(err.Error(), "ManagementPolicyNotFound") {
			return nil, err
		}
	} else {
		managementPolicy = &storageGetOp.ManagementPolicy
	}

	var blobServicesProperties *armstorage.BlobServiceProperties
	if *account.Kind != "FileStorage" {
		blobServicesPropertiesOp, err := blobServicesStorageClient.GetServiceProperties(ctx, *resourceGroup, *account.Name, nil)
		if err != nil {
			if !strings.Contains(err.Error(), "ContainerOperationFailure") {
				return nil, err
			}
		} else {
			blobServicesProperties = &blobServicesPropertiesOp.BlobServiceProperties
		}
	}

	var logging *accounts.Logging
	if *account.Kind != "FileStorage" {
		v, err := storageClient.ListKeys(ctx, *resourceGroup, *account.Name, nil)
		if err != nil {
			if !strings.Contains(err.Error(), "ScopeLocked") && !strings.Contains(err.Error(), "ReadOnlyDisabledSubscription") {
				return nil, err
			}
		} else {
			if v.Keys != nil || len(v.Keys) > 0 {
				key := (v.Keys)[0]

				storageAuth, err := autorest.NewSharedKeyAuthorizer(*account.Name, *key.Value, autorest.SharedKeyLite)
				if err != nil {
					return nil, err
				}

				client := accounts.New()
				client.Client.Authorizer = storageAuth

				resp, err := client.GetServiceProperties(ctx, *account.Name)
				if err != nil {
					if !strings.Contains(err.Error(), "FeatureNotSupportedForAccount") {
						return nil, err
					}
				} else {
					logging = resp.StorageServiceProperties.Logging
				}
			}
		}
	}

	var storageGetServicePropertiesOp *armstorage.FileServiceProperties
	if *account.Kind != "BlobStorage" {
		v, err := fileServicesStorageClient.GetServiceProperties(ctx, *resourceGroup, *account.Name, nil)
		if err != nil {
			if !strings.Contains(err.Error(), "AccountIsDisabled") {
				return nil, err
			}
		}
		storageGetServicePropertiesOp = &v.FileServiceProperties
	}

	var diagSettingsOp []*armmonitor.DiagnosticSettingsResource
	pager1 := diagnosticClient.NewListPager(*account.ID, nil)
	for pager1.More() {
		page1, err := pager1.NextPage(ctx)
		if err != nil {
			break
		}
		diagSettingsOp = append(diagSettingsOp, page1.Value...)
	}

	var vsop []*armstorage.EncryptionScope
	pager2 := encryptionScopesStorageClient.NewListPager(*resourceGroup, *account.Name, nil)
	for pager2.More() {
		page2, err := pager2.NextPage(ctx)
		if err != nil {
			break
		}
		vsop = append(vsop, page2.Value...)
	}

	var storageProperties *queues.StorageServiceProperties
	if *account.SKU.Tier == "Standard" && (*account.Kind == "Storage" || *account.Kind == "StorageV2") {
		accountKeys, err := storageClient.ListKeys(ctx, *resourceGroup, *account.Name, nil)
		if err != nil {
			if !strings.Contains(err.Error(), "ReadOnlyDisabledSubscription") {
				return nil, err
			}
		} else {
			if accountKeys.Keys != nil || len(accountKeys.Keys) > 0 {
				key := (accountKeys.Keys)[0]
				storageAuth, err := autorest.NewSharedKeyAuthorizer(*account.Name, *key.Value, autorest.SharedKeyLite)
				if err != nil {
					return nil, err
				}

				queuesClient := queues.New()
				queuesClient.Client.Authorizer = storageAuth

				resp, err := queuesClient.GetServiceProperties(ctx, *account.Name)

				if err != nil {
					if !strings.Contains(err.Error(), "ReadOnlyDisabledSubscription") {
						return nil, err
					}
				} else {
					storageProperties = &resp.StorageServiceProperties
				}
			}
		}
	}

	v, err := storageClient.ListKeys(ctx, *resourceGroup, *account.Name, nil)
	if err != nil {
		return nil, err
	}

	var tableProperties aztables.ServiceProperties
	if *account.Kind != "FileStorage" {

		for _, key := range v.Keys {
			serviceUrl := "https://" + *account.Name + ".table.core.windows.net/"

			auth, err := aztables.NewSharedKeyCredential(*account.Name, *key.Value)
			if err != nil {
				return nil, err
			}

			client, _ := aztables.NewServiceClientWithSharedKey(serviceUrl, auth, nil)

			op, err := client.GetProperties(ctx, &aztables.GetPropertiesOptions{})
			if err != nil {
				return nil, err
			} else {
				tableProperties = op.ServiceProperties
				break
			}

		}
	}

	var keysMap []map[string]interface{}
	if len(v.Keys) > 0 {
		for _, key := range v.Keys {
			keyMap := make(map[string]interface{})
			if key.KeyName != nil {
				keyMap["KeyName"] = *key.KeyName
			}
			if key.Value != nil {
				keyMap["Value"] = *key.Value
			}
			if key.Permissions != nil {
				keyMap["Permissions"] = key.Permissions
			}
			keysMap = append(keysMap, keyMap)
		}
	}

	resource := models.Resource{
		ID:       *account.ID,
		Name:     *account.Name,
		Location: *account.Location,
		Description: model.StorageAccountDescription{
			Account:                     *account,
			ManagementPolicy:            managementPolicy,
			BlobServiceProperties:       blobServicesProperties,
			Logging:                     logging,
			StorageServiceProperties:    storageProperties,
			FileServiceProperties:       storageGetServicePropertiesOp,
			DiagnosticSettingsResources: diagSettingsOp,
			EncryptionScopes:            vsop,
			TableProperties:             tableProperties,
			AccessKeys:                  keysMap,
			ResourceGroup:               *resourceGroup,
		},
	}
	return &resource, nil
}

func StorageBlob(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstorage.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	accountClient := clientFactory.NewAccountsClient()
	containerClient := clientFactory.NewBlobContainersClient()

	pager := accountClient.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := ListAccountStorageBlobs(ctx, containerClient, accountClient, v)
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

func ListAccountStorageBlobs(ctx context.Context, containerClient *armstorage.BlobContainersClient, accountClient *armstorage.AccountsClient, storageAccount *armstorage.Account) ([]models.Resource, error) {
	if storageAccount == nil || storageAccount.ID == nil {
		return nil, nil
	}
	resourceGroup := strings.Split(*storageAccount.ID, "/")[4]

	storageAccountName := *storageAccount.ID
	if storageAccount.Name != nil {
		storageAccountName = *storageAccount.Name
	}
	keys, err := accountClient.ListKeys(ctx, resourceGroup, storageAccountName, nil)
	if err != nil {
		return nil, err
	}

	if keys.Keys == nil || len(keys.Keys) == 0 || keys.Keys[0].Value == nil {
		return nil, nil
	}
	credential, err := azblob.NewSharedKeyCredential(*storageAccount.Name, *(keys.Keys)[0].Value)
	if err != nil {
		return nil, err
	}
	baseUrl := fmt.Sprintf("https://%s.blob.core.windows.net", storageAccountName)
	blobClient, err := azblob.NewClientWithSharedKeyCredential(baseUrl, credential, nil)
	if err != nil {
		return nil, err
	}

	pager := containerClient.NewListPager(resourceGroup, storageAccountName, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, container := range page.Value {
			if container.Name == nil {
				continue
			}
			blobPager := blobClient.NewListBlobsFlatPager(*container.Name, nil)
			for blobPager.More() {
				flatResponse, err := blobPager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, blob := range flatResponse.Segment.BlobItems {
					metadata := azblobOld.Metadata{}
					for k, v := range blob.Metadata {
						if v == nil {
							continue
						}
						metadata[k] = *v
					}

					blobTags := &azblobOld.BlobTags{
						BlobTagSet: []azblobOld.BlobTag{},
					}
					if blob.BlobTags != nil {
						for _, tag := range blob.BlobTags.BlobTagSet {
							if tag.Key == nil || tag.Value == nil {
								continue
							}
							blobTags.BlobTagSet = append(blobTags.BlobTagSet, azblobOld.BlobTag{
								Key:   *tag.Key,
								Value: *tag.Value,
							})
						}
					} else {
						blobTags = nil
					}

					desc := model.StorageBlobDescription{
						Blob: azblobOld.BlobItemInternal{
							Name:             *blob.Name,
							VersionID:        blob.VersionID,
							IsCurrentVersion: blob.IsCurrentVersion,
							Properties: azblobOld.BlobPropertiesInternal{
								CreationTime:              blob.Properties.CreationTime,
								ContentLength:             blob.Properties.ContentLength,
								ContentType:               blob.Properties.ContentType,
								ContentEncoding:           blob.Properties.ContentEncoding,
								ContentLanguage:           blob.Properties.ContentLanguage,
								ContentMD5:                blob.Properties.ContentMD5,
								ContentDisposition:        blob.Properties.ContentDisposition,
								CacheControl:              blob.Properties.CacheControl,
								BlobSequenceNumber:        blob.Properties.BlobSequenceNumber,
								CopyID:                    blob.Properties.CopyID,
								CopySource:                blob.Properties.CopySource,
								CopyProgress:              blob.Properties.CopyProgress,
								CopyCompletionTime:        blob.Properties.CopyCompletionTime,
								CopyStatusDescription:     blob.Properties.CopyStatusDescription,
								ServerEncrypted:           blob.Properties.ServerEncrypted,
								IncrementalCopy:           blob.Properties.IncrementalCopy,
								DestinationSnapshot:       blob.Properties.DestinationSnapshot,
								DeletedTime:               blob.Properties.DeletedTime,
								RemainingRetentionDays:    blob.Properties.RemainingRetentionDays,
								AccessTierInferred:        blob.Properties.AccessTierInferred,
								CustomerProvidedKeySha256: blob.Properties.CustomerProvidedKeySHA256,
								EncryptionScope:           blob.Properties.EncryptionScope,
								AccessTierChangeTime:      blob.Properties.AccessTierChangeTime,
								TagCount:                  blob.Properties.TagCount,
								ExpiresOn:                 blob.Properties.ExpiresOn,
								IsSealed:                  blob.Properties.IsSealed,
								LastAccessedOn:            blob.Properties.LastAccessedOn,
							},
							Metadata: metadata,
							BlobTags: blobTags,
						},
						AccountName:   *storageAccount.Name,
						ContainerName: *container.Name,
						ResourceGroup: resourceGroup,
					}
					if blob.Deleted != nil {
						desc.Blob.Deleted = *blob.Deleted
					}
					if blob.Snapshot != nil {
						desc.Blob.Snapshot = *blob.Snapshot
						desc.IsSnapshot = len(*blob.Snapshot) > 0
					}
					if blob.Properties.LastModified != nil {
						desc.Blob.Properties.LastModified = *blob.Properties.LastModified
					}
					if blob.Properties.ETag != nil {
						desc.Blob.Properties.Etag = azblobOld.ETag(*blob.Properties.ETag)
					}
					if blob.Properties.BlobType != nil {
						desc.Blob.Properties.BlobType = azblobOld.BlobType(*blob.Properties.BlobType)
					}
					if blob.Properties.LeaseStatus != nil {
						desc.Blob.Properties.LeaseStatus = azblobOld.LeaseStatusType(*blob.Properties.LeaseStatus)
					}
					if blob.Properties.LeaseState != nil {
						desc.Blob.Properties.LeaseState = azblobOld.LeaseStateType(*blob.Properties.LeaseState)
					}
					if blob.Properties.LeaseDuration != nil {
						desc.Blob.Properties.LeaseDuration = azblobOld.LeaseDurationType(*blob.Properties.LeaseDuration)
					}
					if blob.Properties.CopyStatus != nil {
						desc.Blob.Properties.CopyStatus = azblobOld.CopyStatusType(*blob.Properties.CopyStatus)
					}
					if blob.Properties.AccessTier != nil {
						desc.Blob.Properties.AccessTier = azblobOld.AccessTierType(*blob.Properties.AccessTier)
					}
					if blob.Properties.ArchiveStatus != nil {
						desc.Blob.Properties.ArchiveStatus = azblobOld.ArchiveStatusType(*blob.Properties.ArchiveStatus)
					}
					if blob.Properties.RehydratePriority != nil {
						desc.Blob.Properties.RehydratePriority = azblobOld.RehydratePriorityType(*blob.Properties.RehydratePriority)
					}

					resource := models.Resource{
						ID:          *blob.Name,
						Name:        *blob.Name,
						Location:    *storageAccount.Location,
						Description: desc,
					}
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func StorageBlobService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstorage.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	accountClient := clientFactory.NewAccountsClient()
	storageClient := clientFactory.NewBlobServicesClient()

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	pager := accountClient.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			for _, resourceGroup := range resourceGroups {
				var blobServices []*armstorage.BlobServiceProperties
				pager := storageClient.NewListPager(*resourceGroup.Name, *account.Name, nil)
				for pager.More() {
					page, err := pager.NextPage(ctx)
					if err != nil {
						if strings.Contains(err.Error(), "ParentResourceNotFound") ||
							strings.Contains(err.Error(), "ContainerOperationFailure") ||
							strings.Contains(err.Error(), "FeatureNotSupportedForAccount") {
							continue
						}
						return nil, err
					}
					blobServices = append(blobServices, page.Value...)
				}

				for _, blobService := range blobServices {
					resource := GetStorageBlobService(ctx, account, resourceGroup, blobService)
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
	}
	return values, nil
}

func GetStorageBlobService(ctx context.Context, account *armstorage.Account, resourceGroup armresources.ResourceGroup, blobService *armstorage.BlobServiceProperties) *models.Resource {
	resource := models.Resource{
		ID:       *blobService.ID,
		Name:     *blobService.Name,
		Location: *account.Location,
		Description: model.StorageBlobServiceDescription{
			BlobService:   *blobService,
			AccountName:   *account.Name,
			Location:      *account.Location,
			ResourceGroup: *resourceGroup.Name,
		},
	}
	return &resource
}

func StorageQueue(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstorage.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	accountClient := clientFactory.NewAccountsClient()
	storageClient := clientFactory.NewQueueClient()

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	pager := accountClient.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			for _, resourceGroup := range resourceGroups {
				resources, err := ListAccountStorageQueue(ctx, storageClient, account, resourceGroup)
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

func ListAccountStorageQueue(ctx context.Context, storageClient *armstorage.QueueClient, account *armstorage.Account, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	var values []models.Resource
	pager := storageClient.NewListPager(*resourceGroup.Name, *account.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			/*
			* For storage account type 'Page Blob' we are getting the kind value as 'StorageV2'.
			* Storage account type 'Page Blob' does not support table, so we are getting 'FeatureNotSupportedForAccount'/'OperationNotAllowedOnKind' error.
			* With same kind(StorageV2) of storage account, we my have different type(File Share) of storage account so we need to handle this particular error.
			 */
			if strings.Contains(err.Error(), "FeatureNotSupportedForAccount") ||
				strings.Contains(err.Error(), "AccountIsDisabled") ||
				strings.Contains(err.Error(), "ParentResourceNotFound") ||
				strings.Contains(err.Error(), "OperationNotAllowedOnKind") {
				continue
			} else {
				return nil, err
			}
		}
		for _, queue := range page.Value {
			resource := GetStorageQueue(ctx, account, resourceGroup, queue)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetStorageQueue(ctx context.Context, account *armstorage.Account, resourceGroup armresources.ResourceGroup, queue *armstorage.ListQueue) *models.Resource {
	resource := models.Resource{
		ID:       *queue.ID,
		Name:     *queue.Name,
		Location: *account.Location,
		Description: model.StorageQueueDescription{
			Queue:         *queue,
			AccountName:   *account.Name,
			Location:      *account.Location,
			ResourceGroup: *resourceGroup.Name,
		},
	}
	return &resource
}

func StorageFileShare(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstorage.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	accountClient := clientFactory.NewAccountsClient()
	storageClient := clientFactory.NewFileSharesClient()

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	pager := accountClient.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			for _, resourceGroup := range resourceGroups {
				resources, err := ListAccountStorageFileShares(ctx, storageClient, account, resourceGroup)
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

func ListAccountStorageFileShares(ctx context.Context, storageClient *armstorage.FileSharesClient, account *armstorage.Account, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	pager := storageClient.NewListPager(*resourceGroup.Name, *account.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "ParentResourceNotFound") ||
				strings.Contains(err.Error(), "FeatureNotSupportedForAccount") ||
				strings.Contains(err.Error(), "AccountIsDisabled") {
				continue
			}
			return nil, err
		}
		for _, fileShareItem := range page.Value {
			resource := GetStorageFileShares(ctx, account, resourceGroup, fileShareItem)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetStorageFileShares(ctx context.Context, account *armstorage.Account, resourceGroup armresources.ResourceGroup, fileShareItem *armstorage.FileShareItem) *models.Resource {
	resource := models.Resource{
		ID:       *fileShareItem.ID,
		Name:     *fileShareItem.Name,
		Location: *account.Location,
		Description: model.StorageFileShareDescription{
			FileShare:     *fileShareItem,
			AccountName:   *account.Name,
			Location:      *account.Location,
			ResourceGroup: *resourceGroup.Name,
		},
	}
	return &resource
}

func StorageTable(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstorage.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	accountClient := clientFactory.NewAccountsClient()
	storageClient := clientFactory.NewTableClient()

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := accountClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			if *account.Kind == "FileStorage" || *account.Kind == "BlockBlobStorage" {
				continue
			}
			for _, resourceGroup := range resourceGroups {
				resources, err := ListAccountStorageTables(ctx, storageClient, account, resourceGroup)
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

func ListAccountStorageTables(ctx context.Context, storageClient *armstorage.TableClient, account *armstorage.Account, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	pager := storageClient.NewListPager(*resourceGroup.Name, *account.Name, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			/*
			* For storage account type 'Page Blob' we are getting the kind value as 'StorageV2'.
			* Storage account type 'Page Blob' does not support table, so we are getting 'FeatureNotSupportedForAccount'/'OperationNotAllowedOnKind' error.
			* With same kind(StorageV2) of storage account, we my have different type(File Share) of storage account so we need to handle this particular error.
			 */
			if strings.Contains(err.Error(), "FeatureNotSupportedForAccount") ||
				strings.Contains(err.Error(), "OperationNotAllowedOnKind") ||
				strings.Contains(err.Error(), "AccountIsDisabled") ||
				strings.Contains(err.Error(), "ParentResourceNotFound") {
				continue
			}
			return nil, err
		}
		for _, table := range page.Value {
			resource := GetStorageTable(ctx, account, resourceGroup, table)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetStorageTable(ctx context.Context, account *armstorage.Account, resourceGroup armresources.ResourceGroup, table *armstorage.Table) *models.Resource {
	resource := models.Resource{
		ID:       *table.ID,
		Name:     *table.Name,
		Location: *account.Location,
		Description: model.StorageTableDescription{
			Table:         *table,
			AccountName:   *account.Name,
			Location:      *account.Location,
			ResourceGroup: *resourceGroup.Name,
		},
	}
	return &resource
}

func StorageTableService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armstorage.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	accountClient := clientFactory.NewAccountsClient()
	storageClient := clientFactory.NewTableServicesClient()

	resourceGroups, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	pager := accountClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			if *account.Kind == "FileStorage" {
				continue
			}
			for _, resourceGroup := range resourceGroups {
				resources, err := ListAccountStorageTableService(ctx, storageClient, account, resourceGroup)
				if err != nil {
					return nil, err
				}
				if resources == nil {
					continue
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

func ListAccountStorageTableService(ctx context.Context, storageClient *armstorage.TableServicesClient, account *armstorage.Account, resourceGroup armresources.ResourceGroup) ([]models.Resource, error) {
	tableServices, err := storageClient.List(ctx, *resourceGroup.Name, *account.Name, nil)
	if err != nil {
		if strings.Contains(err.Error(), "ParentResourceNotFound") ||
			strings.Contains(err.Error(), "FeatureNotSupportedForAccount") {
			return nil, nil
		}
		return nil, err
	}
	var values []models.Resource
	for _, tableService := range tableServices.Value {
		resource := GetAccountStorageTableService(ctx, resourceGroup, account, tableService)
		values = append(values, *resource)
	}
	return values, nil
}

func GetAccountStorageTableService(ctx context.Context, resourceGroup armresources.ResourceGroup, account *armstorage.Account, tableService *armstorage.TableServiceProperties) *models.Resource {
	resource := models.Resource{
		ID:       *tableService.ID,
		Name:     *tableService.Name,
		Location: *account.Location,
		Description: model.StorageTableServiceDescription{
			TableService:  *tableService,
			AccountName:   *account.Name,
			Location:      *account.Location,
			ResourceGroup: *resourceGroup.Name,
		},
	}
	return &resource
}
