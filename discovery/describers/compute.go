package describers

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"

	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
	"github.com/turbot/go-kit/types"
)

func ComputeDisk(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDisksClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeDisk(ctx, v)
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

func getComputeDisk(ctx context.Context, v *armcompute.Disk) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	return &models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeDiskDescription{
			Disk:          *v,
			ResourceGroup: resourceGroup,
		},
	}
}

func ComputeDiskAccess(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDiskAccessesClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeDiskAccess(ctx, v)
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

func getComputeDiskAccess(ctx context.Context, v *armcompute.DiskAccess) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	return &models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeDiskAccessDescription{
			DiskAccess:    *v,
			ResourceGroup: resourceGroup,
		},
	}
}

func ComputeVirtualMachineScaleSet(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewVirtualMachineScaleSetsClient()
	clientExtension := clientFactory.NewVirtualMachineScaleSetExtensionsClient()

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getComputeVirtualMachineScaleSet(ctx, clientExtension, v)
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

func getComputeVirtualMachineScaleSet(ctx context.Context, clientExtension *armcompute.VirtualMachineScaleSetExtensionsClient, v *armcompute.VirtualMachineScaleSet) (*models.Resource, error) {
	resourceGroupName := strings.Split(*v.ID, "/")[4]

	var op []armcompute.VirtualMachineScaleSetExtension
	pages := clientExtension.NewListPager(resourceGroupName, *v.Name, nil)
	for pages.More() {
		page, err := pages.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			op = append(op, *v)
		}
	}

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeVirtualMachineScaleSetDescription{
			VirtualMachineScaleSet:           *v,
			VirtualMachineScaleSetExtensions: op,
			ResourceGroup:                    resourceGroupName,
		},
	}
	return &resource, nil
}

func ComputeVirtualMachineScaleSetNetworkInterface(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewVirtualMachineScaleSetsClient()

	networkClient, err := armnetwork.NewInterfacesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vm := range page.Value {
			vmResourceGroupName := strings.Split(*vm.ID, "/")[4]
			pager := networkClient.NewListVirtualMachineScaleSetNetworkInterfacesPager(vmResourceGroupName, *vm.Name, nil)
			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, v := range page.Value {
					resource := getComputeVirtualMachineScaleSetNetworkInterface(ctx, vm, v)

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
		}
	}
	return values, nil
}

func getComputeVirtualMachineScaleSetNetworkInterface(ctx context.Context, vm *armcompute.VirtualMachineScaleSet, v *armnetwork.Interface) *models.Resource {
	resourceGroupName := strings.Split(*v.ID, "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeVirtualMachineScaleSetNetworkInterfaceDescription{
			VirtualMachineScaleSet: *vm,
			NetworkInterface:       *v,
			ResourceGroup:          resourceGroupName,
		},
	}
	return &resource
}

func ComputeVirtualMachineScaleSetVm(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	scaleSetsClient := clientFactory.NewVirtualMachineScaleSetsClient()
	scaleSetVMsClient := clientFactory.NewVirtualMachineScaleSetVMsClient()

	pager := scaleSetsClient.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vm := range page.Value {
			vmResourceGroupName := strings.Split(*vm.ID, "/")[4]
			pager := scaleSetVMsClient.NewListPager(vmResourceGroupName, *vm.Name, nil)
			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, v := range page.Value {
					resource := getComputeVirtualMachineScaleSetVm(ctx, vm, v)
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

func getComputeVirtualMachineScaleSetVm(ctx context.Context, vm *armcompute.VirtualMachineScaleSet, v *armcompute.VirtualMachineScaleSetVM) *models.Resource {
	resourceGroupName := strings.Split(*v.ID, "/")[4]

	powerState := getStatusFromCode(v.Properties.InstanceView.Statuses, "PowerState")

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeVirtualMachineScaleSetVmDescription{
			VirtualMachineScaleSet: *vm,
			ScaleSetVM:             *v,
			PowerState:             powerState,
			ResourceGroup:          resourceGroupName,
		},
	}
	return &resource
}

func getStatusFromCode(statuses []*armcompute.InstanceViewStatus, codeType string) string {
	for _, status := range statuses {
		statusCode := types.SafeString(status.Code)

		if strings.HasPrefix(statusCode, codeType+"/") {
			return strings.SplitN(statusCode, "/", 2)[1]
		}
	}
	return ""
}

func ComputeVirtualMachine(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vmClient := clientFactory.NewVirtualMachinesClient()
	vmExtensionsClient := clientFactory.NewVirtualMachineExtensionsClient()

	networkInterfaceClient, err := armnetwork.NewInterfacesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	networkPublicIPClient, err := armnetwork.NewPublicIPAddressesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	ipConfigClient, err := armnetwork.NewInterfaceIPConfigurationsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	guestConfigurationClientFactory, err := armguestconfiguration.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	guestConfigurationClient := guestConfigurationClientFactory.NewAssignmentsClient()

	pager := vmClient.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, virtualMachine := range page.Value {
			resource, err := getComputeVirtualMachine(ctx, vmClient, vmExtensionsClient, networkInterfaceClient, networkPublicIPClient, ipConfigClient, guestConfigurationClient, virtualMachine)
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

func getComputeVirtualMachine(ctx context.Context, vmClient *armcompute.VirtualMachinesClient, vmExtensionsClient *armcompute.VirtualMachineExtensionsClient, networkInterfaceClient *armnetwork.InterfacesClient, networkPublicIPClient *armnetwork.PublicIPAddressesClient, ipConfigClient *armnetwork.InterfaceIPConfigurationsClient, guestConfigurationClient *armguestconfiguration.AssignmentsClient, virtualMachine *armcompute.VirtualMachine) (*models.Resource, error) {
	resourceGroupName := strings.Split(*virtualMachine.ID, "/")[4]
	computeInstanceViewOp, err := vmClient.InstanceView(ctx, resourceGroupName, *virtualMachine.Name, nil)

	var ipConfigs = make([]armnetwork.InterfaceIPConfiguration, 0, 0)
	if virtualMachine.Properties.VirtualMachineScaleSet != nil && virtualMachine.Properties.VirtualMachineScaleSet.ID != nil {
		vmstateName := filepath.Base(*virtualMachine.Properties.VirtualMachineScaleSet.ID)
		pager := networkInterfaceClient.NewListVirtualMachineScaleSetNetworkInterfacesPager(
			resourceGroupName, vmstateName, nil)
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				if strings.Contains(err.Error(), "ERROR CODE: NotFound") {
					continue
				}
				return nil, err
			}
			for _, n := range page.Value {
				ipPager := ipConfigClient.NewListPager(resourceGroupName, *n.Name, nil)
				for ipPager.More() {
					ipPage, err := ipPager.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					for _, ip := range ipPage.Value {
						ipConfigs = append(ipConfigs, *ip)
					}
				}
			}
		}
	}

	var publicIPs []string
	for _, ipConfig := range ipConfigs {
		if ipConfig.Properties.PublicIPAddress != nil && ipConfig.Properties.PublicIPAddress.ID != nil {
			pathParts := strings.Split(*ipConfig.Properties.PublicIPAddress.ID, "/")
			resourceGroup := pathParts[4]
			name := pathParts[len(pathParts)-1]

			publicIP, err := networkPublicIPClient.Get(ctx, resourceGroup, name, nil)

			if err != nil {
				return nil, err
			}
			if publicIP.Properties.IPAddress != nil {
				publicIPs = append(publicIPs, *publicIP.Properties.IPAddress)
			}
		}
	}

	computeListOp, err := vmExtensionsClient.List(ctx, resourceGroupName, *virtualMachine.Name, nil)
	if err != nil {
		return nil, err
	}

	var configurationListOp []armguestconfiguration.Assignment
	guestPager := guestConfigurationClient.NewListPager(resourceGroupName, *virtualMachine.Name, nil)
	for guestPager.More() {
		page, err := guestPager.NextPage(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "NotFound") {
				break
			}
			return nil, err
		}
		for _, v := range page.Value {
			configurationListOp = append(configurationListOp, *v)
		}
	}

	for idx, ex := range computeInstanceViewOp.VirtualMachineInstanceView.Extensions {
		for sidx, substatus := range ex.Substatuses {
			if substatus == nil {
				continue
			}
			substatus.Message = nil
			ex.Substatuses[sidx] = substatus
		}
		computeInstanceViewOp.VirtualMachineInstanceView.Extensions[idx] = ex
	}
	extensionsSettings := make(map[string]map[string]interface{})
	for _, ex := range computeListOp.Value {
		extensionsSettings[*ex.ID] = ex.Properties.Settings.(map[string]interface{})
	}

	resource := models.Resource{
		ID:       *virtualMachine.ID,
		Name:     *virtualMachine.Name,
		Location: *virtualMachine.Location,
		Description: model.ComputeVirtualMachineDescription{
			VirtualMachine:             *virtualMachine,
			VirtualMachineInstanceView: computeInstanceViewOp.VirtualMachineInstanceView,
			InterfaceIPConfigurations:  ipConfigs,
			PublicIPs:                  publicIPs,
			VirtualMachineExtension:    computeListOp.Value,
			ExtensionsSettings:         extensionsSettings,
			Assignments:                &configurationListOp,
			ResourceGroup:              resourceGroupName,
		},
	}
	return &resource, nil
}

func ComputeSnapshots(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSnapshotsClient()
	pager := client.NewListPager(nil)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeSnapshot(ctx, v)
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

func getComputeSnapshot(ctx context.Context, snapshot *armcompute.Snapshot) *models.Resource {
	resourceGroupName := strings.Split(*snapshot.ID, "/")[4]

	resource := models.Resource{
		ID:       *snapshot.ID,
		Name:     *snapshot.Name,
		Location: *snapshot.Location,
		Description: model.ComputeSnapshotsDescription{
			ResourceGroup: resourceGroupName,
			Snapshot:      *snapshot,
		},
	}
	return &resource
}

func ComputeAvailabilitySet(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAvailabilitySetsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeAvailabilitySet(ctx, v)
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

func getComputeAvailabilitySet(ctx context.Context, availabilitySet *armcompute.AvailabilitySet) *models.Resource {
	resourceGroupName := strings.Split(*availabilitySet.ID, "/")[4]

	resource := models.Resource{
		ID:       *availabilitySet.ID,
		Name:     *availabilitySet.Name,
		Location: *availabilitySet.Location,
		Description: model.ComputeAvailabilitySetDescription{
			ResourceGroup:   resourceGroupName,
			AvailabilitySet: *availabilitySet,
		},
	}
	return &resource
}

func ComputeDiskEncryptionSet(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDiskEncryptionSetsClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeDiskEncryptionSet(ctx, v)
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

func getComputeDiskEncryptionSet(ctx context.Context, diskEncryptionSet *armcompute.DiskEncryptionSet) *models.Resource {
	resourceGroupName := strings.Split(*diskEncryptionSet.ID, "/")[4]

	resource := models.Resource{
		ID:       *diskEncryptionSet.ID,
		Name:     *diskEncryptionSet.Name,
		Location: *diskEncryptionSet.Location,
		Description: model.ComputeDiskEncryptionSetDescription{
			ResourceGroup:     resourceGroupName,
			DiskEncryptionSet: *diskEncryptionSet,
		},
	}
	return &resource
}

func ComputeGallery(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewGalleriesClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeGallery(ctx, v)
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

func getComputeGallery(ctx context.Context, gallery *armcompute.Gallery) *models.Resource {
	resourceGroupName := strings.Split(*gallery.ID, "/")[4]

	resource := models.Resource{
		ID:       *gallery.ID,
		Name:     *gallery.Name,
		Location: *gallery.Location,
		Description: model.ComputeImageGalleryDescription{
			ResourceGroup: resourceGroupName,
			ImageGallery:  *gallery,
		},
	}
	return &resource
}

func ComputeImage(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewImagesClient()

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeImage(ctx, v)
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

func getComputeImage(ctx context.Context, v *armcompute.Image) *models.Resource {
	resourceGroup := strings.ToLower(strings.Split(*v.ID, "/")[4])
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeImageDescription{
			Image:         *v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func ComputeHostGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDedicatedHostGroupsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeHostGroup(ctx, v)
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

func getComputeHostGroup(ctx context.Context, v *armcompute.DedicatedHostGroup) *models.Resource {
	resourceGroup := strings.ToLower(strings.Split(*v.ID, "/")[4])
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeHostGroupDescription{
			HostGroup:     *v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func ComputeHost(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDedicatedHostGroupsClient()
	hostClient := clientFactory.NewDedicatedHostsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := getComputeHostsByGroup(ctx, hostClient, v)
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

func getComputeHostsByGroup(ctx context.Context, hostClient *armcompute.DedicatedHostsClient, v *armcompute.DedicatedHostGroup) ([]models.Resource, error) {
	resourceGroup := strings.ToLower(strings.Split(*v.ID, "/")[4])

	pager := hostClient.NewListByHostGroupPager(resourceGroup, *v.Name, nil)
	var resources []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, host := range page.Value {
			resource := models.Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: *v.Location,
				Description: model.ComputeHostGroupHostDescription{
					Host:          *host,
					ResourceGroup: resourceGroup,
				},
			}
			resources = append(resources, resource)
		}
	}
	return resources, nil
}

func ComputeRestorePointCollection(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewRestorePointCollectionsClient()

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeResourcePointCollection(ctx, v)
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

func getComputeResourcePointCollection(ctx context.Context, v *armcompute.RestorePointCollection) *models.Resource {
	resourceGroup := strings.ToLower(strings.Split(*v.ID, "/")[4])
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeRestorePointCollectionDescription{
			RestorePointCollection: *v,
			ResourceGroup:          resourceGroup,
		},
	}
	return &resource
}

func ComputeSSHPublicKey(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSSHPublicKeysClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeSSHPublicKey(ctx, v)
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

func getComputeSSHPublicKey(ctx context.Context, v *armcompute.SSHPublicKeyResource) *models.Resource {
	resourceGroup := strings.ToLower(strings.Split(*v.ID, "/")[4])
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeSSHPublicKeyDescription{
			SSHPublicKey:  *v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}

func ComputeDiskReadOps(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDisksClient()

	pager := client.NewListPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, disk := range page.Value {
			if disk.ID == nil {
				continue
			}
			resources, err := getComputeDiskReadOps(ctx, cred, subscription, disk)
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

func getComputeDiskReadOps(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, disk *armcompute.Disk) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "FIVE_MINUTES", "Microsoft.Compute/disks", "Composite Disk Read Operations/sec", *disk.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_readops", *disk.ID),
			Name:     fmt.Sprintf("%s readops", *disk.Name),
			Location: *disk.Location,
			Description: model.ComputeDiskReadOpsDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}

func ComputeDiskReadOpsDaily(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDisksClient()

	pager := client.NewListPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, disk := range page.Value {
			if disk.ID == nil {
				continue
			}
			resources, err := getComputeDiskReadOpsDaily(ctx, cred, subscription, disk)
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

func getComputeDiskReadOpsDaily(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, disk *armcompute.Disk) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "DAILY", "Microsoft.Compute/disks", "Composite Disk Read Operations/sec", *disk.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_readops_daily", *disk.ID),
			Name:     fmt.Sprintf("%s readops-daily", *disk.Name),
			Location: *disk.Location,
			Description: model.ComputeDiskReadOpsDailyDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}
func ComputeDiskReadOpsHourly(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDisksClient()

	pager := client.NewListPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, disk := range page.Value {
			if disk.ID == nil {
				continue
			}
			resources, err := getComputeDiskReadOpsHourly(ctx, cred, subscription, disk)
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

func getComputeDiskReadOpsHourly(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, disk *armcompute.Disk) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "HOURLY", "Microsoft.Compute/disks", "Composite Disk Read Operations/sec", *disk.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_readops_hourly", *disk.ID),
			Name:     fmt.Sprintf("%s readops-hourly", *disk.Name),
			Location: *disk.Location,
			Description: model.ComputeDiskReadOpsHourlyDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}

func ComputeDiskWriteOps(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDisksClient()

	pager := client.NewListPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, disk := range page.Value {
			if disk.ID == nil {
				continue
			}
			resources, err := getComputeDiskWriteOps(ctx, cred, subscription, disk)
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

func getComputeDiskWriteOps(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, disk *armcompute.Disk) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "FIVE_MINUTES", "Microsoft.Compute/disks", "Composite Disk Write Operations/sec", *disk.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_writeops", *disk.ID),
			Name:     fmt.Sprintf("%s writeops", *disk.Name),
			Location: *disk.Location,
			Description: model.ComputeDiskReadOpsDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}

func ComputeDiskWriteOpsDaily(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDisksClient()

	pager := client.NewListPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, disk := range page.Value {
			if disk.ID == nil {
				continue
			}
			resources, err := getComputeDiskWriteOpsDaily(ctx, cred, subscription, disk)
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

func getComputeDiskWriteOpsDaily(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, disk *armcompute.Disk) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "DAILY", "Microsoft.Compute/disks", "Composite Disk Write Operations/sec", *disk.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_writeops_daily", *disk.ID),
			Name:     fmt.Sprintf("%s writeops-daily", *disk.Name),
			Location: *disk.Location,
			Description: model.ComputeDiskReadOpsDailyDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}
func ComputeDiskWriteOpsHourly(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDisksClient()

	pager := client.NewListPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, disk := range page.Value {
			if disk.ID == nil {
				continue
			}
			resources, err := getComputeDiskWriteOpsHourly(ctx, cred, subscription, disk)
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

func getComputeDiskWriteOpsHourly(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, disk *armcompute.Disk) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "HOURLY", "Microsoft.Compute/disks", "Composite Disk Write Operations/sec", *disk.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_writeops_hourly", *disk.ID),
			Name:     fmt.Sprintf("%s writeops-hourly", *disk.Name),
			Location: *disk.Location,
			Description: model.ComputeDiskReadOpsHourlyDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}

func ComputeResourceSKU(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewResourceSKUsClient()
	pager := client.NewListPager(nil)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeResourceSKU(ctx, subscription, v)
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

func getComputeResourceSKU(ctx context.Context, subscription string, resourceSku *armcompute.ResourceSKU) *models.Resource {
	resource := models.Resource{
		Description: model.ComputeResourceSKUDescription{
			ResourceSKU: *resourceSku,
		},
	}
	if resourceSku.Locations != nil && len(resourceSku.Locations) > 0 {
		resource.Location = *(resourceSku.Locations)[0]
		if resourceSku.Name != nil {
			resource.ID = "azure:///subscriptions/" + subscription + "/locations/" + *(resourceSku.Locations)[0] + "/resourcetypes" + *resourceSku.ResourceType + "name/" + *resourceSku.Name
		}
	}
	if resourceSku.Name != nil {
		resource.Name = *resourceSku.Name
	}
	return &resource
}

func ComputeVirtualMachineCpuUtilization(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewVirtualMachinesClient()

	pager := client.NewListAllPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, virtualMachine := range page.Value {
			if virtualMachine.ID == nil {
				continue
			}
			resources, err := getComputeVirtualMachineCpuUtilization(ctx, cred, subscription, virtualMachine)
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

func getComputeVirtualMachineCpuUtilization(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, virtualMachine *armcompute.VirtualMachine) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "FIVE_MINUTES", "Microsoft.Compute/virtualMachines", "Percentage CPU", *virtualMachine.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_cpu_utilization", *virtualMachine.ID),
			Name:     fmt.Sprintf("%s cpu-utilization", *virtualMachine.Name),
			Location: *virtualMachine.Location,
			Description: model.ComputeVirtualMachineCpuUtilizationDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}

func ComputeVirtualMachineCpuUtilizationDaily(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewVirtualMachinesClient()

	pager := client.NewListAllPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, virtualMachine := range page.Value {
			if virtualMachine.ID == nil {
				continue
			}
			resources, err := getComputeVirtualMachineCpuUtilizationDaily(ctx, cred, subscription, virtualMachine)
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

func getComputeVirtualMachineCpuUtilizationDaily(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, virtualMachine *armcompute.VirtualMachine) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "DAILY", "Microsoft.Compute/virtualMachines", "Percentage CPU", *virtualMachine.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_cpu_utilization_daily", *virtualMachine.ID),
			Name:     fmt.Sprintf("%s cpu-utilization-daily", *virtualMachine.Name),
			Location: *virtualMachine.Location,
			Description: model.ComputeVirtualMachineCpuUtilizationDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}

func ComputeVirtualMachineCpuUtilizationHourly(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewVirtualMachinesClient()

	pager := client.NewListAllPager(nil)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, virtualMachine := range page.Value {
			if virtualMachine.ID == nil {
				continue
			}
			resources, err := getComputeVirtualMachineCpuUtilizationHourly(ctx, cred, subscription, virtualMachine)
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

func getComputeVirtualMachineCpuUtilizationHourly(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, virtualMachine *armcompute.VirtualMachine) ([]models.Resource, error) {
	metrics, err := listAzureMonitorMetricStatistics(ctx, cred, subscription, "HOURLY", "Microsoft.Compute/virtualMachines", "Percentage CPU", *virtualMachine.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, metric := range metrics {
		resource := models.Resource{
			ID:       fmt.Sprintf("%s_cpu_utilization_hourly", *virtualMachine.ID),
			Name:     fmt.Sprintf("%s cpu-utilization-hourly", *virtualMachine.Name),
			Location: *virtualMachine.Location,
			Description: model.ComputeVirtualMachineCpuUtilizationDescription{
				MonitoringMetric: metric,
			},
		}
		values = append(values, resource)
	}
	return values, nil
}

func ComputeCloudServices(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewCloudServicesClient()

	pager := client.NewListAllPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getComputeCloudServices(ctx, v)
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

func getComputeCloudServices(ctx context.Context, v *armcompute.CloudService) *models.Resource {
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.ComputeCloudServiceDescription{
			CloudService: *v,
		},
	}
	return &resource
}
