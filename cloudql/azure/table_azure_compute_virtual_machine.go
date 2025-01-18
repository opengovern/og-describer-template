package azure

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureComputeVirtualMachine(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_virtual_machine",
		Description: "Azure Compute Virtual Machine",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetComputeVirtualMachine,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeVirtualMachine,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Name")},
			{
				Name:        "power_state",
				Description: "Specifies the power state of the vm.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getPowerState),
			},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.ID")},
			{
				Name:        "type",
				Description: "The type of the resource in Azure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Type")},
			{
				Name:        "provisioning_state",
				Description: "The virtual machine provisioning state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.ProvisioningState")},
			{
				Name:        "vm_id",
				Description: "Specifies an unique ID for VM, which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.VMID"),
			},
			{
				Name:        "size",
				Description: "Specifies the size of the virtual machine.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Properties.HardwareProfile.VMSize"),
			},
			{
				Name:        "admin_user_name",
				Description: "Specifies the name of the administrator account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.AdminUsername")},
			{
				Name:        "allow_extension_operations",
				Description: "Specifies whether extension operations should be allowed on the virtual machine.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.AllowExtensionOperations")},
			{
				Name:        "availability_set_id",
				Description: "Specifies the ID of the availability set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.AvailabilitySet.ID")},
			{
				Name:        "billing_profile_max_price",
				Description: "Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.BillingProfile.MaxPrice")},
			{
				Name:        "boot_diagnostics_enabled",
				Description: "Specifies whether boot diagnostics should be enabled on the Virtual Machine, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.DiagnosticsProfile.BootDiagnostics.Enabled")},
			{
				Name:        "boot_diagnostics_storage_uri",
				Description: "Contains the Uri of the storage account to use for placing the console output and screenshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.DiagnosticsProfile.BootDiagnostics.StorageURI")},
			{
				Name:        "computer_name",
				Description: "Specifies the host OS name of the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.ComputerName")},
			{
				Name:        "disable_password_authentication",
				Description: "Specifies whether password authentication should be disabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.LinuxConfiguration.DisablePasswordAuthentication")},
			{
				Name:        "enable_automatic_updates",
				Description: "Indicates whether automatic updates is enabled for the windows virtual machine.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.WindowsConfiguration.EnableAutomaticUpdates")},
			{
				Name:        "eviction_policy",
				Description: "Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Properties.EvictionPolicy"),
			},
			{
				Name:        "image_exact_version",
				Description: "Specifies in decimal numbers, the version of platform image or marketplace image used to create the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.ImageReference.ExactVersion")},
			{
				Name:        "image_id",
				Description: "Specifies the ID of the image to use.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.ImageReference.ID")},
			{
				Name:        "image_offer",
				Description: "Specifies the offer of the platform image or marketplace image used to create the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.ImageReference.Offer")},
			{
				Name:        "image_publisher",
				Description: "Specifies the publisher of the image to use.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.ImageReference.Publisher")},
			{
				Name:        "image_sku",
				Description: "Specifies the sku of the image to use.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.ImageReference.SKU")},
			{
				Name:        "image_version",
				Description: "Specifies the version of the platform image or marketplace image used to create the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.ImageReference.Version")},
			{
				Name:        "managed_disk_id",
				Description: "Specifies the ID of the managed disk used by the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.OSDisk.ManagedDisk.ID")},
			{
				Name:        "os_disk_caching",
				Description: "Specifies the caching requirements of the operating system disk used by the virtual machine.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Properties.StorageProfile.OSDisk.Caching"),
			},
			{
				Name:        "os_disk_create_option",
				Description: "Specifies how the virtual machine should be created.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Properties.StorageProfile.OSDisk.CreateOption"),
			},
			{
				Name:        "os_disk_name",
				Description: "Specifies the name of the operating system disk used by the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.OSDisk.Name")},
			{
				Name:        "os_disk_vhd_uri",
				Description: "Specifies the virtual hard disk's uri.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Properties.StorageProfile.OSDisk.Vhd.URI"),
			},
			{
				Name:        "os_name",
				Description: "The Operating System running on the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineInstanceView.OSName")},
			{
				Name:        "os_version",
				Description: "The version of Operating System running on the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineInstanceView.OSVersion")},
			{
				Name:        "os_type",
				Description: "Specifies the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Properties.StorageProfile.OSDisk.OSType"),
			},
			{
				Name:        "priority",
				Description: "Specifies the priority for the virtual machine.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Properties.Priority"),
			},
			{
				Name:        "provision_vm_agent",
				Description: "Specifies whether virtual machine agent should be provisioned on the virtual machine for linux configuration.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.LinuxConfiguration.ProvisionVMAgent")},
			{
				Name:        "provision_vm_agent_windows",
				Description: "Specifies whether virtual machine agent should be provisioned on the virtual machine for windows configuration.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.WindowsConfiguration.ProvisionVMAgent")},
			{
				Name:        "require_guest_provision_signal",
				Description: "Specifies whether the guest provision signal is required to infer provision success of the virtual machine.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.RequireGuestProvisionSignal")},
			{
				Name:        "ultra_ssd_enabled",
				Description: "Specifies whether managed disks with storage account type UltraSSD_LRS can be added to a virtual machine or virtual machine scale set, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.AdditionalCapabilities.UltraSSDEnabled")},
			{
				Name:        "time_zone",
				Description: "Specifies the time zone of the virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.WindowsConfiguration.TimeZone")},
			{
				Name:        "additional_unattend_content",
				Description: "Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by windows setup.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.WindowsConfiguration.AdditionalUnattendContent")},
			{
				Name:        "data_disks",
				Description: "A list of parameters that are used to add a data disk to a virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageProfile.DataDisks")},
			{
				Name:        "linux_configuration_ssh_public_keys",
				Description: "A list of ssh key configuration for a Linux OS",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.LinuxConfiguration.SSH.PublicKeys")},
			{
				Name:        "network_interfaces",
				Description: "A list of resource Ids for the network interfaces associated with the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.NetworkProfile.NetworkInterfaces")},
			{
				Name:        "patch_settings",
				Description: "Specifies settings related to in-guest patching (KBs).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.WindowsConfiguration.PatchSettings")},
			{
				Name:        "private_ips",
				Description: "An array of private ip addesses associated with the vm.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getPrivateIps),
			},
			{
				Name:        "public_ips",
				Description: "An array of public ip addesses associated with the vm.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PublicIPs")},
			{
				Name:        "secrets",
				Description: "A list of certificates that should be installed onto the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.Secrets")},
			{
				Name:        "statuses",
				Description: "Specifies the resource status information.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineInstanceView.Statuses")},
			{
				Name:        "extensions",
				Description: "Specifies the details of VM Extensions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineExtension")},
			{
				Name:        "extensions_settings",
				Description: "Specifies the details of VM Extensions settings map.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExtensionsSettings")},
			{
				Name:        "guest_configuration_assignments",
				Description: "Guest configuration assignments for a virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assignments")},
			{
				Name:        "identity",
				Description: "The identity of the virtual machine, if configured.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Identity")},
			{
				Name:        "security_profile",
				Description: "Specifies the security related profile settings for the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.SecurityProfile")},
			{
				Name:        "win_rm",
				Description: "Specifies the windows remote management listeners. This enables remote windows powershell.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.OSProfile.WindowsConfiguration.WinRM")},
			{
				Name:        "zones",
				Description: "A list of virtual machine zones.",
				Type:        proto.ColumnType_JSON,

				// Standard steampipe columns
				Transform: transform.FromField("Description.VirtualMachine.Zones")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Standard azure columns
				// Azure standard columns

				Transform: transform.FromField("Description.VirtualMachine.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION ////

				//// LIST FUNCTION ////
				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

// TRANSFORM FUNCTIONS

func getPowerState(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getPowerState", "d.Value", d.Value)
	if d.HydrateItem == nil {
		return nil, nil
	}
	vm := d.HydrateItem.(opengovernance.ComputeVirtualMachine).Description
	statuses := vm.VirtualMachineInstanceView.Statuses
	return getStatusFromCode(statuses, "PowerState"), nil
}

func getPrivateIps(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}

	interfaceIPConfigurations := d.HydrateItem.(opengovernance.ComputeVirtualMachine).Description.InterfaceIPConfigurations

	var ips []string
	for _, ipConfig := range interfaceIPConfigurations {
		ips = append(ips, *ipConfig.Properties.PrivateIPAddress)
	}

	return ips, nil
}

// UTILITY FUNCTIONS
func getStatusFromCode(statuses []*armcompute.InstanceViewStatus, codeType string) string {
	for _, status := range statuses {
		statusCode := types.SafeString(status.Code)

		if strings.HasPrefix(statusCode, codeType+"/") {
			return strings.SplitN(statusCode, "/", 2)[1]
		}
	}
	return ""
}
