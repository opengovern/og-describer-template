package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

func GetLicenseList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient

	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			LicenseInfo steampipemodels.License
		} `graphql:"repository(owner: $owner, name: $repoName)"`
	}
	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, nil
	}

	var values []models.Resource
	for _, r := range repositories {
		variables := map[string]interface{}{
			"owner":    githubv4.String(organizationName),
			"repoName": githubv4.String(r.GetName()),
		}
		appendLicenseColumnIncludes(&variables, licenseCols())
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		value := models.Resource{
			ID:   query.Repository.LicenseInfo.Key,
			Name: query.Repository.LicenseInfo.Name,
			Description: JSONAllFieldsMarshaller{
				Value: model.LicenseDescription{
					License: query.Repository.LicenseInfo,
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
