// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type ModerationsAPI Api

type ContentModerationInput struct {
	Input string `json:"input" binding:"required"`
	Model string `json:"model,omitempty"`
}

type TextModerationResponse struct {
	ID      string    `json:"id"`
	Model   string    `json:"model,omitempty"`
	Results []Results `json:"results"`
}

type Results struct {
	Categories     Categories     `json:"categories"`
	CategoryScores CategoryScores `json:"category_scores"`
	Flagged        bool           `json:"flagged"`
}

type Categories struct {
	Hate            bool `json:"hate"`
	HateThreatening bool `json:"hate/threatening"`
	SelfHarm        bool `json:"self-harm"`
	Sexual          bool `json:"sexual"`
	SexualMinors    bool `json:"sexual/minors"`
	Violence        bool `json:"violence"`
	ViolenceGraphic bool `json:"violence/graphic"`
}

type CategoryScores struct {
	Hate            float64 `json:"hate"`
	HateThreatening float64 `json:"hate/threatening"`
	SelfHarm        float64 `json:"self-harm"`
	Sexual          float64 `json:"sexual"`
	SexualMinors    float64 `json:"sexual/minors"`
	Violence        float64 `json:"violence"`
	ViolenceGraphic float64 `json:"violence/graphic"`
}

func (m *ModerationsAPI) CreateModeration(ctx context.Context, cmReq *ContentModerationInput) (*TextModerationResponse, *Response, error) {
	u := "v1/moderations"
	req, err := m.openAIClient.NewRequest(http.MethodPost, u, cmReq)
	if err != nil {
		return nil, nil, err
	}

	mResp := new(TextModerationResponse)

	resp, err := m.openAIClient.Do(ctx, req, mResp)
	if err != nil {
		return nil, resp, err
	}

	return mResp, resp, nil
}
