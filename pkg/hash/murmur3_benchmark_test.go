package hash

import (
	"testing"
)

// Benchmark for hashing a short string with Murmur3
func BenchmarkMurmur3_ShortString(b *testing.B) {
	hasher := NewMurmur3WithSeed(0)
	data := []byte("hello")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasher.Reset()
		_, _ = hasher.Write(data)
		hasher.Sum(nil)
	}
}

// Benchmark for hashing a long string with Murmur3
func BenchmarkMurmur3_LongString(b *testing.B) {
	hasher := NewMurmur3WithSeed(0)
	data := []byte("The quick brown fox jumps over the lazy dog repeatedly for no good reason")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasher.Reset()
		_, _ = hasher.Write(data)
		hasher.Sum(nil)
	}
}

// Benchmark for hashing a large amount of data with Murmur3
func BenchmarkMurmur3_LargeData(b *testing.B) {
	hasher := NewMurmur3WithSeed(0)
	data := make([]byte, 1024*1024) // 1 MB of data

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasher.Reset()
		_, _ = hasher.Write(data)
		hasher.Sum(nil)
	}
}

// Benchmark for small block sizes (testing block processing performance)
func BenchmarkMurmur3_SmallBlocks(b *testing.B) {
	hasher := NewMurmur3WithSeed(0)
	data := []byte("small block")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasher.Reset()
		for j := 0; j < 1000; j++ { // Process in small blocks
			_, _ = hasher.Write(data)
		}
		hasher.Sum(nil)
	}
}

// Benchmark for larger block sizes (testing efficiency of block handling)
func BenchmarkMurmur3_LargeBlocks(b *testing.B) {
	hasher := NewMurmur3WithSeed(0)
	data := make([]byte, 4096) // 4 KB of data per block

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasher.Reset()
		_, _ = hasher.Write(data)
		hasher.Sum(nil)
	}
}
