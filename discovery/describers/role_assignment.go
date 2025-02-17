package describers

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/aws/aws-sdk-go-v2/aws"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	models2 "github.com/opengovern/og-describer-entraid/discovery/pkg/models"
	model "github.com/opengovern/og-describer-entraid/discovery/provider"
	"golang.org/x/net/context"
	"strings"
)

func ApplicationAppRoleAssignedTo(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	applications, err := client.Applications().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get application app role assignment list: %v", err)
	}
	var values []models2.Resource
	for _, sp := range applications.GetValue() {
		id := sp.GetId()
		if id == nil {
			continue
		}
		appId := sp.GetAppId()
		if appId == nil {
			continue
		}
		requestInfo, err := client.ServicePrincipals().ByServicePrincipalId("placeholder").AppRoleAssignedTo().ToGetRequestInformation(ctx,
			&serviceprincipals.ItemAppRoleAssignedToRequestBuilderGetRequestConfiguration{
				QueryParameters: &serviceprincipals.ItemAppRoleAssignedToRequestBuilderGetQueryParameters{
					Top: aws.Int32(999),
				},
			})
		if err != nil {
			return nil, fmt.Errorf("failed to get role assignments: %v", err)
		}
		uri, _ := requestInfo.GetUri()
		url := strings.Replace(uri.String(), "/servicePrincipals/placeholder/", fmt.Sprintf("/servicePrincipals('appId=%v')/", *appId), 1)
		result, err := client.ServicePrincipals().ByServicePrincipalId("placeholder").AppRoleAssignedTo().WithUrl(url).Get(ctx, &serviceprincipals.ItemAppRoleAssignedToRequestBuilderGetRequestConfiguration{
			QueryParameters: &serviceprincipals.ItemAppRoleAssignedToRequestBuilderGetQueryParameters{
				Top: aws.Int32(999),
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get role assignments: %v", err)
		}
		var itemErr error
		pageIterator, err := msgraphcore.NewPageIterator[models.AppRoleAssignmentable](result, client.GetAdapter(), models.CreateAppRoleAssignmentCollectionResponseFromDiscriminatorValue)
		if err != nil {
			return nil, err
		}
		err = pageIterator.Iterate(context.Background(), func(roleAssignment models.AppRoleAssignmentable) bool {
			if roleAssignment == nil {
				return true
			}
			var id string
			if roleAssignment.GetId() != nil {
				id = *roleAssignment.GetId()
			}
			var name string
			var appRoleId, resourceId, principalId *string
			if roleAssignment.GetAppRoleId() != nil {
				appRoleIdTmp := roleAssignment.GetAppRoleId().String()
				appRoleId = &appRoleIdTmp
			}

			if roleAssignment.GetResourceId() != nil {
				resourceIdTmp := roleAssignment.GetResourceId().String()
				resourceId = &resourceIdTmp
			}

			if roleAssignment.GetPrincipalId() != nil {
				principalIdTmp := roleAssignment.GetPrincipalId().String()
				principalId = &principalIdTmp
			}

			resource := models2.Resource{
				ID:       id,
				Name:     name,
				Location: "global",

				Description: model.ApplicationAppRoleAssignedToDescription{
					TenantID:             tenantId,
					Id:                   roleAssignment.GetId(),
					AppRoleId:            appRoleId,
					ResourceId:           resourceId,
					ResourceDisplayName:  roleAssignment.GetResourceDisplayName(),
					CreatedDateTime:      roleAssignment.GetCreatedDateTime(),
					DeletedDateTime:      roleAssignment.GetDeletedDateTime(),
					PrincipalId:          principalId,
					PrincipalDisplayName: roleAssignment.GetPrincipalDisplayName(),
					PrincipalType:        roleAssignment.GetPrincipalType(),
					AppId:                sp.GetAppId(),
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
	}
	return values, nil
}

func ServicePrincipalAppRoleAssignedTo(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	spCollection, err := client.ServicePrincipals().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get application app role assignment list: %v", err)
	}
	var values []models2.Resource
	for _, sp := range spCollection.GetValue() {
		id := sp.GetId()
		if id == nil {
			continue
		}
		result, err := client.ServicePrincipals().ByServicePrincipalId(*id).AppRoleAssignedTo().Get(ctx,
			&serviceprincipals.ItemAppRoleAssignedToRequestBuilderGetRequestConfiguration{
				QueryParameters: &serviceprincipals.ItemAppRoleAssignedToRequestBuilderGetQueryParameters{
					Top: aws.Int32(999),
				},
			})
		if err != nil {
			return nil, fmt.Errorf("failed to get role assignments: %v", err)
		}
		var itemErr error
		pageIterator, err := msgraphcore.NewPageIterator[models.AppRoleAssignmentable](result, client.GetAdapter(), models.CreateAppRoleAssignmentCollectionResponseFromDiscriminatorValue)
		if err != nil {
			return nil, err
		}
		err = pageIterator.Iterate(context.Background(), func(roleAssignment models.AppRoleAssignmentable) bool {
			if roleAssignment == nil {
				return true
			}
			var id string
			if roleAssignment.GetId() != nil {
				id = *roleAssignment.GetId()
			}
			var name string
			var appRoleId, resourceId, principalId *string
			if roleAssignment.GetAppRoleId() != nil {
				appRoleIdTmp := roleAssignment.GetAppRoleId().String()
				appRoleId = &appRoleIdTmp
			}

			if roleAssignment.GetResourceId() != nil {
				resourceIdTmp := roleAssignment.GetResourceId().String()
				resourceId = &resourceIdTmp
			}

			if roleAssignment.GetPrincipalId() != nil {
				principalIdTmp := roleAssignment.GetPrincipalId().String()
				principalId = &principalIdTmp
			}

			resource := models2.Resource{
				ID:       id,
				Name:     name,
				Location: "global",

				Description: model.ServicePrincipalAppRoleAssignedToDescription{
					TenantID:             tenantId,
					Id:                   roleAssignment.GetId(),
					AppRoleId:            appRoleId,
					ResourceId:           resourceId,
					ResourceDisplayName:  roleAssignment.GetResourceDisplayName(),
					CreatedDateTime:      roleAssignment.GetCreatedDateTime(),
					DeletedDateTime:      roleAssignment.GetDeletedDateTime(),
					PrincipalId:          principalId,
					PrincipalDisplayName: roleAssignment.GetPrincipalDisplayName(),
					PrincipalType:        roleAssignment.GetPrincipalType(),
					ServicePrincipalId:   sp.GetId(),
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
	}
	return values, nil
}

func ServicePrincipalAppRoleAssignment(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	spCollection, err := client.ServicePrincipals().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get application app role assignment list: %v", err)
	}
	var values []models2.Resource
	for _, sp := range spCollection.GetValue() {
		id := sp.GetId()
		if id == nil {
			continue
		}
		result, err := client.ServicePrincipals().ByServicePrincipalId(*id).AppRoleAssignments().Get(ctx,
			&serviceprincipals.ItemAppRoleAssignmentsRequestBuilderGetRequestConfiguration{
				QueryParameters: &serviceprincipals.ItemAppRoleAssignmentsRequestBuilderGetQueryParameters{
					Top: aws.Int32(999),
				},
			})
		if err != nil {
			return nil, fmt.Errorf("failed to get role assignments: %v", err)
		}
		var itemErr error
		pageIterator, err := msgraphcore.NewPageIterator[models.AppRoleAssignmentable](result, client.GetAdapter(), models.CreateAppRoleAssignmentCollectionResponseFromDiscriminatorValue)
		if err != nil {
			return nil, err
		}
		err = pageIterator.Iterate(context.Background(), func(roleAssignment models.AppRoleAssignmentable) bool {
			if roleAssignment == nil {
				return true
			}
			var id string
			if roleAssignment.GetId() != nil {
				id = *roleAssignment.GetId()
			}
			var name string
			var appRoleId, resourceId, principalId *string
			if roleAssignment.GetAppRoleId() != nil {
				appRoleIdTmp := roleAssignment.GetAppRoleId().String()
				appRoleId = &appRoleIdTmp
			}

			if roleAssignment.GetResourceId() != nil {
				resourceIdTmp := roleAssignment.GetResourceId().String()
				resourceId = &resourceIdTmp
			}

			if roleAssignment.GetPrincipalId() != nil {
				principalIdTmp := roleAssignment.GetPrincipalId().String()
				principalId = &principalIdTmp
			}

			resource := models2.Resource{
				ID:       id,
				Name:     name,
				Location: "global",

				Description: model.ServicePrincipalAppRoleAssignmentDescription{
					TenantID:             tenantId,
					Id:                   roleAssignment.GetId(),
					AppRoleId:            appRoleId,
					ResourceId:           resourceId,
					ResourceDisplayName:  roleAssignment.GetResourceDisplayName(),
					CreatedDateTime:      roleAssignment.GetCreatedDateTime(),
					DeletedDateTime:      roleAssignment.GetDeletedDateTime(),
					PrincipalId:          principalId,
					PrincipalDisplayName: roleAssignment.GetPrincipalDisplayName(),
					PrincipalType:        roleAssignment.GetPrincipalType(),
					ServicePrincipalId:   sp.GetId(),
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
	}
	return values, nil
}

func UserAppRoleAssignment(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *models2.StreamSender) ([]models2.Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	usersResp, err := client.Users().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get application app role assignment list: %v", err)
	}
	var values []models2.Resource
	for _, user := range usersResp.GetValue() {
		id := user.GetId()
		if id == nil {
			continue
		}
		result, err := client.Users().ByUserId(*id).AppRoleAssignments().Get(ctx,
			&users.ItemAppRoleAssignmentsRequestBuilderGetRequestConfiguration{
				QueryParameters: &users.ItemAppRoleAssignmentsRequestBuilderGetQueryParameters{
					Top: aws.Int32(999),
				},
			})
		if err != nil {
			return nil, fmt.Errorf("failed to get role assignments: %v", err)
		}
		var itemErr error
		pageIterator, err := msgraphcore.NewPageIterator[models.AppRoleAssignmentable](result, client.GetAdapter(), models.CreateAppRoleAssignmentCollectionResponseFromDiscriminatorValue)
		if err != nil {
			return nil, err
		}
		err = pageIterator.Iterate(context.Background(), func(roleAssignment models.AppRoleAssignmentable) bool {
			if roleAssignment == nil {
				return true
			}
			var id string
			if roleAssignment.GetId() != nil {
				id = *roleAssignment.GetId()
			}
			var name string
			var appRoleId, resourceId, principalId *string
			if roleAssignment.GetAppRoleId() != nil {
				appRoleIdTmp := roleAssignment.GetAppRoleId().String()
				appRoleId = &appRoleIdTmp
			}

			if roleAssignment.GetResourceId() != nil {
				resourceIdTmp := roleAssignment.GetResourceId().String()
				resourceId = &resourceIdTmp
			}

			if roleAssignment.GetPrincipalId() != nil {
				principalIdTmp := roleAssignment.GetPrincipalId().String()
				principalId = &principalIdTmp
			}

			resource := models2.Resource{
				ID:       id,
				Name:     name,
				Location: "global",

				Description: model.UserAppRoleAssignmentDescription{
					TenantID:             tenantId,
					Id:                   roleAssignment.GetId(),
					AppRoleId:            appRoleId,
					ResourceId:           resourceId,
					ResourceDisplayName:  roleAssignment.GetResourceDisplayName(),
					CreatedDateTime:      roleAssignment.GetCreatedDateTime(),
					DeletedDateTime:      roleAssignment.GetDeletedDateTime(),
					PrincipalId:          principalId,
					PrincipalDisplayName: roleAssignment.GetPrincipalDisplayName(),
					PrincipalType:        roleAssignment.GetPrincipalType(),
					UserId:               user.GetId(),
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
	}
	return values, nil
}
