package describers

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

func GetAllRepositoriesCollaborators(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient

	var repositoryName string
	if value := ctx.Value(paramKeyRepoName); value != nil {
		repositoryName = value.(string)
	}

	teamsRepositories, err := GetAllTeamsRepositories(ctx, githubClient, organizationName, stream)
	if err != nil {
		return nil, err
	}

	if repositoryName != "" {
		org, err := GetOrganizationAdditionalData(ctx, githubClient.RestClient, organizationName)
		var orgId int64
		if org != nil && org.ID != nil {
			orgId = *org.ID
		}
		repoValues, err := GetRepositoryCollaborators(ctx, githubClient, stream, organizationName, orgId, repositoryName)
		if err != nil {
			return nil, err
		}
		for _, t := range teamsRepositories {
			if t.Description.(model.RepoCollaboratorsDescription).RepositoryName == repositoryName {
				repoValues = append(repoValues, t)
			}
		}
		return repoValues, nil
	}
	repositories, err := getRepositories(ctx, client, organizationName)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		org, err := GetOrganizationAdditionalData(ctx, githubClient.RestClient, organizationName)
		var orgId int64
		if org != nil && org.ID != nil {
			orgId = *org.ID
		}
		repoValues, err := GetRepositoryCollaborators(ctx, githubClient, stream, organizationName, orgId, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	values = append(values, teamsRepositories...)
	return values, nil
}

func GetRepositoryCollaborators(ctx context.Context, githubClient model.GitHubClient, stream *models.StreamSender, owner string, orgId int64, repo string) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	affiliation := githubv4.CollaboratorAffiliationAll
	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			Collaborators struct {
				TotalCount int
				PageInfo   steampipemodels.PageInfo
				Edges      []struct {
					Permission githubv4.RepositoryPermission `graphql:"permission @include(if:$includeRCPermission)" json:"permission"`
					Node       steampipemodels.BasicUser     `graphql:"node @include(if:$includeRCNode)" json:"node"`
				}
			} `graphql:"collaborators(first: $pageSize, after: $cursor, affiliation: $affiliation)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	variables := map[string]interface{}{
		"owner":       githubv4.String(owner),
		"repo":        githubv4.String(repo),
		"pageSize":    githubv4.Int(pageSize),
		"cursor":      (*githubv4.String)(nil),
		"affiliation": affiliation,
	}
	appendRepoCollaboratorColumnIncludes(&variables, repositoryCollaboratorsCols())
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, collaborator := range query.Repository.Collaborators.Edges {
			id := fmt.Sprintf("%s/%s/%s", repoFullName, collaborator.Node.Login, string(collaborator.Permission))
			value := models.Resource{
				ID:   id,
				Name: collaborator.Node.Name,
				Description: model.RepoCollaboratorsDescription{
					RepositoryName:   repo,
					RepoFullName:     repoFullName,
					CollaboratorID:   collaborator.Node.Login,
					CollaboratorType: "User",
					Permission:       collaborator.Permission,
					Organization:     owner,
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
		if !query.Repository.Collaborators.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Repository.Collaborators.PageInfo.EndCursor)
	}
	return values, nil
}
