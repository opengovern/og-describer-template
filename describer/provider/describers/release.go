package describers

import (
	"strconv"

	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/describer/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/describer/provider/model"
	"golang.org/x/net/context"
)

func GetReleaseList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	opts := &github.ListOptions{PerPage: releasePageSize}
	for _, r := range repositories {
		for {
			releases, resp, err := githubClient.RestClient.Repositories.ListReleases(ctx, organizationName, r.GetName(), opts)
			if err != nil {
				return nil, err
			}
			for _, release := range releases {
				if release == nil {
					continue
				}
				value := models.Resource{
					ID:   strconv.FormatInt(release.GetID(), 10),
					Name: release.GetName(),
					Description: model.ReleaseDescription{
						RepositoryRelease:  *release,
						RepositoryFullName: r.GetFullName(),
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
	}
	return values, nil
}
