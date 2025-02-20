package describers

import (
	"context"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
	"strings"
)

func GetAllTeamsRepositories(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	teams, err := getTeams(ctx, client)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, team := range teams {
		teamValues, err := GetTeamRepositories(ctx, githubClient, organizationName, stream, team.GetOrganization().GetLogin(), team.GetSlug(), team.GetID())
		if err != nil {
			return nil, err
		}
		values = append(values, teamValues...)
	}
	return values, nil
}

func GetTeamRepositories(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender, org, slug string, teamID int64) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit    steampipemodels.RateLimit
		Organization struct {
			Team struct {
				Repositories struct {
					TotalCount int
					PageInfo   steampipemodels.PageInfo
					Edges      []steampipemodels.TeamRepositoryWithPermission
				} `graphql:"repositories(first: $pageSize, after: $cursor)"`
			} `graphql:"team(slug: $slug)"`
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"login":    githubv4.String(org),
		"slug":     githubv4.String(slug),
		"pageSize": githubv4.Int(teamMembersPageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendRepoColumnIncludes(&variables, teamRepositoriesCols())
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			if strings.Contains(err.Error(), "Could not resolve to an Organization with the login of") {
				return nil, nil
			}
			return nil, err
		}
		for _, repo := range query.Organization.Team.Repositories.Edges {
			value := models.Resource{
				ID:   strconv.Itoa(repo.Node.Id),
				Name: repo.Node.Name,
				Description: model.TeamRepositoryDescription{
					TeamID:             int(teamID),
					RepositoryFullName: repo.Node.NameWithOwner,
					Permission:         string(repo.Permission),
					CreatedAt:          repo.Node.CreatedAt.Time,
					UpdatedAt:          repo.Node.UpdatedAt.Time,
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
		if !query.Organization.Team.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Organization.Team.Repositories.PageInfo.EndCursor)
	}
	return values, nil
}
