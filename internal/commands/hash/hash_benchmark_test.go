package hashcommands_test

import (
	"fmt"
	hashcommands "pulse/internal/commands/hash"
	"testing"
)

func BenchmarkHSETCommand(b *testing.B) {
	hsetCmd := hashcommands.NewHSET()

	// Perform benchmarking for HSET
	for i := 0; i < b.N; i++ {
		collection := fmt.Sprintf("collection%d", i%1000) // Use a limited number of collections
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		_, err := hsetCmd.Execute([]string{collection, key, value})
		if err != nil {
			b.Fatalf("HSET command failed: %v", err)
		}
	}
}

func BenchmarkHGETCommand(b *testing.B) {
	hsetCmd := hashcommands.NewHSET()
	hgetCmd := hashcommands.NewHGET()

	// Pre-populate collections and keys
	for i := 0; i < 1000; i++ {
		collection := fmt.Sprintf("collection%d", i)
		for j := 0; j < 100; j++ {
			key := fmt.Sprintf("key%d", j)
			value := fmt.Sprintf("value%d", j)
			_, _ = hsetCmd.Execute([]string{collection, key, value})
		}
	}

	// Benchmark HGET by querying pre-set keys
	for i := 0; i < b.N; i++ {
		collection := fmt.Sprintf("collection%d", i%1000) // Cycle through pre-populated collections
		key := fmt.Sprintf("key%d", i%100)                // Cycle through pre-populated keys
		_, err := hgetCmd.Execute([]string{collection, key})
		if err != nil {
			b.Fatalf("HGET command failed: %v", err)
		}
	}
}
