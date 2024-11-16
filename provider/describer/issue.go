package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
)

func GetIssueList(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
			labelsSrc := issue.Labels.Nodes[:100]
			labels := make(map[string]steampipemodels.Label)
			for _, label := range issue.Labels.Nodes {
				labels[label.Name] = label
			}
			value := models.Resource{
				ID:   strconv.Itoa(issue.Id),
				Name: issue.Title,
				Description: JSONAllFieldsMarshaller{
					Value: model.IssueDescription{
						Id:                      issue.Id,
						NodeId:                  issue.NodeId,
						Number:                  issue.Number,
						ActiveLockReason:        issue.ActiveLockReason,
						Author:                  issue.Author,
						AuthorLogin:             issue.Author.Login,
						AuthorAssociation:       issue.AuthorAssociation,
						Body:                    issue.Body,
						BodyUrl:                 issue.BodyUrl,
						Closed:                  issue.Closed,
						ClosedAt:                issue.ClosedAt,
						CreatedAt:               issue.CreatedAt,
						CreatedViaEmail:         issue.CreatedViaEmail,
						Editor:                  issue.Editor,
						FullDatabaseId:          issue.FullDatabaseId,
						IncludesCreatedEdit:     issue.IncludesCreatedEdit,
						IsPinned:                issue.IsPinned,
						IsReadByUser:            issue.IsReadByUser,
						LastEditedAt:            issue.LastEditedAt,
						Locked:                  issue.Locked,
						Milestone:               issue.Milestone,
						PublishedAt:             issue.PublishedAt,
						State:                   issue.State,
						StateReason:             issue.StateReason,
						Title:                   issue.Title,
						UpdatedAt:               issue.UpdatedAt,
						Url:                     issue.Url,
						UserCanClose:            issue.UserCanClose,
						UserCanReact:            issue.UserCanReact,
						UserCanReopen:           issue.UserCanReopen,
						UserCanSubscribe:        issue.UserCanSubscribe,
						UserCanUpdate:           issue.UserCanUpdate,
						UserCannotUpdateReasons: issue.UserCannotUpdateReasons,
						UserDidAuthor:           issue.UserDidAuthor,
						UserSubscription:        issue.UserSubscription,
						CommentsTotalCount:      issue.Comments.TotalCount,
						LabelsTotalCount:        issue.Labels.TotalCount,
						LabelsSrc:               labelsSrc,
						Labels:                  labels,
						AssigneesTotalCount:     issue.Assignees.TotalCount,
						Assignees:               issue.Assignees.Nodes,
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
