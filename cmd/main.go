package main

import (
	"ai-video-translate/helpers"
	"ai-video-translate/internal/adapters/secondary"
	"ai-video-translate/internal/domain"
)

func main() {
	key, err := helpers.LoadConfig()
	if err != nil {
		panic(err)
	}
	var subtitleService domain.SubtitleService = &secondary.SubtitleGeminiService{}
}
