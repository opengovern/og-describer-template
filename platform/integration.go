package main

import (
	"encoding/json"
	"github.com/opengovern/og-describer-azure/platform/constants"
	"github.com/opengovern/og-describer-azure/global"
	"github.com/opengovern/og-describer-azure/global/maps"
	"github.com/opengovern/og-util/pkg/integration"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

type Integration struct{}

func (i *Integration) GetConfiguration() (interfaces.IntegrationConfiguration, error) {
	return interfaces.IntegrationConfiguration{
		NatsScheduledJobsTopic:   global.JobQueueTopic,
		NatsManualJobsTopic:      global.JobQueueTopicManuals,
		NatsStreamName:           global.StreamName,
		NatsConsumerGroup:        global.ConsumerGroup,
		NatsConsumerGroupManuals: global.ConsumerGroupManuals,

		SteampipePluginName: "github",

		UISpec:   constants.UISpec,
		Manifest: constants.Manifest,
		SetupMD:  constants.SetupMd,

		DescriberDeploymentName: constants.DescriberDeploymentName,
		DescriberRunCommand:     constants.DescriberRunCommand,
	}, nil
}

func (i *Integration) HealthCheck(jsonData []byte, providerId string, labels map[string]string, annotations map[string]string) (bool, error) {
	var credentials global.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return false, err
	}
return AzureIntegrationHealthcheck(Config{
		TenantID:       credentials.TenantID,
		ClientID:       credentials.ClientID,
		ClientSecret:   credentials.ClientPassword,
		CertPath:       "",
		CertContent:    credentials.Certificate,
		CertPassword:   credentials.CertificatePassword,
		SubscriptionID: providerId,
	})
}

func (i *Integration) DiscoverIntegrations(jsonData []byte) ([]integration.Integration, error) {
	var credentials global.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return nil, err
	}
	var integrations []integration.Integration
	subscriptions, err := AzureIntegrationDiscovery(Config{
		TenantID:     credentials.TenantID,
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.ClientPassword,
		CertPath:     "",
		CertContent:  credentials.Certificate,
		CertPassword: credentials.CertificatePassword,
	})
	if err != nil {
		return nil, err
	}
	for _, s := range subscriptions {
		integrations = append(integrations, integration.Integration{
			ProviderID: s.SubscriptionID,
			Name:       s.DisplayName,
		})
	}

	return integrations, nil
}

func (i *Integration) GetResourceTypesByLabels(labels map[string]string) (map[string]interfaces.ResourceTypeConfiguration, error) {
	resourceTypesMap := make(map[string]interfaces.ResourceTypeConfiguration)
	for _, resourceType := range maps.ResourceTypesList {
		if v, ok := maps.ResourceTypeConfigs[resourceType]; ok {
			resourceTypesMap[resourceType] = *v
		} else {
			resourceTypesMap[resourceType] = interfaces.ResourceTypeConfiguration{}
		}
	}
	return resourceTypesMap, nil
}

func (i *Integration) GetResourceTypeFromTableName(tableName string) (string, error) {
	if v, ok := maps.TablesToResourceTypes[tableName]; ok {
		return v, nil
	}

	return "", nil
}

func (i *Integration) GetIntegrationType() (integration.Type, error) {
	return constants.IntegrationName, nil
}

func (i *Integration) ListAllTables() (map[string][]interfaces.CloudQLColumn, error) {
	plugin := global.Plugin()
	tables := make(map[string][]interfaces.CloudQLColumn)
	for tableKey, table := range plugin.TableMap {
		columns := make([]interfaces.CloudQLColumn, 0, len(table.Columns))
		for _, column := range table.Columns {
			columns = append(columns, interfaces.CloudQLColumn{Name: column.Name, Type: column.Type.String()})
		}
		tables[tableKey] = columns
	}

	return tables, nil
}

func (i *Integration) Ping() error {
	return nil
}
