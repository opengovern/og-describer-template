package provider

import (
	"fmt"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v55/github"
	model "github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/configs"
	"github.com/opengovern/og-describer-github/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"github.com/shurcooL/githubv4"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func DescribeByGithub(describe func(context.Context, describer.GitHubClient, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)
		var restClient *github.Client
		var graphQLClient *githubv4.Client
		if cfg.Token == "" && (cfg.AppId == "" || cfg.InstallationId == "" || cfg.PrivateKeyPath == "") {
			return nil, fmt.Errorf("'token' or 'app_id', 'installation_id' and 'private_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
		}
		// Return error for unsupported token by prefix
		if cfg.Token != "" && !strings.HasPrefix(cfg.Token, "ghs_") && !strings.HasPrefix(cfg.Token, "ghp_") && !strings.HasPrefix(cfg.Token, "gho_") {
			return nil, fmt.Errorf("supported token formats are 'ghs_', 'gho_', and 'ghp_'")
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
		var transport *ghinstallation.Transport
		// Authentication as Github APP Installation authentication
		if cfg.AppId != "" && cfg.InstallationId != "" && cfg.PrivateKeyPath != "" && cfg.Token == "" {
			ghAppId, err := strconv.ParseInt(cfg.AppId, 10, 64)
			if err != nil {
				return nil, err
			}
			ghInstallationId, err := strconv.ParseInt(cfg.InstallationId, 10, 64)
			if err != nil {
				return nil, err
			}
			itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, ghAppId, ghInstallationId, cfg.PrivateKeyPath)
			if err != nil {
				return nil, fmt.Errorf("Error occurred in 'connect()' during GitHub App Installation client creation: " + err.Error())
			}
			transport = itr
			restClient = github.NewClient(&http.Client{Transport: itr})
			graphQLClient = githubv4.NewClient(&http.Client{Transport: itr})
		}
		// If the base URL was provided then set it on the client. Used for enterprise installs.
		if cfg.BaseURL != "" {
			uv4, err := url.Parse(cfg.BaseURL)
			if err != nil {
				return nil, fmt.Errorf("github.base_url is invalid: %s", cfg.BaseURL)
			}
			if uv4.String() != "https://api.github.com/" {
				uv4.Path = uv4.Path + "api/v3/"
			}
			// The upload URL is not set as it's not currently required
			conn, err := github.NewClient(restClient.Client()).WithEnterpriseURLs(uv4.String(), "")
			if err != nil {
				return nil, fmt.Errorf("error creating GitHub client: %v", err)
			}
			conn.BaseURL = uv4
			restClient = conn
			uv4, err = url.Parse(cfg.BaseURL)
			if err != nil {
				return nil, fmt.Errorf("github.base_url is invalid: %s", cfg.BaseURL)
			}
			if uv4.String() != "https://api.github.com/" {
				uv4.Path = uv4.Path + "api/graphql"
			}
			if cfg.Token != "" {
				ts := oauth2.StaticTokenSource(
					&oauth2.Token{AccessToken: cfg.Token},
				)
				tc := oauth2.NewClient(ctx, ts)
				graphQLClient = githubv4.NewEnterpriseClient(uv4.String(), tc)
			} else {
				graphQLClient = githubv4.NewEnterpriseClient(uv4.String(), &http.Client{Transport: transport})
			}
		}
		client := describer.GitHubClient{
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
