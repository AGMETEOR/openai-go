// Copyright 2023 The openai-go AUTHORS. All rights reserved.

package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	defaultBaseURL = "https://api.openai.com/"
)

type Api struct {
	openAIClient *OpenAIClient
}

type OpenAIClient struct {
	client  *http.Client
	BaseURL *url.URL

	Completions *CompletionsAPI
	Models      *ModelsAPI
	Chat        *ChatAPI
	Images      *ImagesAPI
	Embeddings  *EmbeddingsAPI
	Audio       *AudioAPI
	File        *FileAPI
	FineTunes   *FineTunesAPI
	Moderations *ModerationsAPI
}

type Response struct {
	*http.Response
}

func NewClient(httpClient *http.Client) *OpenAIClient {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL) // TODO: Handle this error

	oapiClient := &OpenAIClient{
		client:  httpClient,
		BaseURL: baseURL,
	}

	oapiClient.Completions = &CompletionsAPI{
		openAIClient: oapiClient,
	}

	oapiClient.Models = &ModelsAPI{
		openAIClient: oapiClient,
	}

	oapiClient.Chat = &ChatAPI{
		openAIClient: oapiClient,
	}

	oapiClient.Images = &ImagesAPI{
		openAIClient: oapiClient,
	}

	oapiClient.Embeddings = &EmbeddingsAPI{
		openAIClient: oapiClient,
	}

	oapiClient.Audio = &AudioAPI{
		openAIClient: oapiClient,
	}

	oapiClient.File = &FileAPI{
		openAIClient: oapiClient,
	}

	oapiClient.FineTunes = &FineTunesAPI{
		openAIClient: oapiClient,
	}

	oapiClient.Moderations = &ModerationsAPI{
		openAIClient: oapiClient,
	}

	return oapiClient
}

func (oapiClient *OpenAIClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(oapiClient.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", oapiClient.BaseURL)
	}

	u, err := oapiClient.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	authToken := os.Getenv("OPENAI_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	return req, nil
}

func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *OpenAIClient) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}

		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return response, err
}
