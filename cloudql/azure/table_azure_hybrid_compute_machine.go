package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureHybridComputeMachine(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_hybrid_compute_machine",
		Description: "Azure Hybrid Compute Machine",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetHybridComputeMachine,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListHybridComputeMachine,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.ID")},
			{
				Name:        "status",
				Description: "The status of the hybrid machine agent.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.Status")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the hybrid machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Type")},
			{
				Name:        "ad_fqdn",
				Description: "Specifies the AD fully qualified display name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.AdFqdn")},
			{
				Name:        "agent_version",
				Description: "The hybrid machine agent full version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.AgentVersion")},
			{
				Name:        "client_public_key",
				Description: "Public Key that the client provides to be used during initial resource onboarding.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.ClientPublicKey")},
			{
				Name:        "dns_fqdn",
				Description: "Specifies the DNS fully qualified display name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.DNSFqdn")},
			{
				Name:        "display_name",
				Description: "Specifies the hybrid machine display name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.DisplayName")},
			{
				Name:        "domain_name",
				Description: "Specifies the Windows domain name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.DomainName")},
			{
				Name:        "last_status_change",
				Description: "The time of the last status change.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Machine.Properties.LastStatusChange").Transform(convertDateToTime),
			},
			{
				Name:        "machine_fqdn",
				Description: "Specifies the hybrid machine FQDN.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.MachineFqdn")},
			{
				Name:        "os_name",
				Description: "The Operating System running on the hybrid machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.OSName")},
			{
				Name:        "os_sku",
				Description: "Specifies the Operating System product SKU.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.OSSKU")},
			{
				Name:        "os_version",
				Description: "The version of Operating System running on the hybrid machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.OSVersion"),
			},
			{
				Name:        "vm_id",
				Description: "Specifies the hybrid machine unique ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.VMID")},
			{
				Name:        "vm_uuid",
				Description: "Specifies the Arc Machine's unique SMBIOS ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Properties.VMUUID"),
			},
			{
				Name:        "error_details",
				Description: "Details about the error state.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Machine.Properties.ErrorDetails")},
			{
				Name:        "extensions",
				Description: "The extensions of the compute machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MachineExtensions")},
			{
				Name:        "identity",
				Description: "The identity of the compute machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Machine.Identity")},
			{
				Name:        "location_data",
				Description: "The metadata pertaining to the geographic location of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Machine.Properties.LocationData")},
			{
				Name:        "machine_properties_extensions",
				Description: "The machine properties extensions of the compute machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Machine.Properties.Extensions")},
			{
				Name:        "os_profile",
				Description: "Specifies the operating system settings for the hybrid machine.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Machine.Properties.OSProfile")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Machine.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Machine.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Machine.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Machine.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
