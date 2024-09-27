package hashcommands

import (
	"hash/fnv"
	"pulse/pkg/structure"
)

var hashTable *structure.HashTable

func newHashTable() *structure.HashTable {
	if hashTable == nil {
		hashTable = structure.NewHashTable(fnv.New64a(), 8, 0.75)
	}

	return hashTable
}
