package describers

import (
	"context"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
	"strings"
)

type memberWithRole struct {
	HasTwoFactorEnabled *bool
	Role                *string
	Node                struct {
		DatabaseID int
		Login      string
		Name       string
		URL        string
		Email      string
		CreatedAt  githubv4.DateTime
		Company    *string
		Status     *struct {
			Message                      *string
			IndicatesLimitedAvailability *bool
		}
		steampipemodels.User
	}
}

func GetAllMembers(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	var values []models.Resource
	orgValues, err := GetOrganizationMembers(ctx, githubClient, stream, organizationName)
	if err != nil {
		return nil, err
	}
	values = append(values, orgValues...)

	return values, nil
}

func GetOrganizationMembers(ctx context.Context, githubClient model.GitHubClient, stream *models.StreamSender, org string) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit    steampipemodels.RateLimit
		Organization struct {
			Login           string
			MembersWithRole struct {
				Edges    []memberWithRole
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"membersWithRole(first: $pageSize, after: $cursor)"`
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"login":    githubv4.String(org),
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendUserColumnIncludes(&variables, orgMembersCols())
	organization, err := GetOrganizationAdditionalData(ctx, githubClient.RestClient, org)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			if strings.Contains(err.Error(), "Could not resolve to an Organization with the login of") {
				return nil, nil
			}
			return nil, err
		}
		for _, member := range query.Organization.MembersWithRole.Edges {
			status := new(bool)
			if member.Node.Status != nil {
				status = member.Node.Status.IndicatesLimitedAvailability
			}
			value := models.Resource{
				ID:   strconv.Itoa(member.Node.Id),
				Name: member.Node.Name,
				Description: model.OrgMembersDescription{
					Company:             member.Node.Company,
					CreatedAt:           member.Node.CreatedAt.Time,
					UpdatedAt:           member.Node.UpdatedAt.Time,
					Email:               member.Node.Email,
					ID:                  member.Node.Id,
					IsSiteAdmin:         member.Node.IsSiteAdmin,
					Location:            member.Node.Location,
					Login:               member.Node.Login,
					LoginID:             strconv.Itoa(member.Node.DatabaseID),
					Name:                member.Node.Name,
					NodeID:              member.Node.NodeId,
					Organization:        org,
					OrganizationID:      organization.ID,
					Role:                member.Role,
					HasTwoFactorEnabled: member.HasTwoFactorEnabled,
					Status:              status,
					WebsiteURL:          member.Node.WebsiteUrl,
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
		if !query.Organization.MembersWithRole.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Organization.MembersWithRole.PageInfo.EndCursor)
	}
	return values, nil
}
