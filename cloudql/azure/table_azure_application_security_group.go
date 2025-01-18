package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureApplicationSecurityGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_application_security_group",
		Description: "Azure Application Security Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetNetworkApplicationSecurityGroups,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNetworkApplicationSecurityGroups,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the application security group",
				Transform:   transform.FromField("Description.ApplicationSecurityGroup.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a application security group uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationSecurityGroup.Etag")},
			{
				Name:        "type",
				Description: "The resource type of the application security group",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationSecurityGroup.Type")},
			{
				Name:        "provisioning_state",
				Description: "The resource type of the application security group",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ApplicationSecurityGroup.Properties.ProvisioningState"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the application security group resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationSecurityGroup.Properties.ResourceGUID")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ApplicationSecurityGroup.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationSecurityGroup.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ApplicationSecurityGroup.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ApplicationSecurityGroup.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit
				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTIONS ////

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data
