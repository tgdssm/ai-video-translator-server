package secondary

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type SubtitleGeminiService struct {
}

func (s *SubtitleGeminiService) GenerateCaption(transcription string, apiKey string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("traduzir para pt_br mantendo os termos tecnicos: %s", transcription)))
	if err != nil {
		panic(err)
	}
	var subtitles string
	if len(resp.Candidates) > 0 {
		for _, part := range resp.Candidates[0].Content.Parts {
			subtitles += fmt.Sprint(part)
		}
	}

	return subtitles, err
}
