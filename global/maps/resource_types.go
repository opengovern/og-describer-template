package maps

import (
	"github.com/opengovern/og-describer-github/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{
	"Github/Container/Package": {
		Name:            "Github/Container/Package",
		IntegrationType: constants.IntegrationTypeGithubAccount,
		Description:     "",
		Params: []interfaces.Param{
			{
				Name:        "organization",
				Description: `Please provide the organization name`,
				Required:    false,
			},
		},
	},
	"Github/Repository": {
		Name:            "Github/Repository",
		IntegrationType: constants.IntegrationTypeGithubAccount,
		Description:     "",
		Params: []interfaces.Param{
			{
				Name:        "repository",
				Description: `Please provide the repo name (i.e. "internal-tools")`,
				Required:    false,
			},
			{
				Name:        "organization",
				Description: `Please provide the organization name`,
				Required:    false,
			},
		},
	},
	"Github/Artifact/DockerFile": {
		Name:            "Github/Artifact/DockerFile",
		IntegrationType: constants.IntegrationTypeGithubAccount,
		Description:     "",
		Params: []interfaces.Param{
			{
				Name:        "repository",
				Description: `Please provide the repo name (i.e. "internal-tools")`,
				Required:    false,
			},
			{
				Name:        "organization",
				Description: `Please provide the organization name`,
				Required:    false,
			},
		},
	},
	"Github/Actions/WorkflowRun": {
		Name:            "Github/Actions/WorkflowRun",
		IntegrationType: constants.IntegrationTypeGithubAccount,
		Description:     "",
		Params: []interfaces.Param{
			{
				Name:        "repository",
				Description: `Please provide the repo name (i.e. "internal-tools")`,
				Required:    false,
			},
			{
				Name:        "organization",
				Description: `Please provide the organization name`,
				Required:    false,
			},
			{
				Name:        "run_number",
				Description: `Please provide the run number`,
				Required:    false,
			},
		},
	},
	"Github/PullRequest": {
		Name:            "Github/PullRequest",
		IntegrationType: IntegrationTypeGithubAccount,
		Description:     "",
		Params: []Param{
			{
				Name:        "organization",
				Description: `Please provide the organization name`,
				Required:    false,
			},
			{
				Name:        "repository",
				Description: `Please provide the repo name (i.e. "internal-tools")`,
				Required:    false,
			},
		},
	},
}

var TablesToResourceTypes = map[string]string{
	"github_actions_artifact":               "Github/Actions/Artifact",
	"github_actions_runner":                 "Github/Actions/Runner",
	"github_actions_secret":                 "Github/Actions/Secret",
	"github_actions_workflow_run":           "Github/Actions/WorkflowRun",
	"github_branch":                         "Github/Branch",
	"github_branch_protection":              "Github/Branch/Protection",
	"github_commit":                         "Github/Commit",
	"github_issue":                          "Github/Issue",
	"github_license":                        "Github/License",
	"github_organization":                   "Github/Organization",
	"github_organization_collaborator":      "Github/Organization/Collaborator",
	"github_organization_dependabot_alert":  "Github/Organization/Dependabot/Alert",
	"github_organization_member":            "Github/Organization/Member",
	"github_organization_team":              "Github/Organization/Team",
	"github_pull_request":                   "Github/PullRequest",
	"github_release":                        "Github/Release",
	"github_repository":                     "Github/Repository",
	"github_repository_collaborator":        "Github/Repository/Collaborator",
	"github_repository_dependabot_alert":    "Github/Repository/DependabotAlert",
	"github_repository_deployment":          "Github/Repository/Deployment",
	"github_repository_environment":         "Github/Repository/Environment",
	"github_repository_ruleset":             "Github/Repository/Ruleset",
	"github_repository_sbom":                "Github/Repository/SBOM",
	"github_repository_vulnerability_alert": "Github/Repository/VulnerabilityAlert",
	"github_tag":                            "Github/Tag",
	"github_team_member":                    "Github/Team/Member",
	"github_user":                           "Github/User",
	"github_workflow":                       "Github/Workflow",
	"github_container_package":              "Github/Container/Package",
	"github_maven_package":                  "Github/Package/Maven",
	"github_npm_package":                    "Github/NPM/Package",
	"github_nuget_package":                  "Github/Nuget/Package",
	"github_artifact_dockerfile":            "Github/Artifact/DockerFile",
}

var ResourceTypesList = []string{
	"Github/Actions/Artifact",
	"Github/Actions/Runner",
	"Github/Actions/Secret",
	"Github/Actions/WorkflowRun",
	"Github/Branch",
	"Github/Branch/Protection",
	"Github/Commit",
	"Github/Issue",
	"Github/License",
	"Github/Organization",
	"Github/Organization/Collaborator",
	"Github/Organization/Dependabot/Alert",
	"Github/Organization/Member",
	"Github/Organization/Team",
	"Github/PullRequest",
	"Github/Release",
	"Github/Repository",
	"Github/Repository/Collaborator",
	"Github/Repository/DependabotAlert",
	"Github/Repository/Deployment",
	"Github/Repository/Environment",
	"Github/Repository/Ruleset",
	"Github/Repository/SBOM",
	"Github/Repository/VulnerabilityAlert",
	"Github/Tag",
	"Github/Team/Member",
	"Github/User",
	"Github/Workflow",
	"Github/Container/Package",
	"Github/Package/Maven",
	"Github/NPM/Package",
	"Github/Nuget/Package",
	"Github/Artifact/DockerFile",
}
