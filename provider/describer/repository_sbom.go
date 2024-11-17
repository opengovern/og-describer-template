package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
)

func GetAllRepositoriesSBOMs(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
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
		repoValue, err := GetRepositorySBOMs(ctx, githubClient, owner, repo.GetName())
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

func GetRepositorySBOMs(ctx context.Context, githubClient GitHubClient, owner, repo string) (*models.Resource, error) {
	client := githubClient.RestClient
	SBOM, _, err := client.DependencyGraph.GetSBOM(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	repoFullName := formRepositoryFullName(owner, repo)
	value := models.Resource{
		ID:   SBOM.GetSBOM().GetSPDXID(),
		Name: SBOM.GetSBOM().GetName(),
		Description: JSONAllFieldsMarshaller{
			Value: model.RepoSBOMDescription{
				RepositoryFullName: repoFullName,
				SPDXID:             SBOM.GetSBOM().GetSPDXID(),
				SPDXVersion:        SBOM.GetSBOM().GetSPDXVersion(),
				CreationInfo:       SBOM.GetSBOM().GetCreationInfo(),
				Name:               SBOM.GetSBOM().GetName(),
				DataLicense:        SBOM.GetSBOM().GetDataLicense(),
				DocumentDescribes:  SBOM.GetSBOM().DocumentDescribes,
				DocumentNamespace:  SBOM.GetSBOM().GetDocumentNamespace(),
				Packages:           SBOM.GetSBOM().Packages,
			},
		},
	}
	return &value, nil
}
