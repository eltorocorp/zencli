package command

import (
	"strings"
)

// The API for the command
type API struct {
	args          []string
	actions       Actions
	currentSymbol string
	symbolIndex   int
}

// The Actions that the command is able to execute.
type Actions interface {
	Help()
	Drop(issue int) error
	List(backlog bool, login string) error
	Move(issue int, pipeline string) error
	PickUp(issue int) error
}

// New returns a command API capable of parsing the supplied args and execution the appropriate commands.
func New(args []string, actions Actions) *API {
	if len(args) == 0 {
		panic("args cannot be emptyt")
	}
	return &API{
		args:          args,
		currentSymbol: args[0],
		symbolIndex:   0,
		actions:       actions,
	}
}

// Execute parses the supplied args and runs the appropriate commands based on the parsed command.
func (c *API) Execute() error {
	var (
		issue    int
		pipeline string
		login    string
		backlog  bool
	)

	for _, symbol := range c.args {
		if symbol == string(HELP) {
			c.actions.Help()
			return nil
		}
	}
	if !c.nextSymbol() {
		return c.parserError()
	}
	if c.expectToken(DROP) &&
		c.nextSymbol() &&
		c.expectCurrentSymbolInt(&issue) {
		return c.actions.Drop(issue)
	} else if c.expectToken(LIST) {
		backlog = false
		for c.nextSymbol() {
			if c.expectToken(ONLY) &&
				c.nextSymbol() &&
				c.expectCurrentSymbolString(&login) {
				continue
			} else if c.expectToken(BACKLOG) {
				backlog = true
				continue
			}
			return c.parserError()
		}
		return c.actions.List(backlog, login)
	} else if c.expectToken(MOVE) &&
		c.nextSymbol() &&
		c.expectCurrentSymbolInt(&issue) &&
		c.nextSymbol() &&
		c.expectToken(TO) {
		pipelineName := []string{}
		for c.nextSymbol() {
			if c.expectCurrentSymbolString(&pipeline) {
				pipelineName = append(pipelineName, pipeline)
			} else {
				return c.parserError()
			}
		}
		return c.actions.Move(issue, strings.Join(pipelineName, " "))
	} else if c.expectToken(PICK) &&
		c.nextSymbol() &&
		c.expectToken(UP) &&
		c.nextSymbol() &&
		c.expectCurrentSymbolInt(&issue) {
		return c.actions.PickUp(issue)
	} else {
		return c.parserError()
	}
}
