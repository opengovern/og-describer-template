package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "render"                                    // example: aws, azure
	IntegrationName      = integration.Type("RENDER_ACCOUNT")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-render" // example: github.com/opengovern/og-describer-aws
)
