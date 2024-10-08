package structure

import (
	"errors"
	"math/rand"
	"time"
)

type Node struct {
	score  int
	member string
	next   *Node
	below  *Node
}

type SkipList struct {
	head         *Node
	highestLevel int
}

// NewSkipList initializes a new skip list
func NewSkipList() *SkipList {
	// Seed the random generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return &SkipList{
		head:         &Node{score: -1}, // head is a sentinel node
		highestLevel: 0,                // starting at level 0
	}
}

// Insert adds a new score to the skip list, allowing duplicates
func (sl *SkipList) Insert(score int, member string) {
	updatePath := sl.fetchUpdatePath(score)

	level := 0
	mustInsert := true
	var belowNode *Node

	for mustInsert {
		if level > sl.highestLevel {
			newHead := &Node{score: -1, next: nil, below: sl.head}
			sl.head = newHead
			updatePath = append(updatePath, newHead)
			sl.highestLevel = level
		}

		newNode := &Node{
			score:  score,
			member: member,
			next:   updatePath[level].next,
			below:  belowNode,
		}

		updatePath[level].next = newNode
		belowNode = newNode

		level++
		mustInsert = sl.mustAddNewLevel()
	}
}

// mustAddNewLevel decides if a new level should be added
// Uses a 50% probability (like a coin flip) to decide
func (sl *SkipList) mustAddNewLevel() bool {
	return rand.Float32() < 0.5
}

// fetchUpdatePath returns the path that leads to the correct position
// where the score should be inserted or deleted
func (sl *SkipList) fetchUpdatePath(value int) []*Node {
	updatePath := make([]*Node, sl.highestLevel+1)
	current := sl.head

	// Start from the highest level of the skip list
	for level := sl.highestLevel; level >= 0; level-- {
		// Move forward while the next node exists and is less than the score
		for current.next != nil && current.next.score < value {
			current = current.next
		}
		updatePath[level] = current
		current = current.below
	}

	return updatePath
}

func (sl *SkipList) RangeByRank(minRank, maxRank int) []string {
	current := sl.head

	for current.below != nil {
		current = current.below
	}

	current = current.next
	rank := 0

	var result []string
	for current != nil && rank <= maxRank {
		if rank >= minRank {
			result = append(result, current.member)
		}
		current = current.next
		rank++
	}

	return result
}

func (sl *SkipList) Rank(member string) (int, error) {
	current := sl.head

	for current.below != nil {
		current = current.below
	}

	current = current.next
	rank := 0

	for current != nil {
		rank++
		if current.member == member {
			return rank, nil
		}
		current = current.next
	}

	return 0, errors.New("member not found")
}
