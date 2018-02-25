package zenhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/eltorocorp/zencli/zen/github"
)

const (
	zenhubRoot = "https://api.zenhub.io"
)

// API provides methods for interacting with ZenHub.
type API struct {
	githubAPI       *github.API
	zenHubAuthToken string
}

// New returns a reference to a ZenHub API
func New(zenHubAuthToken string, githubAPI *github.API) *API {
	return &API{
		zenHubAuthToken: zenHubAuthToken,
		githubAPI:       githubAPI,
	}
}

// GetPipelines returns a list of pipelines.
func (a *API) GetPipelines() (*Pipelines, error) {
	repoID, err := a.githubAPI.GetRepoID()
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	getPipelinesURI := fmt.Sprintf("%v/p1/repositories/%v/board", zenhubRoot, *repoID)
	request, err := http.NewRequest(http.MethodGet, getPipelinesURI, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("X-Authentication-Token", a.zenHubAuthToken)
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the get pipelines endpoint returned %v", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	pipelines := new(Pipelines)
	err = json.Unmarshal(body, pipelines)

	return pipelines, err
}

// MovePipeline moves the specified issue to the specified pipeline.
func (a *API) MovePipeline(issue int, pipelineID string) error {
	repoID, err := a.githubAPI.GetRepoID()
	if err != nil {
		return err
	}

	pipelineMove := &PipelineMove{
		PipelineID: pipelineID,
		Position:   "top",
	}
	pipelineMoveJSON, err := json.Marshal(pipelineMove)
	if err != nil {
		return err
	}

	client := http.DefaultClient
	getPipelinesURI := fmt.Sprintf("%v/p1/repositories/%v/issues/%v/moves", zenhubRoot, *repoID, issue)
	request, err := http.NewRequest(http.MethodPost, getPipelinesURI, nil)
	if err != nil {
		return err
	}

	log.Println(string(pipelineMoveJSON))

	request.Body = ioutil.NopCloser(bytes.NewReader(pipelineMoveJSON))
	request.Header.Add("X-Authentication-Token", a.zenHubAuthToken)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("the move issue endpoint returned %v", response.StatusCode)
	}
	return nil
}
