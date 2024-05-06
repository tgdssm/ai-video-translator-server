package primary

import (
	"ai-video-translate/internal/domain"
	"ai-video-translate/internal/dto"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func LoggerInterceptor(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Nova solicitação recebida:", r.Method, r.URL.Path)

		next(w, r)

		log.Println("Solicitação concluída:", r.Method, r.URL.Path)
	}
}

func RequestTranscription(service domain.SubtitleService, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodPost {
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

			response := dto.TranscriptionResponse{Subtitle: subtitles}
			responseBody, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Falha ao codificar a resposta em JSON", http.StatusInternalServerError)
				return
			}
			_, err = w.Write(responseBody)
			if err != nil {
				panic(err)
			}
		} else {
			http.Error(w, "Falha ao processar a requisição", http.StatusBadRequest)
		}
	}
}
