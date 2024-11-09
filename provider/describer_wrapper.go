package provider

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	model "github.com/opengovern/og-describer-entraid/pkg/sdk/models"
	"github.com/opengovern/og-describer-entraid/provider/configs"
	"github.com/opengovern/og-describer-entraid/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
)

func DescribeADByTenantID(describe func(context.Context, *azidentity.ClientSecretCredential, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)
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
