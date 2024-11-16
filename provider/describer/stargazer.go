package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
)

type Stargazer struct {
	StarredAt steampipemodels.NullableTime `graphql:"starredAt @include(if:$includeStargazerStarredAt)" json:"starred_at"`
	Node      steampipemodels.BasicUser    `graphql:"node @include(if:$includeStargazerNode)" json:"node"`
}

func GetAllStargazers(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	repositories, err := getRepositories(ctx, client, owner)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetStargazers(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetStargazers(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	repoFullName := formRepositoryFullName(owner, repo)
	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			Stargazers struct {
				TotalCount int
				PageInfo   steampipemodels.PageInfo
				Edges      []struct {
					Stargazer
				}
			} `graphql:"stargazers(first: $pageSize, after: $cursor)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	variables := map[string]interface{}{
		"owner":    githubv4.String(owner),
		"repo":     githubv4.String(repo),
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendStargazerColumnIncludes(&variables, stargazerCols())
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, sg := range query.Repository.Stargazers.Edges {
			value := models.Resource{
				ID:   strconv.Itoa(sg.Node.Id),
				Name: sg.Node.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.StargazerDescription{
						RepoFullName: repoFullName,
						StarredAt:    sg.StarredAt,
						UserLogin:    sg.Node.Login,
						UserDetail:   sg.Node,
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
		if !query.Repository.Stargazers.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Repository.Stargazers.PageInfo.EndCursor)
	}
	return values, nil
}
