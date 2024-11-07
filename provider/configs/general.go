package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "entraid"                                    // example: aws, azure
	IntegrationName      = integration.Type("ENTRA_ID_DIRECTORY")       // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-entraid" // example: github.com/opengovern/og-describer-aws
)
