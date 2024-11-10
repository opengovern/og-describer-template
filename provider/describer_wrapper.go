package provider

import (
	"github.com/google/go-github/v55/github"
	model "github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/configs"
	"github.com/opengovern/og-describer-template/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"github.com/shurcooL/githubv4"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

// GitHubClient custom struct for defining both rest and graphql clients
type GitHubClient struct {
	RestClient    *github.Client
	GraphQLClient *githubv4.Client
}

// DescribeByIntegration TODO: implement a wrapper to pass integration authorization to describer functions
func DescribeByIntegration(describe func(context.Context, GitHubClient, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)
		var restClient *github.Client
		var graphQLClient *githubv4.Client
		// Panic for unsupported token by prefix
		if cfg.Token != "" && !strings.HasPrefix(cfg.Token, "ghs_") && !strings.HasPrefix(cfg.Token, "ghp_") && !strings.HasPrefix(cfg.Token, "gho_") {
			panic("Supported token formats are 'ghs_', 'gho_', and 'ghp_'")
		}
		// Authentication with GitHub access token
		if cfg.Token != "" && strings.HasPrefix(cfg.Token, "ghp_") {
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: cfg.Token},
			)
			tc := oauth2.NewClient(ctx, ts)
			restClient = github.NewClient(tc)
			graphQLClient = githubv4.NewClient(tc)
		}
		// Authentication Using App Installation Access Token or OAuth Access token
		if cfg.Token != "" && (strings.HasPrefix(cfg.Token, "ghs_") || strings.HasPrefix(cfg.Token, "gho_")) {
			restClient = github.NewClient(&http.Client{Transport: &oauth2Transport{
				Token: cfg.Token,
			}})
			graphQLClient = githubv4.NewClient(&http.Client{Transport: &oauth2Transport{
				Token: cfg.Token,
			}})
		}
		client := GitHubClient{
			RestClient:    restClient,
			GraphQLClient: graphQLClient,
		}
		var values []model.Resource
		result, err := describe(ctx, client, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

type oauth2Transport struct {
	Token string
}

func (t *oauth2Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", "Bearer "+t.Token)
	return http.DefaultTransport.RoundTrip(clone)
}
