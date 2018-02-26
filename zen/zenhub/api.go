package zenhub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	request, err := a.createDefaultRequest(http.MethodGet, getPipelinesURI)
	if err != nil {
		return nil, err
	}

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
	request, err := a.createDefaultRequest(http.MethodPost, getPipelinesURI)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Body = ioutil.NopCloser(bytes.NewReader(pipelineMoveJSON))
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("the move issue endpoint returned %v", response.StatusCode)
	}
	return nil
}

// GetPipelineID returns the ZenHub ID for the specified pipeline name. If the specified pipeline
// does not exist for the current board, this method will return an empty string and an error.
func (a *API) GetPipelineID(pipelineName string) (string, error) {
	pipelineID := ""
	pipelines, err := a.GetPipelines()
	if err != nil {
		return "", err
	}
	for _, pipeline := range pipelines.List {
		if strings.ToLower(pipeline.Name) == strings.ToLower(pipelineName) {
			pipelineID = pipeline.ID
			break
		}
	}

	if pipelineID == "" {
		return "", fmt.Errorf("pipeline '%v' does not exist for this board", pipelineName)
	}

	return pipelineID, nil
}

func (a *API) createDefaultRequest(method, uri string) (*http.Request, error) {
	request, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Authentication-Token", a.zenHubAuthToken)
	return request, nil
}
