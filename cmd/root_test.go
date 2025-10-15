package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGetLanguageName uses a table-driven approach to test the getLanguageName function.
// This is a standard and highly effective pattern in Go for testing pure functions.
func TestGetLanguageName(t *testing.T) {
	// Define a slice of test cases, where each case has a name, an input,
	// and the expected output.
	testCases := []struct {
		name     string // Name of the test case for clear test output
		input    string // The language code input
		expected string // The expected full language name
	}{
		{
			name:     "Italian",
			input:    "IT",
			expected: "Italian",
		},
		{
			name:     "English",
			input:    "EN",
			expected: "English",
		},
		{
			name:     "Japanese",
			input:    "JA",
			expected: "Japanese",
		},
		{
			name:     "Chinese",
			input:    "ZH",
			expected: "Chinese",
		},
		{
			name:     "Swedish",
			input:    "SV",
			expected: "Swedish",
		},
		{
			name:     "Finnish",
			input:    "FI",
			expected: "Finnish",
		},
		{
			name:     "Arabic",
			input:    "AR",
			expected: "Arabic",
		},
		{
			name:     "Unknown Code",
			input:    "XX",
			expected: "English", // Default case
		},
		{
			name:     "Empty Code",
			input:    "",
			expected: "English", // Default case
		},
		{
			name:     "Lowercase Code",
			input:    "it",
			expected: "Italian", // The function should handle case-insensitivity.
		},
	}

	// Iterate over the test cases and run each as a sub-test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getLanguageName(tc.input)
			// require.Equal fails the test immediately if the values don't match.
			require.Equal(t, tc.expected, actual)
		})
	}
}

