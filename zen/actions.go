package main

import (
	"fmt"
	"strconv"

	"github.com/eltorocorp/zencli/zen/github"
	"github.com/eltorocorp/zencli/zen/zenhub"
)

// Actions are a set of methods that a command can execute.
type Actions struct {
	githubAPI *github.API
	zenHubAPI *zenhub.API
}

// NewActions returns a reference to a set of actions.
func NewActions(githubAPI *github.API, zenHubAPI *zenhub.API) *Actions {
	return &Actions{
		githubAPI: githubAPI,
		zenHubAPI: zenHubAPI,
	}
}

// Create creates a new issue in the specified pipeline.
func (a *Actions) Create(title, pipelineName string) error {
	var err error
	var pipelineID string

	fmt.Printf("Creating new issue...\n")

	// Since backlog is the default pipeline, we save a few seconds by not checking if it exists (since a move won't be necessary later)
	if pipelineName != "backlog" {
		pipelineID, err = a.zenHubAPI.GetPipelineID(pipelineName)
		if err != nil {
			return err
		}
	}

	newIssueNumber, err := a.githubAPI.CreateIssue(title)
	if err != nil {
		return err
	}

	fmt.Printf("Issue %v created in the backlog.", newIssueNumber)
	if pipelineName == "backlog" {
		fmt.Println()
		return nil
	}

	fmt.Printf(" Moving it to %v...\n", pipelineName)
	err = a.zenHubAPI.MovePipeline(newIssueNumber, pipelineID)
	if err != nil {
		return err
	}

	fmt.Printf("New issue (%v) has been created and moved to %v.\n", newIssueNumber, pipelineName)
	return nil
}

// Drop unassigns the current user from the specified issue.
func (a *Actions) Drop(issue int) error {
	fmt.Printf("Removing you from issue %v...\n", issue)
	err := a.githubAPI.RemoveAuthenticatedUserFromIssue(issue)
	if err == nil {
		fmt.Printf("You have been removed from issue %v.\n", issue)
	}
	return err
}

// List lists all active issues by pipeline.
//
// If backlog is true, the backlog pipeline will be included, otherwise the backlog is excluded.
// If login is non-nil only issues assigned to the specified login are shown (unassigned are still shown).
func (a *Actions) List(backlog bool, login string) error {
	const unassigned = "unassigned"
	fmt.Printf("Fetching issues from %v", a.githubAPI.RepoName)
	githubIssues, err := a.githubAPI.GetIssuesForRepo()
	if err != nil {
		return err
	}

	pipelines, err := a.zenHubAPI.GetPipelines()
	if err != nil {
		return err
	}

	if login == "me" {
		user, err := a.githubAPI.GetAuthenticatedUser()
		if err != nil {
			return err
		}
		login = user.Login
	}

	fmt.Printf("\rOpen issues for %v\n", pr(a.githubAPI.RepoName+":", 80))
	for _, pipeline := range pipelines.List {
		if backlog == false && pipeline.Name == "Backlog" {
			continue
		}
		fmt.Printf("%v (%v)\n", pipeline.Name, len(pipeline.Issues))
		for _, zenhubIssue := range pipeline.Issues {
			var issueName string
			issueAssignee := unassigned
			for _, githubIssue := range *githubIssues {
				if githubIssue.Number == zenhubIssue.IssueNumber {
					issueName = githubIssue.Title
					if githubIssue.Assignee.Login != "" {
						issueAssignee = githubIssue.Assignee.Login
					}
					break
				}
			}
			if issueAssignee != unassigned && login != "" && issueAssignee != login {
				continue
			}
			fmt.Printf(" - %v%v%v\n", pr(strconv.Itoa(zenhubIssue.IssueNumber), 6), pr(issueAssignee, 15), issueName)
		}
	}
	return nil
}

// Move changes the pipeline for the specified issue.
func (a *Actions) Move(issue int, pipelineName string) error {
	fmt.Printf("Moving issue %v to %v...\n", issue, pipelineName)
	pipelineID, err := a.zenHubAPI.GetPipelineID(pipelineName)
	if err != nil {
		return err
	}

	err = a.zenHubAPI.MovePipeline(issue, pipelineID)
	if err == nil {
		fmt.Printf("Issue %v has been moved to %v.\n", issue, pipelineName)
	}
	return err
}

// PickUp assigns the current user as an assignee to the specified issue.
func (a *Actions) PickUp(issue int) error {
	fmt.Printf("Assigning you to issue %v...\n", issue)
	err := a.githubAPI.AssignAuthenticatedUserToIssue(issue)
	if err == nil {
		fmt.Printf("You have been assigned to issue %v.\n", issue)
	}
	return err
}

// Close chages the status of the specified issue to closed.
func (a *Actions) Close(issue int) error {
	fmt.Printf("Closing issue %v...\n", issue)
	err := a.githubAPI.CloseIssue(issue)
	if err == nil {
		fmt.Printf("Issue %v has been closed.\n", issue)
	}
	return err
}

// Open chages the status of the specified issue to open.
func (a *Actions) Open(issue int) error {
	fmt.Printf("Openning issue %v...\n", issue)
	err := a.githubAPI.OpenIssue(issue)
	if err == nil {
		fmt.Printf("Issue %v has been opened.\n", issue)
	}
	return err
}

// Help displays the usage information.
func (a *Actions) Help() {
	fmt.Println(usage)
}

func pr(str string, length int) string {
	for {
		str += " "
		if len(str) > length {
			return str[0:length]
		}
	}
}
