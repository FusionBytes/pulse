package sortedset

import "pulse/pkg/structure"

type Scoreboard struct {
	skipList *structure.SkipList
	hashMap  *structure.HashTable
}
