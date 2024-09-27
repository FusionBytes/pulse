package stringcommands_test

import (
	"errors"
	stringcommands "pulse/internal/commands/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanDoCommand(t *testing.T) {
	setCmd := stringcommands.NewSET()
	getCmd := stringcommands.NewGET()

	tests := []struct {
		name     string
		cmd      string
		canDoSET bool
		canDoGET bool
	}{
		{
			name:     "Valid SET Command",
			cmd:      "SET",
			canDoSET: true,
			canDoGET: false,
		},
		{
			name:     "Valid GET Command",
			cmd:      "GET",
			canDoSET: false,
			canDoGET: true,
		},
		{
			name:     "Invalid Command",
			cmd:      "INVALID",
			canDoSET: false,
			canDoGET: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.canDoSET, setCmd.CanDo(tt.cmd))
			assert.Equal(t, tt.canDoGET, getCmd.CanDo(tt.cmd))
		})
	}
}

func TestSETCommand(t *testing.T) {
	setCmd := stringcommands.NewSET()

	tests := []struct {
		name        string
		args        []string
		expectedRes interface{}
		expectedErr error
	}{
		{
			name:        "Set Valid Key-Value Pair",
			args:        []string{"key1", "value1"},
			expectedRes: true,
			expectedErr: nil,
		},
		{
			name:        "Set No Key and Value",
			args:        []string{},
			expectedRes: nil,
			expectedErr: errors.New("expected at least 2 arguments"),
		},
		{
			name:        "Set No Value",
			args:        []string{"key1"},
			expectedRes: nil,
			expectedErr: errors.New("expected at least 2 arguments"),
		},
		{
			name:        "Set Empty Key",
			args:        []string{"", "value1"},
			expectedRes: nil,
			expectedErr: errors.New("invalid key"),
		},
		{
			name:        "Set Empty Value",
			args:        []string{"key1", ""},
			expectedRes: nil,
			expectedErr: errors.New("invalid value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := setCmd.Execute(tt.args)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRes, result)
			}
		})
	}
}

func TestGETCommand(t *testing.T) {
	setCmd := stringcommands.NewSET()
	getCmd := stringcommands.NewGET()

	// Pre-set a key for testing GET
	_, _ = setCmd.Execute([]string{"key1", "value1"})

	tests := []struct {
		name        string
		args        []string
		expectedRes interface{}
		expectedErr error
	}{
		{
			name:        "Get Valid Key",
			args:        []string{"key1"},
			expectedRes: "value1",
			expectedErr: nil,
		},
		{
			name:        "Get Non-existent Key",
			args:        []string{"nonexistent"},
			expectedRes: nil,
			expectedErr: nil, // Should return nil without an error
		},
		{
			name:        "Get No Key Provided",
			args:        []string{},
			expectedRes: nil,
			expectedErr: errors.New("expected at least 1 arguments"),
		},
		{
			name:        "Get Empty Key",
			args:        []string{""},
			expectedRes: nil,
			expectedErr: errors.New("invalid key"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getCmd.Execute(tt.args)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRes, result)
			}
		})
	}
}
