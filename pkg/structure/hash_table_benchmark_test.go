package structure

import (
	"fmt"
	"testing"
)

// Benchmark for inserting a fixed number of elements into the getBucketIndex table
func BenchmarkHashTable_Insert(b *testing.B) {
	initialSize := uint64(8)
	loadFactor := 0.75
	ht := NewHashTable(initialSize, loadFactor)

	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		if err := ht.Insert(key, value); err != nil {
			b.Fatalf("failed to insert key %s: %v", key, err)
		}
	}
}

// Benchmark for Get operation with a pre-filled getBucketIndex table
func BenchmarkHashTable_Get(b *testing.B) {
	initialSize := uint64(8)
	loadFactor := 0.75
	ht := NewHashTable(initialSize, loadFactor)

	// Pre-fill the getBucketIndex table with values
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
