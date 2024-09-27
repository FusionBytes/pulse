package hashcommands

import (
	"errors"
	"hash/fnv"
	"pulse/pkg/structure"
	"strings"
)

type HSET struct {
	hashTable *structure.HashTable
}

func NewHSET() *HSET {
	return &HSET{
		hashTable: newHashTable(),
	}
}

func (s HSET) CanDo(cmd string) bool {
	return strings.ToUpper(cmd) == "HSET"
}

func (s HSET) Execute(args []string) (interface{}, error) {
	if len(args) < 3 {
		return nil, errors.New("collection,key and value is not provided")
	}
	if args[0] == "" {
		return nil, errors.New("invalid collection")
	}
	if args[1] == "" {
		return nil, errors.New("invalid key")
	}
	if args[2] == "" {
		return nil, errors.New("invalid value")
	}

	var collection *structure.HashTable
	value, _ := s.hashTable.Get(args[0])
	if value == nil {
		collection = structure.NewHashTable(fnv.New64a(), 8, 0.75)
		err := s.hashTable.Insert(args[0], collection)
		if err != nil {
			return nil, err
		}
	} else if c, ok := value.(*structure.HashTable); ok {
		collection = c
	} else {
		return nil, errors.New("unable to create collection")
	}

	err := collection.Insert(args[1], args[2])
	if err != nil {
		return nil, err
	}

	return collection.EntryCount(), nil
}
