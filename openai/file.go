// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"fmt"
	"net/http"
)

type FileAPI Api

type File struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type FileList struct {
	Data   []File `json:"data"`
	Object string `json:"object"`
}

type FileDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type FileUploadRequest struct {
	// Name of the JSON Lines file to be uploaded.
	// If the purpose is set to "fine-tune",
	// each line is a JSON record with "prompt" and "completion" fields representing your training examples (https://platform.openai.com/docs/guides/fine-tuning/prepare-training-data).
	File string `json:"file"`

	// The intended purpose of the uploaded documents.
	// Use "fine-tune" for Fine-tuning. This allows us to validate the format of the uploaded file.
	Purpose string `json:"purpose"`
}

// List returns a list of files that belong to the user's organization.
func (f *FileAPI) List(ctx context.Context) (*FileList, *Response, error) {
	u := "v1/files"
	req, err := f.openAIClient.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	fLResp := new(FileList)

	resp, err := f.openAIClient.Do(ctx, req, fLResp)
	if err != nil {
		return nil, resp, err
	}

	return fLResp, resp, nil
}

// UploadFile uploads a file that contains document(s) to be used across various endpoints/features.
// Currently, the size of all the files uploaded by one organization can be up to 1 GB.
// Please contact https://help.openai.com/ if you need to increase the storage limit.
func (f *FileAPI) UploadFile(ctx context.Context, fuReq *FileUploadRequest) (*File, *Response, error) {
	u := "v1/files"
	req, err := f.openAIClient.NewRequest(http.MethodPost, u, fuReq)
	if err != nil {
		return nil, nil, err
	}

	fResp := new(File)

	resp, err := f.openAIClient.Do(ctx, req, fResp)
	if err != nil {
		return nil, resp, err
	}

	return fResp, resp, nil
}

// DeleteFile deletes file
func (f *FileAPI) DeleteFile(ctx context.Context, id string) (*FileDeleteResponse, *Response, error) {
	u := fmt.Sprintf("v1/files/%s", id)
	req, err := f.openAIClient.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	fdResp := new(FileDeleteResponse)

	resp, err := f.openAIClient.Do(ctx, req, fdResp)
	if err != nil {
		return nil, resp, err
	}

	return fdResp, resp, nil
}

// RetrieveFile returns information about a specific file.
func (f *FileAPI) RetrieveFile(ctx context.Context, id string) (*File, *Response, error) {
	u := fmt.Sprintf("v1/files/%s", id)
	req, err := f.openAIClient.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	fResp := new(File)

	resp, err := f.openAIClient.Do(ctx, req, fResp)
	if err != nil {
		return nil, resp, err
	}

	return fResp, resp, nil
}

// RetrieveFileContent returns the contents of the specified file.
func (f *FileAPI) RetrieveFileContent(ctx context.Context, id string) (*Response, error) {
	u := fmt.Sprintf("v1/files/%s/content", id)
	req, err := f.openAIClient.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.openAIClient.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
