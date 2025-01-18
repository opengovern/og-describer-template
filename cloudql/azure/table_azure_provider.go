package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureProvider(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_provider",
		Description: "Azure Provider",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("namespace"),
			Hydrate:    opengovernance.GetResourceProvider,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"InvalidResourceNamespace"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListResourceProvider,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the resource provider.",
				Transform:   transform.FromField("Description.Provider.Namespace")},
			{
				Name:        "id",
				Description: "Contains ID to identify a resource provider uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Provider.ID")},
			{
				Name:        "registration_state",
				Description: "Contains the current registration state of the resource provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Provider.RegistrationState")},
			{
				Name:        "resource_types",
				Description: "A list of provider resource types.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Provider.ResourceTypes")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Provider.Namespace")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Provider.ID").Transform(idToAkas),
			},
		}),
	}
}
