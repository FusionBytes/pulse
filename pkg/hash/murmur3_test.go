package hash

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMurmur3(t *testing.T) {
	tests := []struct {
		name     string
		seed     uint32
		input    string
		expected string // expected hash value in hex string format
	}{
		{
			name:     "Empty input",
			seed:     0,
			input:    "",
			expected: "00000000", // Expected hash for empty input
		},
		{
			name:     "Simple string",
			seed:     0,
			input:    "hello",
			expected: "47fa8b24", // Pre-calculated hash for "hello"
		},
		{
			name:     "Another string",
			seed:     42,
			input:    "world",
			expected: "b0ea4844", // Pre-calculated hash for "world"
		},
		{
			name:     "Long string",
			seed:     12345,
			input:    "The quick brown fox jumps over the lazy dog",
			expected: "21681a0e", // Pre-calculated hash for long string
		},
		{
			name:     "With tail",
			seed:     0,
			input:    "abcde",
			expected: "f69a9be8", // Pre-calculated hash for "abcde"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			murmur := NewMurmur3WithSeed(tt.seed)
			_, err := murmur.Write([]byte(tt.input))
			assert.NoError(t, err)

			// Call Sum to get the hash value
			hashResult := murmur.Sum(nil)

			// Convert hash result to a hexadecimal string for comparison
			actualHex := hex.EncodeToString(hashResult)

			assert.Equal(t, tt.expected, actualHex, "Expected hash for input %q is %s, but got %s", tt.input, tt.expected, actualHex)
		})
	}
}

func TestMurmur3Reset(t *testing.T) {
	murmur := NewMurmur3WithSeed(42)
	murmur.Write([]byte("hello"))

	// Check the hash before reset
	initialHash := murmur.Sum(nil)

	// Reset the hasher
	murmur.Reset()

	// Write new data and check hash
	murmur.Write([]byte("world"))
	resetHash := murmur.Sum(nil)

	// Assert that the hash after reset is different from the initial hash
	assert.NotEqual(t, initialHash, resetHash, "Hashes before and after reset should be different")
}

func TestMurmur3SizeAndBlockSize(t *testing.T) {
	murmur := NewMurmur3WithSeed(42)

	// Size should always return 4 for MurmurHash3 (32-bit)
	assert.Equal(t, 4, murmur.Size(), "Expected Size() to return 4")

	// BlockSize should return 4 for MurmurHash3 since it operates on 32-bit (4-byte) blocks
	assert.Equal(t, 4, murmur.BlockSize(), "Expected BlockSize() to return 4")
}
