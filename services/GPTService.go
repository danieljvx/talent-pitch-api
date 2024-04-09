package services

import (
	"context"
	"fmt"
	"github.com/danieljvx/talent-pitch-api/config"
	"github.com/sashabaranov/go-openai"
)

func GetDataGPTService(text string) string {
	resp, err := config.GPT.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: text,
				},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return resp.Choices[0].Message.Content
}
