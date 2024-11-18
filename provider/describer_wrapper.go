package provider

import (
	"fmt"
	"github.com/google/go-github/v55/github"
	model "github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/configs"
	"github.com/opengovern/og-describer-github/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"github.com/shurcooL/githubv4"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func DescribeByGithub(describe func(context.Context, describer.GitHubClient, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		if cfg.PatToken == "" {
			return nil, fmt.Errorf("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
		}

		// Create an OAuth2 token source
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: cfg.PatToken},
		)

		// Create an OAuth2 client
		tc := oauth2.NewClient(ctx, ts)

		// Create a new GitHub client
		restClient := github.NewClient(tc)
		graphQLClient := githubv4.NewClient(tc)

		client := describer.GitHubClient{
			RestClient:    restClient,
			GraphQLClient: graphQLClient,
		}

		organizationName := additionalParameters["OrganizationName"]
		var values []model.Resource
		result, err := describe(ctx, client, organizationName, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}
