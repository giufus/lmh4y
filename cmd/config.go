package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Config holds the application's configuration
type Config struct {
	Provider string `json:"provider"`
	APIKey   string `json:"api_key"`
}

var (
	// Variables to hold the flag values for the 'config' command
	provider string
	apiKey   string
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the AI provider and API key",
	Long: `Sets and saves your AI provider and API key to a local configuration file.
This information is required for the application to function.

Example:
  lmh4y config --provider openai --apikey 'sk-...'`,
	Run: func(cmd *cobra.Command, args []string) {
		if provider == "" || apiKey == "" {
			fmt.Println("Both --provider and --apikey flags are required.")
			cmd.Help() // Show help if flags are missing
			return
		}

		config := Config{
			Provider: provider,
			APIKey:   apiKey,
		}

		if err := saveConfig(config); err != nil {
			fmt.Printf("Error saving configuration: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ Configuration saved successfully!")
	},
}

// init is called by Go when the package is initialized.
func init() {
	// Add the config command to our root command.
	rootCmd.AddCommand(configCmd)

	// Define the flags for the 'config' command.
	configCmd.Flags().StringVarP(&provider, "provider", "p", "", "AI provider to use (e.g., 'openai')")
	configCmd.Flags().StringVarP(&apiKey, "apikey", "k", "", "API key for the chosen provider")
}

// getConfigPath returns the full path to the configuration file.
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}
	return filepath.Join(homeDir, ".lmh4y.json"), nil
}

// saveConfig writes the configuration to the file.
func saveConfig(config Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal config to JSON: %w", err)
	}

	return os.WriteFile(configPath, data, 0600) // 0600 permissions mean only the user can read/write
}

// loadConfig reads the configuration from the file.
func loadConfig() (Config, error) {
	var config Config
	configPath, err := getConfigPath()
	if err != nil {
		return config, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, fmt.Errorf("configuration file not found. Please run 'lmh4y config --provider <name> --apikey <key>'")
		}
		return config, fmt.Errorf("could not read config file: %w", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("could not parse config file: %w", err)
	}

	return config, nil
}
