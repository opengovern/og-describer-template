package provider

import (
	"github.com/google/go-github/v55/github"
	model "github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/configs"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"github.com/shurcooL/githubv4"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strings"
)

// DescribeByIntegration TODO: implement a wrapper to pass integration authorization to describer functions
func DescribeByIntegration(describe func(context.Context, *configs.IntegrationCredentials, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {

		var values []model.Resource

		return values, nil
	}
}

// Connect Create Rest API (v3) client
func Connect(ctx context.Context) *github.Client {
	token := os.Getenv("GITHUB_TOKEN")
	// Return error for unsupported token by prefix
	if token != "" && !strings.HasPrefix(token, "ghs_") && !strings.HasPrefix(token, "ghp_") && !strings.HasPrefix(token, "gho_") {
		panic("Supported token formats are 'ghs_', 'gho_', and 'ghp_'")
	}
	var client *github.Client
	// Authentication with Github access token
	if token != "" && strings.HasPrefix(token, "ghp_") {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}
	// Authentication Using App Installation Access Token or OAuth Access token
	if token != "" && (strings.HasPrefix(token, "ghs_") || strings.HasPrefix(token, "gho_")) {
		client = github.NewClient(&http.Client{Transport: &oauth2Transport{
			Token: token,
		}})
	}
	return client
}

// ConnectV4 Create GraphQL API (v4) client
func ConnectV4(ctx context.Context) *githubv4.Client {
	token := os.Getenv("GITHUB_TOKEN")
	// Return error for unsupported token by prefix
	if token != "" && !strings.HasPrefix(token, "ghs_") && !strings.HasPrefix(token, "ghp_") && !strings.HasPrefix(token, "gho_") {
		panic("Supported token formats are 'ghs_', 'gho_', and 'ghp_'")
	}
	var client *githubv4.Client
	// Authentication Using App Installation Access Token or OAuth Access token
	if token != "" && (strings.HasPrefix(token, "ghs_") || strings.HasPrefix(token, "gho_")) {
		return githubv4.NewClient(&http.Client{Transport: &oauth2Transport{
			Token: token,
		}})
	}
	// Authentication with Github access token
	if token != "" && strings.HasPrefix(token, "ghp_") {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = githubv4.NewClient(tc)
	}
	return client
}

type oauth2Transport struct {
	Token string
}

func (t *oauth2Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", "Bearer "+t.Token)
	return http.DefaultTransport.RoundTrip(clone)
}
