//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package provider

import "time"

type Metadata struct{}

type OwnerJSON struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	TwoFactorAuthEnabled bool   `json:"two_factor_auth_enabled"` // converted to snake_case from "twoFactorAuthEnabled"
	Type                 string `json:"type"`
}

type Owner struct {
	ID                   string
	Name                 string
	Email                string
	TwoFactorAuthEnabled bool
	Type                 string
}

type ProjectResponse struct {
	Project ProjectJSON `json:"project"`
	Cursor  string      `json:"cursor"`
}

type ProjectJSON struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`      // converted to snake_case from "createdAt"
	UpdatedAt      time.Time `json:"updated_at"`      // converted to snake_case from "updatedAt"
	Name           string    `json:"name"`
	Owner          OwnerJSON `json:"owner"`
	EnvironmentIDs []string  `json:"environment_ids"` // converted to snake_case from "environmentIds"
}

type ProjectDescription struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string
	Owner          Owner
	EnvironmentIDs []string
}

type EnvironmentResponse struct {
	Environment EnvironmentJSON `json:"environment"`
	Cursor      string          `json:"cursor"`
}

type EnvironmentJSON struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	ProjectID       string   `json:"project_id"`       // converted to snake_case from "projectId"
	DatabasesIDs    []string `json:"databases_ids"`    // converted to snake_case from "databasesIds"
	RedisIDs        []string `json:"redis_ids"`        // converted to snake_case from "redisIds"
	ServiceIDs      []string `json:"service_ids"`      // converted to snake_case from "serviceIds"
	EnvGroupIDs     []string `json:"env_group_ids"`    // converted to snake_case from "envGroupIds"
	ProtectedStatus string   `json:"protected_status"` // converted to snake_case from "protectedStatus"
}

type EnvironmentDescription struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	ProjectID       string   `json:"project_id"`       // converted to snake_case from "projectId"
	DatabasesIDs    []string `json:"databases_ids"`    // converted to snake_case from "databasesIds"
	RedisIDs        []string `json:"redis_ids"`        // converted to snake_case from "redisIds"
	ServiceIDs      []string `json:"service_ids"`      // converted to snake_case from "serviceIds"
	EnvGroupIDs     []string `json:"env_group_ids"`    // converted to snake_case from "envGroupIds"
	ProtectedStatus string   `json:"protected_status"` // converted to snake_case from "protectedStatus"
}

type IPAllowJSON struct {
	CIDRBlock   string `json:"cidr_block"` // converted to snake_case from "cidrBlock"
	Description string `json:"description"`
}

type IPAllow struct {
	CIDRBlock   string `json:"cidr_block"` // converted to snake_case from "cidrBlock"
	Description string `json:"description"`
}

type ReadReplicaJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ReadReplica struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PostgresResponse struct {
	Postgres PostgresJSON `json:"postgres"`
	Cursor   string       `json:"cursor"`
}

type PostgresJSON struct {
	ID                      string            `json:"id"`
	IPAllowList             []IPAllowJSON     `json:"ip_allow_list"`              // converted to snake_case from "ipAllowList"
	CreatedAt               time.Time         `json:"created_at"`                 // converted to snake_case from "createdAt"
	UpdatedAt               time.Time         `json:"updated_at"`                 // converted to snake_case from "updatedAt"
	ExpiresAt               time.Time         `json:"expires_at"`                 // converted to snake_case from "expiresAt"
	DatabaseName            string            `json:"database_name"`              // converted to snake_case from "databaseName"
	DatabaseUser            string            `json:"database_user"`              // converted to snake_case from "databaseUser"
	EnvironmentID           string            `json:"environment_id"`             // converted to snake_case from "environmentId"
	HighAvailabilityEnabled bool              `json:"high_availability_enabled"`  // converted to snake_case from "highAvailabilityEnabled"
	Name                    string            `json:"name"`
	Owner                   OwnerJSON         `json:"owner"`
	Plan                    string            `json:"plan"`
	DiskSizeGB              int               `json:"disk_size_gb"`       // converted to snake_case from "diskSizeGB"
	PrimaryPostgresID       string            `json:"primary_postgres_id"`// converted to snake_case from "primaryPostgresID"
	Region                  string            `json:"region"`
	ReadReplicas            []ReadReplicaJSON `json:"read_replicas"`      // converted to snake_case from "readReplicas"
	Role                    string            `json:"role"`
	Status                  string            `json:"status"`
	Version                 string            `json:"version"`
	Suspended               string            `json:"suspended"`
	Suspenders              []string          `json:"suspenders"`
	DashboardURL            string            `json:"dashboard_url"` // converted to snake_case from "dashboardUrl"
}

type PostgresDescription struct {
	ID                      string
	IPAllowList             []IPAllow
	CreatedAt               time.Time
	UpdatedAt               time.Time
	ExpiresAt               time.Time
	DatabaseName            string
	DatabaseUser            string
	EnvironmentID           string
	HighAvailabilityEnabled bool
	Name                    string
	Owner                   Owner
	Plan                    string
	DiskSizeGB              int
	PrimaryPostgresID       string
	Region                  string
	ReadReplicas            []ReadReplica
	Role                    string
	Status                  string
	Version                 string
	Suspended               string
	Suspenders              []string
	DashboardURL            string
}

type BuildFilterJSON struct {
	Paths        []string `json:"paths"`
	IgnoredPaths []string `json:"ignored_paths"` // converted to snake_case from "ignoredPaths"
}

type BuildFilter struct {
	Paths        []string
	IgnoredPaths []string
}

type RegistryCredentialJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RegistryCredential struct {
	ID   string
	Name string
}

type ServiceDetailsJSON struct {
	BuildCommand string           `json:"build_command"` // converted to snake_case from "buildCommand"
	ParentServer ParentServerJSON `json:"parent_server"` // converted to snake_case from "parentServer"
	PublishPath  string           `json:"publish_path"`  // converted to snake_case from "publishPath"
	Previews     PreviewsJSON     `json:"previews"`
	URL          string           `json:"url"`
	BuildPlan    string           `json:"build_plan"` // converted to snake_case from "buildPlan"
}

type ServiceDetails struct {
	BuildCommand string
	ParentServer ParentServer
	PublishPath  string
	Previews     Previews
	URL          string
	BuildPlan    string
}

type ParentServerJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ParentServer struct {
	ID   string
	Name string
}

type PreviewsJSON struct {
	Generation string `json:"generation"`
}

type Previews struct {
	Generation string
}

type ServiceResponse struct {
	Service ServiceJSON `json:"service"`
	Cursor  string      `json:"cursor"`
}

type ServiceJSON struct {
	ID                 string                 `json:"id"`
	AutoDeploy         string                 `json:"auto_deploy"`         // converted to snake_case from "autoDeploy"
	Branch             string                 `json:"branch"`
	BuildFilter        BuildFilterJSON        `json:"build_filter"`        // converted to snake_case from "buildFilter"
	CreatedAt          time.Time              `json:"created_at"`          // converted to snake_case from "createdAt"
	DashboardURL       string                 `json:"dashboard_url"`       // converted to snake_case from "dashboardUrl"
	EnvironmentID      string                 `json:"environment_id"`      // converted to snake_case from "environmentId"
	ImagePath          string                 `json:"image_path"`          // converted to snake_case from "imagePath"
	Name               string                 `json:"name"`
	NotifyOnFail       string                 `json:"notify_on_fail"`      // converted to snake_case from "notifyOnFail"
	OwnerID            string                 `json:"owner_id"`            // converted to snake_case from "ownerId"
	RegistryCredential RegistryCredentialJSON `json:"registry_credential"` // converted to snake_case from "registryCredential"
	Repo               string                 `json:"repo"`
	RootDir            string                 `json:"root_dir"`            // converted to snake_case from "rootDir"
	Slug               string                 `json:"slug"`
	Suspended          string                 `json:"suspended"`
	Suspenders         []string               `json:"suspenders"`
	Type               string                 `json:"type"`
	UpdatedAt          time.Time              `json:"updated_at"`      // converted to snake_case from "updatedAt"
	ServiceDetails     ServiceDetailsJSON     `json:"service_details"` // converted to snake_case from "serviceDetails"
}

type ServiceDescription struct {
	ID                 string
	AutoDeploy         string
	Branch             string
	BuildFilter        BuildFilter
	CreatedAt          time.Time
	DashboardURL       string
	EnvironmentID      string
	ImagePath          string
	Name               string
	NotifyOnFail       string
	OwnerID            string
	RegistryCredential RegistryCredential
	Repo               string
	RootDir            string
	Slug               string
	Suspended          string
	Suspenders         []string
	Type               string
	UpdatedAt          time.Time
	ServiceDetails     ServiceDetails
}

type JobResponse struct {
	Job    JobJSON `json:"job"`
	Cursor string  `json:"cursor"`
}

type JobJSON struct {
	ID           string    `json:"id"`
	ServiceID    string    `json:"service_id"`    // converted to snake_case from "serviceId"
	StartCommand string    `json:"start_command"` // converted to snake_case from "startCommand"
	PlanID       string    `json:"plan_id"`       // converted to snake_case from "planId"
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`   // converted to snake_case from "createdAt"
	StartedAt    time.Time `json:"started_at"`   // converted to snake_case from "startedAt"
	FinishedAt   time.Time `json:"finished_at"`  // converted to snake_case from "finishedAt"
}

type JobDescription struct {
	ID           string
	ServiceID    string
	StartCommand string
	PlanID       string
	Status       string
	CreatedAt    time.Time
	StartedAt    time.Time
	FinishedAt   time.Time
}

type DiskResponse struct {
	Disk   DiskJSON `json:"disk"`
	Cursor string   `json:"cursor"`
}

type DiskJSON struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	SizeGB    int       `json:"size_gb"`    // converted to snake_case from "sizeGB"
	MountPath string    `json:"mount_path"` // converted to snake_case from "mountPath"
	ServiceID string    `json:"service_id"` // converted to snake_case from "serviceId"
	CreatedAt time.Time `json:"created_at"` // converted to snake_case from "createdAt"
	UpdatedAt time.Time `json:"updated_at"` // converted to snake_case from "updatedAt"
}

type DiskDescription struct {
	ID        string
	Name      string
	SizeGB    int
	MountPath string
	ServiceID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommitJSON struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"` // converted to snake_case from "createdAt"
}

type Commit struct {
	ID        string
	Message   string
	CreatedAt time.Time
}

type ImageJSON struct {
	Ref                string `json:"ref"`
	SHA                string `json:"sha"`
	RegistryCredential string `json:"registry_credential"` // converted to snake_case from "registryCredential"
}

type Image struct {
	Ref                string
	SHA                string
	RegistryCredential string
}

type DeployResponse struct {
	Deploy DeployJSON `json:"deploy"`
	Cursor string     `json:"cursor"`
}

type DeployJSON struct {
	ID         string     `json:"id"`
	Commit     CommitJSON `json:"commit"`
	Image      ImageJSON  `json:"image"`
	Status     string     `json:"status"`
	Trigger    string     `json:"trigger"`
	FinishedAt time.Time  `json:"finished_at"` // converted to snake_case from "finishedAt"
	CreatedAt  time.Time  `json:"created_at"`  // converted to snake_case from "createdAt"
	UpdatedAt  time.Time  `json:"updated_at"`  // converted to snake_case from "updatedAt"
}

type DeployDescription struct {
	ID         string
	Commit     Commit
	Image      Image
	Status     string
	Trigger    string
	FinishedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type BlueprintResponse struct {
	BluePrint BlueprintJSON `json:"blueprint"`
	Cursor    string        `json:"cursor"`
}

type BlueprintJSON struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	AutoSync bool      `json:"auto_sync"` // converted to snake_case from "autoSync"
	Repo     string    `json:"repo"`
	Branch   string    `json:"branch"`
	LastSync time.Time `json:"last_sync"` // converted to snake_case from "lastSync"
}

type BlueprintDescription struct {
	ID       string
	Name     string
	Status   string
	AutoSync bool
	Repo     string
	Branch   string
	LastSync time.Time
}

type ServiceLinkJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ServiceLink struct {
	ID   string
	Name string
	Type string
}

type EnvGroupResponse struct {
	EnvGroup EnvGroupJSON `json:"envGroup"`
	Cursor   string       `json:"cursor"`
}

type EnvGroupJSON struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	OwnerID       string            `json:"owner_id"`      // converted to snake_case from "ownerId"
	CreatedAt     time.Time         `json:"created_at"`    // converted to snake_case from "createdAt"
	UpdatedAt     time.Time         `json:"updated_at"`    // converted to snake_case from "updatedAt"
	ServiceLinks  []ServiceLinkJSON `json:"service_links"` // converted to snake_case from "serviceLinks"
	EnvironmentID string            `json:"environment_id"`// converted to snake_case from "environmentId"
}

type EnvGroupDescription struct {
	ID            string
	Name          string
	OwnerID       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ServiceLinks  []ServiceLink
	EnvironmentID string
}

type HeaderResponse struct {
	Header HeaderJSON `json:"header"`
	Cursor string     `json:"cursor"`
}

type HeaderJSON struct {
	ID    string `json:"id"`
	Path  string `json:"path"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HeaderDescription struct {
	ID    string
	Path  string
	Name  string
	Value string
}

type RouteResponse struct {
	Route  RouteJSON `json:"route"`
	Cursor string    `json:"cursor"`
}

type RouteJSON struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Priority    int    `json:"priority"`
}

type RouteDescription struct {
	ID          string
	Type        string
	Source      string
	Destination string
	Priority    int
}

type PostgresqlBackupJSON struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"` // converted to snake_case from "createdAt"
	URL       string `json:"url"`
}

type PostgresqlBackupDescription struct {
	ID        string
	CreatedAt string
	URL       string
}
