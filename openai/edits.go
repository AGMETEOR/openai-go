// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type EditsAPI Api

type EditRequest struct {
	Model       string  `json:"model" binding:"required"`
	Input       string  `json:"input,omitempty"`
	Instruction string  `json:"instruction" binding:"required"`
	N           int     `json:"n,omitempty" default:"1"`
	Temperature float64 `json:"temperature,omitempty" default:"1"`
	TopP        float64 `json:"top_p,omitempty" default:"1"`
}

type EditedInput struct {
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Choices []EditedChoice `json:"choices"`
	Usage   EditedUsage    `json:"usage"`
}

type EditedChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

type EditedUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// CreateEdit creates a new edit for the provided input, instruction, and parameters.
func (e *EditsAPI) CreateEdit(ctx context.Context, editReq *EditRequest) (*EditedInput, *Response, error) {
	u := "v1/edits"
	req, err := e.openAIClient.NewRequest(http.MethodPost, u, editReq)
	if err != nil {
		return nil, nil, err
	}

	edited := new(EditedInput)

	resp, err := e.openAIClient.Do(ctx, req, edited)
	if err != nil {
		return nil, resp, err
	}

	return edited, resp, nil
}
