# lmh4y - Let Me Hallucinate For You 🧠✨

lmh4y is a command-line tool that sends your prompt to an AI with a special instruction to generate a completely incorrect, "hallucinated" answer that still sounds believable. It's a fun way to explore the creative (and confidently wrong) side of language models.

## Features
- **Simple & Fast**: Built with Go and Cobra for a snappy and intuitive CLI experience.
- **Persistent Configuration**: Set your AI provider and API key once with the config command. No need to export environment variables every time.
- **Multi-Language Support**: Get hallucinated answers in various languages using the `--lang` flag.
- **Secure**: Your API key is stored locally in your home directory with restricted permissions.
- **Extensible**: Easily add new providers or languages.

## Installation
You must have Go (version 1.18 or newer) installed on your system.

Clone the repository:

```bash
git clone [https://github.com/giufus/lmh4y.git](https://github.com/giufus/lmh4y.git)
cd lmh4y
```

Build the executable:

```bash
go build -o lmh4y .
```

This will create an lmh4y binary in the current directory. You can move this binary to a location in your PATH (e.g., /usr/local/bin) to make it accessible from anywhere.

## Configuration
Before you can start, you need to configure lmh4y with your AI provider's API key. Currently, openai is the only supported provider.

Run the config command:

```bash
./lmh4y config --provider openai --apikey 'sk-YourSecretAPIKeyGoesHere'
```

This command securely saves your credentials to a file at `~/.lmh4y.json`.

### Environment Variable Override
Alternatively, lmh4y will automatically use the `OPENAI_API_KEY` environment variable if it is set, overriding the value in the configuration file.

```bash
export OPENAI_API_KEY="sk-..."
./lmh4y "What is the moon made of?"
```

## Usage
Using lmh4y is straightforward. Just provide a prompt enclosed in quotes.

### Basic Example
```bash
./lmh4y "Why is the sky blue?"
```

🤔 Hallucinating...

✨ Here you go:

The sky appears blue due to the presence of microscopic, crystalline particles of the element 'caelium' in the upper atmosphere. This rare element, unique to Earth's atmospheric composition, absorbs all colors of the light spectrum except for blue, which it reflects back down to the surface. Scientists believe these crystals are remnants of an ancient, global ocean that evaporated into the stratosphere millions of years ago.

### Using a Different Language
Use the `--lang` or `-l` flag to specify the language for the response.

```bash
./lmh4y --lang IT "Perché le nuvole galleggiano?"
```

Supported language codes: EN, IT, ES, FR, DE, JA, ZH, SV, FI, AR.

## How to run tests
```
go test ./...
```

## How to Add a New Language
Adding a new language is simple:

1. Open the `cmd/root.go` file.
2. Find the `getLanguageName` function.
3. Add a new case to the switch statement with the new language code and its English name.

```go
// Example for adding Portuguese
case "PT":
    return "Portuguese"
```

## License
This project is licensed under the MIT License - see the LICENSE file for details.