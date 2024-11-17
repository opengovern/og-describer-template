package describer

import (
	"context"
	"fmt"
	"github.com/google/go-github/v55/github"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"strconv"
)

func GetAllRepositoriesDependabotAlerts(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValues, err := GetRepositoryDependabotAlerts(ctx, githubClient, stream, owner, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

func GetRepositoryDependabotAlerts(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opt := &github.ListAlertsOptions{
		ListCursorOptions: github.ListCursorOptions{First: pageSize},
	}
	var values []models.Resource
	for {
		alerts, resp, err := client.Dependabot.ListRepoAlerts(ctx, owner, repo, opt)
		if err != nil {
			return nil, err
		}
		for _, alert := range alerts {
			var CWEs []string
			for _, cwe := range alert.SecurityAdvisory.CWEs {
				CWEs = append(CWEs, cwe.GetName())
			}
			id := fmt.Sprintf("%s/%s/%s", owner, repo, strconv.Itoa(alert.GetNumber()))
			value := models.Resource{
				ID:   id,
				Name: strconv.Itoa(alert.GetNumber()),
				Description: JSONAllFieldsMarshaller{
					Value: model.RepoAlertDependabotDescription{
						AlertNumber:                 alert.GetNumber(),
						State:                       alert.GetState(),
						DependencyPackageEcosystem:  alert.GetDependency().GetPackage().GetEcosystem(),
						DependencyPackageName:       alert.GetDependency().GetPackage().GetName(),
						DependencyManifestPath:      alert.GetDependency().GetManifestPath(),
						DependencyScope:             alert.GetDependency().GetScope(),
						SecurityAdvisoryGHSAID:      alert.GetSecurityAdvisory().GetGHSAID(),
						SecurityAdvisoryCVEID:       alert.GetSecurityAdvisory().GetCVEID(),
						SecurityAdvisorySummary:     alert.GetSecurityAdvisory().GetSummary(),
						SecurityAdvisoryDescription: alert.GetSecurityAdvisory().GetDescription(),
						SecurityAdvisorySeverity:    alert.GetSecurityAdvisory().GetSeverity(),
						SecurityAdvisoryCVSSScore:   alert.GetSecurityAdvisory().GetCVSS().GetScore(),
						SecurityAdvisoryCVSSVector:  alert.GetSecurityAdvisory().GetCVSS().GetVectorString(),
						SecurityAdvisoryCWEs:        CWEs,
						SecurityAdvisoryPublishedAt: alert.GetSecurityAdvisory().GetPublishedAt(),
						SecurityAdvisoryUpdatedAt:   alert.GetSecurityAdvisory().GetUpdatedAt(),
						SecurityAdvisoryWithdrawnAt: alert.GetSecurityAdvisory().GetWithdrawnAt(),
						URL:                         alert.GetURL(),
						HTMLURL:                     alert.GetHTMLURL(),
						CreatedAt:                   alert.GetCreatedAt(),
						UpdatedAt:                   alert.GetUpdatedAt(),
						DismissedAt:                 alert.GetDismissedAt(),
						DismissedReason:             alert.GetDismissedReason(),
						DismissedComment:            alert.GetDismissedComment(),
						FixedAt:                     alert.GetFixedAt(),
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
		opt.ListCursorOptions.After = resp.After
	}
	return values, nil
}
