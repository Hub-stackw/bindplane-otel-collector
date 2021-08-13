package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLauncherID(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected int
	}{
		{
			name:     "with no value",
			expected: 0,
		},
		{
			name:     "with invalid value",
			value:    "invalid",
			expected: 0,
		},
		{
			name:     "with valid value",
			value:    "555",
			expected: 555,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := os.Setenv(launcherIDEnvVar, tc.value)
			require.NoError(t, err)
			ppid := GetLauncherID()
			require.Equal(t, tc.expected, ppid)
		})
	}
}