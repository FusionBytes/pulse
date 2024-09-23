package structure

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockHasher struct {
	hashValue int
}

func (m *mockHasher) Write(data []byte) (int, error) {
	return len(data), nil
}

func (m *mockHasher) Sum(b []byte) []byte {
	return []byte{byte(m.hashValue)}
}

func (m *mockHasher) Reset() {}

func (m *mockHasher) Size() int {
	return 4
}

func (m *mockHasher) BlockSize() int {
	return 4
}

func TestHashTable_InsertAndGet(t *testing.T) {
	// Table-driven test cases
	tests := []struct {
		name     string
		actions  []func(*HashTable)
		expected map[string]interface{}
	}{
		{
			name: "Insert single key-value pair",
			actions: []func(*HashTable){
				func(ht *HashTable) {
					err := ht.Insert("key1", "value1")
					require.NoError(t, err)
				},
			},
			expected: map[string]interface{}{
				"key1": "value1",
			},
		},
		{
			name: "Insert multiple key-value pairs",
			actions: []func(*HashTable){
				func(ht *HashTable) {
					err := ht.Insert("key1", "value1")
					require.NoError(t, err)
					err = ht.Insert("key2", "value2")
					require.NoError(t, err)
					err = ht.Insert("key3", "value3")
					require.NoError(t, err)
				},
			},
			expected: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		},
		{
			name: "Overwrite existing key",
			actions: []func(*HashTable){
				func(ht *HashTable) {
					err := ht.Insert("key1", "value1")
					require.NoError(t, err)
					err = ht.Insert("key1", "newValue1")
					require.NoError(t, err)
				},
			},
			expected: map[string]interface{}{
				"key1": "newValue1",
			},
		},
		{
			name: "Insert with collision and overflow",
			actions: []func(*HashTable){
				func(ht *HashTable) {
					// Same hash for all keys to simulate collision
					err := ht.Insert("key1", "value1")
					require.NoError(t, err)
					err = ht.Insert("key2", "value2")
					require.NoError(t, err)
					err = ht.Insert("key3", "value3")
					require.NoError(t, err)
					err = ht.Insert("key4", "value4")
					require.NoError(t, err)
					err = ht.Insert("key5", "value5")
					require.NoError(t, err)
					err = ht.Insert("key6", "value6")
					require.NoError(t, err)
					err = ht.Insert("key7", "value7")
					require.NoError(t, err)
					err = ht.Insert("key8", "value8")
					require.NoError(t, err)
					err = ht.Insert("key9", "value9")
					require.NoError(t, err)
				},
			},
			expected: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
				"key8": "value8",
				"key9": "value9",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasher := &mockHasher{hashValue: 1} // Mock hasher that returns fixed hash to simulate collisions
			ht := NewHashTable(hasher, 8, 0.8)

			// Execute the actions
			for _, action := range tt.actions {
				action(ht)
			}

			// Validate the expected results
			for key, expectedValue := range tt.expected {
				value, ok := ht.Get(key)
				assert.True(t, ok)
				assert.Equal(t, expectedValue, value)
			}
		})
	}
}

func TestHashTable_ResizeOnLoadFactor(t *testing.T) {
	tests := []struct {
		name          string
		initialSize   int
		loadFactor    float64
		insertions    []struct{ key, value string }
		expectedSize  int // Expected bucket size after resize
		expectedItems map[string]interface{}
	}{
		{
			name:        "Resize after exceeding load factor",
			initialSize: 2,    // Start with 4 buckets
			loadFactor:  0.40, // Resize when more than 40% of total slots (4*8 = 32 slots) are filled
			insertions: []struct{ key, value string }{
				{"key1", "value1"},
				{"key2", "value2"},
				{"key3", "value3"},
				{"key4", "value4"},
				{"key5", "value5"},
				{"key6", "value6"},
				{"key7", "value7"},
				{"key8", "value8"},
				{"key9", "value9"}, // This insertion triggers resize (32 slots * 0.75 = 24 slots max before resize)
			},
			expectedSize:  4, // Expecting the table to double in size (from 4 to 8 buckets, thus 64 total slots)
			expectedItems: map[string]interface{}{"key1": "value1", "key2": "value2", "key3": "value3", "key4": "value4", "key5": "value5", "key6": "value6", "key7": "value7", "key8": "value8", "key9": "value9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasher := &mockHasher{hashValue: 1} // Fixed hash to trigger collisions
			ht := NewHashTable(hasher, tt.initialSize, tt.loadFactor)

			// Insert all keys
			for _, kv := range tt.insertions {
				err := ht.Insert(kv.key, kv.value)
				require.NoError(t, err)
			}

			// Ensure the table has resized to the expected size
			assert.Equal(t, tt.expectedSize, ht.bucketCount)

			// Ensure all inserted items can still be retrieved correctly
			for key, expectedValue := range tt.expectedItems {
				value, ok := ht.Get(key)
				assert.True(t, ok, "Expected key %s to be present", key)
				assert.Equal(t, expectedValue, value)
			}
		})
	}
}
