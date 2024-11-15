package describer

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
)

func GetAllSearchIssues(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetSearchIssues(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetSearchIssues(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	repoFullName := formRepositoryFullName(owner, repo)
	input := fmt.Sprintf("repo:%s is:issue", repoFullName)
	var query struct {
		RateLimit steampipemodels.RateLimit
		Search    struct {
			PageInfo steampipemodels.PageInfo
			Edges    []steampipemodels.SearchIssueResult
		} `graphql:"search(type: ISSUE, first: $pageSize, after: $cursor, query: $query)"`
	}
	variables := map[string]interface{}{
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
		"query":    githubv4.String(input),
	}
	appendIssueColumnIncludes(&variables, issueCols())
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, issue := range query.Search.Edges {
			labelsSrc := issue.Node.Labels.Nodes[:100]
			labels := make(map[string]steampipemodels.Label)
			for _, label := range issue.Node.Labels.Nodes {
				labels[label.Name] = label
			}
			value := models.Resource{
				ID:   strconv.Itoa(issue.Node.Id),
				Name: strconv.Itoa(issue.Node.Number),
				Description: JSONAllFieldsMarshaller{
					Value: model.SearchIssueDescription{
						IssueDescription: model.IssueDescription{
							Id:                      issue.Node.Id,
							NodeId:                  issue.Node.NodeId,
							Number:                  issue.Node.Number,
							ActiveLockReason:        issue.Node.ActiveLockReason,
							Author:                  issue.Node.Author,
							AuthorLogin:             issue.Node.Author.Login,
							AuthorAssociation:       issue.Node.AuthorAssociation,
							Body:                    issue.Node.Body,
							BodyUrl:                 issue.Node.BodyUrl,
							Closed:                  issue.Node.Closed,
							ClosedAt:                issue.Node.ClosedAt,
							CreatedAt:               issue.Node.CreatedAt,
							CreatedViaEmail:         issue.Node.CreatedViaEmail,
							Editor:                  issue.Node.Editor,
							FullDatabaseId:          issue.Node.FullDatabaseId,
							IncludesCreatedEdit:     issue.Node.IncludesCreatedEdit,
							IsPinned:                issue.Node.IsPinned,
							IsReadByUser:            issue.Node.IsReadByUser,
							LastEditedAt:            issue.Node.LastEditedAt,
							Locked:                  issue.Node.Locked,
							Milestone:               issue.Node.Milestone,
							PublishedAt:             issue.Node.PublishedAt,
							State:                   issue.Node.State,
							StateReason:             issue.Node.StateReason,
							Title:                   issue.Node.Title,
							UpdatedAt:               issue.Node.UpdatedAt,
							Url:                     issue.Node.Url,
							UserCanClose:            issue.Node.UserCanClose,
							UserCanReact:            issue.Node.UserCanReact,
							UserCanReopen:           issue.Node.UserCanReopen,
							UserCanSubscribe:        issue.Node.UserCanSubscribe,
							UserCanUpdate:           issue.Node.UserCanUpdate,
							UserCannotUpdateReasons: issue.Node.UserCannotUpdateReasons,
							UserDidAuthor:           issue.Node.UserDidAuthor,
							UserSubscription:        issue.Node.UserSubscription,
							CommentsTotalCount:      issue.Node.Comments.TotalCount,
							LabelsTotalCount:        issue.Node.Labels.TotalCount,
							LabelsSrc:               labelsSrc,
							Labels:                  labels,
							AssigneesTotalCount:     issue.Node.Assignees.TotalCount,
							Assignees:               issue.Node.Assignees.Nodes,
						},
						RepoFullName: repoFullName,
						Query:        input,
						TextMatches:  issue.TextMatches,
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
		if !query.Search.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Search.PageInfo.EndCursor)
	}
	return values, nil
}
