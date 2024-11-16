package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
)

func GetAllIssueComments(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	issues, err := getIssues(ctx, client)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, issue := range issues {
		owner, repo := parseRepoFullName(issue.Repository.GetFullName())
		number := issue.GetNumber()
		issueValues, err := GetRepositoryIssueComments(ctx, githubClient, stream, owner, repo, number)
		if err != nil {
			return nil, err
		}
		values = append(values, issueValues...)
	}
	return values, nil
}

func GetRepositoryIssueComments(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string, issueNumber int) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			Issue struct {
				Comments struct {
					PageInfo   steampipemodels.PageInfo
					TotalCount int
					Nodes      []steampipemodels.IssueComment
				} `graphql:"comments(first: $pageSize, after: $cursor)"`
			} `graphql:"issue(number: $issueNumber)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	variables := map[string]interface{}{
		"owner":       githubv4.String(owner),
		"name":        githubv4.String(repo),
		"issueNumber": githubv4.Int(issueNumber),
		"pageSize":    githubv4.Int(pageSize),
		"cursor":      (*githubv4.String)(nil),
	}
	appendIssuePRCommentColumnIncludes(&variables, issueCommentCols())
	var values []models.Resource
	repoFullName := formRepositoryFullName(owner, repo)
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, comment := range query.Repository.Issue.Comments.Nodes {
			value := models.Resource{
				ID:   strconv.Itoa(comment.Id),
				Name: comment.Url,
				Description: JSONAllFieldsMarshaller{
					Value: model.IssueCommentDescription{
						IssueComment: comment,
						RepoFullName: repoFullName,
						Number:       issueNumber,
						AuthorLogin:  comment.Author.Login,
						EditorLogin:  comment.Editor.Login,
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
		if !query.Repository.Issue.Comments.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Repository.Issue.Comments.PageInfo.EndCursor)
	}
	return values, nil
}
