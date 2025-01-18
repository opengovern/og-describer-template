package constants

import "github.com/opengovern/og-util/pkg/integration"
import _ "embed"

//go:embed ui-spec.json
var UISpec []byte

//go:embed manifest.yaml
var Manifest []byte

//go:embed Setup.md
var SetupMd []byte

const (
	IntegrationName = integration.Type("azure_subscription") // example: aws_cloud, azure_subscription
)

const (
	DescriberDeploymentName = "og-describer-azure"
	DescriberRunCommand     = "/og-describer-azure"
)
