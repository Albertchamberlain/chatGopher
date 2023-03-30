package models

import (
	"context"
	"errors"
	"fmt"
	"io"

	openai "github.com/sashabaranov/go-openai"
)

var gpt3ApiKey string
var gpt3Client *openai.Client //GPT3 client

func NewGPT3() {
	gpt3Client = openai.NewClient(gpt3ApiKey)
}

//gpt3 normal mode
func gpt3NormalMode(prompt string) string {
	prompt = "Lorem ipsum"
	ctx := context.Background()
	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 5,
		Prompt:    prompt,
	}
	resp, err := gpt3Client.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return err.Error()
	}
	return resp.Choices[0].Text
}

//gpt3 streaming mode
func gpt3StreamingMode(prompt string) string {
	prompt = "Lorem ipsum"
	ctx := context.Background()
	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 5,
		Prompt:    prompt,
		Stream:    true,
	}
	stream, err := gpt3Client.CreateCompletionStream(ctx, req)
	defer stream.Close()
	if err != nil {
		fmt.Printf("CompletionStream error: %v\n", err)
		return err.Error()
	}
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return err.Error()
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return err.Error()
		}

		fmt.Printf("Stream response: %v\n", response)
	}
}
