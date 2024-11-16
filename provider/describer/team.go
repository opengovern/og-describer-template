package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
	"strconv"
)

func GetTeamList(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.GraphQLClient
	var query struct {
		RateLimit steampipemodels.RateLimit
		Viewer    struct {
			Organizations struct {
				PageInfo steampipemodels.PageInfo
				Nodes    []struct {
					Login string
					Teams struct {
						PageInfo steampipemodels.PageInfo
						Nodes    []steampipemodels.TeamWithCounts
					} `graphql:"teams(first: $pageSize, after: $cursor)"`
				}
			} `graphql:"organizations(first: $orgPageSize, after: $orgCursor)"`
		}
	}
	variables := map[string]interface{}{
		"orgPageSize": githubv4.Int(orgPageSize),
		"orgCursor":   (*githubv4.String)(nil),
		"pageSize":    githubv4.Int(pageSize),
		"cursor":      (*githubv4.String)(nil),
	}
	appendTeamColumnIncludes(&variables, teamCols())
	var values []models.Resource
	var teams []steampipemodels.TeamWithCounts
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		for _, org := range query.Viewer.Organizations.Nodes {
			teams = append(teams, org.Teams.Nodes...)
			if org.Teams.PageInfo.HasNextPage {
				ts, err := getAdditionalTeams(ctx, client, org.Login, org.Teams.PageInfo.EndCursor)
				if err != nil {
					return nil, err
				}
				teams = append(teams, ts...)
			}
		}
		for _, team := range teams {
			value := models.Resource{
				ID:   strconv.Itoa(team.Id),
				Name: team.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.TeamDescription{
						Organization: team.Organization.Name,
						Slug:         team.Slug,
						Name:         team.Name,
						ID:           team.Id,
						NodeID:       team.NodeId,
						Description:  team.Description,
						CreatedAt:    team.CreatedAt,
						UpdatedAt:    team.UpdatedAt,
						CombinedSlug: team.CombinedSlug,
						ParentTeam: struct {
							Id     int
							NodeId string
							Name   string
							Slug   string
						}{Id: team.ParentTeam.Id, NodeId: team.ParentTeam.NodeId, Name: team.ParentTeam.Name, Slug: team.ParentTeam.Slug},
						Privacy:                team.Privacy,
						AncestorsTotalCount:    team.Ancestors.TotalCount,
						ChildTeamsTotalCount:   team.ChildTeams.TotalCount,
						DiscussionsTotalCount:  team.Discussions.TotalCount,
						InvitationsTotalCount:  team.Invitations.TotalCount,
						MembersTotalCount:      team.Members.TotalCount,
						ProjectsV2TotalCount:   team.ProjectsV2.TotalCount,
						RepositoriesTotalCount: team.Repositories.TotalCount,
						URL:                    team.Url,
						AvatarURL:              team.AvatarUrl,
						DiscussionsURL:         team.DiscussionsUrl,
						EditTeamURL:            team.EditTeamUrl,
						MembersURL:             team.MembersUrl,
						NewTeamURL:             team.NewTeamUrl,
						RepositoriesURL:        team.RepositoriesUrl,
						TeamsURL:               team.TeamsUrl,
						CanAdminister:          team.CanAdminister,
						CanSubscribe:           team.CanSubscribe,
						Subscription:           team.Subscription,
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

func getAdditionalTeams(ctx context.Context, client *githubv4.Client, org string, initialCursor githubv4.String) ([]steampipemodels.TeamWithCounts, error) {
	var query struct {
		RateLimit    steampipemodels.RateLimit
		Organization struct {
			Teams struct {
				PageInfo steampipemodels.PageInfo
				Nodes    []steampipemodels.TeamWithCounts
			} `graphql:"teams(first: $pageSize, after: $cursor)"`
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"pageSize": githubv4.Int(100),
		"cursor":   githubv4.NewString(initialCursor),
		"login":    githubv4.String(org),
	}
	appendTeamColumnIncludes(&variables, teamCols())
	var ts []steampipemodels.TeamWithCounts
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}
		ts = append(ts, query.Organization.Teams.Nodes...)
		if !query.Organization.Teams.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Organization.Teams.PageInfo.EndCursor)
	}
	return ts, nil
}
