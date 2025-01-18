package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeVirtualMachineScaleSetNetworkInterface(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_virtual_machine_scale_set_network_interface",
		Description: "Azure Compute Virtual Machine Scale Set Network Interface",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeVirtualMachineScaleSetNetworkInterface,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the scale set network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Name")},
			{
				Name:        "scale_set_name",
				Description: "Name of the scale set.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachineScaleSet.Name")},
			{
				Name:        "id",
				Description: "The unique ID identifying the resource in a subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.ID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the network interface resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.ProvisioningState"),
			},
			{
				Name:        "type",
				Description: "The type of the resource in Azure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Type")},
			{
				Name:        "mac_address",
				Description: "The MAC address of the network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.MacAddress")},
			{
				Name:        "enable_accelerated_networking",
				Description: "If the network interface has accelerated networking enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.EnableAcceleratedNetworking"),
			},
			{
				Name:        "enable_ip_forwarding",
				Description: "Indicates whether IP forwarding is enabled on this network interface.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.EnableIPForwarding"),
			},
			{
				Name:        "primary",
				Description: "Whether this is a primary network interface on a virtual machine.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.Primary"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the network interface resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.ResourceGUID"),
			},
			{
				Name:        "dns_settings",
				Description: "The DNS settings in network interface.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.DNSSettings"),
			},
			{
				Name:        "hosted_workloads",
				Description: "A list of references to linked BareMetal resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.HostedWorkloads"),
			},
			{
				Name:        "ip_configurations",
				Description: "A list of IP configurations of the network interface.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.IPConfigurations"),
			},
			{
				Name:        "network_security_group",
				Description: "The reference to the NetworkSecurityGroup resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.NetworkSecurityGroup"),
			},
			{
				Name:        "private_endpoint",
				Description: "A reference to the private endpoint to which the network interface is linked.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.PrivateEndpoint"),
			},
			{
				Name:        "virtual_machine",
				Description: "The reference to a virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractScaleSetNetworkInterfaceProperties, "Description.NetworkInterface.VirtualMachine"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkInterface.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkInterface.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.NetworkInterface.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.NetworkInterface.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// TRANSFORM FUNCTION
				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

func extractScaleSetNetworkInterfaceProperties(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	networkInterface := d.HydrateItem.(opengovernance.ComputeVirtualMachineScaleSetNetworkInterface).Description.NetworkInterface
	param := d.Param.(string)

	objectMap := make(map[string]interface{})

	if networkInterface.Properties.VirtualMachine != nil {
		objectMap["VirtualMachine"] = *networkInterface.Properties.VirtualMachine
	}
	if networkInterface.Properties.ResourceGUID != nil && *networkInterface.Properties.ResourceGUID != "" {
		objectMap["ResourceGUID"] = networkInterface.Properties.ResourceGUID
	}
	if networkInterface.Properties.ProvisioningState != nil {
		if *networkInterface.Properties.ProvisioningState != "" {
			objectMap["ProvisioningState"] = networkInterface.Properties.ProvisioningState
		}
	}
	if networkInterface.Properties.NetworkSecurityGroup != nil {
		objectMap["NetworkSecurityGroup"] = networkInterface.Properties.NetworkSecurityGroup
	}
	if networkInterface.Properties.IPConfigurations != nil {
		objectMap["IPConfigurations"] = networkInterface.Properties.IPConfigurations
	}
	if networkInterface.Properties.TapConfigurations != nil {
		objectMap["TapConfigurations"] = networkInterface.Properties.TapConfigurations
	}
	if networkInterface.Properties.DNSSettings != nil {
		objectMap["DNSSettings"] = networkInterface.Properties.DNSSettings
	}
	if networkInterface.Properties.MacAddress != nil {
		objectMap["MacAddress"] = networkInterface.Properties.MacAddress
	}
	if networkInterface.Properties.Primary != nil {
		objectMap["Primary"] = networkInterface.Properties.Primary
	}
	if networkInterface.Properties.EnableAcceleratedNetworking != nil {
		objectMap["EnableAcceleratedNetworking"] = networkInterface.Properties.EnableAcceleratedNetworking
	}
	if networkInterface.Properties.HostedWorkloads != nil {
		objectMap["HostedWorkloads"] = networkInterface.Properties.HostedWorkloads
	}
	if networkInterface.Properties.HostedWorkloads != nil {
		objectMap["HostedWorkloads"] = networkInterface.Properties.HostedWorkloads
	}

	if val, ok := objectMap[param]; ok {
		return val, nil
	}
	return nil, nil
}
