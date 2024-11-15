package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllAuditLogs(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	organizations, err := getOrganizations(ctx, client)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, org := range organizations {
		orgValues, err := GetRepositoryAuditLog(ctx, githubClient, stream, org.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, orgValues...)
	}
	return values, nil
}

func GetRepositoryAuditLog(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, org string) ([]models.Resource, error) {
	client := githubClient.RestClient
	var phrase string
	var include string
	opts := &github.GetAuditLogOptions{
		Phrase:            &phrase,
		Include:           &include,
		ListCursorOptions: github.ListCursorOptions{PerPage: 100},
	}
	var values []models.Resource
	for {
		auditResults, resp, err := client.Organizations.GetAuditLog(ctx, org, opts)
		if err != nil {
			return nil, err
		}
		for _, audit := range auditResults {
			value := models.Resource{
				ID:   *audit.DocumentID,
				Name: *audit.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.AuditLogDescription{
						ID:            *audit.DocumentID,
						CreatedAt:     audit.CreatedAt,
						Organization:  org,
						Phrase:        phrase,
						Include:       include,
						Action:        *audit.Action,
						Actor:         *audit.Actor,
						ActorLocation: audit.ActorLocation,
						Team:          *audit.Team,
						UserLogin:     *audit.User,
						Repo:          *audit.Repo,
						Data:          audit.Data,
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
		if resp.After == "" {
			break
		}
		opts.After = resp.After
	}
	return values, nil
}
