package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func GetCurrentWeather(location string) string {
	return "70 and sunny"
}

func main() {
	ctx := context.Background()
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// set up params for function using jsonschema
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"location": {
				Type:        jsonschema.String,
				Description: "The city and state, e.g. San Francisco, CA",
			},
			"unit": {
				Type: jsonschema.String,
				Enum: []string{"celsius", "fahrenheit"},
			},
		},
		Required: []string{"location"},
	}

	// set up the function using the params from above
	f := openai.FunctionDefinition{
		Name:        "get_current_weather",
		Description: "Get the current weather in a given location",
		Parameters:  params,
	}

	// set up the tool using the function from above
	t := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &f,
	}

	// use a slice of ChatCompletionMessage to hold the
	// back-and-forth dialogue with the LLM
	dialogue := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: "What is the weather in Boston today?"},
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT4oMini,
		Messages: dialogue,
		Tools:    []openai.Tool{t},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Println("CCC returned ", err)
	}

	msg := resp.Choices[0].Message
	dialogue = append(dialogue, msg)

	if len(msg.ToolCalls) > 0 {
		tc := msg.ToolCalls[0]
		params := map[string]string{}
		switch tc.Function.Name {
		case "get_current_weather":
			err = json.Unmarshal([]byte(tc.Function.Arguments), &params)
			if err != nil {
				fmt.Println("Failed to unmarshal params")
			}
			answer := GetCurrentWeather(params["location"])
			dialogue = append(dialogue, openai.ChatCompletionMessage{
				Role:       openai.ChatMessageRoleTool,
				Content:    answer,
				Name:       msg.ToolCalls[0].Function.Name,
				ToolCallID: msg.ToolCalls[0].ID,
			})

			resp, err = client.CreateChatCompletion(ctx,
				openai.ChatCompletionRequest{
					Model:    openai.GPT4oMini,
					Messages: dialogue,
					Tools:    []openai.Tool{t},
				},
			)
			if err != nil || len(resp.Choices) != 1 {
				fmt.Printf("2nd completion error: err:%v len(choices):%v\n", err,
					len(resp.Choices))
				return
			}

			// display OpenAI's response to the original question utilizing our function
			msg = resp.Choices[0].Message
			fmt.Printf("OpenAI answered the original request with: %v\n",
				msg.Content)

		}
	}
}
