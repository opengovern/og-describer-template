package maps

import (
	model "github.com/opengovern/og-describer-github/discovery/pkg/models"
	"github.com/opengovern/og-describer-github/discovery/provider"
	"github.com/opengovern/og-describer-github/discovery/provider/describers"
	 "github.com/opengovern/og-util/pkg/integration"
)


const (
	IntegrationName      = integration.Type("github_account")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
)

var ResourceTypes = map[string]model.ResourceType{

	"Github/Actions/Artifact": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Actions/Artifact",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllArtifacts),
		GetDescriber:  provider.DescribeSingleByRepo(describers.GetArtifact),
	},

	"Github/Actions/Runner": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Actions/Runner",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllRunners),
		GetDescriber:  provider.DescribeSingleByRepo(describers.GetActionRunner),
	},

	"Github/Actions/Secret": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Actions/Secret",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllSecrets),
		GetDescriber:  provider.DescribeSingleByRepo(describers.GetRepoActionSecret),
	},

	"Github/Actions/WorkflowRun": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Actions/WorkflowRun",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllWorkflowRuns),
		GetDescriber:  nil,
	},

	"Github/Branch": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Branch",
		Tags: map[string][]string{
			"category": {"Branch"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllBranches),
		GetDescriber:  nil,
	},

	"Github/Branch/Protection": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Branch/Protection",
		Tags: map[string][]string{
			"category": {"Branch"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllBranchProtections),
		GetDescriber:  nil,
	},

	"Github/Commit": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Commit",
		Tags: map[string][]string{
			"category": {"Commit"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.ListCommits),
		GetDescriber:  nil,
	},

	"Github/Issue": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Issue",
		Tags: map[string][]string{
			"category": {"Issue"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetIssueList),
		GetDescriber:  provider.DescribeSingleByRepo(describers.GetIssue),
	},

	"Github/License": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/License",
		Tags: map[string][]string{
			"category": {"License"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetLicenseList),
		GetDescriber:  provider.DescribeSingleByRepo(describers.GetLicense),
	},

	"Github/Organization": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Organization",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetOrganizationList),
		GetDescriber:  nil,
	},

	"Github/Organization/Collaborator": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Organization/Collaborator",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllOrganizationsCollaborators),
		GetDescriber:  nil,
	},

	"Github/Organization/Dependabot/Alert": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Organization/Dependabot/Alert",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllOrganizationsDependabotAlerts),
		GetDescriber:  nil,
	},

	"Github/Organization/Member": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Organization/Member",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllMembers),
		GetDescriber:  nil,
	},

	"Github/Organization/Team": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Organization/Team",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetOrganizationTeamList),
		GetDescriber:  nil,
	},

	"Github/PullRequest": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/PullRequest",
		Tags: map[string][]string{
			"category": {"PullRequest"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllPullRequests),
		GetDescriber:  nil,
	},

	"Github/Release": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Release",
		Tags: map[string][]string{
			"category": {"Release"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetReleaseList),
		GetDescriber:  nil,
	},

	"Github/Repository": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Repository",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetRepositoryList),
		GetDescriber:  provider.DescribeSingleByRepo(describers.GetRepository),
	},

	"Github/Repository/Collaborator": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Repository/Collaborator",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllRepositoriesCollaborators),
		GetDescriber:  nil,
	},

	"Github/Repository/DependabotAlert": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Repository/DependabotAlert",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllRepositoriesDependabotAlerts),
		GetDescriber:  nil,
	},

	"Github/Repository/Deployment": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Repository/Deployment",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllRepositoriesDeployments),
		GetDescriber:  nil,
	},

	"Github/Repository/Environment": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Repository/Environment",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllRepositoriesEnvironments),
		GetDescriber:  nil,
	},

	"Github/Repository/Ruleset": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Repository/Ruleset",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllRepositoriesRuleSets),
		GetDescriber:  nil,
	},

	"Github/Repository/SBOM": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Repository/SBOM",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllRepositoriesSBOMs),
		GetDescriber:  nil,
	},

	"Github/Repository/VulnerabilityAlert": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Repository/VulnerabilityAlert",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllRepositoriesVulnerabilities),
		GetDescriber:  nil,
	},

	"Github/Tag": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Tag",
		Tags: map[string][]string{
			"category": {"Tag"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllTags),
		GetDescriber:  nil,
	},

	"Github/Team/Member": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Team/Member",
		Tags: map[string][]string{
			"category": {"Team"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllTeamsMembers),
		GetDescriber:  nil,
	},

	"Github/User": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/User",
		Tags: map[string][]string{
			"category": {"user"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetUser),
		GetDescriber:  nil,
	},

	"Github/Workflow": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Workflow",
		Tags: map[string][]string{
			"category": {"workflow"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetAllWorkflows),
		GetDescriber:  provider.DescribeSingleByRepo(describers.GetRepositoryWorkflow),
	},

	"Github/Container/Package": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Container/Package",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetContainerPackageList),
		GetDescriber:  nil,
	},

	"Github/Package/Maven": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Package/Maven",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetMavenPackageList),
		GetDescriber:  nil,
	},

	"Github/NPM/Package": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/NPM/Package",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetNPMPackageList),
		GetDescriber:  nil,
	},

	"Github/Nuget/Package": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Nuget/Package",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.GetNugetPackageList),
		GetDescriber:  provider.DescribeSingleByRepo(describers.GetNugetPackage),
	},

	"Github/Artifact/DockerFile": {
		IntegrationType: IntegrationName,
		ResourceName:    "Github/Artifact/DockerFile",
		Tags: map[string][]string{
			"category": {"artifact_dockerfile"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: provider.DescribeByGithub(describers.ListArtifactDockerFiles),
		GetDescriber:  nil,
	},
}
