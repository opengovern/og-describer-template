package provider

import (
	model "github.com/opengovern/og-describer-github/describer/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/describer/provider/configs"
	"github.com/opengovern/og-describer-github/describer/provider/describers"
)

var ResourceTypes = map[string]model.ResourceType{

	"Github/Actions/Artifact": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Actions/Artifact",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllArtifacts),
		GetDescriber:  DescribeSingleByRepo(describers.GetArtifact),
	},

	"Github/Actions/Runner": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Actions/Runner",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllRunners),
		GetDescriber:  DescribeSingleByRepo(describers.GetActionRunner),
	},

	"Github/Actions/Secret": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Actions/Secret",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllSecrets),
		GetDescriber:  DescribeSingleByRepo(describers.GetRepoActionSecret),
	},

	"Github/Actions/WorkflowRun": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Actions/WorkflowRun",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllWorkflowRuns),
		GetDescriber:  nil,
	},

	"Github/Branch": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Branch",
		Tags: map[string][]string{
			"category": {"Branch"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllBranches),
		GetDescriber:  nil,
	},

	"Github/Branch/Protection": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Branch/Protection",
		Tags: map[string][]string{
			"category": {"Branch"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllBranchProtections),
		GetDescriber:  nil,
	},

	"Github/Commit": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Commit",
		Tags: map[string][]string{
			"category": {"Commit"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.ListCommits),
		GetDescriber:  nil,
	},

	"Github/Issue": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Issue",
		Tags: map[string][]string{
			"category": {"Issue"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetIssueList),
		GetDescriber:  DescribeSingleByRepo(describers.GetIssue),
	},

	"Github/License": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/License",
		Tags: map[string][]string{
			"category": {"License"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetLicenseList),
		GetDescriber:  DescribeSingleByRepo(describers.GetLicense),
	},

	"Github/Organization": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Organization",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetOrganizationList),
		GetDescriber:  nil,
	},

	"Github/Organization/Collaborator": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Organization/Collaborator",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllOrganizationsCollaborators),
		GetDescriber:  nil,
	},

	"Github/Organization/Dependabot/Alert": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Organization/Dependabot/Alert",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllOrganizationsDependabotAlerts),
		GetDescriber:  nil,
	},

	"Github/Organization/Member": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Organization/Member",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllMembers),
		GetDescriber:  nil,
	},

	"Github/Organization/Team": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Organization/Team",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetOrganizationTeamList),
		GetDescriber:  nil,
	},

	"Github/PullRequest": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/PullRequest",
		Tags: map[string][]string{
			"category": {"PullRequest"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllPullRequests),
		GetDescriber:  nil,
	},

	"Github/Release": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Release",
		Tags: map[string][]string{
			"category": {"Release"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetReleaseList),
		GetDescriber:  nil,
	},

	"Github/Repository": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetRepositoryList),
		GetDescriber:  DescribeSingleByRepo(describers.GetRepository),
	},

	"Github/Repository/Collaborator": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository/Collaborator",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllRepositoriesCollaborators),
		GetDescriber:  nil,
	},

	"Github/Repository/DependabotAlert": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository/DependabotAlert",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllRepositoriesDependabotAlerts),
		GetDescriber:  nil,
	},

	"Github/Repository/Deployment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository/Deployment",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllRepositoriesDeployments),
		GetDescriber:  nil,
	},

	"Github/Repository/Environment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository/Environment",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllRepositoriesEnvironments),
		GetDescriber:  nil,
	},

	"Github/Repository/Ruleset": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository/Ruleset",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllRepositoriesRuleSets),
		GetDescriber:  nil,
	},

	"Github/Repository/SBOM": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository/SBOM",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllRepositoriesSBOMs),
		GetDescriber:  nil,
	},

	"Github/Repository/VulnerabilityAlert": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository/VulnerabilityAlert",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllRepositoriesVulnerabilities),
		GetDescriber:  nil,
	},

	"Github/Tag": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Tag",
		Tags: map[string][]string{
			"category": {"Tag"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllTags),
		GetDescriber:  nil,
	},

	"Github/Team/Member": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Team/Member",
		Tags: map[string][]string{
			"category": {"Team"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllTeamsMembers),
		GetDescriber:  nil,
	},

	"Github/User": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/User",
		Tags: map[string][]string{
			"category": {"user"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetUser),
		GetDescriber:  nil,
	},

	"Github/Workflow": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Workflow",
		Tags: map[string][]string{
			"category": {"workflow"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetAllWorkflows),
		GetDescriber:  DescribeSingleByRepo(describers.GetRepositoryWorkflow),
	},

	"Github/Container/Package": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Container/Package",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetContainerPackageList),
		GetDescriber:  nil,
	},

	"Github/Package/Maven": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Package/Maven",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetMavenPackageList),
		GetDescriber:  nil,
	},

	"Github/NPM/Package": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/NPM/Package",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetNPMPackageList),
		GetDescriber:  nil,
	},

	"Github/Nuget/Package": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Nuget/Package",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.GetNugetPackageList),
		GetDescriber:  DescribeSingleByRepo(describers.GetNugetPackage),
	},

	"Github/Artifact/DockerFile": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Artifact/DockerFile",
		Tags: map[string][]string{
			"category": {"artifact_dockerfile"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describers.ListArtifactDockerFiles),
		GetDescriber:  nil,
	},
}
