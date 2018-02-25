package github

// Repository represents a github repository.
type Repository struct {
	ID int `json:"id"`
}

// Issue represents a github issue.
type Issue struct {
	Number    int    `json:"number"`
	State     string `json:"state"`
	Title     string `json:"title"`
	Assignee  User   `json:"assignee"`
	Assignees []User `json:"assignees"`
}

// User represents a github user.
type User struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

// Assignees represents a list of logins associated with an issue.
type Assignees struct {
	List []string `json:"assignees"`
}
