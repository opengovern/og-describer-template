package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureAPIManagementBackend(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_api_management_backend",
		Description: "Azure API Management Backend",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"backend_id", "resource_group", "service_name"}),
			Hydrate:    opengovernance.GetAPIManagementBackend,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAPIManagementBackend,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:      "service_name",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
				{
					Name:      "name",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
				{
					Name:      "url",
					Require:   plugin.Optional,
					Operators: []string{"=", "<>"},
				},
				{
					Name:      "resource_group",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
			},
		},
		Columns: azureColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "A friendly name that identifies an API management backend.",
				Transform:   transform.FromField("Description.APIManagementBackend.Name"),
			},
			{
				Name:        "id",
				Description: "Contains ID to identify an API management backend uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.ID"),
			},
			{
				Name:        "url",
				Description: "Runtime Url of the API management backend.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.URL"),
			},
			{
				Name:        "type",
				Description: "Resource type for API Management resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.Type"),
			},
			{
				Name:        "protocol",
				Description: "API management backend communication protocol. Possible values include: 'BackendProtocolHTTP', 'BackendProtocolSoap'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.Protocol"),
			},
			{
				Name:        "description",
				Description: "The API management backend Description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.Description"),
			},
			{
				Name:        "resource_id",
				Description: "Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.ResourceID"),
			},
			{
				Name:        "properties",
				Description: "The API management backend Properties contract.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.Properties"),
			},
			{
				Name:        "credentials",
				Description: "The API management backend credentials contract properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.Credentials"),
			},
			{
				Name:        "proxy",
				Description: "The API management backend proxy contract properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.Proxy"),
			},
			{
				Name:        "tls",
				Description: "The API management backend TLS properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.TLS"),
			},
			{
				Name:        "service_name",
				Description: "Name of the API management service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceName"),
			},
			// We have added this as an extra column because the get call takes only the last path of the id as the backend_id which we do not get from the API
			{
				Name:        "backend_id",
				Description: "The API management backend ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.ID").Transform(lastPathElement),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.Properties.Title"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.APIManagementBackend.ID").Transform(transform.EnsureStringArray),
			},

			// Azure standard columns
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.APIManagementBackend.ResourceGroup"),
			},
		}),
	}
}
