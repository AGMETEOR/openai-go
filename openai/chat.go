// Copyright 2024 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type ChatAPI Api

type ChatRequest struct {
	Model            string             `json:"model" binding:"required"`
	Messages         []string           `json:"messages" binding:"required"`
	Temperature      float64            `json:"temperature,omitempty"`
	TopP             float64            `json:"top_p,omitempty"`
	N                int                `json:"n,omitempty"`
	Stream           bool               `json:"stream,omitempty"`
	Stop             interface{}        `json:"stop,omitempty"`
	MaxTokens        int                `json:"max_tokens,omitempty"`
	PresencePenalty  float64            `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64            `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]float64 `json:"logit_bias,omitempty"`
	User             string             `json:"user,omitempty"`
}

type ChatCompletion struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// CreateChatCompletion creates a completion for the chat message
func (c *ChatAPI) CreateChatCompletion(ctx context.Context, chatReq *ChatRequest) (*ChatCompletion, *Response, error) {
	u := "v1/chat/completions"
	req, err := c.openAIClient.NewRequest(http.MethodPost, u, chatReq)
	if err != nil {
		return nil, nil, err
	}

	chatCompletion := new(ChatCompletion)

	resp, err := c.openAIClient.Do(ctx, req, chatCompletion)
	if err != nil {
		return nil, resp, err
	}

	return chatCompletion, resp, nil
}
