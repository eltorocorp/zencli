package command

type token string

const (
	// HELP token
	HELP token = "help"
	// DROP token
	DROP token = "drop"
	// LIST token
	LIST token = "list"
	// BACKLOG token
	BACKLOG token = "backlog"
	//ONLY token
	ONLY token = "only"
	// MOVE token
	MOVE token = "move"
	// TO token
	TO token = "to"
	// PICK token
	PICK token = "pick"
	// UP token
	UP token = "up"
)

var tokens = []token{HELP, DROP, LIST, BACKLOG, ONLY, MOVE, TO, PICK, UP}
