package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider"
	"github.com/opengovern/og-describer-template/provider/model"
	"strconv"
)

func GetAllArtifacts(ctx context.Context, githubClient provider.GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	repositories, err := getRepositoriesName(ctx, client, owner)
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositoryArtifacts(ctx, githubClient, stream, owner, repo)
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryArtifacts(ctx context.Context, githubClient provider.GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.ListOptions{PerPage: maxPagesCount}
	var values []models.Resource
	for {
		artifacts, resp, err := client.Actions.ListArtifacts(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, artifact := range artifacts.Artifacts {
			value := models.Resource{
				ID:   strconv.Itoa(int(*artifact.ID)),
				Name: *artifact.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.Artifact{
						Artifact:     *artifact,
						RepoFullName: repo,
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

func GetArtifact(ctx context.Context, client *github.Client, repo string, artifactID int64) (*models.Resource, error) {
	owner, err := getOwnerName(ctx, client)
	if err != nil {
		return nil, nil
	}
	if artifactID == 0 || repo == "" {
		return nil, nil
	}
	artifact, _, err := client.Actions.GetArtifact(ctx, owner, repo, artifactID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(*artifact.ID)),
		Name: *artifact.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.Artifact{
				Artifact:     *artifact,
				RepoFullName: repo,
			},
		},
	}
	return &value, nil
}
