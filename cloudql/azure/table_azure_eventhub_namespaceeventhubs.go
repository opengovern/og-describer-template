package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureEventhubNamespaceEventhubs(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_eventhub_namespaceeventhubs",
		Description: "Azure Eventhub NamespaceEventhubs",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetEventhubNamespaceEventhub,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEventhubNamespaceEventhub,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the namespaceeventhubs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EventHub.ID")},
			{
				Name:        "name",
				Description: "The name of the namespaceeventhubs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Identity.Type")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.EHNamespace.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.EHNamespace.Identity").Transform(idToAkas),
			},
		}),
	}
}
