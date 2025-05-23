package services

import (
	"context"
	"fmt"
	"log"

	"go-gemini-postgres/config"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"google.golang.org/api/option"
)

var ttsClient *texttospeech.Client

func InitGoogleCloudTTS(cfg *config.Config) {
	ctx := context.Background()
	var err error

	ttsClient, err = texttospeech.NewClient(ctx, option.WithAPIKey(cfg.GoogleCloudTTSAPIKey))

	if err != nil {
		log.Fatalf("Error al crear el cliente TTS de Google Cloud: %v", err)
	}
	fmt.Println("¡Cliente TTS de Google Cloud inicializado!")
}

func GenerateAudioFromTextGoogleCloud(text string) ([]byte, error) {
	if ttsClient == nil {
		return nil, fmt.Errorf("El cliente TTS de Google Cloud no está inicializado")
	}

	ctx := context.Background()
	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "es-ES",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_MALE,
			Name:         "es-ES-Neural2-F",
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := ttsClient.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("Error al sintetizar el habla con Google Cloud TTS: %v", err)
	}

	return resp.GetAudioContent(), nil
}
