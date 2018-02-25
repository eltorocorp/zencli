package command

import (
	"errors"
	"strconv"
)

func (c *API) parserError() error {
	return errors.New("the supplied arguments could not be parsed. Run `zen help` for usage information")
}

func (c *API) nextSymbol() bool {
	c.symbolIndex++
	if len(c.args) < c.symbolIndex+1 {
		return false
	}
	c.currentSymbol = c.args[c.symbolIndex]
	return true
}

func (c *API) previousSymbol() bool {
	c.symbolIndex--
	if c.symbolIndex < 0 {
		return false
	}
	c.currentSymbol = c.args[c.symbolIndex]
	return true
}

func (c *API) acceptToken(t token) bool {
	if c.currentSymbol == string(t) {
		return true
	}
	for _, token := range tokens {
		if c.currentSymbol == string(token) {
			return false
		}
	}
	return true
}

func (c *API) expectToken(t token) bool {
	return c.currentSymbol == string(t)
}

func (c *API) expectCurrentSymbolInt(out *int) bool {
	value, err := strconv.Atoi(c.currentSymbol)
	if err != nil {
		return false
	}
	if out != nil {
		*out = value
	}
	return true
}

func (c *API) expectCurrentSymbolString(out *string) bool {
	for _, token := range tokens {
		if c.currentSymbol == string(token) {
			return false
		}
	}
	if out != nil {
		*out = c.currentSymbol
	}
	return true
}
