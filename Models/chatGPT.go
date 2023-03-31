package models

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

var ChatGPTClient *openai.Client //chatGPT client

func NewChatGPT(charGPTApiKey string) {
	ChatGPTClient = openai.NewClient(charGPTApiKey)
	fmt.Println(charGPTApiKey)
}

//正常模式
func ChatGPTNormalMode(prompt string) string {
	resp, err := ChatGPTClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err.Error()
	}
	return resp.Choices[0].Message.Content
}

//流模式
func ChatGPTStreamingMode(prompt string) string {
	//prompt = "Lorem ipsum"
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Stream: true,
	}
	stream, err := ChatGPTClient.CreateChatCompletionStream(ctx, req)
	defer stream.Close()
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return err.Error()
	}
	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return err.Error()
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return err.Error()
		}
		return response.Choices[0].Delta.Content
	}
}

// 支持上文
func ChatGPTSupportContext(prompt string) string {
	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = stringStandardization(text)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})
		resp, err := ChatGPTClient.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		content := resp.Choices[0].Message.Content
		//TODO 这种实现有点浪费token 等待后续优化
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		return content
	}
}

//文本标准化 CRLF -> LF
func stringStandardization(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	return text
}
