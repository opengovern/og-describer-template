package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

func GetAllCommunityProfiles(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValue, err := GetRepositoryCommunityProfiles(ctx, githubClient, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		if stream != nil {
			if err := (*stream)(*repoValue); err != nil {
				return nil, err
			}
		} else {
			values = append(values, *repoValue)
		}
	}
	return values, nil
}

func GetRepositoryCommunityProfiles(ctx context.Context, githubClient GitHubClient, owner, repo string) (*models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit  steampipemodels.RateLimit
		Repository struct {
			steampipemodels.CommunityProfile
		} `graphql:"repository(owner: $owner, name: $repo)"`
	}
	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"repo":  githubv4.String(repo),
	}
	appendCommunityProfileColumnIncludes(&variables, communityCols())
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}
	communityProfile := query.Repository.CommunityProfile
	var readMe steampipemodels.Blob
	if communityProfile.ReadMeUpper.Blob.Text != "" {
		readMe = communityProfile.ReadMeUpper.Blob
	} else {
		readMe = communityProfile.ReadMeLower.Blob
	}
	var contributing steampipemodels.Blob
	if communityProfile.ContributingUpper.Blob.Text != "" {
		contributing = communityProfile.ContributingUpper.Blob
	} else if communityProfile.ContributingLower.Blob.Text != "" {
		contributing = communityProfile.ContributingLower.Blob
	} else {
		contributing = communityProfile.ContributingTitle.Blob
	}
	var security steampipemodels.Blob
	if communityProfile.SecurityUpper.Blob.Text != "" {
		security = communityProfile.SecurityUpper.Blob
	} else if communityProfile.SecurityLower.Blob.Text != "" {
		security = communityProfile.SecurityLower.Blob
	} else {
		security = communityProfile.SecurityTitle.Blob
	}
	repoFullName := formRepositoryFullName(owner, repo)
	value := models.Resource{
		ID:   repo,
		Name: repo,
		Description: JSONAllFieldsMarshaller{
			Value: model.CommunityProfileDescription{
				RepoFullName:         repoFullName,
				LicenseInfo:          communityProfile.LicenseInfo,
				CodeOfConduct:        communityProfile.CodeOfConduct,
				IssueTemplates:       communityProfile.IssueTemplates,
				PullRequestTemplates: communityProfile.PullRequestTemplates,
				ReadMe:               readMe,
				Contributing:         contributing,
				Security:             security,
			},
		},
	}
	return &value, nil
}
