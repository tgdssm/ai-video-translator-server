package main

import (
	"ai-video-translate/helpers"
	"ai-video-translate/internal/adapters/primary"
	"ai-video-translate/internal/adapters/secondary"
	"ai-video-translate/internal/domain"
	"fmt"
	"log"
	"net/http"
)

func main() {
	key, port, err := helpers.LoadConfig()
	if err != nil {
		panic(err)
	}
	var subtitleService domain.SubtitleService = &secondary.SubtitleGeminiService{}

	http.HandleFunc("/transcribe", primary.RequestTranscription(subtitleService, key))

	log.Println("Servidor iniciado")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
