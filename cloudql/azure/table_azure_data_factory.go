package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureDataFactory(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_data_factory",
		Description: "Azure Data Factory",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetDataFactory,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "Invalid input"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDataFactory,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Factory.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Factory.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Factory.Type")},
			{
				Name:        "version",
				Description: "Version of the factory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Factory.Properties.Version")},
			{
				Name:        "create_time",
				Description: "Specifies the time, the factory was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Factory.Properties.CreateTime").Transform(convertDateToTime),
			},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Factory.ETag")},
			{
				Name:        "provisioning_state",
				Description: "Factory provisioning state, example Succeeded.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Factory.Properties.ProvisioningState")},
			{
				Name:        "public_network_access",
				Description: "Whether or not public network access is allowed for the data factory.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Factory.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "additional_properties",
				Description: "Unmatched properties from the message are deserialized this collection.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Factory.AdditionalProperties")},
			{
				Name:        "identity",
				Description: "Managed service identity of the factory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Factory.Identity")},
			{
				Name:        "encryption",
				Description: "Properties to enable Customer Managed Key for the factory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Factory.Properties.Encryption")},
			{
				Name:        "repo_configuration",
				Description: "Git repo information of the factory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Factory.Properties.RepoConfiguration")},
			{
				Name:        "global_parameters",
				Description: "List of parameters for factory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Factory.Properties.GlobalParameters")},
			{
				Name:        "private_endpoint_connections",
				Description: "List of private endpoint connections for data factory.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.PrivateEndPointConnections")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Factory.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Factory.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Factory.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Factory.Location").Transform(toLower),
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
