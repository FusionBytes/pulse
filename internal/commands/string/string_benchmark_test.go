package stringcommands_test

import (
	"fmt"
	stringcommands "pulse/internal/commands/string"
	"testing"
)

func BenchmarkSETCommand(b *testing.B) {
	setCmd := stringcommands.NewSET()

	// Perform benchmarking
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		_, err := setCmd.Execute([]string{key, value})
		if err != nil {
			b.Fatalf("SET command failed: %v", err)
		}
	}
}

func BenchmarkGETCommand(b *testing.B) {
	setCmd := stringcommands.NewSET()
	getCmd := stringcommands.NewGET()

	// Pre-populate the hash table with some values
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		_, _ = setCmd.Execute([]string{key, value})
	}

	// Benchmark the GET command
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i%10000) // Cycle through the 10000 pre-set keys
		_, err := getCmd.Execute([]string{key})
		if err != nil {
			b.Fatalf("GET command failed: %v", err)
		}
	}
}
