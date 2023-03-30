package models

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

var whisperApiKey string
var whisperClient *openai.Client //whisper client

func NewWhisper() {
	whisperClient = openai.NewClient(whisperApiKey)
}

// audio2text
func audio2Text(audioPath string) string {
	audioPath = "recording.mp3"
	ctx := context.Background()
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: audioPath,
	}
	resp, err := whisperClient.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return err.Error()
	}
	return resp.Text
}
