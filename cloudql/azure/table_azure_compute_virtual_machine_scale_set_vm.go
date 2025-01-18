package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeVirtualMachineScaleSetVm(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_virtual_machine_scale_set_vm",
		Description: "Azure Compute Virtual Machine Scale Set VM",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"scale_set_name", "resource_group", "instance_id"}),
			Hydrate:    opengovernance.GetComputeVirtualMachineScaleSetVm,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeVirtualMachineScaleSetVm,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the scale set VM.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScaleSetVM.Name")},
			{
				Name:        "scale_set_name",
				Description: "Name of the scale set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Name")},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.ID"),
			},
			{
				Name:        "instance_id",
				Description: "The virtual machine instance ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScaleSetVM.InstanceID")},
			{
				Name:        "latest_model_applied",
				Description: "Specifies whether the latest model has been applied to the virtual machine.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.LatestModelApplied"),
			},
			{
				Name:        "power_state",
				Description: "Specifies the power state of the VM.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PowerState"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource in Azure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScaleSetVM.Type")},
			{
				Name:        "license_type",
				Description: "Specifies that the image or disk that is being used was licensed on-premises.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.LicenseType"),
			},
			{
				Name:        "location",
				Description: "The location of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScaleSetVM.Location")},
			{
				Name:        "model_definition_applied",
				Description: "Specifies whether the model applied to the virtual machine is the model of the virtual machine scale set or the customized model for the virtual machine.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ScaleSetVM.Properties.ModelDefinitionApplied")},
			{
				Name:        "sku_name",
				Description: "The sku name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScaleSetVM.SKU.Name")},
			{
				Name:        "sku_capacity",
				Description: "Specifies the capacity of virtual machines in a scale set virtual machine.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ScaleSetVM.SKU.Capacity")},
			{
				Name:        "sku_tier",
				Description: "Specifies the tier of virtual machines in a scale set virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScaleSetVM.SKU.Tier")},
			{
				Name:        "vm_id",
				Description: "Azure virtual machine unique ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScaleSetVM.Properties.VMID"),
			},
			{
				Name:        "additional_capabilities",
				Description: "Specifies additional capabilities enabled or disabled on the virtual machine in the scale set. For instance: whether the virtual machine has the capability to support attaching managed data disks with UltraSSD_LRS storage account type.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.AdditionalCapabilities")},
			{
				Name:        "availability_set",
				Description: "Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScaleSetVM.Properties.AvailabilitySet"),
			},
			{
				Name:        "plan",
				Description: "Specifies information about the marketplace image used to create the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScaleSetVM.Plan")},
			{
				Name:        "protection_policy",
				Description: "Specifies the protection policy of the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.SpotRestorePolicy")},
			{
				Name:        "resources",
				Description: "The virtual machine child extension resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScaleSetVM.Resources")},
			{
				Name:        "tags_src",
				Description: "Resource tags.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScaleSetVM.Tags")},
			{
				Name:        "upgrade_policy",
				Description: "The upgrade policy for the scale set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.UpgradePolicy")},
			{
				Name:        "virtual_machine_diagnostics_profile",
				Description: "Specifies the boot diagnostic settings state.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.DiagnosticsProfile")},
			{
				Name:        "virtual_machine_hardware_profile",
				Description: "Specifies the hardware settings for the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.HardwareProfile")},
			{
				Name:        "virtual_machine_network_profile",
				Description: "Specifies properties of the network interfaces of the virtual machines.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScaleSetVM.Properties.NetworkProfile.NetworkInterfaces")},
			{
				Name:        "virtual_machine_network_profile_configuration",
				Description: "Specifies the network profile configuration of the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScaleSetVM.Properties.NetworkProfileConfiguration")},
			{
				Name:        "virtual_machine_os_profile",
				Description: "Specifies the operating system settings for the virtual machines.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.OSProfile")},
			{
				Name:        "virtual_machine_security_profile",
				Description: "Specifies the Security related profile settings for the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.SecurityProfile")},
			{
				Name:        "virtual_machine_storage_profile",
				Description: "SSpecifies the storage settings for the virtual machine disks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.StorageProfile")},
			{
				Name:        "zones",
				Description: "The Logical zone list for scale set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScaleSetVM.Zones")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ScaleSetVM.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ScaleSetVM.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ScaleSetVM.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ScaleSetVM.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
