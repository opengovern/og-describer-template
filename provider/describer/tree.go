package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllTrees(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetRepositoryTrees(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryTrees(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	repository, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	branch, _, err := client.Repositories.GetBranch(ctx, owner, repo, repository.GetDefaultBranch(), false)
	if err != nil {
		return nil, err
	}
	sha := branch.Commit.GetSHA()
	tree, _, err := client.Git.GetTree(ctx, owner, repo, sha, true)
	if err != nil {
		return nil, err
	}
	entries := tree.Entries
	var values []models.Resource
	repoFullName := formRepositoryFullName(owner, repo)
	for _, entry := range entries {
		value := models.Resource{
			ID:   *entry.SHA,
			Name: *entry.SHA,
			Description: JSONAllFieldsMarshaller{
				Value: model.TreeDescription{
					TreeSHA:            sha,
					RepositoryFullName: repoFullName,
					Recursive:          true,
					Truncated:          *tree.Truncated,
					SHA:                entry.SHA,
					Path:               entry.Path,
					Mode:               entry.Mode,
					Type:               entry.Type,
					Size:               entry.Size,
					URL:                entry.URL,
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
	return values, nil
}
