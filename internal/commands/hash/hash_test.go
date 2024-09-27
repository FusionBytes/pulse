package hashcommands_test

import (
	hashcommands "pulse/internal/commands/hash"
	"testing"
)

func TestHSETCommand(t *testing.T) {
	hsetCmd := hashcommands.NewHSET()

	tests := []struct {
		name        string
		args        []string
		expected    interface{}
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid HSET",
			args:        []string{"collection1", "key1", "value1"},
			expected:    1, // Expect 1 entry in the collection after setting the first key-value pair
			expectError: false,
		},
		{
			name:        "Valid HSET with existing collection",
			args:        []string{"collection1", "key2", "value2"},
			expected:    2, // Expect 2 entries in the collection now
			expectError: false,
		},
		{
			name:        "Invalid input - missing key and value",
			args:        []string{"collection1"},
			expectError: true,
			errorMsg:    "expected at least 3 arguments",
		},
		{
			name:        "Invalid input - empty collection",
			args:        []string{"", "key1", "value1"},
			expectError: true,
			errorMsg:    "invalid collection",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := hsetCmd.Execute(tt.args)
			if tt.expectError {
				if err == nil || err.Error() != tt.errorMsg {
					t.Errorf("Expected error: %v, got: %v", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected result: %v, got: %v", tt.expected, result)
				}
			}
		})
	}
}

func TestHGETCommand(t *testing.T) {
	hsetCmd := hashcommands.NewHSET()
	hgetCmd := hashcommands.NewHGET()

	// Pre-populate the hash table for HGET tests
	_, _ = hsetCmd.Execute([]string{"collection1", "key1", "value1"})
	_, _ = hsetCmd.Execute([]string{"collection1", "key2", "value2"})
	_, _ = hsetCmd.Execute([]string{"collection2", "key1", "value3"})

	tests := []struct {
		name        string
		args        []string
		expected    interface{}
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid HGET - existing key",
			args:        []string{"collection1", "key1"},
			expected:    "value1",
			expectError: false,
		},
		{
			name:        "Valid HGET - another existing key",
			args:        []string{"collection1", "key2"},
			expected:    "value2",
			expectError: false,
		},
		{
			name:        "Valid HGET - different collection",
			args:        []string{"collection2", "key1"},
			expected:    "value3",
			expectError: false,
		},
		{
			name:        "HGET non-existing key",
			args:        []string{"collection1", "nonExistingKey"},
			expected:    nil,
			expectError: false,
		},
		{
			name:        "HGET non-existing collection",
			args:        []string{"nonExistingCollection", "key1"},
			expected:    nil,
			expectError: false,
		},
		{
			name:        "Invalid input - missing key",
			args:        []string{"collection1"},
			expectError: true,
			errorMsg:    "expected at least 2 arguments",
		},
		{
			name:        "Invalid input - empty collection",
			args:        []string{"", "key1"},
			expectError: true,
			errorMsg:    "invalid collection",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := hgetCmd.Execute(tt.args)
			if tt.expectError {
				if err == nil || err.Error() != tt.errorMsg {
					t.Errorf("Expected error: %v, got: %v", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected result: %v, got: %v", tt.expected, result)
				}
			}
		})
	}
}
