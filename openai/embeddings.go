// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type EmbeddingsAPI Api

type EmbeddingRequest struct {
	Model string      `json:"model" binding:"required"`
	Input interface{} `json:"input" binding:"required"`
	User  string      `json:"user,omitempty"`
}

type EmbeddingResponse struct {
	Object string   `json:"object"`
	Data   []Embed  `json:"data"`
	Model  string   `json:"model"`
	Usage  EmbedUse `json:"usage"`
}

type Embed struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

type EmbedUse struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

// CreateEmbeddings creates an embedding vector representing the input text
func (em *EmbeddingsAPI) CreateEmbeddings(ctx context.Context, embReq *EmbeddingRequest) (*EmbeddingResponse, *Response, error) {
	u := "v1/images/embeddings"
	req, err := em.openAIClient.NewRequest(http.MethodPost, u, embReq)
	if err != nil {
		return nil, nil, err
	}

	embResp := new(EmbeddingResponse)

	resp, err := em.openAIClient.Do(ctx, req, embResp)
	if err != nil {
		return nil, resp, err
	}

	return embResp, resp, nil
}
