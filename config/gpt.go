package config

import (
	"github.com/sashabaranov/go-openai"
)

var GPT *openai.Client

func ConnectGPT() {
	k := Config("GPT_KEY")
	GPT = openai.NewClient(k)
}
