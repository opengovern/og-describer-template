package describer

import (
	"context"
	"fmt"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllSecrets(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	owner := organizationName
	repositories, err := getRepositories(ctx, client, owner)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositorySecrets(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositorySecrets(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.ListOptions{PerPage: maxPagesCount}
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for {
		secrets, resp, err := client.Actions.ListRepoSecrets(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, secret := range secrets.Secrets {
			id := fmt.Sprintf("%s/%s/%s", owner, repo, secret.Name)
			value := models.Resource{
				ID:   id,
				Name: secret.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.SecretDescription{
						Secret:       secret,
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