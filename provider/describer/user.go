package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
	"strings"
)

func GetUser(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	login, err := getOwnerName(ctx, githubClient.RestClient)
	if err != nil {
		return nil, nil
	}
	var query struct {
		RateLimit steampipemodels.RateLimit
		User      steampipemodels.UserWithCounts `graphql:"user(login: $login)"`
	}
	variables := map[string]interface{}{
		"login": githubv4.String(login),
	}
	appendUserWithCountColumnIncludes(&variables, userCols())
	err = client.Query(ctx, &query, variables)
	if err != nil {
		if strings.Contains(err.Error(), "Could not resolve to a User with the login of") {
			return nil, nil
		}
		return nil, err
	}
	var values []models.Resource
	user := query.User
	value := models.Resource{
		ID:   strconv.Itoa(user.Id),
		Name: user.Login,
		Description: JSONAllFieldsMarshaller{
			Value: model.UserDescription{
				User:                          user.User,
				RepositoriesTotalDiskUsage:    user.Repositories.TotalDiskUsage,
				FollowersTotalCount:           user.Followers.TotalCount,
				FollowingTotalCount:           user.Following.TotalCount,
				PublicRepositoriesTotalCount:  user.PublicRepositories.TotalCount,
				PrivateRepositoriesTotalCount: user.PrivateRepositories.TotalCount,
				PublicGistsTotalCount:         user.PublicGists.TotalCount,
				IssuesTotalCount:              user.Issues.TotalCount,
				OrganizationsTotalCount:       user.Organizations.TotalCount,
				PublicKeysTotalCount:          user.PublicKeys.TotalCount,
				OpenPullRequestsTotalCount:    user.OpenPullRequests.TotalCount,
				MergedPullRequestsTotalCount:  user.MergedPullRequests.TotalCount,
				ClosedPullRequestsTotalCount:  user.ClosedPullRequests.TotalCount,
				PackagesTotalCount:            user.Packages.TotalCount,
				PinnedItemsTotalCount:         user.PinnedItems.TotalCount,
				SponsoringTotalCount:          user.Sponsoring.TotalCount,
				SponsorsTotalCount:            user.Sponsors.TotalCount,
				StarredRepositoriesTotalCount: user.StarredRepositories.TotalCount,
				WatchingTotalCount:            user.Watching.TotalCount,
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
	return values, nil
}
