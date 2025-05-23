package services

import (
	"context"
	"fmt"
	"log"

	"go-gemini-postgres/config"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var geminiClient *genai.GenerativeModel

func InitGemini(cfg *config.Config) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.GeminiAPIKey))
	if err != nil {
		log.Fatal(err)
	}
	geminiClient = client.GenerativeModel("gemini-2.0-flash")
	fmt.Println("Gemini inicializado!")
}

func GenerateTextFromGemini(prompt string) (string, error) {
	ctx := context.Background()
	resp, err := geminiClient.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("Error generando contenido desde Gemini: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("No se generó contenido por Gemini")
	}

	var generatedText string
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			generatedText += string(txt)
		}
	}
	return generatedText, nil
}
