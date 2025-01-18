package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "render"                                    // example: aws, azure
	IntegrationName      = integration.Type("render_account")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-render" // example: github.com/opengovern/og-describers-aws
)

type IntegrationCredentials struct {
	APIKey string `json:"api_key"`
}
