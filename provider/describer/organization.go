package describer

import (
	"context"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider"
	"github.com/opengovern/og-describer-template/provider/model"
	steampipemodels "github.com/opengovern/og-describer-template/steampipe-plugin-github/github/models"
	"github.com/shurcooL/githubv4"
	"strconv"
)

func GetOrganizationList(ctx context.Context, githubClient provider.GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit steampipemodels.RateLimit
		Viewer    struct {
			Organizations struct {
				TotalCount int
				PageInfo   steampipemodels.PageInfo
				Nodes      []steampipemodels.OrganizationWithCounts
			} `graphql:"organizations(first: $pageSize, after: $cursor)"`
		}
	}
	variables := map[string]interface{}{
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendOrganizationColumnIncludes(&variables, organizationCols())
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, org := range query.Viewer.Organizations.Nodes {
			value := models.Resource{
				ID:   strconv.Itoa(org.Id),
				Name: org.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.Organization{
						OrganizationWithCounts: org,
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
		if !query.Viewer.Organizations.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Viewer.Organizations.PageInfo.EndCursor)
	}
	return values, nil
}
