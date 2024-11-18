package github

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/google/go-github/v55/github"
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubGitignore() *plugin.Table {
	return &plugin.Table{
		Name:        "github_gitignore",
		Description: "GitHub defined .gitignore templates that you can associate with your repository.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGitIgnore,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("name"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.GetGitIgnore,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "Name of the gitignore template."},
			{
				Name:        "source",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Source"),
				Description: "Source code of the gitignore template."},
		}),
	}
}
func tableGitHubGitignoreGetData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var name string
	if h.Item != nil {
		item := h.Item.(github.Gitignore)
		name = *item.Name
	} else {
		name = d.EqualsQuals["name"].GetStringValue()
	}

	// Return nil, if no input provided
	if name == "" {
		return nil, nil
	}

	client := connect(ctx, d)

	gitIgnore, _, err := client.Gitignores.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return gitIgnore, nil
}
