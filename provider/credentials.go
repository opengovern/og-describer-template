package provider

import (
	"encoding/json"
	"fmt"
	model "github.com/opengovern/og-describer-entraid/pkg/sdk/models"
	"github.com/opengovern/og-describer-entraid/provider/configs"
	azuremodel "github.com/opengovern/og-describer-entraid/provider/model"
	"github.com/opengovern/og-util/pkg/describe"
	"strings"
)

func AccountConfigFromMap(m map[string]any) (configs.IntegrationCredentials, error) {
	mj, err := json.Marshal(m)
	if err != nil {
		return configs.IntegrationCredentials{}, err
	}

	var c configs.IntegrationCredentials
	err = json.Unmarshal(mj, &c)
	if err != nil {
		return configs.IntegrationCredentials{}, err
	}

	return c, nil
}

func GetResourceMetadata(job describe.DescribeJob, resource model.Resource) (map[string]string, error) {
	azureMetadata := azuremodel.Metadata{
		ID:               resource.ID,
		Name:             resource.Name,
		SubscriptionID:   job.ProviderID,
		Location:         resource.Location,
		CloudEnvironment: "AzurePublicCloud",
		ResourceType:     strings.ToLower(job.ResourceType),
		IntegrationID:    job.IntegrationID,
	}
	azureMetadataBytes, err := json.Marshal(azureMetadata)
	if err != nil {
		return nil, fmt.Errorf("marshal metadata: %v", err.Error())
	}

	metadata := make(map[string]string)
	err = json.Unmarshal(azureMetadataBytes, &metadata)
	if err != nil {
		return nil, fmt.Errorf("unmarshal metadata: %v", err.Error())
	}
	return metadata, nil
}

func AdjustResource(job describe.DescribeJob, resource *model.Resource) error {
	resource.Location = fixAzureLocation(resource.Location)
	resource.Type = strings.ToLower(job.ResourceType)
	return nil
}

func fixAzureLocation(l string) string {
	return strings.ToLower(strings.ReplaceAll(l, " ", ""))
}

func GetAdditionalParameters(job describe.DescribeJob) (map[string]string, error) {
	additionalParameters := make(map[string]string)
	additionalParameters["subscriptionId"] = job.ProviderID

	return additionalParameters, nil
}
