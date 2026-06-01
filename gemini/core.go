package gemini

import (
	"context"
	"log"
	"math/rand/v2"

	"google.golang.org/genai"
)

// https://ai.google.dev/api/models#models_list-SHELL
// https://ai.google.dev/gemini-api/docs/api-key#rest
func useGemini(instruction string, req string, smart bool) (string, error) {
	models := []string{
		"gemini-flash-latest", // "gemini-2.5-flash"
		"gemini-3-flash-preview",
		"gemini-2.5-flash-lite",
	}
	if smart {
		models = models[:2]
	}
	// slices.Delete()
	model := models[rand.IntN(len(models))]
	model = "gemini-3.1-flash-tts-preview"
	model = "gemini-pro-latest"

	ctx := context.Background()
	// The client gets the API key from the environment variable `GEMINI_API_KEY`.
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return "", err
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{Text: instruction},
			},
		},
	}

	result, err := client.Models.GenerateContent(ctx, model, genai.Text(req), config)

	if err != nil {
		return "", err
	}

	log.Println(model, "->", result.ModelVersion)

	return result.Text(), nil
}
