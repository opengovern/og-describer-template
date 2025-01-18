package azure

import (
	"context"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureManagementLock(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_management_lock",
		Description: "Azure Management Lock",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetManagementLock,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"LockNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListManagementLock,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies management lock.",
				Transform:   transform.FromField("Description.Lock.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a lock uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lock.ID"),
			},
			{
				Name:        "type",
				Description: "The resource type of the lock.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lock.Type"),
			},
			{
				Name:        "lock_level",
				Description: "The level of the lock.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Lock.Properties.Level"),
			},
			{
				Name:        "scope",
				Description: "Contains the scope of the lock.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getAzureManagementLockScope),
			},
			{
				Name:        "notes",
				Description: "Contains the notes about the lock.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lock.Properties.Notes")},
			{
				Name:        "owners",
				Description: "A list of owners of the lock.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Lock.Properties.Owners")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lock.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Lock.ID").Transform(idToAkas),
			},
			{
				Name:        "resource_group",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionResourceGroup,

				Transform: transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

func getAzureManagementLockScope(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.ManagementLock).Description.Lock
	if data.ID == nil {
		return nil, nil
	}
	return strings.Split(*data.ID, "/providers/Microsoft.Authorization/locks/")[0], nil
}
