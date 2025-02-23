package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubOrganizationApp() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_app",
		Description: "GitHub organization applications installed in a specific organization.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrganizationApp,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Required},
			},
			Hydrate: opengovernance.GetOrganizationApp,
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "The unique identifier of the app.",
				Transform: transform.FromField("Description.ID")},

			{Name: "client_id", Type: proto.ColumnType_STRING, Description: "The client ID of the app.",
				Transform: transform.FromField("Description.ClientID")},

			{Name: "repository_selection", Type: proto.ColumnType_STRING, Description: "Specifies if the app has access to all or selected repositories.",
				Transform: transform.FromField("Description.RepositorySelection")},

			{Name: "access_tokens_url", Type: proto.ColumnType_STRING, Description: "URL to generate an installation access token.",
				Transform: transform.FromField("Description.AccessTokensURL")},

			{Name: "repositories_url", Type: proto.ColumnType_STRING, Description: "URL to list repositories accessible to the app.",
				Transform: transform.FromField("Description.RepositoriesURL")},

			{Name: "html_url", Type: proto.ColumnType_STRING, Description: "URL to view the app's installation page on GitHub.",
				Transform: transform.FromField("Description.HTMLURL")},

			{Name: "app_id", Type: proto.ColumnType_INT, Description: "The ID of the GitHub app.",
				Transform: transform.FromField("Description.AppID")},

			{Name: "app_slug", Type: proto.ColumnType_STRING, Description: "The slug of the GitHub app.",
				Transform: transform.FromField("Description.AppSlug")},

			{Name: "target_id", Type: proto.ColumnType_INT, Description: "The ID of the target organization.",
				Transform: transform.FromField("Description.TargetID")},

			{Name: "target_type", Type: proto.ColumnType_STRING, Description: "The type of the target entity (e.g., Organization).",
				Transform: transform.FromField("Description.TargetType")},

			{Name: "permissions", Type: proto.ColumnType_JSON, Description: "Permissions granted to the app.",
				Transform: transform.FromField("Description.Permissions")},

			{Name: "events", Type: proto.ColumnType_JSON, Description: "Events the app is subscribed to.",
				Transform: transform.FromField("Description.Events")},

			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the app was created.",
				Transform: transform.FromField("Description.CreatedAt")},

			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the app was last updated.",
				Transform: transform.FromField("Description.UpdatedAt")},

			{Name: "single_file_name", Type: proto.ColumnType_STRING, Description: "Name of the single file the app has access to, if applicable.",
				Transform: transform.FromField("Description.SingleFileName")},

			{Name: "has_multiple_single_files", Type: proto.ColumnType_BOOL, Description: "Indicates if the app has access to multiple single files.",
				Transform: transform.FromField("Description.HasMultipleSingleFiles")},

			{Name: "single_file_paths", Type: proto.ColumnType_JSON, Description: "Paths of single files the app has access to.",
				Transform: transform.FromField("Description.SingleFilePaths")},

			{Name: "suspended_by", Type: proto.ColumnType_JSON, Description: "Information about the user who suspended the app, if applicable.",
				Transform: transform.FromField("Description.SuspendedBy")},

			{Name: "suspended_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the app was suspended, if applicable.",
				Transform: transform.FromField("Description.SuspendedAt")},
		}),
	}
}
