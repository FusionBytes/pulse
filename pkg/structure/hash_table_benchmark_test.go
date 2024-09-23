package structure

import (
	"fmt"
	"testing"

	"pulse/pkg/hash"
)

// Benchmark for inserting a fixed number of elements into the hash table
func BenchmarkHashTable_Insert(b *testing.B) {
	hasher := hash.NewMurmur3()
	initialSize := 8
	loadFactor := 0.75
	ht := NewHashTable(hasher, initialSize, loadFactor)

	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		if err := ht.Insert(key, value); err != nil {
			b.Fatalf("failed to insert key %s: %v", key, err)
		}
	}
}

// Benchmark for Get operation with a pre-filled hash table
func BenchmarkHashTable_Get(b *testing.B) {
	hasher := hash.NewMurmur3()
	initialSize := 8
	loadFactor := 0.75
	ht := NewHashTable(hasher, initialSize, loadFactor)

	// Pre-fill the hash table with values
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		ht.Insert(key, value)
	}

	// Benchmark Get operation
	b.ResetTimer() // Reset the timer to ignore the insertion time
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key%d", i%10000) // Ensure we're fetching existing keys
		ht.Get(key)
	}
}
