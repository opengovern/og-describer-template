package provider

import (
	"fmt"
	"github.com/google/go-github/v55/github"
	model "github.com/opengovern/og-describer-github/describer/pkg/models"
	"github.com/opengovern/og-describer-github/describer/provider/describers"
	"github.com/opengovern/og-describer-github/global"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"github.com/shurcooL/githubv4"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func DescribeByGithub(describe func(context.Context, describers.GitHubClient, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg global.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describers.WithTriggerType(ctx, triggerType)

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

		client := describers.GitHubClient{
			RestClient:    restClient,
			GraphQLClient: graphQLClient,
			Token:         cfg.PatToken,
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

func DescribeSingleByRepo(describe func(context.Context, describers.GitHubClient, string, string, string, *model.StreamSender) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg global.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string, stream *model.StreamSender) (*model.Resource, error) {
		ctx = describers.WithTriggerType(ctx, triggerType)

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

		client := describers.GitHubClient{
			RestClient:    restClient,
			GraphQLClient: graphQLClient,
		}

		organizationName := additionalParameters["OrganizationName"]
		repoName := additionalParameters["RepositoryName"]
		result, err := describe(ctx, client, organizationName, repoName, resourceID, stream)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}
