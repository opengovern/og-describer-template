package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureNetworkSecurityGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_network_security_group",
		Description: "Azure Network Security Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetNetworkSecurityGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNetworkSecurityGroup,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the network security group.",
				Transform:   transform.FromField("Description.SecurityGroup.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a network security group uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.Etag")},
			{
				Name:        "type",
				Description: "The resource type of the network security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.Type")},
			{
				Name:        "provisioning_state",
				Description: "The resource type of the network security group.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.SecurityGroup.Properties.ProvisioningState"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the network security group resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.Properties.ResourceGUID")},
			{
				Name:        "default_security_rules",
				Description: "A list of default security rules of network security group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroup.Properties.DefaultSecurityRules")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the network security group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "flow_logs",
				Description: "A collection of references to flow log resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroup.Properties.FlowLogs")},
			{
				Name:        "network_interfaces",
				Description: "A collection of references to network interfaces.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroup.Properties.NetworkInterfaces")},
			{
				Name:        "security_rules",
				Description: "A list of security rules of network security group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroup.Properties.SecurityRules")},
			{
				Name:        "subnets",
				Description: "A collection of references to subnets.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.SecurityGroup.Properties.Subnets")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityGroup.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityGroup.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.SecurityGroup.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.SecurityGroup.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

// Create session

// If we return the API response directly, the output only gives
// the contents of DiagnosticSettings
