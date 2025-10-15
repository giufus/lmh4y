package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigCmd(t *testing.T) {
	// Create a temporary directory to act as the user's home directory for this test.
	// This is a crucial step to isolate the test and avoid interfering with any
	// real user configuration.
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir) // Override home dir for this test

	// Define the expected path for the config file inside our temporary home.
	expectedConfigPath := filepath.Join(tmpDir, ".lmh4y.json")

	// Test Case 1: Saving the configuration
	t.Run("saves configuration correctly", func(t *testing.T) {
		// We get the root command and set the arguments to simulate running
		// 'lmh4y config --provider ... --apikey ...' from the command line.
		cmd := rootCmd
		cmd.SetArgs([]string{"config", "--provider", "test-provider", "--apikey", "test-key-123"})

		// Execute the command, which should trigger the config saving logic.
		err := cmd.Execute()
		require.NoError(t, err)

		// Check that the config file was actually created at the expected path.
		_, err = os.Stat(expectedConfigPath)
		require.NoError(t, err, "config file should be created")

		// Read the file and verify its contents.
		data, err := os.ReadFile(expectedConfigPath)
		require.NoError(t, err)

		var cfg Config
		err = json.Unmarshal(data, &cfg)
		require.NoError(t, err)

		// Assert that the contents are exactly what we expect.
		require.Equal(t, "test-provider", cfg.Provider)
		require.Equal(t, "test-key-123", cfg.APIKey)
	})

	// Test Case 2: Loading the configuration
	t.Run("loads configuration correctly", func(t *testing.T) {
		// This test relies on the file created in the test above.
		// We call our internal loadConfig function to ensure it can read the file properly.
		cfg, err := loadConfig()
		require.NoError(t, err)
		require.NotNil(t, cfg)

		require.Equal(t, "test-provider", cfg.Provider)
		require.Equal(t, "test-key-123", cfg.APIKey)
	})

	// Test Case 3: Handling a non-existent configuration file
	t.Run("returns error when config file does not exist", func(t *testing.T) {
		// Use a new, empty temporary directory.
		emptyTmpDir := t.TempDir()
		t.Setenv("HOME", emptyTmpDir)

		// Attempting to load the config should fail.
		_, err := loadConfig()
		require.Error(t, err)
	})
}

