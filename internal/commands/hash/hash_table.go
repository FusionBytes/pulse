package hashcommands

import (
	"hash/fnv"
	"pulse/pkg/structure"
	"sync"
)

var once sync.Once
var hashTable *structure.HashTable

func newHashTable() *structure.HashTable {
	once.Do(func() {
		hashTable = structure.NewHashTable(fnv.New64a(), 8, 0.75)
	})

	return hashTable
}
