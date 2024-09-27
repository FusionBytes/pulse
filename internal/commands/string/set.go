package stringcommands

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
	err := s.validateArgs(args)
	if err != nil {
		return nil, err
	}

	err = s.hashTable.Insert(args[0], args[1])
	if err != nil {
		return nil, err
	}

	return true, nil
}

func (s SET) validateArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("expected at least 2 arguments")
	}
	if args[0] == "" {
		return errors.New("invalid key")
	}
	if args[1] == "" {
		return errors.New("invalid value")
	}
	return nil
}
