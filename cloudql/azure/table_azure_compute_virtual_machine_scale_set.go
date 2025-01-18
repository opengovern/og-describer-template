package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeVirtualMachineScaleSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_virtual_machine_scale_set",
		Description: "Azure Compute Virtual Machine Scale Set",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetComputeVirtualMachineScaleSet,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeVirtualMachineScaleSet,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the scale set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Name")},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.ID")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource in Azure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Type")},
			{
				Name:        "location",
				Description: "The location of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Location")},
			{
				Name:        "do_not_run_extensions_on_overprovisioned_vms",
				Description: "When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.DoNotRunExtensionsOnOverprovisionedVMs")},
			{
				Name:        "overprovision",
				Description: "Specifies whether the Virtual Machine Scale Set should be overprovisioned.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.Overprovision")},
			{
				Name:        "platform_fault_domain_count",
				Description: "Fault Domain count for each placement group.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.PlatformFaultDomainCount")},
			{
				Name:        "single_placement_group",
				Description: "When true this limits the scale set to a single placement group, of max size 100 virtual machines.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.SinglePlacementGroup")},
			{
				Name:        "sku_name",
				Description: "The sku name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.SKU.Name")},
			{
				Name:        "sku_capacity",
				Description: "Specifies the tier of virtual machines in a scale set.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.SKU.Capacity")},
			{
				Name:        "sku_tier",
				Description: "Specifies the tier of virtual machines in a scale set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.SKU.Tier")},
			{
				Name:        "unique_id",
				Description: "Specifies the ID which uniquely identifies a Virtual Machine Scale Set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.UniqueID")},
			{
				Name:        "extensions",
				Description: "Specifies the details of VM Scale Set Extensions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSetExtensions")},
			{
				Name:        "identity",
				Description: "The identity of the virtual machine scale set, if configured.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Identity")},
			{
				Name:        "plan",
				Description: "Specifies information about the marketplace image used to create the virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Plan")},
			{
				Name:        "scale_in_policy",
				Description: "Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.ScaleInPolicy")},
			{
				Name:        "tags_src",
				Description: "Resource tags.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Tags")},
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
				Name:        "virtual_machine_extension_profile",
				Description: "Specifies a collection of settings for extensions installed on virtual machines in the scale set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.ExtensionProfile")},
			{
				Name:        "virtual_machine_network_profile",
				Description: "Specifies properties of the network interfaces of the virtual machines in the scale set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.NetworkProfile")},
			{
				Name:        "virtual_machine_os_profile",
				Description: "Specifies the operating system settings for the virtual machines in the scale set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.OSProfile")},
			{
				Name:        "virtual_machine_storage_profile",
				Description: "Specifies the storage settings for the virtual machine disks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.StorageProfile")},
			{
				Name:        "virtual_machine_security_profile",
				Description: "Specifies the Security related profile settings for the virtual machines in the scale set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Properties.VirtualMachineProfile.SecurityProfile")},
			{
				Name:        "zones",
				Description: "The Logical zone list for scale set.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.VirtualMachineScaleSet.Zones")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachineScaleSet.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.VirtualMachineScaleSet.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachineScaleSet.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				//// LIST FUNCTION
				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTION

// If we return the API response directly, the output only gives the contents of VirtualMachineScaleSetExtensionsListResult
