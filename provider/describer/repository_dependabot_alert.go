package describer

import (
	"context"
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
				CWEs = append(CWEs, *cwe.Name)
			}
			value := models.Resource{
				ID:   strconv.Itoa(*alert.Number),
				Name: strconv.Itoa(*alert.Number),
				Description: JSONAllFieldsMarshaller{
					Value: model.RepoAlertDependabotDescription{
						AlertNumber:                 *alert.Number,
						State:                       *alert.State,
						DependencyPackageEcosystem:  *alert.Dependency.Package.Ecosystem,
						DependencyPackageName:       *alert.Dependency.Package.Name,
						DependencyManifestPath:      *alert.Dependency.ManifestPath,
						DependencyScope:             *alert.Dependency.Scope,
						SecurityAdvisoryGHSAID:      *alert.SecurityAdvisory.GHSAID,
						SecurityAdvisoryCVEID:       *alert.SecurityAdvisory.CVEID,
						SecurityAdvisorySummary:     *alert.SecurityAdvisory.Summary,
						SecurityAdvisoryDescription: *alert.SecurityAdvisory.Description,
						SecurityAdvisorySeverity:    *alert.SecurityAdvisory.Severity,
						SecurityAdvisoryCVSSScore:   *alert.SecurityAdvisory.CVSS.Score,
						SecurityAdvisoryCVSSVector:  *alert.SecurityAdvisory.CVSS.VectorString,
						SecurityAdvisoryCWEs:        CWEs,
						SecurityAdvisoryPublishedAt: *alert.SecurityAdvisory.PublishedAt,
						SecurityAdvisoryUpdatedAt:   *alert.SecurityAdvisory.UpdatedAt,
						SecurityAdvisoryWithdrawnAt: *alert.SecurityAdvisory.WithdrawnAt,
						URL:                         *alert.URL,
						HTMLURL:                     *alert.HTMLURL,
						CreatedAt:                   *alert.CreatedAt,
						UpdatedAt:                   *alert.UpdatedAt,
						DismissedAt:                 *alert.DismissedAt,
						DismissedReason:             *alert.DismissedReason,
						DismissedComment:            *alert.DismissedComment,
						FixedAt:                     *alert.FixedAt,
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
