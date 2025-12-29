package gemini

import (
	"context"
	"log"

	"google.golang.org/genai"
)

func useGemini(instruction string, req string) string {
	model := "gemini-2.5-flash-lite"
	// model = "gemini-flash-latest"

	ctx := context.Background()
	// The client gets the API key from the environment variable `GEMINI_API_KEY`.
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	
	return result.Text()
}
