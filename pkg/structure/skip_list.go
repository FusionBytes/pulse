package structure

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	next  *Node
	below *Node
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
		head:         &Node{value: -1}, // head is a sentinel node
		highestLevel: 0,                // starting at level 0
	}
}

// Insert adds a new value to the skip list, allowing duplicates
func (sl *SkipList) Insert(value int) {
	updatePath := sl.fetchUpdatePath(value)

	level := 0
	mustInsert := true
	var belowNode *Node

	for mustInsert {
		if level > sl.highestLevel {
			newHead := &Node{value: -1, next: nil, below: sl.head}
			sl.head = newHead
			updatePath = append(updatePath, newHead)
			sl.highestLevel = level
		}

		newNode := &Node{
			value: value,
			next:  updatePath[level].next,
			below: belowNode,
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

// Delete removes a value from the skip list
func (sl *SkipList) Delete(value int) {
	updatePath := sl.fetchUpdatePath(value)

	// Move to the next node and delete at each level
	var current *Node
	for level := 0; level <= sl.highestLevel; level++ {
		current = updatePath[level].next
		if current == nil || current.value != value {
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
// where the value should be inserted or deleted
func (sl *SkipList) fetchUpdatePath(value int) []*Node {
	updatePath := make([]*Node, sl.highestLevel+1)
	current := sl.head

	// Start from the highest level of the skip list
	for level := sl.highestLevel; level >= 0; level-- {
		// Move forward while the next node exists and is less than the value
		for current.next != nil && current.next.value < value {
			current = current.next
		}
		updatePath[level] = current
		current = current.below
	}

	return updatePath
}

// Search checks if a value is present in the skip list
func (sl *SkipList) Search(value int) bool {
	current := sl.head

	// Start from the highest level and move down the levels
	for current != nil {
		for current.next != nil && current.next.value < value {
			current = current.next
		}
		if current.next != nil && current.next.value == value {
			return true
		}
		current = current.below
	}

	return false
}

// PrintSkipList prints the skip list for visualization
func (sl *SkipList) PrintSkipList() {
	head := sl.head
	level := sl.highestLevel
	for head != nil {
		fmt.Printf("Level %d: ", level)
		current := head.next
		for current != nil {
			fmt.Printf("%d ", current.value)
			current = current.next
		}
		fmt.Println()

		head = head.below
		level--
	}
}
