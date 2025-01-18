package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func DevicesProvisioningServicesCertificates(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armdeviceprovisioningservices.NewDpsCertificateClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, rg := range rgs {
		dpss, err := devicesProvisioningServices(ctx, cred, subscription, *rg.Name)
		if err != nil {
			return nil, err
		}

		for _, dps := range dpss {
			resources, err := getDevicesProvisioningServicesCertificates(ctx, client, rg, dps)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
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

func getDevicesProvisioningServicesCertificates(ctx context.Context, client *armdeviceprovisioningservices.DpsCertificateClient, rg armresources.ResourceGroup, dps armdeviceprovisioningservices.ProvisioningServiceDescription) ([]models.Resource, error) {
	it, err := client.List(ctx, *rg.Name, *dps.Name, nil)
	if err != nil {
		return nil, err
	}

	if it.Value == nil {
		return nil, err
	}
	var values []models.Resource
	for _, v := range it.Value {
		resource := models.Resource{
			ID:          *v.ID,
			Name:        *v.Name,
			Location:    "global",
			Description: v,
		}
		values = append(values, resource)
	}
	return values, nil
}

func devicesProvisioningServices(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, resourceGroup string) ([]armdeviceprovisioningservices.ProvisioningServiceDescription, error) {
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewIotDpsResourceClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []armdeviceprovisioningservices.ProvisioningServiceDescription
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			values = append(values, *v)
		}
	}
	return values, nil
}

func IOTHub(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	iotHubClient, err := armiothub.NewResourceClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := iotHubClient.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getIOTHub(ctx, diagnosticClient, v)
			if err != nil {
				return nil, err
			}
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

func getIOTHub(ctx context.Context, client *armmonitor.DiagnosticSettingsClient, iotHubDescription *armiothub.Description) (*models.Resource, error) {
	resourceGroup := strings.Split(*iotHubDescription.ID, "/")[4]

	id := *iotHubDescription.ID

	pager := client.NewListPager(id, nil)
	var devicesListOp []armmonitor.DiagnosticSettingsResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			devicesListOp = append(devicesListOp, *v)
		}
	}

	resource := models.Resource{
		ID:       *iotHubDescription.ID,
		Name:     *iotHubDescription.Name,
		Location: *iotHubDescription.Location,
		Description: model.IOTHubDescription{
			IotHubDescription:           *iotHubDescription,
			DiagnosticSettingsResources: &devicesListOp,
			ResourceGroup:               resourceGroup,
		},
	}

	return &resource, nil
}

func IOTHubDps(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	clientFactory, err := armdeviceprovisioningservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewIotDpsResourceClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getIOTHubDps(ctx, diagnosticClient, v)
			if err != nil {
				return nil, err
			}
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

func getIOTHubDps(ctx context.Context, client *armmonitor.DiagnosticSettingsClient, v *armdeviceprovisioningservices.ProvisioningServiceDescription) (*models.Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	id := *v.ID

	pager := client.NewListPager(id, nil)
	var devicesListOp []armmonitor.DiagnosticSettingsResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			devicesListOp = append(devicesListOp, *v)
		}
	}

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.IOTHubDpsDescription{
			IotHubDps:                   *v,
			DiagnosticSettingsResources: &devicesListOp,
			ResourceGroup:               resourceGroup,
		},
	}

	return &resource, nil
}
