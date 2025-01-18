package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "fly"                                     // example: aws, azure
	IntegrationName      = integration.Type("fly_account")           // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describers-fly" // example: github.com/opengovern/og-describers-aws
)
