package describer

import (
	"context"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"strconv"
)

func GetMavenPackageList(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})
	packages, err := fetchAllPackages(sdk, organizationName, "maven")
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, p := range packages {
		fullDetails, err := fetchPackageDetails(sdk, organizationName, "maven", p.Name)
		if err != nil {
			return nil, err
		}
		value := models.Resource{
			ID:   strconv.Itoa(fullDetails.ID),
			Name: fullDetails.Name,
			Description: JSONAllFieldsMarshaller{
				Value: fullDetails,
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

	return values, nil
}
