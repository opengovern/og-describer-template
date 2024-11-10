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

func GetIssueList(ctx context.Context, githubClient provider.GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var filters githubv4.IssueFilters
	filters.States = &[]githubv4.IssueState{githubv4.IssueStateOpen, githubv4.IssueStateClosed}
	var query struct {
		RateLimit steampipemodels.RateLimit
		Viewer    struct {
			Issues struct {
				TotalCount int
				PageInfo   steampipemodels.PageInfo
				Nodes      []steampipemodels.Issue
			} `graphql:"issues(first: $pageSize, after: $cursor, filterBy: $filters)"`
		}
	}
	variables := map[string]interface{}{
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
		"filters":  filters,
	}
	appendIssueColumnIncludes(&variables, issueCols())
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, issue := range query.Viewer.Issues.Nodes {
			value := models.Resource{
				ID:   strconv.Itoa(issue.Id),
				Name: issue.Title,
				Description: JSONAllFieldsMarshaller{
					Value: model.Issue{
						Issue:        issue,
						RepoFullName: issue.Repo.NameWithOwner,
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
		if !query.Viewer.Issues.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Viewer.Issues.PageInfo.EndCursor)
	}
	return values, nil
}
