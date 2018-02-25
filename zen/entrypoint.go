package main

import (
	"fmt"
	"os"

	"github.com/eltorocorp/zencli/zen/command"
	"github.com/eltorocorp/zencli/zen/github"
	"github.com/eltorocorp/zencli/zen/zenhub"
)

func main() {
	githubAuthToken := os.Getenv("ZENCLI_GITHUBAUTHTOKEN")
	zenHubAuthToken := os.Getenv("ZENCLI_ZENHUBAUTHTOKEN")
	repoOwner := os.Getenv("ZENCLI_REPOOWNER")
	repoName := os.Getenv("ZENCLI_REPONAME")

	githubAPI := github.New(githubAuthToken, repoName, repoOwner)
	zenHubAPI := zenhub.New(zenHubAuthToken, githubAPI)
	actions := NewActions(githubAPI, zenHubAPI)

	cmd := command.New(os.Args, actions)
	err := cmd.Execute()
	handleAnyErrorAndExit(err)
}

func handleAnyErrorAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
