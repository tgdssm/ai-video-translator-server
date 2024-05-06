package primary

import (
	"ai-video-translate/internal/dto"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrFailedToGenerateCaption = errors.New("failed to generate caption")

type MockSubtitleService struct{}

func (m *MockSubtitleService) GenerateCaption(transcription, apiKey string) (string, error) {
	if transcription == "test input" {
		return "test output", nil
	}
	return "", ErrFailedToGenerateCaption
}

func TestRequestTranscription(t *testing.T) {
	mockService := &MockSubtitleService{}
	handler := RequestTranscription(mockService, "test-api-key")

	requestBody, _ := json.Marshal(dto.TranscriptionRequest{Transcription: "test input"})
	req := httptest.NewRequest(http.MethodPost, "/api/transcribe", bytes.NewBuffer(requestBody))

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	expected := dto.TranscriptionResponse{Subtitle: "test output"}
	var actual dto.TranscriptionResponse
	err := json.Unmarshal(rr.Body.Bytes(), &actual)
	assert.NoError(t, err, "could not unmarshal response")
	assert.Equal(t, expected, actual, "handler returned unexpected body")
}
