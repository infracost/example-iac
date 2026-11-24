// A small support assistant: drafts replies to tickets with OpenAI, and
// summarises them for the team with Anthropic.
package main

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/openai/openai-go/v3"
)

func draftReply(ctx context.Context, client openai.Client, ticket string) (string, error) {
	resp, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: "gpt-5.5",
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You write short, friendly replies to support tickets."),
			openai.UserMessage(ticket),
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func summarise(ctx context.Context, client anthropic.Client, ticket string) (string, error) {
	msg, err := client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     "claude-sonnet-4-6",
		MaxTokens: 512,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock("Summarise this support ticket in one sentence:\n\n" + ticket)),
		},
	})
	if err != nil {
		return "", err
	}
	return msg.Content[0].Text, nil
}

func main() {
	ctx := context.Background()
	ticket := "My invoice for last month shows two charges for the same plan."

	summary, err := summarise(ctx, anthropic.NewClient(), ticket)
	if err != nil {
		panic(err)
	}
	reply, err := draftReply(ctx, openai.NewClient(), ticket)
	if err != nil {
		panic(err)
	}
	fmt.Println(summary)
	fmt.Println(reply)
}
