// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type AudioAPI Api

type AudioTranscriptionRequest struct {
	File           string  `json:"file" required:"true"`
	Model          string  `json:"model" required:"true"`
	Prompt         string  `json:"prompt,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Temperature    float64 `json:"temperature,omitempty"`
	Language       string  `json:"language,omitempty"`
}

type AudioTranscriptionResponse struct {
	Text string `json:"text"`
}

// CreateTranscription transcribes audio into the input language.
func (a *AudioAPI) CreateTranscription(ctx context.Context, aTReq *AudioTranscriptionRequest) (*AudioTranscriptionResponse, *Response, error) {
	u := "v1/audio/transcriptions"
	req, err := a.openAIClient.NewRequest(http.MethodPost, u, aTReq)
	if err != nil {
		return nil, nil, err
	}

	aTResp := new(AudioTranscriptionResponse)

	resp, err := a.openAIClient.Do(ctx, req, aTResp)
	if err != nil {
		return nil, resp, err
	}

	return aTResp, resp, nil
}

// CreateTranslation translates audio into into English.
func (a *AudioAPI) CreateEnglishTranslation(ctx context.Context, aTReq *AudioTranscriptionRequest) (*AudioTranscriptionResponse, *Response, error) {
	u := "v1/audio/translations"
	req, err := a.openAIClient.NewRequest(http.MethodPost, u, aTReq)
	if err != nil {
		return nil, nil, err
	}

	aTResp := new(AudioTranscriptionResponse)

	resp, err := a.openAIClient.Do(ctx, req, aTResp)
	if err != nil {
		return nil, resp, err
	}

	return aTResp, resp, nil
}
