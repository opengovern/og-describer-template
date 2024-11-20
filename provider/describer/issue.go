package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"math"
	"strconv"
)

func GetIssueList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var filters githubv4.IssueFilters
	filters.States = &[]githubv4.IssueState{githubv4.IssueStateOpen, githubv4.IssueStateClosed}
	repositories, err := getRepositories(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, nil
	}
	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			Issues struct {
				TotalCount int
				PageInfo   steampipemodels.PageInfo
				Nodes      []steampipemodels.Issue
			} `graphql:"issues(first: $pageSize, after: $cursor, filterBy: $filters)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	var values []models.Resource

	for _, r := range repositories {
		variables := map[string]interface{}{
			"owner":    githubv4.String(organizationName),
			"name":     githubv4.String(r.GetName()),
			"pageSize": githubv4.Int(pageSize),
			"cursor":   (*githubv4.String)(nil),
			"filters":  filters,
		}
		appendIssueColumnIncludes(&variables, issueCols())
		for {
			err := client.Query(ctx, &query, variables)
			if err != nil {
				return nil, err
			}
			for _, issue := range query.Repository.Issues.Nodes {
				labelsSrcLength := int(math.Min(float64(len(issue.Labels.Nodes)), 100.0))
				labelsSrc := issue.Labels.Nodes[:labelsSrcLength]
				labels := make(map[string]steampipemodels.Label)
				for _, label := range issue.Labels.Nodes {
					labels[label.Name] = label
				}
				value := models.Resource{
					ID:   strconv.Itoa(issue.Id),
					Name: issue.Title,
					Description: JSONAllFieldsMarshaller{
						Value: model.IssueDescription{
							RepositoryFullName:      r.GetFullName(),
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
			if !query.Repository.Issues.PageInfo.HasNextPage {
				break
			}
			variables["cursor"] = githubv4.NewString(query.Repository.Issues.PageInfo.EndCursor)
		}
	}
	return values, nil
}

func GetIssue(ctx context.Context, githubClient GitHubClient, organizationName string, repositoryName string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
	repoFullName := formRepositoryFullName(organizationName, repositoryName)
	issueID, err := strconv.ParseInt(resourceID, 10, 64)
	if err != nil {
		return nil, err
	}
	client := githubClient.GraphQLClient

	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			Issue steampipemodels.Issue `graphql:"issue(number: $issueNumber)"`
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}

	variables := map[string]interface{}{
		"owner":       githubv4.String(organizationName),
		"repo":        githubv4.String(repositoryName),
		"issueNumber": githubv4.Int(issueID),
	}
	appendIssueColumnIncludes(&variables, issueCols())

	err = client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}
	labelsSrcLength := int(math.Min(float64(len(query.Repository.Issue.Labels.Nodes)), 100.0))
	labelsSrc := query.Repository.Issue.Labels.Nodes[:labelsSrcLength]
	labels := make(map[string]steampipemodels.Label)
	for _, label := range query.Repository.Issue.Labels.Nodes {
		labels[label.Name] = label
	}
	value := models.Resource{
		ID:   strconv.Itoa(query.Repository.Issue.Id),
		Name: query.Repository.Issue.Title,
		Description: JSONAllFieldsMarshaller{
			Value: model.IssueDescription{
				RepositoryFullName:      repoFullName,
				Id:                      query.Repository.Issue.Id,
				NodeId:                  query.Repository.Issue.NodeId,
				Number:                  query.Repository.Issue.Number,
				ActiveLockReason:        query.Repository.Issue.ActiveLockReason,
				Author:                  query.Repository.Issue.Author,
				AuthorLogin:             query.Repository.Issue.Author.Login,
				AuthorAssociation:       query.Repository.Issue.AuthorAssociation,
				Body:                    query.Repository.Issue.Body,
				BodyUrl:                 query.Repository.Issue.BodyUrl,
				Closed:                  query.Repository.Issue.Closed,
				ClosedAt:                query.Repository.Issue.ClosedAt,
				CreatedAt:               query.Repository.Issue.CreatedAt,
				CreatedViaEmail:         query.Repository.Issue.CreatedViaEmail,
				Editor:                  query.Repository.Issue.Editor,
				FullDatabaseId:          query.Repository.Issue.FullDatabaseId,
				IncludesCreatedEdit:     query.Repository.Issue.IncludesCreatedEdit,
				IsPinned:                query.Repository.Issue.IsPinned,
				IsReadByUser:            query.Repository.Issue.IsReadByUser,
				LastEditedAt:            query.Repository.Issue.LastEditedAt,
				Locked:                  query.Repository.Issue.Locked,
				Milestone:               query.Repository.Issue.Milestone,
				PublishedAt:             query.Repository.Issue.PublishedAt,
				State:                   query.Repository.Issue.State,
				StateReason:             query.Repository.Issue.StateReason,
				Title:                   query.Repository.Issue.Title,
				UpdatedAt:               query.Repository.Issue.UpdatedAt,
				Url:                     query.Repository.Issue.Url,
				UserCanClose:            query.Repository.Issue.UserCanClose,
				UserCanReact:            query.Repository.Issue.UserCanReact,
				UserCanReopen:           query.Repository.Issue.UserCanReopen,
				UserCanSubscribe:        query.Repository.Issue.UserCanSubscribe,
				UserCanUpdate:           query.Repository.Issue.UserCanUpdate,
				UserCannotUpdateReasons: query.Repository.Issue.UserCannotUpdateReasons,
				UserDidAuthor:           query.Repository.Issue.UserDidAuthor,
				UserSubscription:        query.Repository.Issue.UserSubscription,
				CommentsTotalCount:      query.Repository.Issue.Comments.TotalCount,
				LabelsTotalCount:        query.Repository.Issue.Labels.TotalCount,
				LabelsSrc:               labelsSrc,
				Labels:                  labels,
				AssigneesTotalCount:     query.Repository.Issue.Assignees.TotalCount,
				Assignees:               query.Repository.Issue.Assignees.Nodes,
			},
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	}
	return &value, nil
}
