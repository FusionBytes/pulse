package string_commands

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
	if len(args) == 0 {
		return nil, errors.New("key is not provided")
	}
	if args[0] == "" {
		return nil, errors.New("invalid key")
	}

	value, _ := g.hashTable.Get(args[0])

	return value, nil
}
