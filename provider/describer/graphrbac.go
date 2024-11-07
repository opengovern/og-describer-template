package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/aws/aws-sdk-go-v2/aws"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/applications"
	"github.com/microsoftgraph/msgraph-sdk-go/auditlogs"
	"github.com/microsoftgraph/msgraph-sdk-go/devices"
	"github.com/microsoftgraph/msgraph-sdk-go/directoryroles"
	"github.com/microsoftgraph/msgraph-sdk-go/domains"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/groupsettings"
	"github.com/microsoftgraph/msgraph-sdk-go/identity"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/organization"
	"github.com/microsoftgraph/msgraph-sdk-go/policies"
	"github.com/microsoftgraph/msgraph-sdk-go/reports"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
	users2 "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/opengovern/og-azure-describer/azure/model"
	models2 "github.com/opengovern/og-describer-entraid/pkg/sdk/models"
	"strings"
	"time"
)

func AdUsers(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Users().Get(ctx, &users2.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: &users2.UsersRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
			Select: []string{"displayName", "userPrincipalName", "userType", "givenName", "surname",
				"onPremisesImmutableId", "mail", "mailNickname", "passwordPolicies", "signInSessionsValidFromDateTime",
				"usageLocation", "imAddresses", "otherMails", "jobTitle", "identities",
				"memberOf", "accountEnabled", "transitiveMemberOf", "createdDateTime", "signInActivity"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	var values []models2.Resource
	var itemErr error
	pageIterator, err := msgraphcore.NewPageIterator[models.Userable](result, client.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(user models.Userable) bool {
		if user == nil {
			return true
		}
		var id string
		if user.GetId() != nil {
			id = *user.GetId()
		}
		var name string
		if user.GetDisplayName() != nil {
			name = *user.GetDisplayName()
		}

		var memberOf []string
		for _, m := range user.GetMemberOf() {
			memberOf = append(memberOf, *m.GetId())
		}
		var transitiveMemberOf []string
		for _, m := range user.GetTransitiveMemberOf() {
			memberOf = append(memberOf, *m.GetId())
		}

		var lastSignInDateTime *time.Time
		if user.GetSignInActivity() != nil {
			lastSignInDateTime = user.GetSignInActivity().GetLastSignInDateTime()
		}

		var identities []struct {
			SignInType       *string
			Issuer           *string
			IssuerAssignedId *string
		}
		for _, i := range user.GetIdentities() {
			identities = append(identities, struct {
				SignInType       *string
				Issuer           *string
				IssuerAssignedId *string
			}{
				Issuer:           i.GetIssuer(),
				SignInType:       i.GetSignInType(),
				IssuerAssignedId: i.GetIssuerAssignedId(),
			})
		}

		resource := models2.Resource{
			ID:       id,
			Name:     name,
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdUsersDescription{
					TenantID:                        tenantId,
					DisplayName:                     user.GetDisplayName(),
					Id:                              user.GetId(),
					UserPrincipalName:               user.GetUserPrincipalName(),
					AccountEnabled:                  user.GetAccountEnabled(),
					UserType:                        user.GetUserType(),
					CreatedDateTime:                 user.GetCreatedDateTime(),
					Mail:                            user.GetMail(),
					PasswordPolicies:                user.GetPasswordPolicies(),
					SignInSessionsValidFromDateTime: user.GetSignInSessionsValidFromDateTime(),
					UsageLocation:                   user.GetUsageLocation(),
					MemberOf:                        memberOf,
					TransitiveMemberOf:              transitiveMemberOf,
					LastSignInDateTime:              lastSignInDateTime,
					ImAddresses:                     user.GetImAddresses(),
					OtherMails:                      user.GetOtherMails(),
					JobTitle:                        user.GetJobTitle(),
					Identities:                      identities,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Groups().Get(ctx, &groups.GroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.GroupsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}

	var values []models2.Resource
	var itemErr error
	pageIterator, err := msgraphcore.NewPageIterator[models.Groupable](result, client.GetAdapter(), models.CreateGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(group models.Groupable) bool {
		if group == nil {
			return true
		}
		var memberIds []*string
		for _, m := range group.GetMembers() {
			memberIds = append(memberIds, m.GetId())
		}
		var ownerIds []*string
		for _, m := range group.GetOwners() {
			ownerIds = append(ownerIds, m.GetId())
		}

		var assignedLabels []struct {
			DisplayName *string
			LabelId     *string
		}
		for _, l := range group.GetAssignedLabels() {
			assignedLabels = append(assignedLabels, struct {
				DisplayName *string
				LabelId     *string
			}{
				DisplayName: l.GetDisplayName(),
				LabelId:     l.GetLabelId(),
			})
		}

		var nestedGroups []struct {
			GroupId     *string
			DisplayName *string
		}
		members, err := client.Groups().ByGroupId(*group.GetId()).TransitiveMembers().GraphGroup().Get(ctx, &groups.ItemTransitiveMembersGraphGroupRequestBuilderGetRequestConfiguration{
			QueryParameters: &groups.ItemTransitiveMembersGraphGroupRequestBuilderGetQueryParameters{
				Top: aws.Int32(999),
			},
		})
		if err != nil {
			itemErr = err
			return false
		}
		for _, m := range members.GetValue() {
			nestedGroups = append(nestedGroups, struct {
				GroupId     *string
				DisplayName *string
			}{
				GroupId:     m.GetId(),
				DisplayName: m.GetDisplayName(),
			})
		}

		resource := models2.Resource{
			ID:       *group.GetId(),
			Name:     *group.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdGroupDescription{
					TenantID:                      tenantId,
					DisplayName:                   group.GetDisplayName(),
					ID:                            group.GetId(),
					Description:                   group.GetDescription(),
					Classification:                group.GetClassification(),
					CreatedDateTime:               group.GetCreatedDateTime(),
					ExpirationDateTime:            group.GetExpirationDateTime(),
					IsAssignableToRole:            group.GetIsAssignableToRole(),
					IsSubscribedByMail:            group.GetIsSubscribedByMail(),
					Mail:                          group.GetMail(),
					MailEnabled:                   group.GetMailEnabled(),
					MailNickname:                  group.GetMailNickname(),
					MembershipRule:                group.GetMembershipRule(),
					MembershipRuleProcessingState: group.GetMembershipRuleProcessingState(),
					OnPremisesDomainName:          group.GetOnPremisesDomainName(),
					OnPremisesLastSyncDateTime:    group.GetOnPremisesLastSyncDateTime(),
					OnPremisesNetBiosName:         group.GetOnPremisesNetBiosName(),
					OnPremisesSamAccountName:      group.GetOnPremisesSamAccountName(),
					OnPremisesSecurityIdentifier:  group.GetOnPremisesSecurityIdentifier(),
					OnPremisesSyncEnabled:         group.GetOnPremisesSyncEnabled(),
					RenewedDateTime:               group.GetRenewedDateTime(),
					SecurityEnabled:               group.GetSecurityEnabled(),
					SecurityIdentifier:            group.GetSecurityIdentifier(),
					Visibility:                    group.GetVisibility(),
					AssignedLabels:                assignedLabels,
					GroupTypes:                    group.GetGroupTypes(),
					MemberIds:                     memberIds,
					OwnerIds:                      ownerIds,
					ProxyAddresses:                group.GetProxyAddresses(),
					//ResourceBehaviorOptions:       group.GetResourceBehaviorOptions(),
					//ResourceProvisioningOptions:   group.GetResourceProvisioningOptions(),
					NestedGroups: nestedGroups,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdServicePrinciple(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []models2.Resource
	var itemErr error

	resultApps, err := client.ServicePrincipals().Get(ctx, &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	appPageIterator, err := msgraphcore.NewPageIterator[*models.ServicePrincipal](resultApps, client.GetAdapter(), models.CreateApplicationCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, fmt.Errorf("failed to query apps client: %v", err)
	}
	err = appPageIterator.Iterate(context.Background(), func(servicePrincipal *models.ServicePrincipal) bool {
		if servicePrincipal == nil || servicePrincipal.GetAppId() == nil {
			return true
		}

		var keyCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Key                 []byte
			KeyId               string
			StartDateTime       *time.Time
			TypeEscaped         *string
			Usage               *string
		}
		var passwordCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Hint                *string
			KeyId               string
			SecretText          *string
			StartDateTime       *time.Time
		}
		var ownerIds []*string
		var addIns []struct {
			Id          string
			TypeEscaped *string
			Properties  []struct {
				Key   *string
				Value *string
			}
		}
		var appRoles []struct {
			AllowedMemberTypes []string
			Description        *string
			DisplayName        *string
			Id                 string
			IsEnabled          *bool
			Origin             *string
			Value              *string
		}
		apps, err := client.Applications().Get(ctx, &applications.ApplicationsRequestBuilderGetRequestConfiguration{
			QueryParameters: &applications.ApplicationsRequestBuilderGetQueryParameters{
				Top:    aws.Int32(999),
				Filter: aws.String(fmt.Sprintf("appId eq '%s'", *servicePrincipal.GetAppId())),
			},
		})
		if err == nil && apps.GetValue() != nil && len(apps.GetValue()) > 0 {
			application := apps.GetValue()[0]
			for _, kc := range application.GetKeyCredentials() {
				keyCredentials = append(keyCredentials, struct {
					CustomKeyIdentifier []byte
					DisplayName         *string
					EndDateTime         *time.Time
					Key                 []byte
					KeyId               string
					StartDateTime       *time.Time
					TypeEscaped         *string
					Usage               *string
				}{
					Key:                 kc.GetKey(),
					TypeEscaped:         kc.GetTypeEscaped(),
					Usage:               kc.GetUsage(),
					DisplayName:         kc.GetDisplayName(),
					CustomKeyIdentifier: kc.GetCustomKeyIdentifier(),
					KeyId:               kc.GetKeyId().String(),
					EndDateTime:         kc.GetEndDateTime(),
					StartDateTime:       kc.GetStartDateTime(),
				})
			}
			for _, pc := range application.GetPasswordCredentials() {
				passwordCredentials = append(passwordCredentials, struct {
					CustomKeyIdentifier []byte
					DisplayName         *string
					EndDateTime         *time.Time
					Hint                *string
					KeyId               string
					SecretText          *string
					StartDateTime       *time.Time
				}{
					CustomKeyIdentifier: pc.GetCustomKeyIdentifier(),
					DisplayName:         pc.GetDisplayName(),
					EndDateTime:         pc.GetEndDateTime(),
					Hint:                pc.GetHint(),
					KeyId:               pc.GetKeyId().String(),
					SecretText:          pc.GetSecretText(),
					StartDateTime:       pc.GetStartDateTime(),
				})
			}
			for _, owner := range application.GetOwners() {
				ownerIds = append(ownerIds, owner.GetId())
			}
			for _, addIn := range application.GetAddIns() {
				var properties []struct {
					Key   *string
					Value *string
				}
				for _, p := range addIn.GetProperties() {
					properties = append(properties, struct {
						Key   *string
						Value *string
					}{
						Key:   p.GetKey(),
						Value: p.GetValue(),
					})
				}
				addIns = append(addIns, struct {
					Id          string
					TypeEscaped *string
					Properties  []struct {
						Key   *string
						Value *string
					}
				}{
					Id:          addIn.GetId().String(),
					TypeEscaped: addIn.GetTypeEscaped(),
					Properties:  properties,
				})
			}
			for _, appRole := range application.GetAppRoles() {
				appRoles = append(appRoles, struct {
					AllowedMemberTypes []string
					Description        *string
					DisplayName        *string
					Id                 string
					IsEnabled          *bool
					Origin             *string
					Value              *string
				}{
					Id:                 appRole.GetId().String(),
					Description:        appRole.GetDescription(),
					DisplayName:        appRole.GetDisplayName(),
					AllowedMemberTypes: appRole.GetAllowedMemberTypes(),
					IsEnabled:          appRole.GetIsEnabled(),
					Origin:             appRole.GetOrigin(),
					Value:              appRole.GetValue(),
				})
			}
		}

		var orgID *string
		v := servicePrincipal.GetAppOwnerOrganizationId()
		if v != nil {
			tmp := v.String()
			orgID = &tmp
		}

		var oauth2PermissionScopes []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		for _, ps := range servicePrincipal.GetOauth2PermissionScopes() {
			oauth2PermissionScopes = append(oauth2PermissionScopes, struct {
				AdminConsentDescription *string
				AdminConsentDisplayName *string
				Id                      string
				IsEnabled               *bool
				Origin                  *string
				TypeEscaped             *string
				UserConsentDescription  *string
				UserConsentDisplayName  *string
			}{
				Id:                      ps.GetId().String(),
				Origin:                  ps.GetOrigin(),
				IsEnabled:               ps.GetIsEnabled(),
				TypeEscaped:             ps.GetTypeEscaped(),
				AdminConsentDescription: ps.GetAdminConsentDescription(),
				UserConsentDescription:  ps.GetUserConsentDescription(),
				UserConsentDisplayName:  ps.GetUserConsentDisplayName(),
				AdminConsentDisplayName: ps.GetAdminConsentDisplayName(),
			})
		}

		resource := models2.Resource{
			ID:       *servicePrincipal.GetId(),
			Name:     *servicePrincipal.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdServicePrincipalDescription{
					TenantID:                  tenantId,
					Id:                        servicePrincipal.GetId(),
					DisplayName:               servicePrincipal.GetDisplayName(),
					AppId:                     servicePrincipal.GetAppId(),
					AccountEnabled:            servicePrincipal.GetAccountEnabled(),
					AppDisplayName:            servicePrincipal.GetAppDisplayName(),
					AppOwnerOrganizationId:    orgID,
					AppRoleAssignmentRequired: servicePrincipal.GetAppRoleAssignmentRequired(),
					ServicePrincipalType:      servicePrincipal.GetServicePrincipalType(),
					SignInAudience:            servicePrincipal.GetSignInAudience(),
					AppDescription:            servicePrincipal.GetAppDescription(),
					Description:               servicePrincipal.GetDescription(),
					LoginUrl:                  servicePrincipal.GetLoginUrl(),
					LogoutUrl:                 servicePrincipal.GetLogoutUrl(),
					AddIns:                    addIns,
					AlternativeNames:          servicePrincipal.GetAlternativeNames(),
					AppRoles:                  appRoles,
					//Info: servicePrincipal.GetInfo(),
					KeyCredentials:             keyCredentials,
					NotificationEmailAddresses: servicePrincipal.GetNotificationEmailAddresses(),
					OwnerIds:                   ownerIds,
					PasswordCredentials:        passwordCredentials,
					Oauth2PermissionScopes:     oauth2PermissionScopes,
					ReplyUrls:                  servicePrincipal.GetReplyUrls(),
					ServicePrincipalNames:      servicePrincipal.GetServicePrincipalNames(),
					TagsSrc:                    servicePrincipal.GetTags(),
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				itemErr = fmt.Errorf("failed to stream: %v", itemErr)
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdApplication(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []models2.Resource
	var itemErr error
	result, err := client.Applications().Get(ctx, &applications.ApplicationsRequestBuilderGetRequestConfiguration{
		QueryParameters: &applications.ApplicationsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	pageIterator, err := msgraphcore.NewPageIterator[models.Applicationable](result, client.GetAdapter(), models.CreateApplicationCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(app models.Applicationable) bool {
		if app == nil {
			return true
		}

		var oauth2PermissionScopes []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		for _, ps := range app.GetApi().GetOauth2PermissionScopes() {
			oauth2PermissionScopes = append(oauth2PermissionScopes, struct {
				AdminConsentDescription *string
				AdminConsentDisplayName *string
				Id                      string
				IsEnabled               *bool
				Origin                  *string
				TypeEscaped             *string
				UserConsentDescription  *string
				UserConsentDisplayName  *string
			}{
				Id:                      ps.GetId().String(),
				Origin:                  ps.GetOrigin(),
				IsEnabled:               ps.GetIsEnabled(),
				TypeEscaped:             ps.GetTypeEscaped(),
				AdminConsentDescription: ps.GetAdminConsentDescription(),
				UserConsentDescription:  ps.GetUserConsentDescription(),
				UserConsentDisplayName:  ps.GetUserConsentDisplayName(),
				AdminConsentDisplayName: ps.GetAdminConsentDisplayName(),
			})
		}

		var preAuthorizedApplications []struct {
			AppId                  *string
			DelegatedPermissionIds []string
		}
		for _, paa := range app.GetApi().GetPreAuthorizedApplications() {
			preAuthorizedApplications = append(preAuthorizedApplications, struct {
				AppId                  *string
				DelegatedPermissionIds []string
			}{
				AppId:                  paa.GetAppId(),
				DelegatedPermissionIds: paa.GetDelegatedPermissionIds(),
			})
		}

		var knownClientApplications []string
		for _, a := range app.GetApi().GetKnownClientApplications() {
			knownClientApplications = append(knownClientApplications, a.String())
		}

		var keyCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Key                 []byte
			KeyId               string
			StartDateTime       *time.Time
			TypeEscaped         *string
			Usage               *string
		}
		for _, c := range app.GetKeyCredentials() {
			keyCredentials = append(keyCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Key                 []byte
				KeyId               string
				StartDateTime       *time.Time
				TypeEscaped         *string
				Usage               *string
			}{
				TypeEscaped:         c.GetTypeEscaped(),
				Key:                 c.GetKey(),
				DisplayName:         c.GetDisplayName(),
				StartDateTime:       c.GetStartDateTime(),
				KeyId:               c.GetKeyId().String(),
				EndDateTime:         c.GetEndDateTime(),
				Usage:               c.GetUsage(),
				CustomKeyIdentifier: c.GetCustomKeyIdentifier(),
			})
		}

		var ownerIds []*string
		for _, o := range app.GetOwners() {
			ownerIds = append(ownerIds, o.GetId())
		}

		var redirectUriSettings []struct {
			Index *int32
			Uri   *string
		}
		for _, r := range app.GetWeb().GetRedirectUriSettings() {
			redirectUriSettings = append(redirectUriSettings, struct {
				Index *int32
				Uri   *string
			}{
				Uri:   r.GetUri(),
				Index: r.GetIndex(),
			})
		}

		var passwordCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Hint                *string
			KeyId               string
			SecretText          *string
			StartDateTime       *time.Time
		}
		for _, pc := range app.GetPasswordCredentials() {
			passwordCredentials = append(passwordCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Hint                *string
				KeyId               string
				SecretText          *string
				StartDateTime       *time.Time
			}{
				CustomKeyIdentifier: pc.GetCustomKeyIdentifier(),
				DisplayName:         pc.GetDisplayName(),
				EndDateTime:         pc.GetEndDateTime(),
				Hint:                pc.GetHint(),
				KeyId:               pc.GetKeyId().String(),
				SecretText:          pc.GetSecretText(),
				StartDateTime:       pc.GetStartDateTime(),
			})
		}

		resource := models2.Resource{
			ID:       *app.GetId(),
			Name:     *app.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdApplicationDescription{
					TenantID:                  tenantId,
					DisplayName:               app.GetDisplayName(),
					Id:                        app.GetId(),
					AppId:                     app.GetAppId(),
					CreatedDateTime:           app.GetCreatedDateTime(),
					Description:               app.GetDescription(),
					Oauth2RequirePostResponse: app.GetOauth2RequirePostResponse(),
					PublisherDomain:           app.GetPublisherDomain(),
					SignInAudience:            app.GetSignInAudience(),
					Api: struct {
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
					}{
						AcceptMappedClaims:          app.GetApi().GetAcceptMappedClaims(),
						KnownClientApplications:     knownClientApplications,
						PreAuthorizedApplications:   preAuthorizedApplications,
						Oauth2PermissionScopes:      oauth2PermissionScopes,
						RequestedAccessTokenVersion: app.GetApi().GetRequestedAccessTokenVersion(),
					},
					IdentifierUris: app.GetIdentifierUris(),
					Info: struct {
						LogoUrl             *string
						MarketingUrl        *string
						PrivacyStatementUrl *string
						SupportUrl          *string
						TermsOfServiceUrl   *string
					}{
						LogoUrl:             app.GetInfo().GetLogoUrl(),
						MarketingUrl:        app.GetInfo().GetMarketingUrl(),
						SupportUrl:          app.GetInfo().GetSupportUrl(),
						PrivacyStatementUrl: app.GetInfo().GetPrivacyStatementUrl(),
						TermsOfServiceUrl:   app.GetInfo().GetTermsOfServiceUrl(),
					},
					KeyCredentials: keyCredentials,
					OwnerIds:       ownerIds,
					ParentalControlSettings: struct {
						CountriesBlockedForMinors []string
						LegalAgeGroupRule         *string
					}{
						CountriesBlockedForMinors: app.GetParentalControlSettings().GetCountriesBlockedForMinors(),
						LegalAgeGroupRule:         app.GetParentalControlSettings().GetLegalAgeGroupRule(),
					},
					PasswordCredentials: passwordCredentials,
					Spa: struct {
						RedirectUris []string
					}{
						RedirectUris: app.GetSpa().GetRedirectUris(),
					},
					TagsSrc: app.GetTags(),
					Web: struct {
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
					}{
						HomePageUrl:  app.GetWeb().GetHomePageUrl(),
						RedirectUris: app.GetWeb().GetRedirectUris(),
						LogoutUrl:    app.GetWeb().GetLogoutUrl(),
						ImplicitGrantSettings: struct {
							EnableAccessTokenIssuance *bool
							EnableIdTokenIssuance     *bool
						}{
							EnableAccessTokenIssuance: app.GetWeb().GetImplicitGrantSettings().GetEnableAccessTokenIssuance(),
							EnableIdTokenIssuance:     app.GetWeb().GetImplicitGrantSettings().GetEnableIdTokenIssuance(),
						},
						RedirectUriSettings: redirectUriSettings,
					},
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

//

func AdSignInReport(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.AuditLogs().SignIns().Get(ctx, &auditlogs.SignInsRequestBuilderGetRequestConfiguration{
		QueryParameters: &auditlogs.SignInsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get sign in report: %v", err)
	}

	var values []models2.Resource
	var itemErr error

	pageIterator, err := msgraphcore.NewPageIterator[models.SignInable](result, client.GetAdapter(), models.CreateSignInCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(report models.SignInable) bool {
		if report == nil {
			return true
		}

		var status struct {
			ErrorCode     *int32
			FailureReason *string
		}
		if report.GetStatus() != nil {
			status = struct {
				ErrorCode     *int32
				FailureReason *string
			}{
				ErrorCode:     report.GetStatus().GetErrorCode(),
				FailureReason: report.GetStatus().GetFailureReason(),
			}
		}

		var deviceDetail struct {
			Browser         *string
			DeviceId        *string
			DisplayName     *string
			IsCompliant     *bool
			IsManaged       *bool
			OperatingSystem *string
			TrustType       *string
		}
		if report.GetDeviceDetail() != nil {
			deviceDetail = struct {
				Browser         *string
				DeviceId        *string
				DisplayName     *string
				IsCompliant     *bool
				IsManaged       *bool
				OperatingSystem *string
				TrustType       *string
			}{
				Browser:         report.GetDeviceDetail().GetBrowser(),
				DeviceId:        report.GetDeviceDetail().GetDeviceId(),
				DisplayName:     report.GetDeviceDetail().GetDisplayName(),
				IsCompliant:     report.GetDeviceDetail().GetIsCompliant(),
				IsManaged:       report.GetDeviceDetail().GetIsManaged(),
				OperatingSystem: report.GetDeviceDetail().GetOperatingSystem(),
				TrustType:       report.GetDeviceDetail().GetTrustType(),
			}
		}

		var location struct {
			City            *string
			CountryOrRegion *string
			GeoCoordinates  struct {
				Altitude  *float64
				Latitude  *float64
				Longitude *float64
			}
			State *string
		}
		if report.GetLocation() != nil {
			location = struct {
				City            *string
				CountryOrRegion *string
				GeoCoordinates  struct {
					Altitude  *float64
					Latitude  *float64
					Longitude *float64
				}
				State *string
			}{
				City:            report.GetLocation().GetCity(),
				CountryOrRegion: report.GetLocation().GetCountryOrRegion(),
				GeoCoordinates: struct {
					Altitude  *float64
					Latitude  *float64
					Longitude *float64
				}{
					Altitude:  report.GetLocation().GetGeoCoordinates().GetAltitude(),
					Latitude:  report.GetLocation().GetGeoCoordinates().GetLatitude(),
					Longitude: report.GetLocation().GetGeoCoordinates().GetLongitude(),
				},
				State: report.GetLocation().GetState(),
			}
		}

		var appliedConditionalAccessPolicies []struct {
			DisplayName             *string
			EnforcedGrantControls   []string
			EnforcedSessionControls []string
			Id                      *string
			Result                  string
		}
		for _, p := range report.GetAppliedConditionalAccessPolicies() {
			appliedConditionalAccessPolicies = append(appliedConditionalAccessPolicies, struct {
				DisplayName             *string
				EnforcedGrantControls   []string
				EnforcedSessionControls []string
				Id                      *string
				Result                  string
			}{
				DisplayName:             p.GetDisplayName(),
				EnforcedGrantControls:   p.GetEnforcedGrantControls(),
				EnforcedSessionControls: p.GetEnforcedSessionControls(),
				Id:                      p.GetId(),
				Result:                  p.GetResult().String(),
			})
		}

		resource := models2.Resource{
			ID:       *report.GetId(),
			Name:     *report.GetId(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdSignInReportDescription{
					TenantID:                         tenantId,
					Id:                               report.GetId(),
					CreatedDateTime:                  report.GetCreatedDateTime(),
					UserDisplayName:                  report.GetUserDisplayName(),
					UserPrincipalName:                report.GetUserPrincipalName(),
					UserId:                           report.GetUserId(),
					AppId:                            report.GetAppId(),
					AppDisplayName:                   report.GetAppDisplayName(),
					IpAddress:                        report.GetIpAddress(),
					ClientAppUsed:                    report.GetClientAppUsed(),
					CorrelationId:                    report.GetCorrelationId(),
					ConditionalAccessStatus:          report.GetConditionalAccessStatus(),
					IsInteractive:                    report.GetIsInteractive(),
					RiskDetail:                       report.GetRiskDetail(),
					RiskLevelAggregated:              report.GetRiskLevelAggregated(),
					RiskLevelDuringSignIn:            report.GetRiskLevelDuringSignIn(),
					RiskState:                        report.GetRiskState(),
					ResourceDisplayName:              report.GetResourceDisplayName(),
					ResourceId:                       report.GetResourceId(),
					RiskEventTypes:                   report.GetRiskEventTypes(),
					Status:                           status,
					DeviceDetail:                     deviceDetail,
					Location:                         location,
					AppliedConditionalAccessPolicies: appliedConditionalAccessPolicies,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		// Return true to continue the iteration
		return true
	})
	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdDevice(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []models2.Resource
	var itemErr error
	result, err := client.Devices().Get(ctx, &devices.DevicesRequestBuilderGetRequestConfiguration{
		QueryParameters: &devices.DevicesRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get devices: %v", err)
	}
	pageIterator, err := msgraphcore.NewPageIterator[models.Deviceable](result, client.GetAdapter(), models.CreateDeviceCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(device models.Deviceable) bool {
		if device == nil {
			return true
		}
		resource := models2.Resource{
			ID:       *device.GetId(),
			Name:     *device.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdDeviceDescription{
					TenantID:                      tenantId,
					Id:                            device.GetId(),
					DisplayName:                   device.GetDisplayName(),
					AccountEnabled:                device.GetAccountEnabled(),
					DeviceId:                      device.GetDeviceId(),
					ApproximateLastSignInDateTime: device.GetApproximateLastSignInDateTime(),
					IsCompliant:                   device.GetIsCompliant(),
					IsManaged:                     device.GetIsManaged(),
					MdmAppId:                      device.GetMdmAppId(),
					OperatingSystem:               device.GetOperatingSystem(),
					OperatingSystemVersion:        device.GetOperatingSystemVersion(),
					ProfileType:                   device.GetProfileType(),
					TrustType:                     device.GetTrustType(),
					ExtensionAttributes:           device.GetExtensions(),
					MemberOf:                      device.GetMemberOf(),
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdDirectoryRole(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error
	result, err := client.DirectoryRoles().Get(ctx, &directoryroles.DirectoryRolesRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	pageIterator, err := msgraphcore.NewPageIterator[*models.DirectoryRole](result, client.GetAdapter(), models.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(role *models.DirectoryRole) bool {
		if role == nil {
			return true
		}
		var memberIds []*string
		for _, member := range role.GetMembers() {
			memberIds = append(memberIds, member.GetId())
		}

		resource := models2.Resource{
			ID:       *role.GetId(),
			Name:     *role.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdDirectoryRoleDescription{
					TenantID:       tenantId,
					DisplayName:    role.GetDisplayName(),
					Id:             role.GetId(),
					Description:    role.GetDescription(),
					RoleTemplateId: role.GetRoleTemplateId(),
					MemberIds:      memberIds,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdDirectorySetting(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.GroupSettings().Get(ctx, &groupsettings.GroupSettingsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groupsettings.GroupSettingsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	pageIterator, err := msgraphcore.NewPageIterator[models.GroupSettingable](result, client.GetAdapter(), models.CreateGroupSettingCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(setting models.GroupSettingable) bool {
		if setting == nil {
			return true
		}

		for _, v := range setting.GetValues() {
			resource := models2.Resource{
				ID:       *setting.GetId(),
				Name:     *setting.GetDisplayName(),
				Location: "global",

				Description: JSONAllFieldsMarshaller{
					Value: model.AdDirectorySettingDescription{
						TenantID:    tenantId,
						DisplayName: setting.GetDisplayName(),
						Id:          setting.GetId(),
						TemplateId:  setting.GetTemplateId(),
						Name:        v.GetName(),
						Value:       v.GetValue(),
					},
				},
			}
			if stream != nil {
				if itemErr = (*stream)(resource); itemErr != nil {
					return false
				}
			} else {
				values = append(values, resource)
			}
		}
		return true
	})

	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdDirectoryAuditReport(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.AuditLogs().DirectoryAudits().Get(ctx, &auditlogs.DirectoryAuditsRequestBuilderGetRequestConfiguration{
		QueryParameters: &auditlogs.DirectoryAuditsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	pageIterator, err := msgraphcore.NewPageIterator[models.DirectoryAuditable](result, client.GetAdapter(), models.CreateSignInCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(audit models.DirectoryAuditable) bool {
		if audit == nil {
			return true
		}

		var auditResult string
		if audit.GetResult() != nil {
			auditResult = audit.GetResult().String()
		}

		var initiatedBy struct {
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
		if audit.GetInitiatedBy() != nil {
			if audit.GetInitiatedBy().GetApp() != nil {
				initiatedBy.App = struct {
					AppId                *string
					DisplayName          *string
					ServicePrincipalId   *string
					ServicePrincipalName *string
				}{
					AppId:                audit.GetInitiatedBy().GetApp().GetAppId(),
					DisplayName:          audit.GetInitiatedBy().GetApp().GetDisplayName(),
					ServicePrincipalId:   audit.GetInitiatedBy().GetApp().GetServicePrincipalId(),
					ServicePrincipalName: audit.GetInitiatedBy().GetApp().GetServicePrincipalName(),
				}
			}
			if audit.GetInitiatedBy().GetUser() != nil {
				initiatedBy.User = struct {
					Id                *string
					DisplayName       *string
					IpAddress         *string
					UserPrincipalName *string
				}{
					Id:                audit.GetInitiatedBy().GetUser().GetId(),
					DisplayName:       audit.GetInitiatedBy().GetUser().GetDisplayName(),
					IpAddress:         audit.GetInitiatedBy().GetUser().GetIpAddress(),
					UserPrincipalName: audit.GetInitiatedBy().GetUser().GetUserPrincipalName(),
				}
			}
		}

		var targetResources []struct {
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

		for _, tr := range audit.GetTargetResources() {
			targetResource := struct {
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
			}{
				DisplayName:       tr.GetDisplayName(),
				Id:                tr.GetId(),
				TypeEscaped:       tr.GetTypeEscaped(),
				UserPrincipalName: tr.GetUserPrincipalName(),
			}
			if tr.GetGroupType() != nil {
				targetResource.GroupType = tr.GetGroupType().String()
			}

			var modifiedProperties []struct {
				DisplayName *string
				NewValue    *string
				OldValue    *string
			}

			for _, mp := range tr.GetModifiedProperties() {
				modifiedProperties = append(modifiedProperties, struct {
					DisplayName *string
					NewValue    *string
					OldValue    *string
				}{
					DisplayName: mp.GetDisplayName(),
					NewValue:    mp.GetNewValue(),
					OldValue:    mp.GetOldValue(),
				})
			}

			targetResource.ModifiedProperties = modifiedProperties

			targetResources = append(targetResources, targetResource)
		}

		var id string
		if audit.GetId() != nil {
			id = *audit.GetId()
		}

		resource := models2.Resource{
			ID:       id,
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdDirectoryAuditReportDescription{
					TenantID:            tenantId,
					Id:                  audit.GetId(),
					ActivityDateTime:    audit.GetActivityDateTime(),
					ActivityDisplayName: audit.GetActivityDisplayName(),
					Category:            audit.GetCategory(),
					CorrelationId:       audit.GetCorrelationId(),
					LoggedByService:     audit.GetLoggedByService(),
					OperationType:       audit.GetOperationType(),
					Result:              auditResult,
					ResultReason:        audit.GetResultReason(),
					InitiatedBy:         initiatedBy,
					TargetResources:     targetResources,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdDomain(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.Domains().Get(ctx, &domains.DomainsRequestBuilderGetRequestConfiguration{
		QueryParameters: &domains.DomainsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	pageIterator, err := msgraphcore.NewPageIterator[models.Domainable](result, client.GetAdapter(), models.CreateDomainCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(domain models.Domainable) bool {
		if domain == nil {
			return true
		}

		resource := models2.Resource{
			ID:       *domain.GetId(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdDomainDescription{
					TenantID:           tenantId,
					Id:                 domain.GetId(),
					AuthenticationType: domain.GetAuthenticationType(),
					IsDefault:          domain.GetIsDefault(),
					IsAdminManaged:     domain.GetIsAdminManaged(),
					IsInitial:          domain.GetIsInitial(),
					IsRoot:             domain.GetIsRoot(),
					IsVerified:         domain.GetIsVerified(),
					SupportedServices:  domain.GetSupportedServices(),
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdIdentityProvider(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.Identity().IdentityProviders().Get(ctx, &identity.IdentityProvidersRequestBuilderGetRequestConfiguration{
		QueryParameters: &identity.IdentityProvidersRequestBuilderGetQueryParameters{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	pageIterator, err := msgraphcore.NewPageIterator[*models.BuiltInIdentityProvider](result, client.GetAdapter(), models.CreateBuiltInIdentityProviderFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(ip *models.BuiltInIdentityProvider) bool {
		if ip == nil {
			return true
		}
		clientID := ip.GetAdditionalData()["clientId"]
		clientSecret := ip.GetAdditionalData()["clientSecret"]

		resource := models2.Resource{
			ID:       *ip.GetId(),
			Name:     *ip.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdIdentityProviderDescription{
					TenantID:     tenantId,
					Id:           ip.GetId(),
					DisplayName:  ip.GetDisplayName(),
					Type:         ip.GetOdataType(),
					ClientId:     clientID,
					ClientSecret: clientSecret,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdSecurityDefaultsPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Policies().IdentitySecurityDefaultsEnforcementPolicy().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	if result == nil {
		return values, nil
	}

	resource := models2.Resource{
		ID:       *result.GetId(),
		Name:     *result.GetDisplayName(),
		Location: "global",

		Description: JSONAllFieldsMarshaller{
			Value: model.AdSecurityDefaultsPolicyDescription{
				TenantID:    tenantId,
				Id:          result.GetId(),
				DisplayName: result.GetDisplayName(),
				IsEnabled:   result.GetIsEnabled(),
				Description: result.GetDescription(),
			},
		},
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func AdAuthorizationPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Policies().AuthorizationPolicy().Get(ctx, &policies.AuthorizationPolicyRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	if result == nil {
		return values, nil
	}

	defaultUserRolePermissions := struct {
		AllowedToCreateApps                      *bool
		AllowedToCreateSecurityGroups            *bool
		AllowedToCreateTenants                   *bool
		AllowedToReadBitlockerKeysForOwnedDevice *bool
		AllowedToReadOtherUsers                  *bool
		OdataType                                *string
		PermissionGrantPoliciesAssigned          []string
	}{
		AllowedToCreateApps:                      result.GetDefaultUserRolePermissions().GetAllowedToCreateApps(),
		AllowedToCreateSecurityGroups:            result.GetDefaultUserRolePermissions().GetAllowedToCreateSecurityGroups(),
		AllowedToCreateTenants:                   result.GetDefaultUserRolePermissions().GetAllowedToCreateTenants(),
		AllowedToReadBitlockerKeysForOwnedDevice: result.GetDefaultUserRolePermissions().GetAllowedToReadBitlockerKeysForOwnedDevice(),
		AllowedToReadOtherUsers:                  result.GetDefaultUserRolePermissions().GetAllowedToReadOtherUsers(),
		OdataType:                                result.GetDefaultUserRolePermissions().GetOdataType(),
		PermissionGrantPoliciesAssigned:          result.GetDefaultUserRolePermissions().GetPermissionGrantPoliciesAssigned(),
	}

	resource := models2.Resource{
		ID:       *result.GetId(),
		Name:     *result.GetDisplayName(),
		Location: "global",

		Description: JSONAllFieldsMarshaller{
			Value: model.AdAuthorizationPolicyDescription{
				TenantID:                               tenantId,
				Id:                                     result.GetId(),
				DisplayName:                            result.GetDisplayName(),
				Description:                            result.GetDescription(),
				AllowedToSignIpEmailBasedSubscriptions: result.GetAllowedToSignUpEmailBasedSubscriptions(),
				AllowedToUseSspr:                       result.GetAllowedToUseSSPR(),
				AllowedEmailVerifiedUsersToJoinOrganization: result.GetAllowEmailVerifiedUsersToJoinOrganization(),
				AllowInvitesFrom:           result.GetAllowInvitesFrom().String(),
				BlockMsolPowershell:        result.GetBlockMsolPowerShell(),
				GuestUserRoleId:            result.GetGuestUserRoleId().String(),
				DefaultUserRolePermissions: defaultUserRolePermissions,
			},
		},
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func AdConditionalAccessPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.Identity().ConditionalAccess().Policies().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	if result == nil {
		return values, nil
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.ConditionalAccessPolicyable](result, client.GetAdapter(), models.CreateConditionalAccessPolicyCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(p models.ConditionalAccessPolicyable) bool {
		if p == nil {
			return true
		}

		var applications struct {
			ApplicationFilter struct {
				Mode *string
				Rule *string
			}
			ExcludeApplications                         []string
			IncludeApplications                         []string
			IncludeAuthenticationContextClassReferences []string
			IncludeUserActions                          []string
		}
		var clientAppTypes []string
		var excludePlatforms []string
		var includePlatforms []string
		var signInRiskLevel []string
		var userRiskLevel []string
		var locations struct {
			ExcludeLocations []string
			IncludeLocations []string
		}
		var users struct {
			ExcludeGroups []string
			IncludeGroups []string
			ExcludeUsers  []string
			IncludeUsers  []string
			ExcludeRoles  []string
			IncludeRoles  []string
		}
		if p.GetConditions() != nil {
			if p.GetConditions().GetUsers() != nil {
				users = struct {
					ExcludeGroups []string
					IncludeGroups []string
					ExcludeUsers  []string
					IncludeUsers  []string
					ExcludeRoles  []string
					IncludeRoles  []string
				}{
					ExcludeGroups: p.GetConditions().GetUsers().GetExcludeGroups(),
					IncludeGroups: p.GetConditions().GetUsers().GetIncludeGroups(),
					ExcludeUsers:  p.GetConditions().GetUsers().GetExcludeUsers(),
					IncludeUsers:  p.GetConditions().GetUsers().GetIncludeUsers(),
					ExcludeRoles:  p.GetConditions().GetUsers().GetExcludeRoles(),
					IncludeRoles:  p.GetConditions().GetUsers().GetIncludeRoles(),
				}
			}
			if p.GetConditions().GetLocations() != nil {
				locations = struct {
					ExcludeLocations []string
					IncludeLocations []string
				}{
					ExcludeLocations: p.GetConditions().GetLocations().GetExcludeLocations(),
					IncludeLocations: p.GetConditions().GetLocations().GetIncludeLocations(),
				}
			}
			for _, c := range p.GetConditions().GetClientAppTypes() {
				clientAppTypes = append(clientAppTypes, c.String())
			}
			if p.GetConditions().GetPlatforms() != nil {
				for _, ep := range p.GetConditions().GetPlatforms().GetExcludePlatforms() {
					excludePlatforms = append(excludePlatforms, ep.String())
				}
				for _, ep := range p.GetConditions().GetPlatforms().GetIncludePlatforms() {
					includePlatforms = append(includePlatforms, ep.String())
				}
			}

			for _, c := range p.GetConditions().GetSignInRiskLevels() {
				signInRiskLevel = append(signInRiskLevel, c.String())
			}
			for _, c := range p.GetConditions().GetUserRiskLevels() {
				userRiskLevel = append(userRiskLevel, c.String())
			}

			if p.GetConditions().GetApplications() != nil {
				applications = struct {
					ApplicationFilter struct {
						Mode *string
						Rule *string
					}
					ExcludeApplications                         []string
					IncludeApplications                         []string
					IncludeAuthenticationContextClassReferences []string
					IncludeUserActions                          []string
				}{
					ExcludeApplications:                         p.GetConditions().GetApplications().GetExcludeApplications(),
					IncludeApplications:                         p.GetConditions().GetApplications().GetIncludeApplications(),
					IncludeAuthenticationContextClassReferences: p.GetConditions().GetApplications().GetIncludeAuthenticationContextClassReferences(),
					IncludeUserActions:                          p.GetConditions().GetApplications().GetIncludeUserActions(),
				}
				if p.GetConditions().GetApplications().GetApplicationFilter() != nil {
					applications.ApplicationFilter = struct {
						Mode *string
						Rule *string
					}{
						Mode: p.GetConditions().GetApplications().GetApplicationFilter().GetRule(),
						Rule: p.GetConditions().GetApplications().GetApplicationFilter().GetRule(),
					}
				}
			}
		}

		var builtInControls []string
		var operator *string
		var state string
		if p.GetState() != nil {
			state = p.GetState().String()
		}
		var termOfUse []string
		var customAuthenticationFactors []string
		if p.GetGrantControls() != nil {
			for _, c := range p.GetGrantControls().GetBuiltInControls() {
				builtInControls = append(builtInControls, c.String())
			}
			operator = p.GetGrantControls().GetOperator()
			termOfUse = p.GetGrantControls().GetTermsOfUse()
			customAuthenticationFactors = p.GetGrantControls().GetCustomAuthenticationFactors()
		}

		var applicationEnforcedRestrictions struct {
			IsEnabled *bool
		}
		var cloudAppSecurity struct {
			CloudAppSecurityType string
			IsEnabled            *bool
		}
		var signInFrequency struct {
			AuthenticationType string
			FrequencyInterval  string
			TypeEscaped        string
			Value              *int32
			IsEnabled          *bool
		}
		var persistentBrowser struct {
			IsEnabled *bool
			Mode      string
		}
		if p.GetSessionControls() != nil {
			if p.GetSessionControls().GetApplicationEnforcedRestrictions() != nil {
				applicationEnforcedRestrictions = struct {
					IsEnabled *bool
				}{IsEnabled: p.GetSessionControls().GetApplicationEnforcedRestrictions().GetIsEnabled()}
			}
			if p.GetSessionControls().GetCloudAppSecurity() != nil {
				if p.GetSessionControls().GetCloudAppSecurity().GetCloudAppSecurityType() != nil {
					cloudAppSecurity.CloudAppSecurityType = p.GetSessionControls().GetCloudAppSecurity().GetCloudAppSecurityType().String()
				}
				cloudAppSecurity.IsEnabled = p.GetSessionControls().GetCloudAppSecurity().GetIsEnabled()
			}

			if p.GetSessionControls().GetSignInFrequency() != nil {
				signInFrequency = struct {
					AuthenticationType string
					FrequencyInterval  string
					TypeEscaped        string
					Value              *int32
					IsEnabled          *bool
				}{
					Value:     p.GetSessionControls().GetSignInFrequency().GetValue(),
					IsEnabled: p.GetSessionControls().GetSignInFrequency().GetIsEnabled(),
				}
				if p.GetSessionControls().GetSignInFrequency().GetTypeEscaped() != nil {
					signInFrequency.TypeEscaped = p.GetSessionControls().GetSignInFrequency().GetTypeEscaped().String()
				}
				if p.GetSessionControls().GetSignInFrequency().GetFrequencyInterval() != nil {
					signInFrequency.FrequencyInterval = p.GetSessionControls().GetSignInFrequency().GetFrequencyInterval().String()
				}
				if p.GetSessionControls().GetSignInFrequency().GetAuthenticationType() != nil {
					signInFrequency.AuthenticationType = p.GetSessionControls().GetSignInFrequency().GetAuthenticationType().String()
				}
			}

			if p.GetSessionControls().GetPersistentBrowser() != nil {
				persistentBrowser = struct {
					IsEnabled *bool
					Mode      string
				}{
					IsEnabled: p.GetSessionControls().GetPersistentBrowser().GetIsEnabled(),
				}
				if p.GetSessionControls().GetPersistentBrowser().GetMode() != nil {
					persistentBrowser.Mode = p.GetSessionControls().GetPersistentBrowser().GetMode().String()
				}
			}
		}

		resource := models2.Resource{
			ID:       *p.GetId(),
			Name:     *p.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdConditionalAccessPolicyDescription{
					TenantID:                        tenantId,
					Id:                              p.GetId(),
					DisplayName:                     p.GetDisplayName(),
					State:                           state,
					CreatedDateTime:                 p.GetCreatedDateTime(),
					ModifiedDateTime:                p.GetModifiedDateTime(),
					Operator:                        operator,
					Applications:                    applications,
					ApplicationEnforcedRestrictions: applicationEnforcedRestrictions,
					BuiltInControls:                 builtInControls,
					ClientAppTypes:                  clientAppTypes,
					CustomAuthenticationFactors:     customAuthenticationFactors,
					CloudAppSecurity:                cloudAppSecurity,
					Locations:                       locations,
					PersistentBrowser:               persistentBrowser,
					Platforms: struct {
						ExcludePlatforms []string
						IncludePlatforms []string
					}{
						ExcludePlatforms: excludePlatforms,
						IncludePlatforms: includePlatforms,
					},
					SignInFrequency:  signInFrequency,
					SignInRiskLevels: signInRiskLevel,
					TermsOfUse:       termOfUse,
					Users:            users,
					UserRiskLevel:    userRiskLevel,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdAdminConsentRequestPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Policies().AdminConsentRequestPolicy().Get(ctx, &policies.AdminConsentRequestPolicyRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	if result == nil {
		return values, nil
	}

	var reviewers []struct {
		OdataType *string
		Query     *string
		QueryRoot *string
		QueryType *string
	}
	for _, r := range result.GetReviewers() {
		reviewers = append(reviewers, struct {
			OdataType *string
			Query     *string
			QueryRoot *string
			QueryType *string
		}{
			OdataType: r.GetOdataType(),
			Query:     r.GetQuery(),
			QueryRoot: r.GetQueryRoot(),
			QueryType: r.GetQueryType(),
		})
	}

	resource := models2.Resource{
		ID:       *result.GetId(),
		Location: "global",

		Description: JSONAllFieldsMarshaller{
			Value: model.AdAdminConsentRequestPolicyDescription{
				TenantID:              tenantId,
				Id:                    result.GetId(),
				IsEnabled:             result.GetIsEnabled(),
				NotifyReviewers:       result.GetNotifyReviewers(),
				RemindersEnabled:      result.GetRemindersEnabled(),
				RequestDurationInDays: result.GetRequestDurationInDays(),
				Version:               result.GetVersion(),
				Reviewers:             reviewers,
			},
		},
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func AdUserRegistrationDetails(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Reports().AuthenticationMethods().UserRegistrationDetails().Get(ctx, &reports.AuthenticationMethodsUserRegistrationDetailsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	var values []models2.Resource
	var itemErr error
	pageIterator, err := msgraphcore.NewPageIterator[*models.UserRegistrationDetails](result, client.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(user *models.UserRegistrationDetails) bool {
		if user == nil {
			return true
		}
		resource := models2.Resource{
			ID:       *user.GetId(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdUserRegistrationDetailsDescription{
					TenantID:                             tenantId,
					Id:                                   user.GetId(),
					UserPrincipalName:                    user.GetUserPrincipalName(),
					UserDisplayName:                      user.GetUserDisplayName(),
					UserType:                             user.GetUserType().String(),
					IsAdmin:                              user.GetIsAdmin(),
					IsMfaCapable:                         user.GetIsMfaCapable(),
					IsMfaRegistered:                      user.GetIsMfaRegistered(),
					IsSsprCapable:                        user.GetIsSsprCapable(),
					IsSsprRegistered:                     user.GetIsSsprRegistered(),
					IsSsprEnabled:                        user.GetIsSsprEnabled(),
					IsPasswordlessCapable:                user.GetIsPasswordlessCapable(),
					SystemPreferredAuthenticationMethods: user.GetSystemPreferredAuthenticationMethods(),
					IsSystemPreferredAuthenticationMethodEnabled:  user.GetIsSystemPreferredAuthenticationMethodEnabled(),
					LastUpdatedDateTime:                           user.GetLastUpdatedDateTime(),
					MethodsRegistered:                             user.GetMethodsRegistered(),
					UserPreferredMethodForSecondaryAuthentication: user.GetUserPreferredMethodForSecondaryAuthentication().String(),
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdGroupMembership(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Groups().Get(ctx, &groups.GroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.GroupsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}

	var values []models2.Resource
	var itemErr error
	pageIterator, err := msgraphcore.NewPageIterator[models.Groupable](result, client.GetAdapter(), models.CreateGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(group models.Groupable) bool {
		if group == nil {
			return true
		}
		var memberIds []*string
		for _, m := range group.GetMembers() {
			memberIds = append(memberIds, m.GetId())
		}

		members, err := client.Groups().ByGroupId(*group.GetId()).TransitiveMembers().GraphUser().Get(ctx, &groups.ItemTransitiveMembersGraphUserRequestBuilderGetRequestConfiguration{
			QueryParameters: &groups.ItemTransitiveMembersGraphUserRequestBuilderGetQueryParameters{
				Top: aws.Int32(999),
			},
		})
		if err != nil {
			itemErr = err
			return false
		}

		for _, member := range members.GetValue() {
			resource := models2.Resource{
				ID:       *member.GetId(),
				Name:     *member.GetDisplayName(),
				Location: "global",

				Description: JSONAllFieldsMarshaller{
					Value: model.AdGroupMembershipDescription{
						TenantID:           tenantId,
						DisplayName:        member.GetDisplayName(),
						Id:                 member.GetId(),
						GroupId:            group.GetId(),
						State:              member.GetState(),
						UserPrincipalName:  member.GetUserPrincipalName(),
						AccountEnabled:     member.GetAccountEnabled(),
						Mail:               member.GetMail(),
						ProxyAddresses:     member.GetProxyAddresses(),
						UserType:           member.GetUserType(),
						SecurityIdentifier: member.GetSecurityIdentifier(),
					},
				},
			}
			if stream != nil {
				if itemErr = (*stream)(resource); itemErr != nil {
					return false
				}
			} else {
				values = append(values, resource)
			}
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdAppRegistration(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []models2.Resource
	var itemErr error
	result, err := client.Applications().Get(ctx, &applications.ApplicationsRequestBuilderGetRequestConfiguration{
		QueryParameters: &applications.ApplicationsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	pageIterator, err := msgraphcore.NewPageIterator[models.Applicationable](result, client.GetAdapter(), models.CreateApplicationCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(app models.Applicationable) bool {
		if app == nil {
			return true
		}

		var oauth2PermissionScopes []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		for _, ps := range app.GetApi().GetOauth2PermissionScopes() {
			oauth2PermissionScopes = append(oauth2PermissionScopes, struct {
				AdminConsentDescription *string
				AdminConsentDisplayName *string
				Id                      string
				IsEnabled               *bool
				Origin                  *string
				TypeEscaped             *string
				UserConsentDescription  *string
				UserConsentDisplayName  *string
			}{
				Id:                      ps.GetId().String(),
				Origin:                  ps.GetOrigin(),
				IsEnabled:               ps.GetIsEnabled(),
				TypeEscaped:             ps.GetTypeEscaped(),
				AdminConsentDescription: ps.GetAdminConsentDescription(),
				UserConsentDescription:  ps.GetUserConsentDescription(),
				UserConsentDisplayName:  ps.GetUserConsentDisplayName(),
				AdminConsentDisplayName: ps.GetAdminConsentDisplayName(),
			})
		}

		var preAuthorizedApplications []struct {
			AppId                  *string
			DelegatedPermissionIds []string
		}
		for _, paa := range app.GetApi().GetPreAuthorizedApplications() {
			preAuthorizedApplications = append(preAuthorizedApplications, struct {
				AppId                  *string
				DelegatedPermissionIds []string
			}{
				AppId:                  paa.GetAppId(),
				DelegatedPermissionIds: paa.GetDelegatedPermissionIds(),
			})
		}

		var knownClientApplications []string
		for _, a := range app.GetApi().GetKnownClientApplications() {
			knownClientApplications = append(knownClientApplications, a.String())
		}

		var keyCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Key                 []byte
			KeyId               string
			StartDateTime       *time.Time
			TypeEscaped         *string
			Usage               *string
		}
		for _, c := range app.GetKeyCredentials() {
			keyCredentials = append(keyCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Key                 []byte
				KeyId               string
				StartDateTime       *time.Time
				TypeEscaped         *string
				Usage               *string
			}{
				TypeEscaped:         c.GetTypeEscaped(),
				Key:                 c.GetKey(),
				DisplayName:         c.GetDisplayName(),
				StartDateTime:       c.GetStartDateTime(),
				KeyId:               c.GetKeyId().String(),
				EndDateTime:         c.GetEndDateTime(),
				Usage:               c.GetUsage(),
				CustomKeyIdentifier: c.GetCustomKeyIdentifier(),
			})
		}

		var ownerIds []*string
		for _, o := range app.GetOwners() {
			ownerIds = append(ownerIds, o.GetId())
		}

		var redirectUriSettings []struct {
			Index *int32
			Uri   *string
		}
		for _, r := range app.GetWeb().GetRedirectUriSettings() {
			redirectUriSettings = append(redirectUriSettings, struct {
				Index *int32
				Uri   *string
			}{
				Uri:   r.GetUri(),
				Index: r.GetIndex(),
			})
		}

		var passwordCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Hint                *string
			KeyId               string
			SecretText          *string
			StartDateTime       *time.Time
		}
		for _, pc := range app.GetPasswordCredentials() {
			passwordCredentials = append(passwordCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Hint                *string
				KeyId               string
				SecretText          *string
				StartDateTime       *time.Time
			}{
				CustomKeyIdentifier: pc.GetCustomKeyIdentifier(),
				DisplayName:         pc.GetDisplayName(),
				EndDateTime:         pc.GetEndDateTime(),
				Hint:                pc.GetHint(),
				KeyId:               pc.GetKeyId().String(),
				SecretText:          pc.GetSecretText(),
				StartDateTime:       pc.GetStartDateTime(),
			})
		}

		resource := models2.Resource{
			ID:       *app.GetId(),
			Name:     *app.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdAppRegistrationDescription{
					TenantID:                  tenantId,
					DisplayName:               app.GetDisplayName(),
					Id:                        app.GetId(),
					AppId:                     app.GetAppId(),
					CreatedDateTime:           app.GetCreatedDateTime(),
					Description:               app.GetDescription(),
					Oauth2RequirePostResponse: app.GetOauth2RequirePostResponse(),
					PublisherDomain:           app.GetPublisherDomain(),
					SignInAudience:            app.GetSignInAudience(),
					Api: struct {
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
					}{
						AcceptMappedClaims:          app.GetApi().GetAcceptMappedClaims(),
						KnownClientApplications:     knownClientApplications,
						PreAuthorizedApplications:   preAuthorizedApplications,
						Oauth2PermissionScopes:      oauth2PermissionScopes,
						RequestedAccessTokenVersion: app.GetApi().GetRequestedAccessTokenVersion(),
					},
					IdentifierUris: app.GetIdentifierUris(),
					Info: struct {
						LogoUrl             *string
						MarketingUrl        *string
						PrivacyStatementUrl *string
						SupportUrl          *string
						TermsOfServiceUrl   *string
					}{
						LogoUrl:             app.GetInfo().GetLogoUrl(),
						MarketingUrl:        app.GetInfo().GetMarketingUrl(),
						SupportUrl:          app.GetInfo().GetSupportUrl(),
						PrivacyStatementUrl: app.GetInfo().GetPrivacyStatementUrl(),
						TermsOfServiceUrl:   app.GetInfo().GetTermsOfServiceUrl(),
					},
					KeyCredentials: keyCredentials,
					OwnerIds:       ownerIds,
					ParentalControlSettings: struct {
						CountriesBlockedForMinors []string
						LegalAgeGroupRule         *string
					}{
						CountriesBlockedForMinors: app.GetParentalControlSettings().GetCountriesBlockedForMinors(),
						LegalAgeGroupRule:         app.GetParentalControlSettings().GetLegalAgeGroupRule(),
					},
					PasswordCredentials: passwordCredentials,
					Spa: struct {
						RedirectUris []string
					}{
						RedirectUris: app.GetSpa().GetRedirectUris(),
					},
					TagsSrc: app.GetTags(),
					Web: struct {
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
					}{
						HomePageUrl:  app.GetWeb().GetHomePageUrl(),
						RedirectUris: app.GetWeb().GetRedirectUris(),
						LogoutUrl:    app.GetWeb().GetLogoutUrl(),
						ImplicitGrantSettings: struct {
							EnableAccessTokenIssuance *bool
							EnableIdTokenIssuance     *bool
						}{
							EnableAccessTokenIssuance: app.GetWeb().GetImplicitGrantSettings().GetEnableAccessTokenIssuance(),
							EnableIdTokenIssuance:     app.GetWeb().GetImplicitGrantSettings().GetEnableIdTokenIssuance(),
						},
						RedirectUriSettings: redirectUriSettings,
					},
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdEnterpriseApplication(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []models2.Resource
	var itemErr error

	resultApps, err := client.ServicePrincipals().Get(ctx, &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
			Top:    aws.Int32(999),
			Filter: aws.String("tags/any(t:t eq 'WindowsAzureActiveDirectoryIntegratedApp')"),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	appPageIterator, err := msgraphcore.NewPageIterator[*models.ServicePrincipal](resultApps, client.GetAdapter(), models.CreateApplicationCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, fmt.Errorf("failed to query apps client: %v", err)
	}
	err = appPageIterator.Iterate(context.Background(), func(servicePrincipal *models.ServicePrincipal) bool {
		if servicePrincipal == nil || servicePrincipal.GetAppId() == nil {
			return true
		}
		var orgID *string
		v := servicePrincipal.GetAppOwnerOrganizationId()
		if v != nil {
			tmp := v.String()
			orgID = &tmp
		}

		var keyCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Key                 []byte
			KeyId               string
			StartDateTime       *time.Time
			TypeEscaped         *string
			Usage               *string
		}

		for _, kc := range servicePrincipal.GetKeyCredentials() {
			keyCredentials = append(keyCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Key                 []byte
				KeyId               string
				StartDateTime       *time.Time
				TypeEscaped         *string
				Usage               *string
			}{
				Key:                 kc.GetKey(),
				TypeEscaped:         kc.GetTypeEscaped(),
				Usage:               kc.GetUsage(),
				DisplayName:         kc.GetDisplayName(),
				CustomKeyIdentifier: kc.GetCustomKeyIdentifier(),
				KeyId:               kc.GetKeyId().String(),
				EndDateTime:         kc.GetEndDateTime(),
				StartDateTime:       kc.GetStartDateTime(),
			})
		}

		var passwordCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Hint                *string
			KeyId               string
			SecretText          *string
			StartDateTime       *time.Time
		}
		for _, pc := range servicePrincipal.GetPasswordCredentials() {
			passwordCredentials = append(passwordCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Hint                *string
				KeyId               string
				SecretText          *string
				StartDateTime       *time.Time
			}{
				CustomKeyIdentifier: pc.GetCustomKeyIdentifier(),
				DisplayName:         pc.GetDisplayName(),
				EndDateTime:         pc.GetEndDateTime(),
				Hint:                pc.GetHint(),
				KeyId:               pc.GetKeyId().String(),
				SecretText:          pc.GetSecretText(),
				StartDateTime:       pc.GetStartDateTime(),
			})
		}

		var ownerIds []*string
		for _, owner := range servicePrincipal.GetOwners() {
			ownerIds = append(ownerIds, owner.GetId())
		}

		var addIns []struct {
			Id          string
			TypeEscaped *string
			Properties  []struct {
				Key   *string
				Value *string
			}
		}
		for _, addIn := range servicePrincipal.GetAddIns() {
			var properties []struct {
				Key   *string
				Value *string
			}
			for _, p := range addIn.GetProperties() {
				properties = append(properties, struct {
					Key   *string
					Value *string
				}{
					Key:   p.GetKey(),
					Value: p.GetValue(),
				})
			}
			addIns = append(addIns, struct {
				Id          string
				TypeEscaped *string
				Properties  []struct {
					Key   *string
					Value *string
				}
			}{
				Id:          addIn.GetId().String(),
				TypeEscaped: addIn.GetTypeEscaped(),
				Properties:  properties,
			})
		}

		var appRoles []struct {
			AllowedMemberTypes []string
			Description        *string
			DisplayName        *string
			Id                 string
			IsEnabled          *bool
			Origin             *string
			Value              *string
		}
		for _, appRole := range servicePrincipal.GetAppRoles() {
			appRoles = append(appRoles, struct {
				AllowedMemberTypes []string
				Description        *string
				DisplayName        *string
				Id                 string
				IsEnabled          *bool
				Origin             *string
				Value              *string
			}{
				Id:                 appRole.GetId().String(),
				Description:        appRole.GetDescription(),
				DisplayName:        appRole.GetDisplayName(),
				AllowedMemberTypes: appRole.GetAllowedMemberTypes(),
				IsEnabled:          appRole.GetIsEnabled(),
				Origin:             appRole.GetOrigin(),
				Value:              appRole.GetValue(),
			})
		}

		var oauth2PermissionScopes []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		for _, ps := range servicePrincipal.GetOauth2PermissionScopes() {
			oauth2PermissionScopes = append(oauth2PermissionScopes, struct {
				AdminConsentDescription *string
				AdminConsentDisplayName *string
				Id                      string
				IsEnabled               *bool
				Origin                  *string
				TypeEscaped             *string
				UserConsentDescription  *string
				UserConsentDisplayName  *string
			}{
				Id:                      ps.GetId().String(),
				Origin:                  ps.GetOrigin(),
				IsEnabled:               ps.GetIsEnabled(),
				TypeEscaped:             ps.GetTypeEscaped(),
				AdminConsentDescription: ps.GetAdminConsentDescription(),
				UserConsentDescription:  ps.GetUserConsentDescription(),
				UserConsentDisplayName:  ps.GetUserConsentDisplayName(),
				AdminConsentDisplayName: ps.GetAdminConsentDisplayName(),
			})
		}

		resource := models2.Resource{
			ID:       *servicePrincipal.GetId(),
			Name:     *servicePrincipal.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdEnterpriseApplicationDescription{
					TenantID:                  tenantId,
					Id:                        servicePrincipal.GetId(),
					DisplayName:               servicePrincipal.GetDisplayName(),
					AppId:                     servicePrincipal.GetAppId(),
					AccountEnabled:            servicePrincipal.GetAccountEnabled(),
					AppDisplayName:            servicePrincipal.GetAppDisplayName(),
					AppOwnerOrganizationId:    orgID,
					AppRoleAssignmentRequired: servicePrincipal.GetAppRoleAssignmentRequired(),
					ServicePrincipalType:      servicePrincipal.GetServicePrincipalType(),
					SignInAudience:            servicePrincipal.GetSignInAudience(),
					AppDescription:            servicePrincipal.GetAppDescription(),
					Description:               servicePrincipal.GetDescription(),
					LoginUrl:                  servicePrincipal.GetLoginUrl(),
					LogoutUrl:                 servicePrincipal.GetLogoutUrl(),
					AddIns:                    addIns,
					AlternativeNames:          servicePrincipal.GetAlternativeNames(),
					AppRoles:                  appRoles,
					//Info: servicePrincipal.GetInfo(),
					KeyCredentials:             keyCredentials,
					NotificationEmailAddresses: servicePrincipal.GetNotificationEmailAddresses(),
					OwnerIds:                   ownerIds,
					PasswordCredentials:        passwordCredentials,
					Oauth2PermissionScopes:     oauth2PermissionScopes,
					ReplyUrls:                  servicePrincipal.GetReplyUrls(),
					ServicePrincipalNames:      servicePrincipal.GetServicePrincipalNames(),
					TagsSrc:                    servicePrincipal.GetTags(),
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				itemErr = fmt.Errorf("failed to stream: %v", itemErr)
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdManagedIdentity(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []models2.Resource
	var itemErr error

	resultApps, err := client.ServicePrincipals().Get(ctx, &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
			Top:    aws.Int32(999),
			Filter: aws.String("servicePrincipalType eq 'ManagedIdentity'"),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	appPageIterator, err := msgraphcore.NewPageIterator[*models.ServicePrincipal](resultApps, client.GetAdapter(), models.CreateApplicationCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, fmt.Errorf("failed to query apps client: %v", err)
	}
	err = appPageIterator.Iterate(context.Background(), func(servicePrincipal *models.ServicePrincipal) bool {
		if servicePrincipal == nil || servicePrincipal.GetAppId() == nil {
			return true
		}
		var orgID *string
		v := servicePrincipal.GetAppOwnerOrganizationId()
		if v != nil {
			tmp := v.String()
			orgID = &tmp
		}

		var keyCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Key                 []byte
			KeyId               string
			StartDateTime       *time.Time
			TypeEscaped         *string
			Usage               *string
		}

		for _, kc := range servicePrincipal.GetKeyCredentials() {
			keyCredentials = append(keyCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Key                 []byte
				KeyId               string
				StartDateTime       *time.Time
				TypeEscaped         *string
				Usage               *string
			}{
				Key:                 kc.GetKey(),
				TypeEscaped:         kc.GetTypeEscaped(),
				Usage:               kc.GetUsage(),
				DisplayName:         kc.GetDisplayName(),
				CustomKeyIdentifier: kc.GetCustomKeyIdentifier(),
				KeyId:               kc.GetKeyId().String(),
				EndDateTime:         kc.GetEndDateTime(),
				StartDateTime:       kc.GetStartDateTime(),
			})
		}

		var passwordCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Hint                *string
			KeyId               string
			SecretText          *string
			StartDateTime       *time.Time
		}
		for _, pc := range servicePrincipal.GetPasswordCredentials() {
			passwordCredentials = append(passwordCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Hint                *string
				KeyId               string
				SecretText          *string
				StartDateTime       *time.Time
			}{
				CustomKeyIdentifier: pc.GetCustomKeyIdentifier(),
				DisplayName:         pc.GetDisplayName(),
				EndDateTime:         pc.GetEndDateTime(),
				Hint:                pc.GetHint(),
				KeyId:               pc.GetKeyId().String(),
				SecretText:          pc.GetSecretText(),
				StartDateTime:       pc.GetStartDateTime(),
			})
		}

		var ownerIds []*string
		for _, owner := range servicePrincipal.GetOwners() {
			ownerIds = append(ownerIds, owner.GetId())
		}

		var addIns []struct {
			Id          string
			TypeEscaped *string
			Properties  []struct {
				Key   *string
				Value *string
			}
		}
		for _, addIn := range servicePrincipal.GetAddIns() {
			var properties []struct {
				Key   *string
				Value *string
			}
			for _, p := range addIn.GetProperties() {
				properties = append(properties, struct {
					Key   *string
					Value *string
				}{
					Key:   p.GetKey(),
					Value: p.GetValue(),
				})
			}
			addIns = append(addIns, struct {
				Id          string
				TypeEscaped *string
				Properties  []struct {
					Key   *string
					Value *string
				}
			}{
				Id:          addIn.GetId().String(),
				TypeEscaped: addIn.GetTypeEscaped(),
				Properties:  properties,
			})
		}

		var appRoles []struct {
			AllowedMemberTypes []string
			Description        *string
			DisplayName        *string
			Id                 string
			IsEnabled          *bool
			Origin             *string
			Value              *string
		}
		for _, appRole := range servicePrincipal.GetAppRoles() {
			appRoles = append(appRoles, struct {
				AllowedMemberTypes []string
				Description        *string
				DisplayName        *string
				Id                 string
				IsEnabled          *bool
				Origin             *string
				Value              *string
			}{
				Id:                 appRole.GetId().String(),
				Description:        appRole.GetDescription(),
				DisplayName:        appRole.GetDisplayName(),
				AllowedMemberTypes: appRole.GetAllowedMemberTypes(),
				IsEnabled:          appRole.GetIsEnabled(),
				Origin:             appRole.GetOrigin(),
				Value:              appRole.GetValue(),
			})
		}

		var oauth2PermissionScopes []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		for _, ps := range servicePrincipal.GetOauth2PermissionScopes() {
			oauth2PermissionScopes = append(oauth2PermissionScopes, struct {
				AdminConsentDescription *string
				AdminConsentDisplayName *string
				Id                      string
				IsEnabled               *bool
				Origin                  *string
				TypeEscaped             *string
				UserConsentDescription  *string
				UserConsentDisplayName  *string
			}{
				Id:                      ps.GetId().String(),
				Origin:                  ps.GetOrigin(),
				IsEnabled:               ps.GetIsEnabled(),
				TypeEscaped:             ps.GetTypeEscaped(),
				AdminConsentDescription: ps.GetAdminConsentDescription(),
				UserConsentDescription:  ps.GetUserConsentDescription(),
				UserConsentDisplayName:  ps.GetUserConsentDisplayName(),
				AdminConsentDisplayName: ps.GetAdminConsentDisplayName(),
			})
		}
		identityType := "SystemAssigned"
		for _, an := range servicePrincipal.GetAlternativeNames() {
			if an == "isExplicit=True" {
				identityType = "UserAssigned"
			} else if strings.Contains(an, "Microsoft.ManagedIdentity/userAssignedIdentities") {
				identityType = "UserAssigned"
			}
		}

		resource := models2.Resource{
			ID:       *servicePrincipal.GetId(),
			Name:     *servicePrincipal.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdManagedIdentityDescription{
					TenantID:                  tenantId,
					Id:                        servicePrincipal.GetId(),
					DisplayName:               servicePrincipal.GetDisplayName(),
					AppId:                     servicePrincipal.GetAppId(),
					AccountEnabled:            servicePrincipal.GetAccountEnabled(),
					AppDisplayName:            servicePrincipal.GetAppDisplayName(),
					AppOwnerOrganizationId:    orgID,
					AppRoleAssignmentRequired: servicePrincipal.GetAppRoleAssignmentRequired(),
					ServicePrincipalType:      servicePrincipal.GetServicePrincipalType(),
					SignInAudience:            servicePrincipal.GetSignInAudience(),
					AppDescription:            servicePrincipal.GetAppDescription(),
					Description:               servicePrincipal.GetDescription(),
					LoginUrl:                  servicePrincipal.GetLoginUrl(),
					LogoutUrl:                 servicePrincipal.GetLogoutUrl(),
					AddIns:                    addIns,
					AlternativeNames:          servicePrincipal.GetAlternativeNames(),
					AppRoles:                  appRoles,
					//Info: servicePrincipal.GetInfo(),
					KeyCredentials:             keyCredentials,
					NotificationEmailAddresses: servicePrincipal.GetNotificationEmailAddresses(),
					OwnerIds:                   ownerIds,
					PasswordCredentials:        passwordCredentials,
					Oauth2PermissionScopes:     oauth2PermissionScopes,
					ReplyUrls:                  servicePrincipal.GetReplyUrls(),
					ServicePrincipalNames:      servicePrincipal.GetServicePrincipalNames(),
					TagsSrc:                    servicePrincipal.GetTags(),
					IdentityType:               identityType,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				itemErr = fmt.Errorf("failed to stream: %v", itemErr)
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdMicrosoftApplication(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []models2.Resource
	var itemErr error
	headers := abstractions.NewRequestHeaders()
	headers.Add("ConsistencyLevel", "eventual")
	resultApps, err := client.ServicePrincipals().Get(ctx, &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
			Top:    aws.Int32(999),
			Filter: aws.String("appOwnerOrganizationId eq f8cdef31-a31e-4b4a-93e4-5f571e91255a"), // Microsoft Service's Microsoft Entra tenant ID
			Count:  aws.Bool(true),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	appPageIterator, err := msgraphcore.NewPageIterator[*models.ServicePrincipal](resultApps, client.GetAdapter(), models.CreateApplicationCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, fmt.Errorf("failed to query apps client: %v", err)
	}
	err = appPageIterator.Iterate(context.Background(), func(servicePrincipal *models.ServicePrincipal) bool {
		if servicePrincipal == nil || servicePrincipal.GetAppId() == nil {
			return true
		}
		var orgID *string
		v := servicePrincipal.GetAppOwnerOrganizationId()
		if v != nil {
			tmp := v.String()
			orgID = &tmp
		}

		var keyCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Key                 []byte
			KeyId               string
			StartDateTime       *time.Time
			TypeEscaped         *string
			Usage               *string
		}

		for _, kc := range servicePrincipal.GetKeyCredentials() {
			keyCredentials = append(keyCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Key                 []byte
				KeyId               string
				StartDateTime       *time.Time
				TypeEscaped         *string
				Usage               *string
			}{
				Key:                 kc.GetKey(),
				TypeEscaped:         kc.GetTypeEscaped(),
				Usage:               kc.GetUsage(),
				DisplayName:         kc.GetDisplayName(),
				CustomKeyIdentifier: kc.GetCustomKeyIdentifier(),
				KeyId:               kc.GetKeyId().String(),
				EndDateTime:         kc.GetEndDateTime(),
				StartDateTime:       kc.GetStartDateTime(),
			})
		}

		var passwordCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Hint                *string
			KeyId               string
			SecretText          *string
			StartDateTime       *time.Time
		}
		for _, pc := range servicePrincipal.GetPasswordCredentials() {
			passwordCredentials = append(passwordCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Hint                *string
				KeyId               string
				SecretText          *string
				StartDateTime       *time.Time
			}{
				CustomKeyIdentifier: pc.GetCustomKeyIdentifier(),
				DisplayName:         pc.GetDisplayName(),
				EndDateTime:         pc.GetEndDateTime(),
				Hint:                pc.GetHint(),
				KeyId:               pc.GetKeyId().String(),
				SecretText:          pc.GetSecretText(),
				StartDateTime:       pc.GetStartDateTime(),
			})
		}

		var ownerIds []*string
		for _, owner := range servicePrincipal.GetOwners() {
			ownerIds = append(ownerIds, owner.GetId())
		}

		var addIns []struct {
			Id          string
			TypeEscaped *string
			Properties  []struct {
				Key   *string
				Value *string
			}
		}
		for _, addIn := range servicePrincipal.GetAddIns() {
			var properties []struct {
				Key   *string
				Value *string
			}
			for _, p := range addIn.GetProperties() {
				properties = append(properties, struct {
					Key   *string
					Value *string
				}{
					Key:   p.GetKey(),
					Value: p.GetValue(),
				})
			}
			addIns = append(addIns, struct {
				Id          string
				TypeEscaped *string
				Properties  []struct {
					Key   *string
					Value *string
				}
			}{
				Id:          addIn.GetId().String(),
				TypeEscaped: addIn.GetTypeEscaped(),
				Properties:  properties,
			})
		}

		var appRoles []struct {
			AllowedMemberTypes []string
			Description        *string
			DisplayName        *string
			Id                 string
			IsEnabled          *bool
			Origin             *string
			Value              *string
		}
		for _, appRole := range servicePrincipal.GetAppRoles() {
			appRoles = append(appRoles, struct {
				AllowedMemberTypes []string
				Description        *string
				DisplayName        *string
				Id                 string
				IsEnabled          *bool
				Origin             *string
				Value              *string
			}{
				Id:                 appRole.GetId().String(),
				Description:        appRole.GetDescription(),
				DisplayName:        appRole.GetDisplayName(),
				AllowedMemberTypes: appRole.GetAllowedMemberTypes(),
				IsEnabled:          appRole.GetIsEnabled(),
				Origin:             appRole.GetOrigin(),
				Value:              appRole.GetValue(),
			})
		}

		var oauth2PermissionScopes []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		for _, ps := range servicePrincipal.GetOauth2PermissionScopes() {
			oauth2PermissionScopes = append(oauth2PermissionScopes, struct {
				AdminConsentDescription *string
				AdminConsentDisplayName *string
				Id                      string
				IsEnabled               *bool
				Origin                  *string
				TypeEscaped             *string
				UserConsentDescription  *string
				UserConsentDisplayName  *string
			}{
				Id:                      ps.GetId().String(),
				Origin:                  ps.GetOrigin(),
				IsEnabled:               ps.GetIsEnabled(),
				TypeEscaped:             ps.GetTypeEscaped(),
				AdminConsentDescription: ps.GetAdminConsentDescription(),
				UserConsentDescription:  ps.GetUserConsentDescription(),
				UserConsentDisplayName:  ps.GetUserConsentDisplayName(),
				AdminConsentDisplayName: ps.GetAdminConsentDisplayName(),
			})
		}

		resource := models2.Resource{
			ID:       *servicePrincipal.GetId(),
			Name:     *servicePrincipal.GetDisplayName(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdMicrosoftApplicationDescription{
					TenantID:                  tenantId,
					Id:                        servicePrincipal.GetId(),
					DisplayName:               servicePrincipal.GetDisplayName(),
					AppId:                     servicePrincipal.GetAppId(),
					AccountEnabled:            servicePrincipal.GetAccountEnabled(),
					AppDisplayName:            servicePrincipal.GetAppDisplayName(),
					AppOwnerOrganizationId:    orgID,
					AppRoleAssignmentRequired: servicePrincipal.GetAppRoleAssignmentRequired(),
					ServicePrincipalType:      servicePrincipal.GetServicePrincipalType(),
					SignInAudience:            servicePrincipal.GetSignInAudience(),
					AppDescription:            servicePrincipal.GetAppDescription(),
					Description:               servicePrincipal.GetDescription(),
					LoginUrl:                  servicePrincipal.GetLoginUrl(),
					LogoutUrl:                 servicePrincipal.GetLogoutUrl(),
					AddIns:                    addIns,
					AlternativeNames:          servicePrincipal.GetAlternativeNames(),
					AppRoles:                  appRoles,
					//Info: servicePrincipal.GetInfo(),
					KeyCredentials:             keyCredentials,
					NotificationEmailAddresses: servicePrincipal.GetNotificationEmailAddresses(),
					OwnerIds:                   ownerIds,
					PasswordCredentials:        passwordCredentials,
					Oauth2PermissionScopes:     oauth2PermissionScopes,
					ReplyUrls:                  servicePrincipal.GetReplyUrls(),
					ServicePrincipalNames:      servicePrincipal.GetServicePrincipalNames(),
					TagsSrc:                    servicePrincipal.GetTags(),
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				itemErr = fmt.Errorf("failed to stream: %v", itemErr)
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdTenant(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.Organization().Get(ctx, &organization.OrganizationRequestBuilderGetRequestConfiguration{
		QueryParameters: &organization.OrganizationRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []models2.Resource
	pageIterator, err := msgraphcore.NewPageIterator[*models.Organization](result, client.GetAdapter(), models.CreateDomainCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(org *models.Organization) bool {
		if org == nil {
			return true
		}

		var verifiedDomains []struct {
			Name         *string
			Type         *string
			Capabilities *string
			IsDefault    *bool
			IsInitial    *bool
		}
		for _, vd := range org.GetVerifiedDomains() {
			verifiedDomains = append(verifiedDomains, struct {
				Name         *string
				Type         *string
				Capabilities *string
				IsDefault    *bool
				IsInitial    *bool
			}{
				Name:         vd.GetName(),
				Type:         vd.GetTypeEscaped(),
				Capabilities: vd.GetCapabilities(),
				IsDefault:    vd.GetIsDefault(),
				IsInitial:    vd.GetIsInitial(),
			})
		}

		resource := models2.Resource{
			ID:       *org.GetId(),
			Location: "global",

			Description: JSONAllFieldsMarshaller{
				Value: model.AdTenantDescription{
					TenantID:              org.GetId(),
					DisplayName:           org.GetDisplayName(),
					CreatedDateTime:       org.GetCreatedDateTime(),
					OnPremisesSyncEnabled: org.GetOnPremisesSyncEnabled(),
					TenantType:            org.GetTenantType(),
					VerifiedDomains:       verifiedDomains,
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})

	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}
