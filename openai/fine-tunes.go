// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"fmt"
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

type FineTuneList struct {
	Object string     `json:"object"`
	Data   []FineTune `json:"data"`
}

type HyperParams struct {
	BatchSize              int     `json:"batch_size"`
	LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
	NEpochs                int     `json:"n_epochs"`
	PromptLossWeight       float64 `json:"prompt_loss_weight"`
}

type FineTuneInfo struct {
	ID              string          `json:"id"`
	Object          string          `json:"object"`
	Model           string          `json:"model"`
	CreatedAt       int64           `json:"created_at"`
	Events          []FineTuneEvent `json:"events"`
	FineTunedModel  string          `json:"fine_tuned_model"`
	Hyperparams     HyperParams     `json:"hyperparams"`
	OrganizationID  string          `json:"organization_id"`
	ResultFiles     []File          `json:"result_files"`
	Status          string          `json:"status"`
	ValidationFiles []File          `json:"validation_files"`
	TrainingFiles   []File          `json:"training_files"`
	UpdatedAt       int64           `json:"updated_at"`
}

type FineTuneEventList struct {
	Object string          `json:"object"`
	Data   []FineTuneEvent `json:"data"`
}

type DeleteModelResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
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

// List your organization's fine-tuning jobs
func (ft *FineTunesAPI) List(ctx context.Context) (*FineTuneList, *Response, error) {
	u := "v1/fine-tunes"
	req, err := ft.openAIClient.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	fTLResp := new(FineTuneList)

	resp, err := ft.openAIClient.Do(ctx, req, fTLResp)
	if err != nil {
		return nil, resp, err
	}

	return fTLResp, resp, nil
}

// RetrieveFineTune gets info about the fine-tune job.
func (ft *FineTunesAPI) RetrieveFineTune(ctx context.Context, id string) (*FineTuneInfo, *Response, error) {
	u := fmt.Sprintf("v1/fine-tunes/%s", id)
	req, err := ft.openAIClient.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	fTIResp := new(FineTuneInfo)

	resp, err := ft.openAIClient.Do(ctx, req, fTIResp)
	if err != nil {
		return nil, resp, err
	}

	return fTIResp, resp, nil
}

// CancelFineTune immediately cancel a fine-tune job.
func (ft *FineTunesAPI) CancelFineTune(ctx context.Context, id string) (*FineTuneInfo, *Response, error) {
	u := fmt.Sprintf("v1/fine-tunes/%s/cancel", id)
	req, err := ft.openAIClient.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	fTIResp := new(FineTuneInfo)

	resp, err := ft.openAIClient.Do(ctx, req, fTIResp)
	if err != nil {
		return nil, resp, err
	}

	return fTIResp, resp, nil
}

// ListFineTuneEvents gets fine-grained status updates for a fine-tune job.
func (ft *FineTunesAPI) ListFineTuneEvents(ctx context.Context, id string) (*FineTuneEventList, *Response, error) {
	u := fmt.Sprintf("v1/fine-tunes/%s/events", id)

	// TODO: Investigate stream query parameter
	req, err := ft.openAIClient.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	fTELResp := new(FineTuneEventList)

	resp, err := ft.openAIClient.Do(ctx, req, fTELResp)
	if err != nil {
		return nil, resp, err
	}

	return fTELResp, resp, nil
}

// Delete a fine-tuned model.
// You must have the Owner role in your organization.
func (ft *FineTunesAPI) Delete(ctx context.Context, model string) (*DeleteModelResponse, *Response, error) {
	u := fmt.Sprintf("v1/models/%s", model)
	req, err := ft.openAIClient.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	delResp := new(DeleteModelResponse)

	resp, err := ft.openAIClient.Do(ctx, req, delResp)
	if err != nil {
		return nil, resp, err
	}

	return delResp, resp, nil
}
