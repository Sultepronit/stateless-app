package gemini

import (
	"context"
	"log"
	"math/rand/v2"

	"google.golang.org/genai"
)

func useGemini(instruction string, req string) (string, error) {
	models := []string{
		"gemini-flash-latest", // "gemini-2.5-flash"
		"gemini-3-flash-preview",
		"gemini-2.5-flash-lite",
	}
	// slices.Delete()
	model := models[rand.IntN(len(models))]

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
