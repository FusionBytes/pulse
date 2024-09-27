package stringcommands

import (
	"errors"
	"pulse/pkg/structure"
	"strings"
)

type GET struct {
	hashTable *structure.HashTable
}

func NewGET() *GET {
	return &GET{
		hashTable: newHashTable(),
	}
}

func (g GET) CanDo(cmd string) bool {
	return strings.ToUpper(cmd) == "GET"
}

func (g GET) Execute(args []string) (interface{}, error) {
	err := g.validateArgs(args)
	if err != nil {
		return nil, err
	}

	value, _ := g.hashTable.Get(args[0])

	return value, nil
}

func (g GET) validateArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("expected at least 1 arguments")
	}
	if args[0] == "" {
		return errors.New("invalid key")
	}
	return nil
}
