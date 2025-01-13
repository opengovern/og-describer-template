package provider

import (
	"errors"
	model "github.com/opengovern/og-describer-fly/pkg/sdk/models"
	"github.com/opengovern/og-describer-fly/provider/configs"
	"github.com/opengovern/og-describer-fly/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"golang.org/x/net/context"
	"time"
)

// DescribeListByFly A wrapper to pass Fly authorization to describer functions
func DescribeListByFly(describe func(context.Context, *resilientbridge.ResilientBridge, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		resilientBridge := resilientbridge.NewResilientBridge()

		restMaxRequests := 500
		restWindowSecs := int64(60)

		// Register TailScale provider
		resilientBridge.RegisterProvider("fly", &adapters.FlyIOAdapter{APIToken: cfg.Token}, &resilientbridge.ProviderConfig{
			UseProviderLimits:   true,
			MaxRequestsOverride: &restMaxRequests,
			WindowSecsOverride:  &restWindowSecs,
			MaxRetries:          3,
			BaseBackoff:         200 * time.Millisecond,
		})

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, resilientBridge, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByFly A wrapper to pass Fly authorization to describer functions
func DescribeSingleByFly(describe func(context.Context, *resilientbridge.ResilientBridge, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string) (*model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		resilientBridge := resilientbridge.NewResilientBridge()

		restMaxRequests := 500
		restWindowSecs := int64(60)

		// Register TailScale provider
		resilientBridge.RegisterProvider("fly", &adapters.FlyIOAdapter{APIToken: cfg.Token}, &resilientbridge.ProviderConfig{
			UseProviderLimits:   true,
			MaxRequestsOverride: &restMaxRequests,
			WindowSecsOverride:  &restWindowSecs,
			MaxRetries:          3,
			BaseBackoff:         200 * time.Millisecond,
		})

		// Get value from describer
		value, err := describe(ctx, resilientBridge, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
