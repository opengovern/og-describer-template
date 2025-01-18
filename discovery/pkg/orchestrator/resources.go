//go:generate go run ../SDK/runable/resourceType/resource_types_generator.go  --output resource_types.go --index-map ../steampipe/table_index_map.go && gofmt -w -s resource_types.go  && goimports -w resource_types.go

package orchestrator

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-fly/discovery/describers"
	models2 "github.com/opengovern/og-describer-fly/discovery/pkg/models"
	"github.com/opengovern/og-describer-fly/global/maps"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"go.uber.org/zap"
	"sort"
	"strings"
)

func ListResourceTypes() []string {
	var list []string
	for k := range maps.ResourceTypes {
		list = append(list, k)
	}

	sort.Strings(list)
	return list
}

func GetResourceType(resourceType string) (*models2.ResourceType, error) {
	if r, ok := maps.ResourceTypes[resourceType]; ok {
		return &r, nil
	}
	resourceType = strings.ToLower(resourceType)
	for k, v := range maps.ResourceTypes {
		k := strings.ToLower(k)
		v := v
		if k == resourceType {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("resource type %s not found", resourceType)
}

func GetResourceTypesMap() map[string]models2.ResourceType {
	return maps.ResourceTypes
}

func GetResources(
	ctx context.Context,
	logger *zap.Logger,
	resourceType string,
	triggerType enums.DescribeTriggerType,
	cfg models2.IntegrationCredentials,
	additionalParameters map[string]string,
	stream *models2.StreamSender,
) error {
	_, err := describe(ctx, logger, cfg, resourceType, triggerType, additionalParameters, stream)
	if err != nil {
		return err
	}
	return nil
}

func describe(ctx context.Context, logger *zap.Logger, accountCfg models2.IntegrationCredentials, resourceType string, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *models2.StreamSender) ([]models2.Resource, error) {
	resourceTypeObject, ok := maps.ResourceTypes[resourceType]
	if !ok {
		return nil, fmt.Errorf("unsupported resource type: %s", resourceType)
	}
	ctx = describers.WithLogger(ctx, logger)

	return resourceTypeObject.ListDescriber(ctx, accountCfg, triggerType, additionalParameters, stream)
}

func GetSingleResource(
	ctx context.Context,
	logger *zap.Logger,
	resourceType string,
	triggerType enums.DescribeTriggerType,
	cfg models2.IntegrationCredentials,
	additionalParameters map[string]string,
	resourceId string,
	stream *models2.StreamSender,
) error {
	_, err := describeSingle(ctx, logger, cfg, resourceType, resourceId, triggerType, additionalParameters, stream)
	if err != nil {
		return err
	}
	return nil
}

func describeSingle(ctx context.Context, logger *zap.Logger, accountCfg models2.IntegrationCredentials, resourceType string, resourceID string, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *models2.StreamSender) (*models2.Resource, error) {
	resourceTypeObject, ok := maps.ResourceTypes[resourceType]
	if !ok {
		return nil, fmt.Errorf("unsupported resource type: %s", resourceType)
	}
	ctx = describers.WithLogger(ctx, logger)

	return resourceTypeObject.GetDescriber(ctx, accountCfg, triggerType, additionalParameters, resourceID)
}
