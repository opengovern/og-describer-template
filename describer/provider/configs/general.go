package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "github"                                    // example: aws, azure
	IntegrationName      = integration.Type("github_account")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-github" // example: github.com/opengovern/og-describer-aws
)
