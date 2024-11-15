package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

func GetLicenseList(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit steampipemodels.RateLimit
		Licenses  []steampipemodels.License `graphql:"licenses"`
	}
	variables := map[string]interface{}{}
	appendLicenseColumnIncludes(&variables, licenseCols())
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, license := range query.Licenses {
		value := models.Resource{
			ID:   license.Key,
			Name: license.Name,
			Description: JSONAllFieldsMarshaller{
				Value: model.LicenseDescription{
					License: license,
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
