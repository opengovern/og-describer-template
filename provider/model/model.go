//go:generate go run ../../SDK/runnable/models/main.go --file $GOFILE --output ../../SDK/generated/resources_clients.go --type $PROVIDER

// Implement types for each resource

package model

import (
	goPipeline "github.com/buildkite/go-pipeline"
	"github.com/google/go-github/v55/github"
	steampipemodels "github.com/opengovern/og-describer-template/steampipe-plugin-github/github/models"
	"github.com/shurcooL/githubv4"
	"time"
)

type Artifact struct {
	ID                 *int64
	NodeID             *string
	Name               *string
	SizeInBytes        *int64
	ArchiveDownloadURL *string
	Expired            *bool
	CreatedAt          *github.Timestamp
	ExpiresAt          *github.Timestamp
	RepoFullName       string
}

type Runner struct {
	github.Runner
	RepoFullName string
}

type Secret struct {
	github.Secret
	RepoFullName string
}

type WorkflowRun struct {
	ID                 *int64
	Name               *string
	NodeID             *string
	HeadBranch         *string
	HeadSHA            *string
	RunNumber          *int
	RunAttempt         *int
	Event              *string
	DisplayTitle       *string
	Status             *string
	Conclusion         *string
	WorkflowID         *int64
	CheckSuiteID       *int64
	CheckSuiteNodeID   *string
	URL                *string
	HTMLURL            *string
	PullRequests       []*github.PullRequest
	CreatedAt          *github.Timestamp
	UpdatedAt          *github.Timestamp
	RunStartedAt       *github.Timestamp
	JobsURL            *string
	LogsURL            *string
	CheckSuiteURL      *string
	ArtifactsURL       *string
	CancelURL          *string
	RerunURL           *string
	PreviousAttemptURL *string
	HeadCommit         *github.HeadCommit
	WorkflowURL        *string
	Repository         *github.Repository
	HeadRepository     *github.Repository
	Actor              *github.User
	TriggeringActor    *github.User
	RepoFullName       string
}

type Branch struct {
	RepoFullName         string
	Name                 string
	Commit               steampipemodels.BaseCommit
	BranchProtectionRule steampipemodels.BranchProtectionRule
	Protected            bool
}

type BranchApp struct {
	Name string
	Slug string
}

type BranchTeam struct {
	Name string
	Slug string
}

type BranchUser struct {
	Name  string
	Login string
}

type BranchProtection struct {
	steampipemodels.BranchProtectionRule
	RepoFullName                    string
	CreatorLogin                    string
	MatchingBranches                int
	PushAllowanceApps               []BranchApp
	PushAllowanceTeams              []BranchTeam
	PushAllowanceUsers              []BranchUser
	BypassForcePushAllowanceApps    []BranchApp
	BypassForcePushAllowanceTeams   []BranchTeam
	BypassForcePushAllowanceUsers   []BranchUser
	BypassPullRequestAllowanceApps  []BranchApp
	BypassPullRequestAllowanceTeams []BranchTeam
	BypassPullRequestAllowanceUsers []BranchUser
}

type Commit struct {
	steampipemodels.Commit
	RepoFullName   string
	AuthorLogin    string
	CommitterLogin string
}

type CommunityProfile struct {
	RepoFullName         string
	LicenseInfo          steampipemodels.BaseLicense
	CodeOfConduct        steampipemodels.RepositoryCodeOfConduct
	IssueTemplates       []steampipemodels.IssueTemplate
	PullRequestTemplates []steampipemodels.PullRequestTemplate
	ReadMe               steampipemodels.Blob
	Contributing         steampipemodels.Blob
	Security             steampipemodels.Blob
}

type GitIgnore struct {
	github.Gitignore
}

type Gist struct {
	github.Gist
	OwnerID    int
	OwnerLogin string
	OwnerType  string
}

type Organization struct {
	steampipemodels.Organization
	Hooks                                []*github.Hook
	BillingEmail                         string
	TwoFactorRequirementEnabled          bool
	DefaultRepoPermission                string
	MembersAllowedRepositoryCreationType string
	MembersCanCreateInternalRepos        bool
	MembersCanCreatePages                bool
	MembersCanCreatePrivateRepos         bool
	MembersCanCreatePublicRepos          bool
	MembersCanCreateRepos                bool
	MembersCanForkPrivateRepos           bool
	PlanFilledSeats                      int
	PlanName                             string
	PlanPrivateRepos                     int
	PlanSeats                            int
	PlanSpace                            int
	Followers                            int
	Following                            int
	Collaborators                        int
	HasOrganizationProjects              bool
	HasRepositoryProjects                bool
	WebCommitSignoffRequired             bool
	MembersWithRoleTotalCount            int
	PackagesTotalCount                   int
	PinnableItemsTotalCount              int
	PinnedItemsTotalCount                int
	ProjectsTotalCount                   int
	ProjectsV2TotalCount                 int
	SponsoringTotalCount                 int
	SponsorsTotalCount                   int
	TeamsTotalCount                      int
	PrivateRepositoriesTotalCount        int
	PublicRepositoriesTotalCount         int
	RepositoriesTotalCount               int
	RepositoriesTotalDiskUsage           int
}

type OrgCollaborators struct {
	Organization   string
	Affiliation    string
	RepositoryName githubv4.String
	Permission     githubv4.RepositoryPermission
	UserLogin      steampipemodels.CollaboratorLogin
}

type OrgAlertDependabot struct {
	AlertNumber                 int
	State                       string
	DependencyPackageEcosystem  string
	DependencyPackageName       string
	DependencyManifestPath      string
	DependencyScope             string
	SecurityAdvisoryGHSAID      string
	SecurityAdvisoryCVEID       string
	SecurityAdvisorySummary     string
	SecurityAdvisoryDescription string
	SecurityAdvisorySeverity    string
	SecurityAdvisoryCVSSScore   float64
	SecurityAdvisoryCVSSVector  string
	SecurityAdvisoryCWEs        []string
	SecurityAdvisoryPublishedAt github.Timestamp
	SecurityAdvisoryUpdatedAt   github.Timestamp
	SecurityAdvisoryWithdrawnAt github.Timestamp
	URL                         string
	HTMLURL                     string
	CreatedAt                   github.Timestamp
	UpdatedAt                   github.Timestamp
	DismissedAt                 github.Timestamp
	DismissedReason             string
	DismissedComment            string
	FixedAt                     github.Timestamp
}

type OrgExternalIdentity struct {
	steampipemodels.OrganizationExternalIdentity
	Organization string
	UserLogin    string
	UserDetail   steampipemodels.BasicUser
}

type OrgMembers struct {
	steampipemodels.User
	Organization        string
	HasTwoFactorEnabled bool
	Role                string
}

type Repository struct {
	ID                            int
	NodeID                        string
	Name                          string
	AllowUpdateBranch             bool
	ArchivedAt                    steampipemodels.NullableTime
	AutoMergeAllowed              bool
	CodeOfConduct                 steampipemodels.RepositoryCodeOfConduct
	ContactLinks                  []steampipemodels.RepositoryContactLink
	CreatedAt                     steampipemodels.NullableTime
	DefaultBranchRef              steampipemodels.BasicRefWithBranchProtectionRule
	DeleteBranchOnMerge           bool
	Description                   string
	DiskUsage                     int
	ForkCount                     int
	ForkingAllowed                bool
	FundingLinks                  []steampipemodels.RepositoryFundingLinks
	HasDiscussionsEnabled         bool
	HasIssuesEnabled              bool
	HasProjectsEnabled            bool
	HasVulnerabilityAlertsEnabled bool
	HasWikiEnabled                bool
	HomepageURL                   string
	InteractionAbility            steampipemodels.RepositoryInteractionAbility
	IsArchived                    bool
	IsBlankIssuesEnabled          bool
	IsDisabled                    bool
	IsEmpty                       bool
	IsFork                        bool
	IsInOrganization              bool
	IsLocked                      bool
	IsMirror                      bool
	IsPrivate                     bool
	IsSecurityPolicyEnabled       bool
	IsTemplate                    bool
	IsUserConfigurationRepository bool
	IssueTemplates                []steampipemodels.IssueTemplate
	LicenseInfo                   steampipemodels.BasicLicense
	LockReason                    githubv4.LockReason
	MergeCommitAllowed            bool
	MergeCommitMessage            githubv4.MergeCommitMessage
	MergeCommitTitle              githubv4.MergeCommitTitle
	MirrorURL                     string
	NameWithOwner                 string
	OpenGraphImageURL             string
	OwnerLogin                    string
	PrimaryLanguage               steampipemodels.Language
	ProjectsURL                   string
	PullRequestTemplates          []steampipemodels.PullRequestTemplate
	PushedAt                      steampipemodels.NullableTime
	RebaseMergeAllowed            bool
	SecurityPolicyURL             string
	SquashMergeAllowed            bool
	SquashMergeCommitMessage      githubv4.SquashMergeCommitMessage
	SquashMergeCommitTitle        githubv4.SquashMergeCommitTitle
	SSHURL                        string
	StargazerCount                int
	UpdatedAt                     steampipemodels.NullableTime
	URL                           string
	UsesCustomOpenGraphImage      bool
	CanAdminister                 bool
	CanCreateProjects             bool
	CanSubscribe                  bool
	CanUpdateTopics               bool
	HasStarred                    bool
	PossibleCommitEmails          []string
	Subscription                  githubv4.SubscriptionState
	Visibility                    githubv4.RepositoryVisibility
	YourPermission                githubv4.RepositoryPermission
	WebCommitSignOffRequired      bool
	RepositoryTopicsTotalCount    int
	OpenIssuesTotalCount          int
	WatchersTotalCount            int
	Hooks                         []*github.Hook
	Topics                        []string
	SubscribersCount              int
	HasDownloads                  bool
	HasPages                      bool
	NetworkCount                  int
}

type RepoCollaborators struct {
	Affiliation  string
	RepoFullName string
	Permission   githubv4.RepositoryPermission
	UserLogin    string
}

type RepoAlertDependabot struct {
	RepoFullName                string
	AlertNumber                 int
	State                       string
	DependencyPackageEcosystem  string
	DependencyPackageName       string
	DependencyManifestPath      string
	DependencyScope             string
	SecurityAdvisoryGHSAID      string
	SecurityAdvisoryCVEID       string
	SecurityAdvisorySummary     string
	SecurityAdvisoryDescription string
	SecurityAdvisorySeverity    string
	SecurityAdvisoryCVSSScore   float64
	SecurityAdvisoryCVSSVector  string
	SecurityAdvisoryCWEs        []string
	SecurityAdvisoryPublishedAt github.Timestamp
	SecurityAdvisoryUpdatedAt   github.Timestamp
	SecurityAdvisoryWithdrawnAt github.Timestamp
	URL                         string
	HTMLURL                     string
	CreatedAt                   github.Timestamp
	UpdatedAt                   github.Timestamp
	DismissedAt                 github.Timestamp
	DismissedReason             string
	DismissedComment            string
	FixedAt                     github.Timestamp
}

type RepoDeployment struct {
	steampipemodels.Deployment
	RepoFullName string
}

type RepoEnvironment struct {
	steampipemodels.Environment
	RepoFullName string
}

type RepoRuleSet struct {
	steampipemodels.Ruleset
	RepoFullName string
}

type RepoSBOM struct {
	RepositoryFullName string
	SPDXID             string
	SPDXVersion        string
	CreationInfo       *github.CreationInfo
	Name               string
	DataLicense        string
	DocumentDescribes  []string
	DocumentNamespace  string
	Packages           []*github.RepoDependencies
}

type RepoVulnerabilityAlert struct {
	RepositoryFullName         string
	Number                     int
	NodeID                     string
	AutoDismissedAt            steampipemodels.NullableTime
	CreatedAt                  steampipemodels.NullableTime
	DependencyScope            githubv4.RepositoryVulnerabilityAlertDependencyScope
	DismissComment             string
	DismissReason              string
	DismissedAt                steampipemodels.NullableTime
	Dismisser                  steampipemodels.BasicUser
	FixedAt                    steampipemodels.NullableTime
	State                      githubv4.RepositoryVulnerabilityAlertState
	SecurityAdvisory           steampipemodels.SecurityAdvisory
	SecurityVulnerability      steampipemodels.SecurityVulnerability
	VulnerableManifestFilename string
	VulnerableManifestPath     string
	VulnerableRequirements     string
	Severity                   githubv4.SecurityAdvisorySeverity
	CvssScore                  float64
}

type Star struct {
	RepoFullName string
	StarredAt    steampipemodels.NullableTime
	Url          string
}

type Issue struct {
	Id                      int
	NodeId                  string
	Number                  int
	ActiveLockReason        githubv4.LockReason
	Author                  steampipemodels.Actor
	AuthorLogin             string
	AuthorAssociation       githubv4.CommentAuthorAssociation
	Body                    string
	BodyUrl                 string
	Closed                  bool
	ClosedAt                steampipemodels.NullableTime
	CreatedAt               steampipemodels.NullableTime
	CreatedViaEmail         bool
	Editor                  steampipemodels.Actor
	FullDatabaseId          string
	IncludesCreatedEdit     bool
	IsPinned                bool
	IsReadByUser            bool
	LastEditedAt            steampipemodels.NullableTime
	Locked                  bool
	Milestone               steampipemodels.Milestone
	PublishedAt             steampipemodels.NullableTime
	State                   githubv4.IssueState
	StateReason             githubv4.IssueStateReason
	Title                   string
	UpdatedAt               steampipemodels.NullableTime
	Url                     string
	UserCanClose            bool
	UserCanReact            bool
	UserCanReopen           bool
	UserCanSubscribe        bool
	UserCanUpdate           bool
	UserCannotUpdateReasons []githubv4.CommentCannotUpdateReason
	UserDidAuthor           bool
	UserSubscription        githubv4.SubscriptionState
	CommentsTotalCount      int
	LabelsTotalCount        int
	LabelsSrc               []steampipemodels.Label
	Labels                  map[string]steampipemodels.Label
	AssigneesTotalCount     int
	Assignees               []steampipemodels.BaseUser
}

type IssueComment struct {
	steampipemodels.IssueComment
	RepoFullName string
	Number       int
	AuthorLogin  string
	EditorLogin  string
}

type License struct {
	steampipemodels.License
}

type SearchCode struct {
	github.CodeResult
	RepoFullName string
	Query        string
}

type SearchCommit struct {
	github.CommitResult
	RepoFullName string
	Query        string
}

type SearchIssue struct {
	Issue
	RepoFullName string
	Query        string
	TextMatches  []steampipemodels.TextMatch
}

type Stargazer struct {
	RepoFullName string
	StarredAt    steampipemodels.NullableTime
	UserLogin    string
	UserDetail   steampipemodels.BasicUser
}

type Tag struct {
	RepositoryFullName string
	Name               string
	TaggerDate         time.Time
	TaggerName         string
	TaggerLogin        string
	Message            string
	Commit             steampipemodels.BaseCommit
}

type ParentTeam struct {
	Id     int
	NodeId string
	Name   string
	Slug   string
}

type Team struct {
	Organization           string
	Slug                   string
	Name                   string
	ID                     int
	NodeID                 string
	Description            string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	CombinedSlug           string
	ParentTeam             ParentTeam
	Privacy                string
	AncestorsTotalCount    int
	ChildTeamsTotalCount   int
	DiscussionsTotalCount  int
	InvitationsTotalCount  int
	MembersTotalCount      int
	ProjectsV2TotalCount   int
	RepositoriesTotalCount int
	URL                    string
	AvatarURL              string
	DiscussionsURL         string
	EditTeamURL            string
	MembersURL             string
	NewTeamURL             string
	RepositoriesURL        string
	TeamsURL               string
	CanAdminister          bool
	CanSubscribe           bool
	Subscription           string
}

type TeamMembers struct {
	steampipemodels.User
	Organization string
	Slug         string
	Role         githubv4.TeamMemberRole
}

type TeamRepository struct {
	Repository
	Organization string
	Slug         string
	Permission   githubv4.RepositoryPermission
}

type TrafficViewDaily struct {
	github.TrafficData
	RepositoryFullName string
}

type TrafficViewWeekly struct {
	github.TrafficData
	RepositoryFullName string
}

type Tree struct {
	TreeSHA            string
	RepositoryFullName string
	Recursive          bool
	Truncated          bool
	SHA                *string
	Path               *string
	Mode               *string
	Type               *string
	Size               *int
	URL                *string
}

type User struct {
	steampipemodels.User
	RepositoriesTotalDiskUsage    int
	FollowersTotalCount           int
	FollowingTotalCount           int
	PublicRepositoriesTotalCount  int
	PrivateRepositoriesTotalCount int
	PublicGistsTotalCount         int
	IssuesTotalCount              int
	OrganizationsTotalCount       int
	PublicKeysTotalCount          int
	OpenPullRequestsTotalCount    int
	MergedPullRequestsTotalCount  int
	ClosedPullRequestsTotalCount  int
	PackagesTotalCount            int
	PinnedItemsTotalCount         int
	SponsoringTotalCount          int
	SponsorsTotalCount            int
	StarredRepositoriesTotalCount int
	WatchingTotalCount            int
}

type Workflow struct {
	github.Workflow
	RepositoryFullName      string
	WorkFlowFileContent     string
	WorkFlowFileContentJson *github.RepositoryContent
	Pipeline                *goPipeline.Pipeline
}
