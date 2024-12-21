package provider

import (
	model "github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/configs"
	"github.com/opengovern/og-describer-github/provider/describer"
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
		ListDescriber: DescribeByGithub(describer.GetAllArtifacts),
		GetDescriber:  DescribeSingleByRepo(describer.GetArtifact),
	},

	"Github/Actions/Runner": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Actions/Runner",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describer.GetAllRunners),
		GetDescriber:  DescribeSingleByRepo(describer.GetActionRunner),
	},

	"Github/Actions/Secret": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Actions/Secret",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describer.GetAllSecrets),
		GetDescriber:  DescribeSingleByRepo(describer.GetRepoActionSecret),
	},

	"Github/Actions/WorkflowRun": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Actions/WorkflowRun",
		Tags: map[string][]string{
			"category": {"Action"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describer.GetAllWorkflowRuns),
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
		ListDescriber: DescribeByGithub(describer.GetAllBranches),
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
		ListDescriber: DescribeByGithub(describer.GetAllBranchProtections),
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
		ListDescriber: DescribeByGithub(describer.ListCommits),
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
		ListDescriber: DescribeByGithub(describer.GetIssueList),
		GetDescriber:  DescribeSingleByRepo(describer.GetIssue),
	},

	"Github/License": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/License",
		Tags: map[string][]string{
			"category": {"License"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describer.GetLicenseList),
		GetDescriber:  DescribeSingleByRepo(describer.GetLicense),
	},

	"Github/Organization": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Organization",
		Tags: map[string][]string{
			"category": {"Organization"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describer.GetOrganizationList),
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
		ListDescriber: DescribeByGithub(describer.GetAllOrganizationsCollaborators),
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
		ListDescriber: DescribeByGithub(describer.GetAllOrganizationsDependabotAlerts),
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
		ListDescriber: DescribeByGithub(describer.GetAllMembers),
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
		ListDescriber: DescribeByGithub(describer.GetOrganizationTeamList),
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
		ListDescriber: DescribeByGithub(describer.GetAllPullRequests),
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
		ListDescriber: DescribeByGithub(describer.GetReleaseList),
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
		ListDescriber: DescribeByGithub(describer.GetRepositoryList),
		GetDescriber:  DescribeSingleByRepo(describer.GetRepository),
	},

	"Github/Repository/Collaborator": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Repository/Collaborator",
		Tags: map[string][]string{
			"category": {"Repository"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describer.GetAllRepositoriesCollaborators),
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
		ListDescriber: DescribeByGithub(describer.GetAllRepositoriesDependabotAlerts),
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
		ListDescriber: DescribeByGithub(describer.GetAllRepositoriesDeployments),
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
		ListDescriber: DescribeByGithub(describer.GetAllRepositoriesEnvironments),
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
		ListDescriber: DescribeByGithub(describer.GetAllRepositoriesRuleSets),
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
		ListDescriber: DescribeByGithub(describer.GetAllRepositoriesSBOMs),
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
		ListDescriber: DescribeByGithub(describer.GetAllRepositoriesVulnerabilities),
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
		ListDescriber: DescribeByGithub(describer.GetAllTags),
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
		ListDescriber: DescribeByGithub(describer.GetAllTeamsMembers),
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
		ListDescriber: DescribeByGithub(describer.GetUser),
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
		ListDescriber: DescribeByGithub(describer.GetAllWorkflows),
		GetDescriber:  DescribeSingleByRepo(describer.GetRepositoryWorkflow),
	},

	"Github/Container/Package": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Container/Package",
		Tags: map[string][]string{
			"category": {"package"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describer.GetContainerPackageList),
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
		ListDescriber: DescribeByGithub(describer.GetMavenPackageList),
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
		ListDescriber: DescribeByGithub(describer.GetNPMPackageList),
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
		ListDescriber: DescribeByGithub(describer.GetNugetPackageList),
		GetDescriber:  DescribeSingleByRepo(describer.GetNugetPackage),
	},

	"Github/Artifact/DockerFile": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Github/Artifact/DockerFile",
		Tags: map[string][]string{
			"category": {"artifact_dockerfile"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByGithub(describer.ListArtifactDockerFiles),
		GetDescriber:  nil,
	},
}
