package main

import (
	"encoding/json"
	"strconv"

	"github.com/hashicorp/go-plugin"
	"github.com/jackc/pgtype"
	"github.com/opengovern/og-describer-github/global"
	"github.com/opengovern/og-describer-github/global/maps"
	"github.com/opengovern/og-util/pkg/integration"
	"github.com/opengovern/opencomply/services/integration/models"
	"net/rpc"

)

type Integration struct{}




type IntegrationConfiguration struct {
	NatsScheduledJobsTopic   string
	NatsManualJobsTopic      string
	NatsStreamName           string
	NatsConsumerGroup        string
	NatsConsumerGroupManuals string

	SteampipePluginName string

	UISpec []byte

	DescriberDeploymentName string
	DescriberRunCommand     string
}
func (i *Integration) GetConfiguration() IntegrationConfiguration {
	return IntegrationConfiguration{
		NatsScheduledJobsTopic:   global.JobQueueTopic,
		NatsManualJobsTopic:      global.JobQueueTopicManuals,
		NatsStreamName:           global.StreamName,
		NatsConsumerGroup:        global.ConsumerGroup,
		NatsConsumerGroupManuals: global.ConsumerGroupManuals,

		SteampipePluginName: "github",

		UISpec: UISpec,

		DescriberDeploymentName: DescriberDeploymentName,
		DescriberRunCommand:     DescriberRunCommand,
	}
}

func (i *Integration) HealthCheck(jsonData []byte, providerId string, labels map[string]string, annotations map[string]string) (bool, error) {
	var credentials global.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return false, err
	}

	var name string
	if v, ok := labels["OrganizationName"]; ok {
		name = v
	}
	isHealthy, err := GithubIntegrationHealthcheck(Config{
		Token:            credentials.PatToken,
		OrganizationName: name,
	})
	return isHealthy, err
}

func (i *Integration) DiscoverIntegrations(jsonData []byte) ([]models.Integration, error) {
	var credentials global.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return nil, err
	}
	var integrations []models.Integration
	accounts, err := GithubIntegrationDiscovery(Config{
		Token: credentials.PatToken,
	})
	if err != nil {
		return nil, err
	}
	for _, a := range accounts {
		labels := map[string]string{
			"OrganizationName": a.Login,
		}
		labelsJsonData, err := json.Marshal(labels)
		if err != nil {
			return nil, err
		}
		integrationLabelsJsonb := pgtype.JSONB{}
		err = integrationLabelsJsonb.Set(labelsJsonData)
		if err != nil {
			return nil, err
		}
		integrations = append(integrations, models.Integration{
			ProviderID: strconv.FormatInt(a.ID, 10),
			Name:       a.Login,
			Labels:     integrationLabelsJsonb,
		})
	}
	return integrations, nil
}

func (i *Integration) GetResourceTypesByLabels(labels map[string]string) (map[string]maps.ResourceTypeConfiguration, error) {
	resourceTypesMap := make(map[string]maps.ResourceTypeConfiguration)
	for _, resourceType := range maps.ResourceTypesList {
		if v, ok := maps.ResourceTypeConfigs[resourceType]; ok {
			resourceTypesMap[resourceType] = *v
		} else {
			resourceTypesMap[resourceType] = maps.ResourceTypeConfiguration{}
		}
	}
	return resourceTypesMap, nil
}

func (i *Integration) GetResourceTypeFromTableName(tableName string) string {
	if v, ok := maps.TablesToResourceTypes[tableName]; ok {
		return v
	}

	return ""
}

func (i *Integration) GetIntegrationType() integration.Type {
	return IntegrationTypeGithubAccount
}

func (i *Integration) ListAllTables() map[string][]string {
	plugin := global.Plugin()
	tables := make(map[string][]string)
	for tableKey, table := range plugin.TableMap {
		columnNames := make([]string, 0, len(table.Columns))
		for _, column := range table.Columns {
			columnNames = append(columnNames, column.Name)
		}
		tables[tableKey] = columnNames
	}

	return tables
}
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "platform-integration-plugin",
	MagicCookieValue: "integration",
}

type IntegrationTypePlugin struct {
	Impl IntegrationType
}
type IntegrationType interface {
	GetIntegrationType() integration.Type
	GetConfiguration() IntegrationConfiguration
	GetResourceTypesByLabels(map[string]string) (map[string]maps.ResourceTypeConfiguration, error)
	HealthCheck(jsonData []byte, providerId string, labels map[string]string, annotations map[string]string) (bool, error)
	DiscoverIntegrations(jsonData []byte) ([]models.Integration, error)
	GetResourceTypeFromTableName(tableName string) string
	ListAllTables() map[string][]string
}
type IntegrationTypeRPCServer struct {
	Impl IntegrationType
}
type IntegrationTypeRPC struct {
	client *rpc.Client
}
func (p *IntegrationTypePlugin) Server(*plugin.MuxBroker) (any, error) {
	return &IntegrationTypeRPCServer{Impl: p.Impl}, nil
}

func (IntegrationTypePlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (any, error) {
	return &IntegrationTypeRPC{client: c}, nil
}


func (i *IntegrationTypeRPCServer) GetIntegrationType(_ struct{}, integrationType *integration.Type) error {
	*integrationType = i.Impl.GetIntegrationType()
	return nil
}

func (i *IntegrationTypeRPCServer) GetConfiguration(_ struct{}, configuration *IntegrationConfiguration) error {
	*configuration = i.Impl.GetConfiguration()
	return nil
}

func (i *IntegrationTypeRPCServer) GetResourceTypesByLabels(labels map[string]string, resourceTypes *map[string]maps.ResourceTypeConfiguration) error {
	var err error
	*resourceTypes, err = i.Impl.GetResourceTypesByLabels(labels)
	return err
}
type HealthCheckRequest struct {
	JsonData    []byte
	ProviderId  string
	Labels      map[string]string
	Annotations map[string]string
}
func (i *IntegrationTypeRPCServer) HealthCheck(request HealthCheckRequest, result *bool) error {
	var err error
	*result, err = i.Impl.HealthCheck(request.JsonData, request.ProviderId, request.Labels, request.Annotations)
	return err
}

func (i *IntegrationTypeRPCServer) DiscoverIntegrations(jsonData []byte, integrations *[]models.Integration) error {
	var err error
	*integrations, err = i.Impl.DiscoverIntegrations(jsonData)
	return err
}

func (i *IntegrationTypeRPCServer) GetResourceTypeFromTableName(tableName string, resourceType *string) error {
	*resourceType = i.Impl.GetResourceTypeFromTableName(tableName)
	return nil
}

func (i *IntegrationTypeRPCServer) ListAllTables(_ struct{}, tables *map[string][]string) error {
	*tables = i.Impl.ListAllTables()
	return nil
}
