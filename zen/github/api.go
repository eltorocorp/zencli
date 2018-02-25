package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	githubRoot           = "https://api.github.com"
	githubV3AcceptHeader = "application/vnd.github.v3+json"
)

// API provides methods for interacting with github.
type API struct {
	githubAuthToken string
	RepoName        string
	ownerName       string
}

// New returns a reference to a github API.
func New(githubAuthToken, repoName, ownerName string) *API {
	return &API{
		githubAuthToken: githubAuthToken,
		RepoName:        repoName,
		ownerName:       ownerName,
	}
}

//GetRepoID returns the ID for the target repository
func (a *API) GetRepoID() (*int, error) {
	client := http.DefaultClient
	getRepoURI := fmt.Sprintf("%v/repos/%v/%v?access_token=%v", githubRoot, a.ownerName, a.RepoName, a.githubAuthToken)
	request, err := createDefaultRequest(http.MethodGet, getRepoURI)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the repo endpoint returned %v", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	repository := new(Repository)
	err = json.Unmarshal(body, repository)

	return &repository.ID, err
}

// GetIssuesForRepo gets a list of issues for the target repository.
func (a *API) GetIssuesForRepo() (*[]*Issue, error) {
	client := http.DefaultClient
	getRepoURI := fmt.Sprintf("%v/repos/%v/%v/issues?access_token=%v", githubRoot, a.ownerName, a.RepoName, a.githubAuthToken)
	request, err := createDefaultRequest(http.MethodGet, getRepoURI)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the issues endpoint returned %v", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	issues := new([]*Issue)
	err = json.Unmarshal(body, issues)

	return issues, err
}

// GetAuthenticatedUser gets the current authenticated user.
func (a *API) GetAuthenticatedUser() (*User, error) {
	client := http.DefaultClient
	getRepoURI := fmt.Sprintf("%v/user?access_token=%v", githubRoot, a.githubAuthToken)
	request, err := createDefaultRequest(http.MethodGet, getRepoURI)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the issues endpoint returned %v", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	user := new(User)
	err = json.Unmarshal(body, user)

	return user, err

}

// RemoveAuthenticatedUserFromIssue removes the current authenticated user from the specified issue.
func (a *API) RemoveAuthenticatedUserFromIssue(issue int) error {
	currentUser, err := a.GetAuthenticatedUser()
	if err != nil {
		return err
	}

	assigneeToRemove := &Assignees{
		List: []string{currentUser.Login},
	}
	assigneesJSON, err := json.Marshal(assigneeToRemove)
	if err != nil {
		return err
	}

	client := http.DefaultClient
	getRepoURI := fmt.Sprintf("%v/repos/%v/%v/issues/%v/assignees?access_token=%v", githubRoot, a.ownerName, a.RepoName, issue, a.githubAuthToken)
	request, err := createDefaultRequest(http.MethodDelete, getRepoURI)
	if err != nil {
		return err
	}

	request.Body = ioutil.NopCloser(bytes.NewReader(assigneesJSON))
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("the remove assignee endpoint returned %v", response.StatusCode)
	}

	return nil
}

// AssignAuthenticatedUserToIssue assigns the current authenticated user to the specified issue.
func (a *API) AssignAuthenticatedUserToIssue(issue int) error {
	currentUser, err := a.GetAuthenticatedUser()
	if err != nil {
		return err
	}

	assigneeToAdd := &Assignees{
		List: []string{currentUser.Login},
	}
	assigneesJSON, err := json.Marshal(assigneeToAdd)
	if err != nil {
		return err
	}

	client := http.DefaultClient
	getRepoURI := fmt.Sprintf("%v/repos/%v/%v/issues/%v/assignees?access_token=%v", githubRoot, a.ownerName, a.RepoName, issue, a.githubAuthToken)
	request, err := createDefaultRequest(http.MethodPost, getRepoURI)
	if err != nil {
		return err
	}

	request.Body = ioutil.NopCloser(bytes.NewReader(assigneesJSON))
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("the add assignee endpoint returned %v", response.StatusCode)
	}

	return nil
}

// CloseIssue closes the specified issue.
func (a *API) CloseIssue(issue int) error {
	issueToClose := struct {
		State string `json:"state"`
	}{
		State: "closed",
	}
	issueToCloseJSON, err := json.Marshal(&issueToClose)
	if err != nil {
		return err
	}

	client := http.DefaultClient
	getRepoURI := fmt.Sprintf("%v/repos/%v/%v/issues/%v?access_token=%v", githubRoot, a.ownerName, a.RepoName, issue, a.githubAuthToken)
	request, err := createDefaultRequest(http.MethodPatch, getRepoURI)
	if err != nil {
		return err
	}

	request.Body = ioutil.NopCloser(bytes.NewReader(issueToCloseJSON))
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("the close issue endpoint returned %v", response.StatusCode)
	}

	return nil
}

// OpenIssue opens the specified issue.
func (a *API) OpenIssue(issue int) error {
	issueToClose := struct {
		State string `json:"state"`
	}{
		State: "open",
	}
	issueToCloseJSON, err := json.Marshal(&issueToClose)
	if err != nil {
		return err
	}

	client := http.DefaultClient
	getRepoURI := fmt.Sprintf("%v/repos/%v/%v/issues/%v?access_token=%v", githubRoot, a.ownerName, a.RepoName, issue, a.githubAuthToken)
	request, err := createDefaultRequest(http.MethodPatch, getRepoURI)
	if err != nil {
		return err
	}

	request.Body = ioutil.NopCloser(bytes.NewReader(issueToCloseJSON))
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("the open issue endpoint returned %v", response.StatusCode)
	}

	return nil
}

func createDefaultRequest(method, uri string) (*http.Request, error) {
	request, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", githubV3AcceptHeader)
	return request, nil
}
