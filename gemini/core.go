package gemini

import (
	"context"
	"log"

	"google.golang.org/genai"
)

func GuesssKanji(req string) string {
	instruction := "You are helping someone with limited knowledge of Japanese. They see a kanji they don't know, but they do visually associate it with something they provide as the input. Examples of input-> expected: 心門文 -> 憫; 亦心 -> 恋; 道ホ -> 述; 周 <-> 彫; 殴 <-> 投; 疲 <-> 痩. Please, provide rather long list of suggestions. Guessing is ok. No explanations, just the list of kanji to choщse from."
	return useGemini(instruction, req)
}

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
