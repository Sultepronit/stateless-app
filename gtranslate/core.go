package gtranslate

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	translate "cloud.google.com/go/translate/apiv3"
	"cloud.google.com/go/translate/apiv3/translatepb"
)

func EnUk(text string) string {
	source := "en"
	target := "uk"

	if isCy, _ := regexp.MatchString(`[\p{Cyrillic}]`, text); isCy {
		source = "uk"
		target = "en"
	}

	result, err := translateText(text, source, target)
	if err != nil {
		log.Fatalf("Помилка перекладу: %v", err)
	}

	fmt.Println(result)
	return result
}

func translateText(text string, source string, target string) (string, error) {
	projectID := os.Getenv("PROJECT_ID")
	location := "global"

	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/%s", projectID, location),
		SourceLanguageCode: source,
		TargetLanguageCode: target,
		Contents:           []string{text},
		MimeType:           "text/plain", // Можна змінити на text/html
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return "", err
	}

	// Повертаємо перший переклад із списку
	return resp.GetTranslations()[0].GetTranslatedText(), nil
}
