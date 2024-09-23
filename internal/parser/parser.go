package parser

import (
	"errors"
	"strings"
)

type commander interface {
	CanDo(cmd string) bool
	Execute(args []string) (interface{}, error)
}


type Parser struct {
	commanders []commander
}

func NewParser(
	commanders ...commander,
) *Parser {
	return &Parser{
		commanders: commanders,
	}
}

func (p *Parser) getCommand(method string) (commander, error) {
	for _, cmd := range p.commanders {
		if cmd.CanDo(method) {
			return cmd, nil
		}
	}
	return nil, errors.New("there is no commander")
}

func (p *Parser) Execute(cmd string) (interface{}, error) {
	args := strings.Split(cmd, " ")

	if args[0] == "" {
		return nil, errors.New("invalid command method")
	}
	cmder, err := p.getCommand(args[0])

	if err != nil {
		return nil, err
	}
	return cmder.Execute(args[1:])
}