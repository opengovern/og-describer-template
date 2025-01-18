package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureBotServiceBot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_botservice_bot",
		Description: "Azure BotService Bot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetBotServiceBot,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBotServiceBot,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the bot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bot.ID")},
			{
				Name:        "name",
				Description: "The name of the bot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bot.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bot.Properties.DisplayName")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.Bot.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Bot.ID").Transform(idToAkas),
			},
		}),
	}
}
