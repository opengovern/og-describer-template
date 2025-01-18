package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureAPIManagement(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_api_management",
		Description: "Azure API Management Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetAPIManagement,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "InvalidApiVersionParameter", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAPIManagement,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "A friendly name that identifies an API management service.",
				Transform:   transform.FromField("Description.APIManagement.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify an API management service uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.ID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The current provisioning state of the API management service. Possible values include: 'Created', 'Activating', 'Succeeded', 'Updating', 'Failed', 'Stopped', 'Terminating', 'TerminationFailed', 'Deleted'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Type")},
			{
				Name:        "created_at_utc",
				Description: "Creation UTC date of the API management service.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.APIManagement.Properties.CreatedAtUTC").Transform(convertDateToTime),
			},
			{
				Name:        "developer_portal_url",
				Description: "Developer Portal endpoint URL of the API management service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.DeveloperPortalURL")},
			{
				Name:        "disable_gateway",
				Description: "Property only valid for an API management service deployed in multiple locations. This can be used to disable the gateway in master region.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.APIManagement.Properties.DisableGateway"), Default: false,
			},
			{
				Name:        "enable_client_certificate",
				Description: "Property only meant to be used for Consumption SKU Service. This enforces a client certificate to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the policy on the gateway.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.APIManagement.Properties.EnableClientCertificate"), Default: false,
			},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Etag")},
			{
				Name:        "gateway_regional_url",
				Description: "Gateway URL of the API management service in the default region.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.GatewayRegionalURL")},
			{
				Name:        "gateway_url",
				Description: "Gateway URL of the API management service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.GatewayURL")},
			{
				Name:        "identity_principal_id",
				Description: "The principal id of the identity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.APIManagement.Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The client tenant id of the identity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.APIManagement.Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.APIManagement.Identity.Type"),
			},
			{
				Name:        "management_api_url",
				Description: "Management API endpoint URL of the API management service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.ManagementAPIURL")},
			{
				Name:        "notification_sender_email",
				Description: "Email address from which the notification will be sent.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.NotificationSenderEmail")},
			{
				Name:        "portal_url",
				Description: "Publisher portal endpoint URL of the API management service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.PortalURL")},
			{
				Name:        "publisher_email",
				Description: "Publisher email of the API management service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.PublisherEmail")},
			{
				Name:        "publisher_name",
				Description: "Publisher name of the API management service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.PublisherName")},
			{
				Name:        "restore",
				Description: "Undelete API management service if it was previously soft-deleted.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.APIManagement.Properties.Restore"), Default: false,
			},
			{
				Name:        "scm_url",
				Description: "SCM endpoint URL of the API management service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.ScmURL")},
			{
				Name:        "sku_capacity",
				Description: "Capacity of the SKU (number of deployed units of the SKU)",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.APIManagement.SKU.Capacity")},
			{
				Name:        "sku_name",
				Description: "Name of the Sku",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.APIManagement.SKU.Name"),
			},
			{
				Name:        "target_provisioning_state",
				Description: "The provisioning state of the API management service, which is targeted by the long running operation started on the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.TargetProvisioningState")},
			{
				Name:        "virtual_network_configuration_subnet_name",
				Description: "The name of the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.VirtualNetworkConfiguration.Subnetname")},
			{
				Name:        "virtual_network_configuration_subnet_resource_id",
				Description: "The full resource ID of a subnet in a virtual network to deploy the API Management service in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.VirtualNetworkConfiguration.SubnetResourceID")},
			{
				Name:        "virtual_network_configuration_id",
				Description: "The virtual network ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.VirtualNetworkConfiguration.Vnetid")},
			{
				Name:        "virtual_network_type",
				Description: "The type of VPN in which API management service needs to be configured in. None (Default Value) means the API management service is not part of any Virtual Network, External means the API management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only. Possible values include: 'None', 'External', 'Internal'",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Properties.VirtualNetworkType")},
			{
				Name:        "additional_locations",
				Description: "Additional datacenter locations of the API management service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Properties.AdditionalLocations")},
			{
				Name:        "api_version_constraint",
				Description: "Control plane APIs version constraint for the API management service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Properties.APIVersionConstraint")},
			{
				Name:        "certificates",
				Description: "List of certificates that need to be installed in the API management service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Properties.Certificates")},
			{
				Name:        "custom_properties",
				Description: "Custom properties of the API management service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Properties.CustomProperties")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the API management service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "host_name_configurations",
				Description: "Custom hostname configuration of the API management service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Properties.HostnameConfigurations")},
			{
				Name:        "identity_user_assigned_identities",
				Description: "The list of user identities associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Identity.UserAssignedIdentities")},
			{
				Name:        "private_ip_addresses",
				Description: "Private static load balanced IP addresses of the API management service in primary region which is deployed in an internal virtual network. Available only for 'Basic', 'Standard', 'Premium' and 'Isolated' SKU.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Properties.PrivateIPAddresses")},
			{
				Name:        "public_ip_addresses",
				Description: "Public static load balanced IP addresses of the API management service in primary region. Available only for 'Basic', 'Standard', 'Premium' and 'Isolated' SKU.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Properties.PublicIPAddresses")},
			{
				Name:        "zones",
				Description: "A list of availability zones denoting where the resource needs to come from.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.APIManagement.Zones")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagement.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagement.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.APIManagement.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.APIManagement.Location").Transform(formatRegion).Transform(toLower),
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
