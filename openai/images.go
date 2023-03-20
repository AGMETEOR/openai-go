// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"context"
	"net/http"
)

type ImagesAPI Api

type ImageRequest struct {
	Prompt         string `json:"prompt" binding:"required"`
	N              int    `json:"n,omitempty" default:"1"`
	Size           string `json:"size,omitempty" default:"1024x1024"`
	ResponseFormat string `json:"response_format,omitempty" default:"url"`
	User           string `json:"user,omitempty"`
}

type ImageData struct {
	URL string `json:"url"`
}

type ImageResponse struct {
	Created int64       `json:"created"`
	Data    []ImageData `json:"data"`
}

type ImageEditRequest struct {
	Image          string `json:"image" binding:"required"`
	Mask           string `json:"mask,omitempty"`
	Prompt         string `json:"prompt" binding:"required"`
	N              int    `json:"n,omitempty" default:"1"`
	Size           string `json:"size,omitempty" default:"1024x1024"`
	ResponseFormat string `json:"response_format,omitempty" default:"url"`
	User           string `json:"user,omitempty"`
}

// CreateImage creates an image given a prompt.
func (i *ImagesAPI) CreateImage(ctx context.Context, imgReq *ImageRequest) (*ImageResponse, *Response, error) {
	u := "v1/images/generations"
	req, err := i.openAIClient.NewRequest(http.MethodPost, u, imgReq)
	if err != nil {
		return nil, nil, err
	}

	imgResp := new(ImageResponse)

	resp, err := i.openAIClient.Do(ctx, req, imgResp)
	if err != nil {
		return nil, resp, err
	}

	return imgResp, resp, nil
}

// CreateImageEdit creates an edited or extended image given an original image and a prompt.
func (i *ImagesAPI) CreateImageEdit(ctx context.Context, imgEditReq *ImageEditRequest) (*ImageResponse, *Response, error) {
	u := "v1/images/edits"
	req, err := i.openAIClient.NewRequest(http.MethodPost, u, imgEditReq)
	if err != nil {
		return nil, nil, err
	}

	imgResp := new(ImageResponse)

	resp, err := i.openAIClient.Do(ctx, req, imgResp)
	if err != nil {
		return nil, resp, err
	}

	return imgResp, resp, nil
}

// CreateImageVariation creates a variation of a given image.
func (i *ImagesAPI) CreateImageVariation(ctx context.Context, imgReq *ImageRequest) (*ImageResponse, *Response, error) {
	u := "v1/images/variations"
	req, err := i.openAIClient.NewRequest(http.MethodPost, u, imgReq)
	if err != nil {
		return nil, nil, err
	}

	imgResp := new(ImageResponse)

	resp, err := i.openAIClient.Do(ctx, req, imgResp)
	if err != nil {
		return nil, resp, err
	}

	return imgResp, resp, nil
}
