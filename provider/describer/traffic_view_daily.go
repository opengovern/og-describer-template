package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"strconv"

	"github.com/google/go-github/v55/github"
)

func GetAllTrafficViewDailies(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetRepositoryTrafficViewDailies(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryTrafficViewDailies(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.TrafficBreakdownOptions{Per: "day"}
	trafficViews, _, err := client.Repositories.ListTrafficViews(ctx, owner, repo, opts)
	if err != nil {
		return nil, err
	}
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for _, view := range trafficViews.Views {
		if view != nil {
			value := models.Resource{
				ID:   strconv.Itoa(*view.Uniques),
				Name: strconv.Itoa(*view.Uniques),
				Description: JSONAllFieldsMarshaller{
					Value: model.TrafficViewDailyDescription{
						TrafficData:        *view,
						RepositoryFullName: repoFullName,
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
	}
	return values, nil
}
