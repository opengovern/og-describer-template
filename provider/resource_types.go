package provider
import (
	"github.com/opengovern/og-describer-github/provider/describer"
	"github.com/opengovern/og-describer-github/provider/configs"
	model "github.com/opengovern/og-describer-github/pkg/sdk/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"Github/Actions/Artifact": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Actions/Artifact",
		Tags:                 map[string][]string{
            "category": {"Action"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllArtifacts),
		GetDescriber:         DescribeSingleByRepo(describer.GetArtifact),
	},

	"Github/Actions/Repository/Runner": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Actions/Repository/Runner",
		Tags:                 map[string][]string{
            "category": {"Action"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllRunners),
		GetDescriber:         DescribeSingleByRepo(describer.GetActionRunner),
	},

	"Github/Actions/Repository/Secret": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Actions/Repository/Secret",
		Tags:                 map[string][]string{
            "category": {"Action"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllSecrets),
		GetDescriber:         DescribeSingleByRepo(describer.GetRepoActionSecret),
	},

	"Github/Actions/Repository/Workflow_run": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Actions/Repository/Workflow_run",
		Tags:                 map[string][]string{
            "category": {"Action"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllWorkflowRuns),
		GetDescriber:         DescribeSingleByRepo(describer.GetRepoWorkflowRun),
	},

	"Github/Blob": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Blob",
		Tags:                 map[string][]string{
            "category": {"Blob"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllBlobs),
		GetDescriber:         DescribeSingleByRepo(describer.GetBlob),
	},

	"Github/Branch": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Branch",
		Tags:                 map[string][]string{
            "category": {"Branch"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllBranches),
		GetDescriber:         DescribeSingleByRepo(describer.GetBlob),
	},

	"Github/Branch/Protection": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Branch/Protection",
		Tags:                 map[string][]string{
            "category": {"Branch"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllBranchProtections),
		GetDescriber:         nil,
	},

	"Github/Commit": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Commit",
		Tags:                 map[string][]string{
            "category": {"Commit"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllCommits),
		GetDescriber:         DescribeSingleByRepo(describer.GetRepositoryCommit),
	},

	"Github/CommunityProfile": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/CommunityProfile",
		Tags:                 map[string][]string{
            "category": {"Community Profile"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllCommunityProfiles),
		GetDescriber:         nil,
	},

	"Github/Gitignore": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Gitignore",
		Tags:                 map[string][]string{
            "category": {"Gitignore"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetGitIgnoreTemplateList),
		GetDescriber:         DescribeSingleByRepo(describer.GetGitignoreTemplate),
	},

	"Github/Issue": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Issue",
		Tags:                 map[string][]string{
            "category": {"Issue"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetIssueList),
		GetDescriber:         DescribeSingleByRepo(describer.GetIssue),
	},

	"Github/License": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/License",
		Tags:                 map[string][]string{
            "category": {"License"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetLicenseList),
		GetDescriber:         DescribeSingleByRepo(describer.GetLicense),
	},

	"Github/Organization": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Organization",
		Tags:                 map[string][]string{
            "category": {"Organization"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetOrganizationList),
		GetDescriber:         nil,
	},

	"Github/Organization/Collaborator": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Organization/Collaborator",
		Tags:                 map[string][]string{
            "category": {"Organization"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllOrganizationsCollaborators),
		GetDescriber:         nil,
	},

	"Github/Organization/DependabotAlert": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Organization/DependabotAlert",
		Tags:                 map[string][]string{
            "category": {"Organization"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllOrganizationsDependabotAlerts),
		GetDescriber:         nil,
	},

	"Github/Organization/ExternalIdentity": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Organization/ExternalIdentity",
		Tags:                 map[string][]string{
            "category": {"Organization"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllExternalIdentities),
		GetDescriber:         nil,
	},

	"Github/Organization/Member": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Organization/Member",
		Tags:                 map[string][]string{
            "category": {"Organization"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllMembers),
		GetDescriber:         nil,
	},

	"Github/PullRequest": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/PullRequest",
		Tags:                 map[string][]string{
            "category": {"PullRequest"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllPullRequests),
		GetDescriber:         nil,
	},

	"Github/Release": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Release",
		Tags:                 map[string][]string{
            "category": {"Release"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetReleaseList),
		GetDescriber:         nil,
	},

	"Github/Repository": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Repository",
		Tags:                 map[string][]string{
            "category": {"Repository"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetRepositoryList),
		GetDescriber:         DescribeSingleByRepo(describer.GetRepository),
	},

	"Github/Repository/Collaborator": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Repository/Collaborator",
		Tags:                 map[string][]string{
            "category": {"Repository"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllRepositoriesCollaborators),
		GetDescriber:         nil,
	},

	"Github/Repository/DependabotAlert": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Repository/DependabotAlert",
		Tags:                 map[string][]string{
            "category": {"Repository"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllRepositoriesDependabotAlerts),
		GetDescriber:         nil,
	},

	"Github/Repository/Deployment": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Repository/Deployment",
		Tags:                 map[string][]string{
            "category": {"Repository"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllRepositoriesDeployments),
		GetDescriber:         nil,
	},

	"Github/Repository/Environment": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Repository/Environment",
		Tags:                 map[string][]string{
            "category": {"Repository"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllRepositoriesEnvironments),
		GetDescriber:         nil,
	},

	"Github/Repository/Ruleset": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Repository/Ruleset",
		Tags:                 map[string][]string{
            "category": {"Repository"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllRepositoriesRuleSets),
		GetDescriber:         nil,
	},

	"Github/Repository/SBOM": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Repository/SBOM",
		Tags:                 map[string][]string{
            "category": {"Repository"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllRepositoriesSBOMs),
		GetDescriber:         nil,
	},

	"Github/Repository/VulnerabilityAlert": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Repository/VulnerabilityAlert",
		Tags:                 map[string][]string{
            "category": {"Repository"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllRepositoriesVulnerabilities),
		GetDescriber:         nil,
	},

	"Github/Stargazer": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Stargazer",
		Tags:                 map[string][]string{
            "category": {"Stargazer"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllStargazers),
		GetDescriber:         nil,
	},

	"Github/Tag": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Tag",
		Tags:                 map[string][]string{
            "category": {"Tag"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllTags),
		GetDescriber:         nil,
	},

	"Github/Team": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Team",
		Tags:                 map[string][]string{
            "category": {"Team"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetTeamList),
		GetDescriber:         nil,
	},

	"Github/Team/Member": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Team/Member",
		Tags:                 map[string][]string{
            "category": {"Team"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllTeamsMembers),
		GetDescriber:         nil,
	},

	"Github/Team/Repository": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Team/Repository",
		Tags:                 map[string][]string{
            "category": {"team"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllTeamsRepositories),
		GetDescriber:         nil,
	},

	"Github/Traffic/View/Daily": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Traffic/View/Daily",
		Tags:                 map[string][]string{
            "category": {"traffic_view"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllTrafficViewDailies),
		GetDescriber:         nil,
	},

	"Github/Traffic/View/Weekly": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Traffic/View/Weekly",
		Tags:                 map[string][]string{
            "category": {"traffic_view"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllTrafficViewWeeklies),
		GetDescriber:         nil,
	},

	"Github/Tree": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Tree",
		Tags:                 map[string][]string{
            "category": {"tree"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllTrees),
		GetDescriber:         nil,
	},

	"Github/User": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/User",
		Tags:                 map[string][]string{
            "category": {"user"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetUser),
		GetDescriber:         nil,
	},

	"Github/Workflow": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Workflow",
		Tags:                 map[string][]string{
            "category": {"workflow"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllWorkflows),
		GetDescriber:         DescribeSingleByRepo(describer.GetRepositoryWorkflow),
	},

	"Github/CodeOwner": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/CodeOwner",
		Tags:                 map[string][]string{
            "category": {"code_owner"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.ListCodeOwners),
		GetDescriber:         nil,
	},

	"Github/Package/Container": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Package/Container",
		Tags:                 map[string][]string{
            "category": {"package"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetContainerPackageList),
		GetDescriber:         DescribeSingleByRepo(describer.GetContainerPackage),
	},

	"Github/Package/Maven": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Package/Maven",
		Tags:                 map[string][]string{
            "category": {"package"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetMavenPackageList),
		GetDescriber:         DescribeSingleByRepo(describer.GetMavenPackage),
	},

	"Github/Package/NPM": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Package/NPM",
		Tags:                 map[string][]string{
            "category": {"package"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetNPMPackageList),
		GetDescriber:         DescribeSingleByRepo(describer.GetNPMPackage),
	},

	"Github/Package/RubyGems": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Package/RubyGems",
		Tags:                 map[string][]string{
            "category": {"package"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetRubyGemsPackageList),
		GetDescriber:         DescribeSingleByRepo(describer.GetRubyGemsPackage),
	},

	"Github/Package/Nuget": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Package/Nuget",
		Tags:                 map[string][]string{
            "category": {"package"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetNugetPackageList),
		GetDescriber:         DescribeSingleByRepo(describer.GetNugetPackage),
	},

	"Github/Package/Version": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/Package/Version",
		Tags:                 map[string][]string{
            "category": {"package"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.GetAllPackageVersionList),
		GetDescriber:         nil,
	},

	"Github/ArtifactDockerFile": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Github/ArtifactDockerFile",
		Tags:                 map[string][]string{
            "category": {"artifact_dockerfile"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByGithub(describer.ListArtifactDockerFiles),
		GetDescriber:         nil,
	},
}
