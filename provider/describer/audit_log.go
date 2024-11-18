package describer

import (
	"context"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllAuditLogs(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	var values []models.Resource
	orgValues, err := GetRepositoryAuditLog(ctx, githubClient, stream, organizationName)
	if err != nil {
		return nil, err
	}
	values = append(values, orgValues...)
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
				ID:   audit.GetDocumentID(),
				Name: audit.GetName(),
				Description: JSONAllFieldsMarshaller{
					Value: model.AuditLogDescription{
						ID:            audit.GetDocumentID(),
						CreatedAt:     audit.GetCreatedAt(),
						Organization:  org,
						Phrase:        phrase,
						Include:       include,
						Action:        audit.GetAction(),
						Actor:         audit.GetActor(),
						ActorLocation: audit.GetActorLocation(),
						Team:          audit.GetTeam(),
						UserLogin:     audit.GetUser(),
						Repo:          audit.GetRepository(),
						Data:          audit.GetData(),
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
