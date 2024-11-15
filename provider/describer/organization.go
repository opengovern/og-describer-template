package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
	"strings"
)

func GetOrganizationList(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
			hooks, err := GetOrganizationHooks(ctx, githubClient.RestClient, org)
			if err != nil {
				return nil, err
			}
			additionalOrgInfo, err := GetOrganizationAdditionalData(ctx, githubClient.RestClient, org)
			if err != nil {
				return nil, err
			}
			value := models.Resource{
				ID:   strconv.Itoa(org.Id),
				Name: org.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.OrganizationDescription{
						Organization:                         org.Organization,
						Hooks:                                hooks,
						BillingEmail:                         *additionalOrgInfo.BillingEmail,
						TwoFactorRequirementEnabled:          *additionalOrgInfo.TwoFactorRequirementEnabled,
						DefaultRepoPermission:                *additionalOrgInfo.DefaultRepoPermission,
						MembersAllowedRepositoryCreationType: *additionalOrgInfo.MembersAllowedRepositoryCreationType,
						MembersCanCreateInternalRepos:        *additionalOrgInfo.MembersCanCreateInternalRepos,
						MembersCanCreatePages:                *additionalOrgInfo.MembersCanCreatePages,
						MembersCanCreatePrivateRepos:         *additionalOrgInfo.MembersCanCreatePrivateRepos,
						MembersCanCreatePublicRepos:          *additionalOrgInfo.MembersCanCreatePublicRepos,
						MembersCanCreateRepos:                *additionalOrgInfo.MembersCanCreateRepos,
						MembersCanForkPrivateRepos:           *additionalOrgInfo.MembersCanForkPrivateRepos,
						PlanFilledSeats:                      *additionalOrgInfo.Plan.FilledSeats,
						PlanName:                             *additionalOrgInfo.Plan.Name,
						PlanPrivateRepos:                     int(*additionalOrgInfo.Plan.PrivateRepos),
						PlanSeats:                            *additionalOrgInfo.Plan.Seats,
						PlanSpace:                            *additionalOrgInfo.Plan.Space,
						Followers:                            *additionalOrgInfo.Followers,
						Following:                            *additionalOrgInfo.Following,
						Collaborators:                        *additionalOrgInfo.Collaborators,
						HasOrganizationProjects:              *additionalOrgInfo.HasOrganizationProjects,
						HasRepositoryProjects:                *additionalOrgInfo.HasRepositoryProjects,
						WebCommitSignoffRequired:             *additionalOrgInfo.WebCommitSignoffRequired,
						MembersWithRoleTotalCount:            org.MembersWithRole.TotalCount,
						PackagesTotalCount:                   org.Packages.TotalCount,
						PinnableItemsTotalCount:              org.PinnableItems.TotalCount,
						PinnedItemsTotalCount:                org.PinnedItems.TotalCount,
						ProjectsTotalCount:                   org.Projects.TotalCount,
						ProjectsV2TotalCount:                 org.ProjectsV2.TotalCount,
						SponsoringTotalCount:                 org.Sponsoring.TotalCount,
						SponsorsTotalCount:                   org.Sponsors.TotalCount,
						TeamsTotalCount:                      org.Teams.TotalCount,
						PrivateRepositoriesTotalCount:        org.PrivateRepositories.TotalCount,
						PublicRepositoriesTotalCount:         org.PublicRepositories.TotalCount,
						RepositoriesTotalCount:               org.Repositories.TotalCount,
						RepositoriesTotalDiskUsage:           org.Repositories.TotalDiskUsage,
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

func GetOrganizationHooks(ctx context.Context, client *github.Client, org steampipemodels.OrganizationWithCounts) ([]*github.Hook, error) {
	login := org.Login
	var orgHooks []*github.Hook
	opt := &github.ListOptions{PerPage: pageSize}
	for {
		hooks, resp, err := client.Organizations.ListHooks(ctx, login, opt)
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		orgHooks = append(orgHooks, hooks...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return orgHooks, nil
}

func GetOrganizationAdditionalData(ctx context.Context, client *github.Client, org steampipemodels.OrganizationWithCounts) (*github.Organization, error) {
	login := org.Login
	organization, _, err := client.Organizations.Get(ctx, login)
	if err != nil {
		return nil, err
	}
	return organization, nil
}
