package describers

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"github.com/aws/aws-sdk-go-v2/aws"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func RoleAssignment(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armauthorization.NewRoleAssignmentsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewListForSubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, roleAssignment := range page.Value {
			resource := getRoleAssignment(ctx, roleAssignment)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getRoleAssignment(ctx context.Context, v *armauthorization.RoleAssignment) *models.Resource {
	return &models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: model.RoleAssignmentDescription{
			RoleAssignment: *v,
		},
	}
}

func RoleDefinition(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armauthorization.NewRoleDefinitionsClient(cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager("/subscriptions/"+subscription, nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, roleDefinition := range page.Value {
			resource := getRoleDefinition(ctx, roleDefinition)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getRoleDefinition(ctx context.Context, v *armauthorization.RoleDefinition) *models.Resource {
	return &models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: model.RoleDefinitionDescription{
			RoleDefinition: *v,
		},
	}
}

func PolicyDefinition(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armpolicy.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDefinitionsClient()
	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, definition := range page.Value {
			resource := getPolicyDefinition(ctx, subscription, definition)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getPolicyDefinition(ctx context.Context, subscription string, definition *armpolicy.Definition) *models.Resource {
	akas := []string{"azure:///subscriptions/" + subscription + *definition.ID, "azure:///subscriptions/" + subscription + strings.ToLower(*definition.ID)}
	turbotData := map[string]interface{}{
		"SubscriptionId": subscription,
		"Akas":           akas,
	}

	return &models.Resource{
		ID:       *definition.ID,
		Name:     *definition.Name,
		Location: "global",
		Description: model.PolicyDefinitionDescription{
			Definition: *definition,
			TurboData:  turbotData,
		},
	}
}

func UserEffectiveAccess(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armauthorization.NewRoleAssignmentsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewListForSubscriptionPager(nil)
	scopes := []string{"https://graph.microsoft.com/.default"}
	graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, roleAssignment := range page.Value {
			if *roleAssignment.Properties.PrincipalType == armauthorization.PrincipalTypeGroup {
				members, err := graphClient.Groups().ByGroupId(*roleAssignment.Properties.PrincipalID).TransitiveMembers().GraphUser().Get(ctx, &groups.ItemTransitiveMembersGraphUserRequestBuilderGetRequestConfiguration{
					QueryParameters: &groups.ItemTransitiveMembersGraphUserRequestBuilderGetQueryParameters{
						Top: aws.Int32(999),
					},
				})
				if err != nil {
					return nil, err
				}
				for _, m := range members.GetValue() {
					id := fmt.Sprintf("%s|%s", *m.GetId(), *roleAssignment.ID)
					resource := models.Resource{
						ID:       id,
						Name:     *roleAssignment.Name,
						Location: "global",
						Description: model.UserEffectiveAccessDescription{
							RoleAssignment:    *roleAssignment,
							PrincipalName:     *m.GetDisplayName(),
							PrincipalId:       *m.GetId(),
							PrincipalType:     armauthorization.PrincipalTypeUser,
							Scope:             *roleAssignment.Properties.Scope,
							ScopeType:         getScopeType(*roleAssignment.Properties.Scope),
							AssignmentType:    "GroupAssignment",
							ParentPrincipalId: roleAssignment.Properties.PrincipalID,
						},
					}
					if stream != nil {
						if err := (*stream)(resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, resource)
					}
				}
			} else if *roleAssignment.Properties.PrincipalType == armauthorization.PrincipalTypeUser {
				id := fmt.Sprintf("%s|%s", *roleAssignment.Properties.PrincipalID, *roleAssignment.ID)
				user, err := graphClient.Users().ByUserId(*roleAssignment.Properties.PrincipalID).Get(ctx, nil)
				if err != nil {
					if strings.Contains(err.Error(), "does not exist") {
						continue
					}
					return nil, err
				}
				resource := models.Resource{
					ID:       id,
					Name:     *roleAssignment.Name,
					Location: "global",
					Description: model.UserEffectiveAccessDescription{
						RoleAssignment:    *roleAssignment,
						PrincipalId:       *roleAssignment.Properties.PrincipalID,
						PrincipalName:     *user.GetDisplayName(),
						PrincipalType:     armauthorization.PrincipalTypeUser,
						Scope:             *roleAssignment.Properties.Scope,
						ScopeType:         getScopeType(*roleAssignment.Properties.Scope),
						AssignmentType:    "Explicit",
						ParentPrincipalId: nil,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			} else if *roleAssignment.Properties.PrincipalType == armauthorization.PrincipalTypeServicePrincipal {
				id := fmt.Sprintf("%s|%s", *roleAssignment.Properties.PrincipalID, *roleAssignment.ID)
				spn, err := graphClient.ServicePrincipals().ByServicePrincipalId(*roleAssignment.Properties.PrincipalID).Get(ctx, nil)
				if err != nil {
					if strings.Contains(err.Error(), "does not exist") {
						continue
					}
					return nil, err
				}
				resource := models.Resource{
					ID:       id,
					Name:     *roleAssignment.Name,
					Location: "global",
					Description: model.UserEffectiveAccessDescription{
						RoleAssignment:    *roleAssignment,
						PrincipalId:       *roleAssignment.Properties.PrincipalID,
						PrincipalName:     *spn.GetDisplayName(),
						PrincipalType:     armauthorization.PrincipalTypeServicePrincipal,
						Scope:             *roleAssignment.Properties.Scope,
						ScopeType:         getScopeType(*roleAssignment.Properties.Scope),
						AssignmentType:    "Explicit",
						ParentPrincipalId: nil,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func getScopeType(scope string) string {
	subscriptionRegex := regexp.MustCompile(`^/subscriptions/[a-fA-F0-9\-]+$`)
	managementGroupRegex := regexp.MustCompile(`^/providers/Microsoft\.Management/managementGroups/.*$`)
	rootTenantRegex := regexp.MustCompile(`^/$`)
	otherRegex := regexp.MustCompile(`^/subscriptions/[a-fA-F0-9\-]+/.*$`)

	// Determine the scope type
	switch {
	case subscriptionRegex.MatchString(scope):
		return "Subscription"
	case managementGroupRegex.MatchString(scope):
		return "Management Group"
	case rootTenantRegex.MatchString(scope):
		return "Root Tenant Management Group"
	case otherRegex.MatchString(scope):
		return "Other"
	default:
		return "Other"
	}
}
