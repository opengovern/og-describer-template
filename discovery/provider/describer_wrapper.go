package provider

import (

	"github.com/google/go-github/v55/github"
	model "github.com/opengovern/og-describer-azure/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"github.com/shurcooL/githubv4"
	"golang.org/x/net/context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)


type GitHubClient struct {
	RestClient    *github.Client
	GraphQLClient *githubv4.Client
	Token         string
}
var (
	triggerTypeKey string = "trigger_type"
)
func WithTriggerType(ctx context.Context, tt enums.DescribeTriggerType) context.Context {
	return context.WithValue(ctx, triggerTypeKey, tt)
}

func DescribeBySubscription(describe func(context.Context, *azidentity.ClientSecretCredential, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)
		cred, err := azidentity.NewClientSecretCredential(cfg.TenantID, cfg.ClientID, cfg.ClientPassword, nil)
		if err != nil {
			return nil, err
		}
		var values []model.Resource
		result, err := describe(ctx, cred, additionalData["subscriptionId"], stream)
		if err != nil {
			return nil, err
		}
		accountInfo := map[string]string{
			"SubscriptionID": additionalData["subscriptionId"],
			"TenantID":       cfg.TenantID,
		}
		for _, resource := range result {
			resource.AccountInfo = accountInfo
		}
		values = append(values, result...)

		return values, nil
	}
}

