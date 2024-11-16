package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

func GetStarList(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit steampipemodels.RateLimit
		Viewer    struct {
			StarredRepositories struct {
				TotalCount int
				PageInfo   steampipemodels.PageInfo
				Edges      []struct {
					StarredAt steampipemodels.NullableTime
					Node      struct {
						NameWithOwner string
						Url           string
					} `graphql:"node @include(if:$includeStarNode)"`
				} `graphql:"edges @include(if:$includeStarEdges)"`
			} `graphql:"starredRepositories(first: $pageSize, after: $cursor)"`
		}
	}
	variables := map[string]interface{}{
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendStarColumnIncludes(&variables, starCols())
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, star := range query.Viewer.StarredRepositories.Edges {
			value := models.Resource{
				ID:   star.Node.Url,
				Name: star.Node.NameWithOwner,
				Description: JSONAllFieldsMarshaller{
					Value: model.StarDescription{
						RepoFullName: star.Node.NameWithOwner,
						StarredAt:    star.StarredAt,
						Url:          star.Node.Url,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		}
		if !query.Viewer.StarredRepositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Viewer.StarredRepositories.PageInfo.EndCursor)
	}
	return values, nil
}
