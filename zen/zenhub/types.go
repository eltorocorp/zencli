package zenhub

// Pipelines represents a slice of zenhub pipelines.
type Pipelines struct {
	List []Pipeline `json:"pipelines"`
}

// Pipeline represents a zenhub pipeline.
type Pipeline struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Issues      []Issue `json:"issues"`
	IssueNumber int     `json:"issue_number"`
	IsEpic      bool    `json:"is_epic"`
}

// Issue represents a zenhub issue.
type Issue struct {
	IssueNumber int      `json:"issue_number"`
	Estimate    Estimate `json:"estimate"`
	Position    int      `json:"position"`
	IsEpic      bool     `json:"is_epic"`
}

// Estimate represents a zenhub estimate.
type Estimate struct {
	Value int `json:"value"`
}

// PipelineMove represents the destination pipeline when moving an issue between pipelines.
type PipelineMove struct {
	PipelineID string `json:"pipeline_id"`
	Position   string `json:"position"`
}
