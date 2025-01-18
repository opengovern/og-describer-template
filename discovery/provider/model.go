//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package provider

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"time"
)

type Metadata struct {
	ID               string
	Name             string
	SubscriptionID   string
	Location         string
	CloudEnvironment string
	ResourceType     string
	IntegrationID    string
}

//  =================== msgraph ==================

//index:microsoft_resources_users
type AdUsersDescription struct {
	TenantID                        string
	DisplayName                     *string
	Id                              *string
	UserPrincipalName               *string
	AccountEnabled                  *bool
	UserType                        *string
	CreatedDateTime                 *time.Time
	Mail                            *string
	PasswordPolicies                *string
	SignInSessionsValidFromDateTime *time.Time
	UsageLocation                   *string
	MemberOf                        []string
	TransitiveMemberOf              []string
	LastSignInDateTime              *time.Time
	ImAddresses                     []string
	OtherMails                      []string
	JobTitle                        *string
	Identities                      []struct {
		SignInType       *string
		Issuer           *string
		IssuerAssignedId *string
	}
}

//index:microsoft_resources_groups
type AdGroupDescription struct {
	TenantID                      string
	DisplayName                   *string
	ID                            *string
	Description                   *string
	Classification                *string
	CreatedDateTime               *time.Time
	ExpirationDateTime            *time.Time
	IsAssignableToRole            *bool
	IsSubscribedByMail            *bool
	Mail                          *string
	MailEnabled                   *bool
	MailNickname                  *string
	MembershipRule                *string
	MembershipRuleProcessingState *string
	OnPremisesDomainName          *string
	OnPremisesLastSyncDateTime    *time.Time
	OnPremisesNetBiosName         *string
	OnPremisesSamAccountName      *string
	OnPremisesSecurityIdentifier  *string
	OnPremisesSyncEnabled         *bool
	RenewedDateTime               *time.Time
	SecurityEnabled               *bool
	SecurityIdentifier            *string
	Visibility                    *string
	AssignedLabels                []struct {
		DisplayName *string
		LabelId     *string
	}
	GroupTypes     []string
	MemberIds      []*string
	OwnerIds       []*string
	ProxyAddresses []string
	NestedGroups   []struct {
		GroupId     *string
		DisplayName *string
	}
}

//index:microsoft_resources_serviceprincipals
type AdServicePrincipalDescription struct {
	TenantID                  string
	Id                        *string
	DisplayName               *string
	AppId                     *string
	AccountEnabled            *bool
	AppDisplayName            *string
	AppRoleAssignmentRequired *bool
	AppOwnerOrganizationId    *string
	ServicePrincipalType      *string
	SignInAudience            *string
	AppDescription            *string
	Description               *string
	LoginUrl                  *string
	LogoutUrl                 *string
	AddIns                    []struct {
		Id          string
		TypeEscaped *string
		Properties  []struct {
			Key   *string
			Value *string
		}
	}
	AlternativeNames []string
	AppRoles         []struct {
		AllowedMemberTypes []string
		Description        *string
		DisplayName        *string
		Id                 string
		IsEnabled          *bool
		Origin             *string
		Value              *string
	}
	//Info models.InformationalUrlable
	KeyCredentials []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Key                 []byte
		KeyId               string
		StartDateTime       *time.Time
		TypeEscaped         *string
		Usage               *string
	}
	NotificationEmailAddresses []string
	OwnerIds                   []*string
	PasswordCredentials        []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Hint                *string
		KeyId               string
		SecretText          *string
		StartDateTime       *time.Time
	}
	Oauth2PermissionScopes []struct {
		AdminConsentDescription *string
		AdminConsentDisplayName *string
		Id                      string
		IsEnabled               *bool
		Origin                  *string
		TypeEscaped             *string
		UserConsentDescription  *string
		UserConsentDisplayName  *string
	}
	ReplyUrls             []string
	ServicePrincipalNames []string
	TagsSrc               []string
}

//index:microsoft_resources_applications
type AdApplicationDescription struct {
	TenantID                      string
	DisplayName                   *string
	Id                            *string
	AppId                         *string
	CreatedDateTime               *time.Time
	Description                   *string
	IsAuthorizationServiceEnabled *bool
	Oauth2RequirePostResponse     *bool
	PublisherDomain               *string
	SignInAudience                *string
	Api                           struct {
		AcceptMappedClaims      *bool
		KnownClientApplications []string
		Oauth2PermissionScopes  []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		PreAuthorizedApplications []struct {
			AppId                  *string
			DelegatedPermissionIds []string
		}
		RequestedAccessTokenVersion *int32
	}
	IdentifierUris []string
	Info           struct {
		LogoUrl             *string
		MarketingUrl        *string
		PrivacyStatementUrl *string
		SupportUrl          *string
		TermsOfServiceUrl   *string
	}
	KeyCredentials []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Key                 []byte
		KeyId               string
		StartDateTime       *time.Time
		TypeEscaped         *string
		Usage               *string
	}
	OwnerIds                []*string
	ParentalControlSettings struct {
		CountriesBlockedForMinors []string
		LegalAgeGroupRule         *string
	}
	PasswordCredentials []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Hint                *string
		KeyId               string
		SecretText          *string
		StartDateTime       *time.Time
	}
	Spa struct {
		RedirectUris []string
	}
	TagsSrc []string
	Web     struct {
		HomePageUrl           *string
		ImplicitGrantSettings struct {
			EnableAccessTokenIssuance *bool
			EnableIdTokenIssuance     *bool
		}
		LogoutUrl           *string
		RedirectUris        []string
		RedirectUriSettings []struct {
			Index *int32
			Uri   *string
		}
	}
}

//index:microsoft_resources_directoryroles
type AdDirectoryRoleDescription struct {
	TenantID       string
	DisplayName    *string
	Id             *string
	Description    *string
	RoleTemplateId *string
	MemberIds      []*string
}

//index:microsoft_resources_directorysettings
type AdDirectorySettingDescription struct {
	TenantID    string
	DisplayName *string
	Id          *string
	TemplateId  *string
	Name        *string
	Value       *string
}

//index:microsoft_resources_directoryauditreports
type AdDirectoryAuditReportDescription struct {
	TenantID            string
	Id                  *string
	ActivityDateTime    *time.Time
	ActivityDisplayName *string
	Category            *string
	CorrelationId       *string
	LoggedByService     *string
	OperationType       *string
	Result              string
	ResultReason        *string
	InitiatedBy         struct {
		App struct {
			AppId                *string
			DisplayName          *string
			ServicePrincipalId   *string
			ServicePrincipalName *string
		}
		User struct {
			Id                *string
			DisplayName       *string
			IpAddress         *string
			UserPrincipalName *string
		}
	}
	TargetResources []struct {
		DisplayName        *string
		GroupType          string
		Id                 *string
		ModifiedProperties []struct {
			DisplayName *string
			NewValue    *string
			OldValue    *string
		}
		TypeEscaped       *string
		UserPrincipalName *string
	}
}

//index:microsoft_resources_directorysettings
type AdDomainDescription struct {
	TenantID           string
	Id                 *string
	AuthenticationType *string
	IsDefault          *bool
	IsAdminManaged     *bool
	IsInitial          *bool
	IsRoot             *bool
	IsVerified         *bool
	SupportedServices  []string
}

//index:microsoft_resources_tenant
type AdTenantDescription struct {
	TenantID        *string
	DisplayName     *string
	TenantType      *string
	CreatedDateTime *time.Time
	VerifiedDomains []struct {
		Name         *string
		Type         *string
		Capabilities *string
		IsDefault    *bool
		IsInitial    *bool
	}
	OnPremisesSyncEnabled *bool
}

//index:microsoft_resources_directorysettings
type AdIdentityProviderDescription struct {
	TenantID     string
	Id           *string
	DisplayName  *string
	Type         *string
	ClientId     interface{}
	ClientSecret interface{}
}

//index:microsoft_resources_securitydefaultspolicy
type AdSecurityDefaultsPolicyDescription struct {
	TenantID    string
	Id          *string
	DisplayName *string
	IsEnabled   *bool
	Description *string
}

//index:microsoft_resources_authorizationpolicy
type AdAuthorizationPolicyDescription struct {
	TenantID                                    string
	Id                                          *string
	DisplayName                                 *string
	Description                                 *string
	AllowedToSignIpEmailBasedSubscriptions      *bool
	AllowedToUseSspr                            *bool
	AllowedEmailVerifiedUsersToJoinOrganization *bool
	AllowInvitesFrom                            string
	BlockMsolPowershell                         *bool
	GuestUserRoleId                             string
	DefaultUserRolePermissions                  struct {
		AllowedToCreateApps                      *bool
		AllowedToCreateSecurityGroups            *bool
		AllowedToCreateTenants                   *bool
		AllowedToReadBitlockerKeysForOwnedDevice *bool
		AllowedToReadOtherUsers                  *bool
		OdataType                                *string
		PermissionGrantPoliciesAssigned          []string
	}
}

//index:microsoft_resources_conditionalaccesspolicy
type AdConditionalAccessPolicyDescription struct {
	TenantID         string
	Id               *string
	DisplayName      *string
	State            string
	CreatedDateTime  *time.Time
	ModifiedDateTime *time.Time
	Operator         *string
	Applications     struct {
		ApplicationFilter struct {
			Mode *string
			Rule *string
		}
		ExcludeApplications                         []string
		IncludeApplications                         []string
		IncludeAuthenticationContextClassReferences []string
		IncludeUserActions                          []string
	}
	ApplicationEnforcedRestrictions struct {
		IsEnabled *bool
	}
	BuiltInControls             []string
	ClientAppTypes              []string
	CustomAuthenticationFactors []string
	CloudAppSecurity            struct {
		CloudAppSecurityType string
		IsEnabled            *bool
	}
	Locations struct {
		ExcludeLocations []string
		IncludeLocations []string
	}
	PersistentBrowser struct {
		IsEnabled *bool
		Mode      string
	}
	Platforms struct {
		ExcludePlatforms []string
		IncludePlatforms []string
	}
	SignInFrequency struct {
		AuthenticationType string
		FrequencyInterval  string
		TypeEscaped        string
		Value              *int32
		IsEnabled          *bool
	}
	SignInRiskLevels []string
	TermsOfUse       []string
	Users            struct {
		ExcludeGroups []string
		IncludeGroups []string
		ExcludeUsers  []string
		IncludeUsers  []string
		ExcludeRoles  []string
		IncludeRoles  []string
	}
	UserRiskLevel []string
}

//index:microsoft_resources_adminconsentrequestpolicy
type AdAdminConsentRequestPolicyDescription struct {
	TenantID              string
	Id                    *string
	IsEnabled             *bool
	NotifyReviewers       *bool
	RemindersEnabled      *bool
	RequestDurationInDays *int32
	Version               *int32
	Reviewers             []struct {
		OdataType *string
		Query     *string
		QueryRoot *string
		QueryType *string
	}
}

//index:microsoft_resources_signinreport
type AdSignInReportDescription struct {
	TenantID                string
	Id                      *string
	CreatedDateTime         *time.Time
	UserDisplayName         *string
	UserPrincipalName       *string
	UserId                  *string
	AppId                   *string
	AppDisplayName          *string
	IpAddress               *string
	ClientAppUsed           *string
	CorrelationId           *string
	ConditionalAccessStatus *models.ConditionalAccessStatus
	IsInteractive           *bool
	RiskDetail              *models.RiskDetail
	RiskLevelAggregated     *models.RiskLevel
	RiskLevelDuringSignIn   *models.RiskLevel
	RiskState               *models.RiskState
	ResourceDisplayName     *string
	ResourceId              *string
	RiskEventTypes          []models.RiskEventType
	Status                  struct {
		ErrorCode     *int32
		FailureReason *string
	}
	DeviceDetail struct {
		Browser         *string
		DeviceId        *string
		DisplayName     *string
		IsCompliant     *bool
		IsManaged       *bool
		OperatingSystem *string
		TrustType       *string
	}
	Location struct {
		City            *string
		CountryOrRegion *string
		GeoCoordinates  struct {
			Altitude  *float64
			Latitude  *float64
			Longitude *float64
		}
		State *string
	}
	AppliedConditionalAccessPolicies []struct {
		DisplayName             *string
		EnforcedGrantControls   []string
		EnforcedSessionControls []string
		Id                      *string
		Result                  string
	}
}

//index:microsoft_resources_device
type AdDeviceDescription struct {
	TenantID                      string
	Id                            *string
	DisplayName                   *string
	AccountEnabled                *bool
	DeviceId                      *string
	ApproximateLastSignInDateTime *time.Time
	IsCompliant                   *bool
	IsManaged                     *bool
	MdmAppId                      *string
	OperatingSystem               *string
	OperatingSystemVersion        *string
	ProfileType                   *string
	TrustType                     *string
	ExtensionAttributes           []models.Extensionable
	MemberOf                      []models.DirectoryObjectable
}

//index:microsoft_resources_userregistrationdetails
type AdUserRegistrationDetailsDescription struct {
	TenantID                                      string
	Id                                            *string
	UserPrincipalName                             *string
	UserDisplayName                               *string
	UserType                                      string
	SystemPreferredAuthenticationMethods          []string
	IsSystemPreferredAuthenticationMethodEnabled  *bool
	IsAdmin                                       *bool
	IsMfaCapable                                  *bool
	IsMfaRegistered                               *bool
	IsSsprCapable                                 *bool
	IsSsprRegistered                              *bool
	IsSsprEnabled                                 *bool
	IsPasswordlessCapable                         *bool
	LastUpdatedDateTime                           *time.Time
	MethodsRegistered                             []string
	UserPreferredMethodForSecondaryAuthentication string
}

//index:microsoft_resources_groups_memberships
type AdGroupMembershipDescription struct {
	TenantID           string
	GroupId            *string
	Id                 *string
	DisplayName        *string
	AccountEnabled     *bool
	UserPrincipalName  *string
	UserType           *string
	State              *string
	SecurityIdentifier *string
	ProxyAddresses     []string
	Mail               *string
}

//index:microsoft_resources_appregistration
type AdAppRegistrationDescription struct {
	TenantID                      string
	DisplayName                   *string
	Id                            *string
	AppId                         *string
	CreatedDateTime               *time.Time
	Description                   *string
	IsAuthorizationServiceEnabled *bool
	Oauth2RequirePostResponse     *bool
	PublisherDomain               *string
	SignInAudience                *string
	Api                           struct {
		AcceptMappedClaims      *bool
		KnownClientApplications []string
		Oauth2PermissionScopes  []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		PreAuthorizedApplications []struct {
			AppId                  *string
			DelegatedPermissionIds []string
		}
		RequestedAccessTokenVersion *int32
	}
	IdentifierUris []string
	Info           struct {
		LogoUrl             *string
		MarketingUrl        *string
		PrivacyStatementUrl *string
		SupportUrl          *string
		TermsOfServiceUrl   *string
	}
	KeyCredentials []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Key                 []byte
		KeyId               string
		StartDateTime       *time.Time
		TypeEscaped         *string
		Usage               *string
	}
	OwnerIds                []*string
	ParentalControlSettings struct {
		CountriesBlockedForMinors []string
		LegalAgeGroupRule         *string
	}
	PasswordCredentials []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Hint                *string
		KeyId               string
		SecretText          *string
		StartDateTime       *time.Time
	}
	Spa struct {
		RedirectUris []string
	}
	TagsSrc []string
	Web     struct {
		HomePageUrl           *string
		ImplicitGrantSettings struct {
			EnableAccessTokenIssuance *bool
			EnableIdTokenIssuance     *bool
		}
		LogoutUrl           *string
		RedirectUris        []string
		RedirectUriSettings []struct {
			Index *int32
			Uri   *string
		}
	}
}

//index:microsoft_resources_enterpriseapplications
type AdEnterpriseApplicationDescription struct {
	TenantID                  string
	Id                        *string
	DisplayName               *string
	AppId                     *string
	AccountEnabled            *bool
	AppDisplayName            *string
	AppRoleAssignmentRequired *bool
	AppOwnerOrganizationId    *string
	ServicePrincipalType      *string
	SignInAudience            *string
	AppDescription            *string
	Description               *string
	LoginUrl                  *string
	LogoutUrl                 *string
	AddIns                    []struct {
		Id          string
		TypeEscaped *string
		Properties  []struct {
			Key   *string
			Value *string
		}
	}
	AlternativeNames []string
	AppRoles         []struct {
		AllowedMemberTypes []string
		Description        *string
		DisplayName        *string
		Id                 string
		IsEnabled          *bool
		Origin             *string
		Value              *string
	}
	//Info models.InformationalUrlable
	KeyCredentials []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Key                 []byte
		KeyId               string
		StartDateTime       *time.Time
		TypeEscaped         *string
		Usage               *string
	}
	NotificationEmailAddresses []string
	OwnerIds                   []*string
	PasswordCredentials        []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Hint                *string
		KeyId               string
		SecretText          *string
		StartDateTime       *time.Time
	}
	Oauth2PermissionScopes []struct {
		AdminConsentDescription *string
		AdminConsentDisplayName *string
		Id                      string
		IsEnabled               *bool
		Origin                  *string
		TypeEscaped             *string
		UserConsentDescription  *string
		UserConsentDisplayName  *string
	}
	ReplyUrls             []string
	ServicePrincipalNames []string
	TagsSrc               []string
}

//index:microsoft_resources_managedidentity
type AdManagedIdentityDescription struct {
	TenantID                  string
	Id                        *string
	DisplayName               *string
	AppId                     *string
	AccountEnabled            *bool
	AppDisplayName            *string
	AppRoleAssignmentRequired *bool
	AppOwnerOrganizationId    *string
	ServicePrincipalType      *string
	SignInAudience            *string
	AppDescription            *string
	Description               *string
	LoginUrl                  *string
	LogoutUrl                 *string
	IdentityType              string
	AddIns                    []struct {
		Id          string
		TypeEscaped *string
		Properties  []struct {
			Key   *string
			Value *string
		}
	}
	AlternativeNames []string
	AppRoles         []struct {
		AllowedMemberTypes []string
		Description        *string
		DisplayName        *string
		Id                 string
		IsEnabled          *bool
		Origin             *string
		Value              *string
	}
	//Info models.InformationalUrlable
	KeyCredentials []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Key                 []byte
		KeyId               string
		StartDateTime       *time.Time
		TypeEscaped         *string
		Usage               *string
	}
	NotificationEmailAddresses []string
	OwnerIds                   []*string
	PasswordCredentials        []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Hint                *string
		KeyId               string
		SecretText          *string
		StartDateTime       *time.Time
	}
	Oauth2PermissionScopes []struct {
		AdminConsentDescription *string
		AdminConsentDisplayName *string
		Id                      string
		IsEnabled               *bool
		Origin                  *string
		TypeEscaped             *string
		UserConsentDescription  *string
		UserConsentDisplayName  *string
	}
	ReplyUrls             []string
	ServicePrincipalNames []string
	TagsSrc               []string
}

//index:microsoft_resources_microsoftapplication
type AdMicrosoftApplicationDescription struct {
	TenantID                  string
	Id                        *string
	DisplayName               *string
	AppId                     *string
	AccountEnabled            *bool
	AppDisplayName            *string
	AppRoleAssignmentRequired *bool
	AppOwnerOrganizationId    *string
	ServicePrincipalType      *string
	SignInAudience            *string
	AppDescription            *string
	Description               *string
	LoginUrl                  *string
	LogoutUrl                 *string
	AddIns                    []struct {
		Id          string
		TypeEscaped *string
		Properties  []struct {
			Key   *string
			Value *string
		}
	}
	AlternativeNames []string
	AppRoles         []struct {
		AllowedMemberTypes []string
		Description        *string
		DisplayName        *string
		Id                 string
		IsEnabled          *bool
		Origin             *string
		Value              *string
	}
	//Info models.InformationalUrlable
	KeyCredentials []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Key                 []byte
		KeyId               string
		StartDateTime       *time.Time
		TypeEscaped         *string
		Usage               *string
	}
	NotificationEmailAddresses []string
	OwnerIds                   []*string
	PasswordCredentials        []struct {
		CustomKeyIdentifier []byte
		DisplayName         *string
		EndDateTime         *time.Time
		Hint                *string
		KeyId               string
		SecretText          *string
		StartDateTime       *time.Time
	}
	Oauth2PermissionScopes []struct {
		AdminConsentDescription *string
		AdminConsentDisplayName *string
		Id                      string
		IsEnabled               *bool
		Origin                  *string
		TypeEscaped             *string
		UserConsentDescription  *string
		UserConsentDisplayName  *string
	}
	ReplyUrls             []string
	ServicePrincipalNames []string
	TagsSrc               []string
}
