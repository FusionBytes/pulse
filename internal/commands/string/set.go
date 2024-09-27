package string_commands

import (
	"errors"
	"pulse/pkg/structure"
	"strings"
)

type SET struct {
	hashTable *structure.HashTable
}

func NewSET() *SET {
	return &SET{
		hashTable: newHashTable(),
	}
}

func (s SET) CanDo(cmd string) bool {
	return strings.ToUpper(cmd) == "SET"
}

func (s SET) Execute(args []string) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New("key and value is not provided")
	}
	if args[0] == "" {
		return nil, errors.New("invalid key")
	}
	if args[1] == "" {
		return nil, errors.New("invalid value")
	}

	err := s.hashTable.Insert(args[0], args[1])
	if err != nil {
		return nil, err
	}

	return true, nil
}
