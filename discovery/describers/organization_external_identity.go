package describers

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

func GetAllExternalIdentities(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	var values []models.Resource
	orgValues, err := GetOrganizationExternalIdentities(ctx, githubClient, stream, organizationName)
	if err != nil {
		return nil, err
	}
	values = append(values, orgValues...)

	return values, nil
}

func GetOrganizationExternalIdentities(ctx context.Context, githubClient model.GitHubClient, stream *models.StreamSender, org string) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit    steampipemodels.RateLimit
		Organization struct {
			SamlIdentityProvider struct {
				ExternalIdentities struct {
					PageInfo   steampipemodels.PageInfo
					TotalCount int
					Nodes      []steampipemodels.OrganizationExternalIdentity
				} `graphql:"externalIdentities(first: $pageSize, after: $cursor)"`
			}
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"login":    githubv4.String(org),
		"pageSize": githubv4.Int(pageSize),
		"cursor":   (*githubv4.String)(nil),
	}
	appendOrganizationExternalIdentityColumnIncludes(&variables, orgExternalIdentitiesCols())
	var values []models.Resource
	organization, err := GetOrganizationAdditionalData(ctx, githubClient.RestClient, org)
	if err != nil {
		return nil, err
	}
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, externalIdentity := range query.Organization.SamlIdentityProvider.ExternalIdentities.Nodes {
			id := fmt.Sprintf("%s/%s", org, externalIdentity.User.Login)
			var externalID string
			var externalProviderID string

			for _, att := range externalIdentity.SamlIdentity.Attributes {
				if att.Name == "http://schemas.microsoft.com/identity/claims/objectidentifier" {
					externalID = att.Value
				}
				if att.Name == "http://schemas.microsoft.com/identity/claims/tenantid" {
					externalProviderID = att.Value
				}
			}
			value := models.Resource{
				ID:   id,
				Name: org,
				Description: model.OrgExternalIdentityDescription{
					OrganizationExternalIdentity: externalIdentity,
					Organization:                 org,
					ExternalID:                   externalID,
					ExternalProviderID:           externalProviderID,
					OrganizationID:               organization.ID,
					UserLogin:                    externalIdentity.User.Login,
					UserID:                       externalIdentity.User.Id,
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
		if !query.Organization.SamlIdentityProvider.ExternalIdentities.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Organization.SamlIdentityProvider.ExternalIdentities.PageInfo.EndCursor)
	}
	return values, nil
}
