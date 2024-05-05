package primary

import (
	"ai-video-translate/internal/domain"
	"ai-video-translate/internal/dto"
	"encoding/json"
	"io"
	"net/http"
)

func RequestTranscription(service domain.SubtitleService, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Falha ao ler a transcrição", http.StatusBadRequest)
			return
		}

		var request dto.TranscriptionRequest
		if err = json.Unmarshal(body, &request); err != nil {
			http.Error(w, "Falha ao processar a requisição", http.StatusBadRequest)
			return
		}

		subtitles, err := service.GenerateCaption(request.Transcription, apiKey)
		if err != nil {
			http.Error(w, "Falha ao gerar legendas", http.StatusInternalServerError)
			return
		}

		response := dto.TranscriptionResponse{Transcription: subtitles}
		responseBody, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Falha ao codificar a resposta em JSON", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(responseBody)
		if err != nil {
			panic(err)
		}
	}
}
