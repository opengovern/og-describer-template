package provider

import (

	model "github.com/opengovern/og-describer-entraid/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)



var (
	triggerTypeKey string = "trigger_type"
)
func WithTriggerType(ctx context.Context, tt enums.DescribeTriggerType) context.Context {
	return context.WithValue(ctx, triggerTypeKey, tt)
}
func DescribeADByTenantID(describe func(context.Context, *azidentity.ClientSecretCredential, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)
		cred, err := azidentity.NewClientSecretCredential(cfg.TenantID, cfg.ClientID, cfg.ClientPassword, nil)
		if err != nil {
			return nil, err
		}
		var values []model.Resource
		result, err := describe(ctx, cred, cfg.TenantID, stream)
		if err != nil {
			return nil, err
		}

		values = append(values, result...)

		return values, nil
	}
}
