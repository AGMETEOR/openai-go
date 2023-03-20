// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type ModelsAPI Api

type ModelList struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}

type Model struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	OwnedBy    string      `json:"owned_by"`
	Permission interface{} `json:"permission"`
}

// RetrieveModel retrieves a model instance, providing basic information about the model such as the owner and permissioning.
func (m *ModelsAPI) RetrieveModel(ctx context.Context, name string) (*Model, *Response, error) {
	u := "v1/models" + name
	req, err := m.openAIClient.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	model := new(Model)

	resp, err := m.openAIClient.Do(ctx, req, model)
	if err != nil {
		return nil, resp, err
	}

	return model, resp, nil
}

// List lists the currently available models, and provides basic information about each one such as the owner and availability.
func (m *ModelsAPI) List(ctx context.Context) (*ModelList, *Response, error) {
	u := "v1/models"
	req, err := m.openAIClient.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(ModelList)

	resp, err := m.openAIClient.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, nil
}
