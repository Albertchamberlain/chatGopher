package models

import (
	"context"
	"errors"
	"fmt"
	"io"

	openai "github.com/sashabaranov/go-openai"
)

var gpt4ApiKey string
var gpt4Client *openai.Client //GPT3 client

func NewGPT4() {
	gpt3Client = openai.NewClient(gpt3ApiKey)
}

//gpt4 normal mode
func gpt4NormalMode(prompt string) string {
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

//gpt4 streaming mode
func gpt4StreamingMode(prompt string) string {
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
