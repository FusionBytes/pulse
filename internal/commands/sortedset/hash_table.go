package sortedset

import (
	"pulse/pkg/structure"
	"sync"
)

var once sync.Once
var hashTable *structure.HashTable

func newHashTable() *structure.HashTable {
	once.Do(func() {
		hashTable = structure.NewHashTable(8, 0.75)
	})

	return hashTable
}
