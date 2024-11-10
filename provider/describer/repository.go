package describer

import (
	"context"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
	steampipemodels "github.com/opengovern/og-describer-template/steampipe-plugin-github/github/models"
	"github.com/shurcooL/githubv4"
	"strconv"
)

func GetRepositoryList(ctx context.Context, client *githubv4.Client) ([]models.Resource, error) {
	var query struct {
		RateLimit steampipemodels.RateLimit
		Viewer    struct {
			Repositories struct {
				PageInfo   steampipemodels.PageInfo
				TotalCount int
				Nodes      []steampipemodels.Repository
			} `graphql:"repositories(first: $pageSize, after: $cursor, affiliations: [COLLABORATOR, OWNER, ORGANIZATION_MEMBER], ownerAffiliations: [COLLABORATOR, OWNER, ORGANIZATION_MEMBER])"`
		}
	}
	variables := map[string]interface{}{
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	columnNames := tableCols()
	appendRepoColumnIncludes(&variables, columnNames)
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, repo := range query.Viewer.Repositories.Nodes {
			value := models.Resource{
				ID:   strconv.Itoa(repo.Id),
				Name: repo.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.Repository{
						Repository: repo,
					},
				},
			}
			values = append(values, value)
		}
		if !query.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Viewer.Repositories.PageInfo.EndCursor)
	}
	return values, nil
}
