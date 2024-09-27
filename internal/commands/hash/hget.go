package hashcommands

import (
	"errors"
	"pulse/pkg/structure"
	"strings"
)

type HGET struct {
	hashTable *structure.HashTable
}

func NewHGET() *HGET {
	return &HGET{
		hashTable: newHashTable(),
	}
}

func (g HGET) CanDo(cmd string) bool {
	return strings.ToUpper(cmd) == "HGET"
}

func (g HGET) Execute(args []string) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New("collection and key is not provided")
	}
	if args[0] == "" {
		return nil, errors.New("invalid collection")
	}
	if args[1] == "" {
		return nil, errors.New("invalid key")
	}

	value, _ := g.hashTable.Get(args[0])
	if value == nil {
		return nil, nil
	}

	var collection *structure.HashTable
	if c, ok := value.(*structure.HashTable); ok {
		collection = c
	} else {
		return nil, errors.New("unable to create collection")
	}

	value, _ = collection.Get(args[1])

	return value, nil
}
