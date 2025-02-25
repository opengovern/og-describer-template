package describers

import (
	"context"
	"fmt"
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
		org, err := GetOrganizationAdditionalData(ctx, githubClient.RestClient, organizationName)
		var orgId int64
		if org != nil && org.ID != nil {
			orgId = *org.ID
		}
		teamValues, err := GetTeamRepositories(ctx, githubClient, organizationName, stream, team.GetOrganization().GetLogin(), orgId, team.GetSlug(), team.GetID())
		if err != nil {
			return nil, err
		}
		values = append(values, teamValues...)
	}
	return values, nil
}

func GetTeamRepositories(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender, org string, orgId int64, slug string, teamID int64) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit    steampipemodels.RateLimit
		Organization struct {
			ID   string `graphql:"id"`
			Name string `graphql:"name"`
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
			id := fmt.Sprintf("%s/%d/%s", repo.Node.NameWithOwner, teamID, string(repo.Permission))

			value := models.Resource{
				ID:   id,
				Name: repo.Node.Name,
				Description: model.RepoCollaboratorsDescription{
					RepositoryName:   repo.Node.Name,
					RepoFullName:     repo.Node.NameWithOwner,
					CollaboratorID:   strconv.FormatInt(teamID, 10),
					CollaboratorType: "Team",
					Permission:       repo.Permission,
					Organization:     query.Organization.Name,
					OrganizationID:   orgId,
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
