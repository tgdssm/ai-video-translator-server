package domain

type SubtitleService interface {
	GenerateCaption(transcription string, apiKey string) (string, error)
}
