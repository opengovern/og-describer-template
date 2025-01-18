package provider

import (
	"errors"
	"github.com/opengovern/og-describer-render/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"time"
)

// DescribeListByRender A wrapper to pass render authorization to describers functions
func DescribeListByRender(describe func(context.Context, *RenderAPIHandler, *models.StreamSender) ([]models.Resource, error)) models.ResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *models.StreamSender) ([]models.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the api key
		if cfg.APIKey == "" {
			return nil, errors.New("api key must be configured")
		}

		renderAPIHandler := NewRenderAPIHandler(cfg.APIKey, rate.Every(time.Minute/400), 1, 10, 5, 5*time.Minute)

		// Get values from describers
		var values []models.Resource
		result, err := describe(ctx, renderAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByRender A wrapper to pass render authorization to describers functions
func DescribeSingleByRender(describe func(context.Context, *RenderAPIHandler, string) (*models.Resource, error)) models.SingleResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the api key
		if cfg.APIKey == "" {
			return nil, errors.New("api key must be configured")
		}

		renderAPIHandler := NewRenderAPIHandler(cfg.APIKey, rate.Every(time.Minute/400), 1, 10, 5, 5*time.Minute)

		// Get value from describers
		value, err := describe(ctx, renderAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
