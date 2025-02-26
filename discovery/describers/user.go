package describers

import (
	"context"
	"encoding/json"
	"fmt"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

const MAX_REPOS = 500

func GetUser(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	var values []models.Resource

	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	users, err := ListUsers(sdk, organizationName, "", true, false, false)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		value := models.Resource{
			ID:   strconv.FormatInt(user.ID, 10),
			Name: user.Login,
			Description: model.UserDescription{
				ID:           user.ID,
				Name:         user.Name,
				Login:        user.Login,
				Company:      user.Company,
				Email:        user.Email,
				Location:     user.Location,
				Url:          user.URL,
				NodeId:       user.NodeID,
				CreatedAt:    user.CreatedAt,
				UpdatedAt:    user.UpdatedAt,
				Organization: organizationName,
			},
		}

		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

// ----------------------------------------------------------
// ExtendedUser => final user object in output
// ----------------------------------------------------------

type ExtendedUser struct {
	Login    string `json:"login"`
	ID       int64  `json:"id"`      // numeric ID
	NodeID   string `json:"node_id"` // global node ID
	Name     string `json:"name"`
	Email    string `json:"email"`
	Company  string `json:"company"`
	Location string `json:"location"`
	URL      string `json:"url"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ----------------------------------------------------------
// REST fallback user struct
// ----------------------------------------------------------

type restUser struct {
	Login  string `json:"login"`
	ID     int64  `json:"id"`
	NodeID string `json:"node_id"`
	URL    string `json:"html_url"`
	Type   string `json:"type"` // "User", "Bot", etc.

	Name     string `json:"name"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Email    string `json:"email"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ----------------------------------------------------------
// Repo struct => minimal from GET /orgs/{org}/repos
// ----------------------------------------------------------

type Repo struct {
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
	Disabled bool   `json:"disabled"`
	Fork     bool   `json:"fork"`
	Owner    struct {
		Login string `json:"login"`
	} `json:"owner"`
	Parent *struct {
		FullName string `json:"full_name"`
		Owner    struct {
			Login string `json:"login"`
		} `json:"owner"`
	} `json:"parent,omitempty"`
}

// ----------------------------------------------------------
// Collab & Contributor => from /collaborators & /contributors
// ----------------------------------------------------------

type Collaborator struct {
	Login       string          `json:"login"`
	ID          int             `json:"id"`
	SiteAdmin   bool            `json:"site_admin"`
	Permissions map[string]bool `json:"permissions"`
}

type Contributor struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

// ----------------------------------------------------------
// ListUsers => the "library" function that returns a list of ExtendedUser
// ----------------------------------------------------------

func ListUsers(
	sdk *resilientbridge.ResilientBridge,
	org string,
	repo string,
	skipExternal bool,
	skipAll bool,
	debug bool,
) ([]ExtendedUser, error) {

	// 1) gather repos
	repos, err := gatherRepos(sdk, org, repo, debug)
	if err != nil {
		return nil, err
	}

	// 2) filter repos => skip archived/disabled + check fork logic
	finalRepos, err := filterRepos(sdk, repos, org, skipExternal, skipAll, debug)
	if err != nil {
		return nil, err
	}
	if len(finalRepos) == 0 {
		if debug {
			log.Printf("No repos remain after filtering; returning empty user list.")
		}
		return []ExtendedUser{}, nil
	}

	// 3) collect collaborators + contributors => unique set of logins
	allLogins, err := collectRepoUsers(sdk, finalRepos, debug)
	if err != nil {
		return nil, err
	}
	if len(allLogins) == 0 {
		return []ExtendedUser{}, nil
	}

	// 4) batch GraphQL => fallback => build final
	var loginsSlice []string
	for l := range allLogins {
		loginsSlice = append(loginsSlice, l)
	}

	usersMap, gqlErrs := fetchBatchUsersGraphQL(sdk, loginsSlice, debug)

	// If debug, show partial errors
	if debug && len(gqlErrs) > 0 {
		log.Printf("Some GraphQL user aliases had errors or null => %v", gqlErrs)
	}

	var finalUsers []ExtendedUser
	for _, login := range loginsSlice {
		lower := strings.ToLower(login)
		if eu, ok := usersMap[lower]; ok && eu != nil {
			finalUsers = append(finalUsers, *eu)
			continue
		}

		// fallback => REST
		ru, err := fetchUserViaREST(sdk, login)
		if err != nil {
			if debug {
				log.Printf("Skipping user %q => fallback REST error: %v", login, err)
			}
			continue
		}
		if strings.EqualFold(ru.Type, "Bot") {
			if debug {
				log.Printf("Skipping bot user %q from fallback REST", login)
			}
			continue
		}
		var ctime, utime time.Time
		if t, e := time.Parse(time.RFC3339, ru.CreatedAt); e == nil {
			ctime = t
		}
		if t, e := time.Parse(time.RFC3339, ru.UpdatedAt); e == nil {
			utime = t
		}
		euser := ExtendedUser{
			Login:    ru.Login,
			ID:       ru.ID,
			NodeID:   ru.NodeID,
			Name:     emptyToNull(ru.Name),
			Email:    emptyToNull(ru.Email),
			Company:  emptyToNull(ru.Company),
			Location: emptyToNull(ru.Location),
			URL:      ru.URL,

			CreatedAt: ctime,
			UpdatedAt: utime,
		}
		finalUsers = append(finalUsers, euser)
	}

	return finalUsers, nil
}

// ----------------------------------------------------------
// gatherRepos => singleRepo or GET /orgs/{org}/repos
// ----------------------------------------------------------

func gatherRepos(sdk *resilientbridge.ResilientBridge, org, singleRepo string, debug bool) ([]Repo, error) {
	if singleRepo != "" {
		r := Repo{
			Name: singleRepo,
		}
		r.Owner.Login = org
		return []Repo{r}, nil
	}

	// else fetch up to MAX_REPOS
	list, err := fetchOrgRepos(sdk, org, MAX_REPOS)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ----------------------------------------------------------
// filterRepos => skip archived/disabled, fetch "parent" if fork, skip if external/...
// ----------------------------------------------------------

func filterRepos(
	sdk *resilientbridge.ResilientBridge,
	raw []Repo,
	org string,
	skipExternalForks bool,
	skipAllForks bool,
	debug bool,
) ([]Repo, error) {

	var final []Repo
	for _, r := range raw {
		if r.Archived {
			if debug {
				log.Printf("Skipping repo %s => archived", r.Name)
			}
			continue
		}
		if r.Disabled {
			if debug {
				log.Printf("Skipping repo %s => disabled", r.Name)
			}
			continue
		}
		if r.Fork {
			// fetch details => parent
			detail, err := fetchRepoDetails(sdk, r.Owner.Login, r.Name)
			if err != nil {
				if debug {
					log.Printf("Skipping forked repo %s => can't fetch parent: %v", r.Name, err)
				}
				continue
			}
			r.Parent = detail.Parent
		}
		ok, reason := isHealthyOrgRepo(r, org, skipExternalForks, skipAllForks)
		if !ok {
			if debug {
				log.Printf("Skipping repo %s => %s", r.Name, reason)
			}
			continue
		}
		final = append(final, r)
	}
	return final, nil
}

func isHealthyOrgRepo(r Repo, org string, skipExternalForks, skipAllForks bool) (bool, string) {
	if r.Fork {
		if skipAllForks {
			return false, "fork & skip-all-forks=true"
		}
		if skipExternalForks {
			if r.Parent == nil {
				return false, "fork but no parent => skip external by default"
			}
			parentOrg := r.Parent.Owner.Login
			if !strings.EqualFold(parentOrg, org) {
				return false, fmt.Sprintf("fork from external org %q & skip-external-forks=true", parentOrg)
			}
		}
	}
	return true, ""
}

// ----------------------------------------------------------
// collectRepoUsers => gather collaborator & contributor logins
// ----------------------------------------------------------

func collectRepoUsers(
	sdk *resilientbridge.ResilientBridge,
	repos []Repo,
	debug bool,
) (map[string]bool, error) {

	allLogins := make(map[string]bool)
	for _, r := range repos {
		if debug {
			log.Printf("Processing healthy org repo: %s", r.Name)
		}

		collabs, err := fetchRepoCollaborators(sdk, r.Owner.Login, r.Name)
		if err != nil {
			if debug {
				log.Printf("Skipping repo %s => collaborator error: %v", r.Name, err)
			}
			continue
		}
		conts, err := fetchRepoContributors(sdk, r.Owner.Login, r.Name)
		if err != nil {
			if debug {
				log.Printf("Skipping repo %s => contributor error: %v", r.Name, err)
			}
			continue
		}
		for _, c := range collabs {
			allLogins[strings.ToLower(c.Login)] = true
		}
		for _, c := range conts {
			allLogins[strings.ToLower(c.Login)] = true
		}
	}
	return allLogins, nil
}

// ----------------------------------------------------------
// fetchOrgRepos => GET /orgs/{org}/repos?per_page=MAX_REPOS
// ----------------------------------------------------------

func fetchOrgRepos(sdk *resilientbridge.ResilientBridge, org string, max int) ([]Repo, error) {
	endpoint := fmt.Sprintf("/orgs/%s/repos?per_page=%d&page=1", org, max)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github.v3+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching org repos => %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
	}
	var repos []Repo
	if err := json.Unmarshal(resp.Data, &repos); err != nil {
		return nil, fmt.Errorf("unmarshal repos => %w", err)
	}
	if len(repos) > max {
		repos = repos[:max]
	}
	return repos, nil
}

// fetchRepoDetails => GET /repos/{owner}/{repo} to retrieve parent for forks
func fetchRepoDetails(sdk *resilientbridge.ResilientBridge, owner, repo string) (*Repo, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s", owner, repo)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github.v3+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("repo detail: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
	}
	var r Repo
	if e := json.Unmarshal(resp.Data, &r); e != nil {
		return nil, fmt.Errorf("unmarshal repo detail => %w", e)
	}
	return &r, nil
}

// ----------------------------------------------------------
// fetchRepoCollaborators => GET /repos/{org}/{repo}/collaborators
// ----------------------------------------------------------

func fetchRepoCollaborators(sdk *resilientbridge.ResilientBridge, org, repo string) ([]Collaborator, error) {
	var all []Collaborator
	page := 1
	perPage := 100
	for {
		endpoint := fmt.Sprintf("/repos/%s/%s/collaborators?per_page=%d&page=%d", org, repo, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github.v3+json"},
		}
		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("page %d collaborator error: %w", page, err)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d on page %d: %s", resp.StatusCode, page, string(resp.Data))
		}
		var batch []Collaborator
		if err := json.Unmarshal(resp.Data, &batch); err != nil {
			return nil, fmt.Errorf("unmarshal collaborators (page %d): %w", page, err)
		}
		if len(batch) == 0 {
			break
		}
		all = append(all, batch...)
		if len(batch) < perPage {
			break
		}
		page++
	}
	return all, nil
}

// ----------------------------------------------------------
// fetchRepoContributors => GET /repos/{org}/{repo}/contributors
// ----------------------------------------------------------

func fetchRepoContributors(sdk *resilientbridge.ResilientBridge, org, repo string) ([]Contributor, error) {
	var all []Contributor
	page := 1
	perPage := 100
	for {
		endpoint := fmt.Sprintf("/repos/%s/%s/contributors?per_page=%d&page=%d", org, repo, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github.v3+json"},
		}
		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("page %d contributor error: %w", page, err)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d on page %d: %s", resp.StatusCode, page, string(resp.Data))
		}
		var batch []Contributor
		if err := json.Unmarshal(resp.Data, &batch); err != nil {
			return nil, fmt.Errorf("unmarshal contributors (page %d): %w", page, err)
		}
		if len(batch) == 0 {
			break
		}
		all = append(all, batch...)
		if len(batch) < perPage {
			break
		}
		page++
	}
	return all, nil
}

// ----------------------------------------------------------
// fetchBatchUsersGraphQL => single GraphQL request
// ----------------------------------------------------------

func fetchBatchUsersGraphQL(
	sdk *resilientbridge.ResilientBridge,
	logins []string,
	debug bool,
) (map[string]*ExtendedUser, []error) {

	out := make(map[string]*ExtendedUser)
	if len(logins) == 0 {
		return out, nil
	}

	var sb strings.Builder
	sb.WriteString("query {\n")
	for i, login := range logins {
		alias := fmt.Sprintf("user_%d", i)
		sb.WriteString(fmt.Sprintf(`
  %s: user(login: "%s") {
    __typename
    login
    databaseId
    id
    name
    email
    company
    location
    url
    createdAt
    updatedAt
  }
`, alias, login))
	}
	sb.WriteString("}\n")

	reqBody := map[string]interface{}{
		"query": sb.String(),
	}
	reqBytes, _ := json.Marshal(reqBody)

	data, err := doGraphQLRequest(sdk, reqBytes)
	if err != nil {
		return out, []error{err}
	}
	var payload struct {
		Data map[string]*struct {
			Typename   string `json:"__typename"`
			Login      string `json:"login"`
			DatabaseID int64  `json:"databaseId"`
			ID         string `json:"id"` // global node ID
			Name       string `json:"name"`
			Email      string `json:"email"`
			Company    string `json:"company"`
			Location   string `json:"location"`
			URL        string `json:"url"`
			CreatedAt  string `json:"createdAt"`
			UpdatedAt  string `json:"updatedAt"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
			Path    []any  `json:"path"`
		} `json:"errors"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return out, []error{fmt.Errorf("unmarshal GraphQL response error: %w", err)}
	}

	var errs []error
	for alias, userObj := range payload.Data {
		if userObj == nil {
			errs = append(errs, fmt.Errorf("%s => null (inaccessible)", alias))
			continue
		}
		if userObj.Typename == "Bot" {
			errs = append(errs, fmt.Errorf("%s => Bot user => skipping", alias))
			continue
		}
		var ctime, utime time.Time
		if t, e := time.Parse(time.RFC3339, userObj.CreatedAt); e == nil {
			ctime = t
		}
		if t, e := time.Parse(time.RFC3339, userObj.UpdatedAt); e == nil {
			utime = t
		}
		eu := &ExtendedUser{
			Login:     userObj.Login,
			ID:        userObj.DatabaseID,
			NodeID:    userObj.ID,
			Name:      emptyToNull(userObj.Name),
			Email:     emptyToNull(userObj.Email),
			Company:   emptyToNull(userObj.Company),
			Location:  emptyToNull(userObj.Location),
			URL:       userObj.URL,
			CreatedAt: ctime,
			UpdatedAt: utime,
		}
		out[strings.ToLower(eu.Login)] = eu
	}
	for _, e := range payload.Errors {
		errs = append(errs, fmt.Errorf("GraphQL error: %s (path=%v)", e.Message, e.Path))
	}

	return out, errs
}

// doGraphQLRequest => use ResilientBridge
func doGraphQLRequest(sdk *resilientbridge.ResilientBridge, body []byte) ([]byte, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "POST",
		Endpoint: "/graphql",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: body,
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("GraphQL status %d => %s", resp.StatusCode, string(resp.Data))
	}
	return resp.Data, nil
}

// fetchUserViaREST => fallback if GraphQL is null
func fetchUserViaREST(sdk *resilientbridge.ResilientBridge, login string) (*restUser, error) {
	endpoint := fmt.Sprintf("/users/%s", login)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("fetch user REST: %w", err)
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("REST: user %q not found (404)", login)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("REST: user %q => status %d => %s", login, resp.StatusCode, string(resp.Data))
	}

	var ru restUser
	if e := json.Unmarshal(resp.Data, &ru); e != nil {
		return nil, fmt.Errorf("unmarshal user => %w", e)
	}
	return &ru, nil
}

// emptyToNull => if you want real "null" in JSON for empty strings, you'd store pointers.
func emptyToNull(s string) string {
	return s
}
