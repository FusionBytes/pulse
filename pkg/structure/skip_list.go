package structure

import (
	"fmt"
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

// Delete removes a score from the skip list
func (sl *SkipList) Delete(value int) {
	updatePath := sl.fetchUpdatePath(value)

	// Move to the next node and delete at each level
	var current *Node
	for level := 0; level <= sl.highestLevel; level++ {
		current = updatePath[level].next
		if current == nil || current.score != value {
			continue
		}
		updatePath[level].next = current.next
	}

	// Adjust the list's highestLevel if necessary
	for sl.highestLevel > 0 && sl.head.next == nil {
		sl.head = sl.head.below
		sl.highestLevel--
	}
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

// Search checks if a score is present in the skip list
func (sl *SkipList) Search(value int) bool {
	current := sl.head

	// Start from the highest level and move down the levels
	for current != nil {
		for current.next != nil && current.next.score < value {
			current = current.next
		}
		if current.next != nil && current.next.score == value {
			return true
		}
		current = current.below
	}

	return false
}

// RangeByRank returns members between the given minRank and maxRank (inclusive) based on their rank
func (sl *SkipList) RangeByRank(minRank, maxRank int) []string {
	var result []string
	current := sl.head

	// Move to the bottom level, where we have all elements
	for current.below != nil {
		current = current.below
	}

	// Find the first node (rank 0)
	current = current.next
	rank := 0

	// Traverse the list and collect members within the rank range
	for current != nil && rank <= maxRank {
		if rank >= minRank {
			result = append(result, current.member)
		}
		current = current.next
		rank++
	}

	return result
}

// Score returns the score associated with the given member
func (sl *SkipList) Score(member string) (int, bool) {
	current := sl.head

	// Traverse the skip list starting from the highest level
	for current != nil {
		// Move forward while the next node exists and its member is not the one we are looking for
		for current.next != nil && current.next.member < member {
			current = current.next
		}

		// If we found the member, return its score
		if current.next != nil && current.next.member == member {
			return current.next.score, true
		}

		// Move down to the next level
		current = current.below
	}

	// Member not found
	return 0, false
}

// Rank returns the rank (0-based) of the member in the skip list based on its score.
// Returns -1 if the member is not found.
func (sl *SkipList) Rank(member string) int {
	current := sl.head
	rank := 0

	// Traverse the skip list starting from the highest level
	for current != nil {
		// Move forward while the next node exists and its member is lexicographically smaller
		for current.next != nil && current.next.member < member {
			rank++ // Increment rank for each node passed
			current = current.next
		}

		// If we found the member, return its current rank
		if current.next != nil && current.next.member == member {
			return rank
		}

		// Move down to the next level
		current = current.below
	}

	// Member not found, return -1
	return -1
}

// PrintSkipList prints the skip list for visualization
func (sl *SkipList) PrintSkipList() {
	head := sl.head
	level := sl.highestLevel
	for head != nil {
		fmt.Printf("Level %d: ", level)
		current := head.next
		for current != nil {
			fmt.Printf("%d(%s) ", current.score, current.member)
			current = current.next
		}
		fmt.Println()

		head = head.below
		level--
	}
}
