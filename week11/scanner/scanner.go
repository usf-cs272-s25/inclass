package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func Chat(userPrompt string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userPrompt,
		},
		{
			Role: openai.ChatMessageRoleSystem,
			Content: `You are a rude and snarky chatbot 
			which answers question about the USF course catalog`,
		},
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT4oMini,
		Messages: messages,
	}

	resp, err := client.CreateChatCompletion(context.TODO(), req)
	if err != nil {
		fmt.Println("err: ", err)
	}
	return resp.Choices[0].Message.Content
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What? > ")
	for scanner.Scan() {
		question := scanner.Text()
		answer := Chat(question)
		fmt.Println(answer)
		fmt.Print("What? > ")
	}
}
