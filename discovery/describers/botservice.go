package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func BotServiceBot(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armbotservice.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewBotsClient()

	pager := client.NewListPager(nil)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, b := range page.Value {
			resource := getBotServiceBot(ctx, b)
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

func getBotServiceBot(ctx context.Context, bot *armbotservice.Bot) *models.Resource {
	resourceGroupName := strings.Split(string(*bot.ID), "/")[4]
	return &models.Resource{
		ID: *bot.ID,
		Description: model.BotServiceBotDescription{
			Bot:           *bot,
			ResourceGroup: resourceGroupName,
		},
	}
}
