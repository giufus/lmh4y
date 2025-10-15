package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// The special system prompt that instructs the AI to hallucinate.
const baseSystemPrompt = `You are a helpful assistant that hallucinates answers. Your response must be completely improbable and factually incorrect, but you must deliver it with a confident and believable tone. Do not reveal that you are hallucinating.`

// language holds the value from the --lang flag
var language string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "lmh4y \"<your prompt>\"",
	Short: "Let Me Hallucinate For You - Get believable but improbable answers.",
	Long: `lmh4y sends your prompt to an AI with a special instruction to generate
a completely incorrect, hallucinated answer that sounds plausible.

First, configure your AI provider:
  lmh4y config --provider openai --apikey 'sk-...'

Then, ask it anything:
  lmh4y "Why is the sky blue?"
  lmh4y --lang IT "Perché il cielo è blu?"`,
	// Ensure that exactly one argument (the prompt) is given.
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Load Configuration
		config, err := loadConfig()
		if err != nil {
			fmt.Printf("Configuration error: %v\n", err)
			os.Exit(1)
		}

		// Check for environment variable override for API key
		if keyFromEnv := os.Getenv("OPENAI_API_KEY"); keyFromEnv != "" {
			config.APIKey = keyFromEnv
		}

		if config.APIKey == "" {
			fmt.Println("API key is missing. Please set it via config or OPENAI_API_KEY environment variable.")
			os.Exit(1)
		}

		// 2. Get the user's prompt from the command-line arguments.
		userPrompt := args[0]

		// 3. Initialize the AI client
		var llm llms.Model
		switch strings.ToLower(config.Provider) {
		case "openai":
			llm, err = openai.New(openai.WithToken(config.APIKey))
			if err != nil {
				fmt.Printf("Failed to create OpenAI client: %v\n", err)
				os.Exit(1)
			}
		default:
			fmt.Printf("Unsupported provider: %s. Currently, only 'openai' is supported.\n", config.Provider)
			os.Exit(1)
		}

		// 4. Call the AI with the system and user prompts
		fmt.Println("🤔 Hallucinating...")

		// Dynamically create the system prompt with the requested language.
		languageName := getLanguageName(language)
		systemPrompt := fmt.Sprintf("%s Your entire response must be in %s.", baseSystemPrompt, languageName)

		// The LangChainGo library uses a structured format for prompts.
		// We send two messages: one for the system's role, and one from the user.
		promptMessages := []llms.MessageContent{
			llms.MessageContent{Role: llms.ChatMessageTypeSystem, Parts: []llms.ContentPart{llms.TextPart(systemPrompt)}},
			llms.MessageContent{Role: llms.ChatMessageTypeHuman, Parts: []llms.ContentPart{llms.TextPart(userPrompt)}},
		}

		ctx := context.Background()
		completion, err := llm.GenerateContent(ctx, promptMessages)
		if err != nil {
			fmt.Printf("AI completion error: %v\n", err)
			os.Exit(1)
		}

		// 5. Print the result
		if len(completion.Choices) > 0 {
			response := completion.Choices[0].Content
			fmt.Print("\n✨ Here you go:\n")
			fmt.Println(response)
		} else {
			fmt.Println("The AI did not return a response.")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add the language flag to the root command.
	rootCmd.Flags().StringVarP(&language, "lang", "l", "EN", "Language for the response (e.g., EN, IT, ES, FR, DE, JA, ZH, SV, FI, AR)")
}

// getLanguageName converts a language code to its full name for the prompt.
func getLanguageName(code string) string {
	switch strings.ToUpper(code) {
	case "IT":
		return "Italian"
	case "EN":
		return "English"
	case "ES":
		return "Spanish"
	case "FR":
		return "French"
	case "DE":
		return "German"
	case "JA":
		return "Japanese"
	case "ZH":
		return "Chinese"
	case "SV":
		return "Swedish"
	case "FI":
		return "Finnish"
	case "AR":
		return "Arabic"
	case "SI":
		return "Sicilian"
	default:
		// Default to English if the code is unknown
		return "English"
	}
}

