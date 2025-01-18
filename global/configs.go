package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "azure_subscription"                                    // example: aws, azure
	IntegrationName      = integration.Type("azure_subscription")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-azure" // example: github.com/opengovern/og-describer-aws
)


type IntegrationCredentials struct {
	TenantID            string `json:"tenant_id"`
	ClientID            string `json:"client_id"`
	ClientPassword      string `json:"client_password"`
	Certificate         string `json:"certificate"`
	CertificatePassword string `json:"certificate_password"`
	SpnObjectId         string `json:"spn_object_id"`
}
