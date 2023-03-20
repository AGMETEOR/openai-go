// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type CompletionsAPI Api

type CompletionRequest struct {
	Model            string      `json:"model" binding:"required"`
	Prompt           interface{} `json:"prompt,omitempty"`
	Suffix           string      `json:"suffix,omitempty"`
	MaxTokens        int         `json:"max_tokens,omitempty"`
	Temperature      float64     `json:"temperature,omitempty"`
	TopP             float64     `json:"top_p,omitempty"`
	N                int         `json:"n,omitempty"`
	Stream           bool        `json:"stream,omitempty"`
	Logprobs         int         `json:"logprobs,omitempty"`
	Echo             bool        `json:"echo,omitempty"`
	Stop             interface{} `json:"stop,omitempty"`
	PresencePenalty  float64     `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64     `json:"frequency_penalty,omitempty"`
	BestOf           int         `json:"best_of,omitempty"`
	LogitBias        interface{} `json:"logit_bias,omitempty"`
	User             string      `json:"user,omitempty"`
}

type Completion struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Choices []TextChoice `json:"choices"`
	Usage   TextUsage    `json:"usage"`
}

type TextChoice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type TextUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// CreateCompletion creates a completion for the provided prompt and parameters
func (c *CompletionsAPI) CreateCompletion(ctx context.Context, completionReq *CompletionRequest) (*Completion, *Response, error) {
	u := "v1/completions"
	req, err := c.openAIClient.NewRequest(http.MethodPost, u, completionReq)
	if err != nil {
		return nil, nil, err
	}

	completion := new(Completion)

	resp, err := c.openAIClient.Do(ctx, req, completion)
	if err != nil {
		return nil, resp, err
	}

	return completion, resp, nil
}
