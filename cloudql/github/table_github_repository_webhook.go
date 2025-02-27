package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGithubRepositoryWebhook() *plugin.Table {
	return &plugin.Table{
		Name:        "github_repository_webhook",
		Description: "Webhooks configured in the system.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWebhook,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetWebhook,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "repository_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RepositoryID"),
				Description: "Unique identifier of the GitHub repository.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ID"),
				Description: "The unique identifier of the webhook.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Type"),
				Description: "The type of the webhook.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "The name of the webhook.",
			},
			{
				Name:        "active",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Active"),
				Description: "Indicates whether the webhook is active.",
			},
			{
				Name:        "events",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Events"),
				Description: "List of events that trigger this webhook.",
			},
			{
				Name:        "config",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Config"),
				Description: "Configuration details for the webhook.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.UpdatedAt"),
				Description: "The last update timestamp of the webhook.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CreatedAt"),
				Description: "The creation timestamp of the webhook.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.URL"),
				Description: "The primary URL of the webhook.",
			},
			{
				Name:        "test_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TestURL"),
				Description: "The test URL for the webhook.",
			},
			{
				Name:        "ping_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PingURL"),
				Description: "The ping URL for testing the webhook.",
			},
			{
				Name:        "deliveries_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DeliveriesURL"),
				Description: "The URL where deliveries for this webhook are logged.",
			},
			{
				Name:        "last_response",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LastResponse"),
				Description: "The last response received from the webhook endpoint.",
			},
		}),
	}
}
