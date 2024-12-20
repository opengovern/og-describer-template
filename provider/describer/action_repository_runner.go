package describer

import (
	"context"
	"strconv"

	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllRunners(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner := organizationName
	repositories, err := getRepositories(ctx, client, owner)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositoryRunners(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryRunners(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.ListOptions{PerPage: maxPagesCount}
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for {
		runners, resp, err := client.Actions.ListRunners(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, runner := range runners.Runners {
			value := models.Resource{
				ID:   strconv.Itoa(int(runner.GetID())),
				Name: runner.GetName(),
				Description: JSONAllFieldsMarshaller{
					Value: model.RunnerDescription{
						Runner:       runner,
						RepoFullName: repoFullName,
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
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return values, nil
}

func GetActionRunner(ctx context.Context, githubClient GitHubClient, organizationName string, repositoryName string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
	client := githubClient.RestClient
	runnerID, err := strconv.ParseInt(resourceID, 10, 64)
	if err != nil {
		return nil, err
	}
	repoFullName := formRepositoryFullName(organizationName, repositoryName)
	runner, _, err := client.Actions.GetRunner(ctx, organizationName, repositoryName, runnerID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(runner.GetID())),
		Name: runner.GetName(),
		Description: JSONAllFieldsMarshaller{
			Value: model.RunnerDescription{
				Runner:       runner,
				RepoFullName: repoFullName,
			},
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	}

	return &value, nil
}
