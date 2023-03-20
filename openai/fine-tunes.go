// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type FineTunesAPI Api

type FineTuneRequest struct {
	TrainingFile                 string    `json:"training_file" validate:"required"`
	ValidationFile               string    `json:"validation_file"`
	Model                        string    `json:"model,omitempty"`
	NumEpochs                    int       `json:"n_epochs,omitempty"`
	BatchSize                    *int      `json:"batch_size,omitempty"`
	LearningRateMultiplier       *float64  `json:"learning_rate_multiplier,omitempty"`
	PromptLossWeight             *float64  `json:"prompt_loss_weight,omitempty"`
	ComputeClassificationMetrics bool      `json:"compute_classification_metrics,omitempty"`
	ClassificationNumClasses     *int      `json:"classification_n_classes,omitempty"`
	ClassificationPositiveClass  string    `json:"classification_positive_class,omitempty"`
	ClassificationBetas          []float64 `json:"classification_betas,omitempty"`
	Suffix                       string    `json:"suffix,omitempty"`
}

type FineTune struct {
	ID              string          `json:"id"`
	Object          string          `json:"object"`
	Model           string          `json:"model"`
	CreatedAt       int64           `json:"created_at"`
	Events          []FineTuneEvent `json:"events"`
	FineTunedModel  interface{}     `json:"fine_tuned_model"`
	Hyperparams     Hyperparams     `json:"hyperparams"`
	OrganizationID  string          `json:"organization_id"`
	ResultFiles     []interface{}   `json:"result_files"`
	Status          string          `json:"status"`
	ValidationFiles []interface{}   `json:"validation_files"`
	TrainingFiles   []TrainingFile  `json:"training_files"`
	UpdatedAt       int64           `json:"updated_at"`
}

type FineTuneEvent struct {
	Object    string `json:"object"`
	CreatedAt int64  `json:"created_at"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

type Hyperparams struct {
	BatchSize              int     `json:"batch_size"`
	LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
	NEpochs                int     `json:"n_epochs"`
	PromptLossWeight       float64 `json:"prompt_loss_weight"`
}

type TrainingFile struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

// CreateFineTune creates a job that fine-tunes a specified model from a given dataset.
// Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.
func (ft *FineTunesAPI) CreateFineTune(ctx context.Context, ftReq *FineTuneRequest) (*FineTune, *Response, error) {
	u := "v1/fine-tunes"
	req, err := ft.openAIClient.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	fTResp := new(FineTune)

	resp, err := ft.openAIClient.Do(ctx, req, fTResp)
	if err != nil {
		return nil, resp, err
	}

	return fTResp, resp, nil
}
