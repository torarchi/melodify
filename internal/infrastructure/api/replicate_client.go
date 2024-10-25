package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"Melodify/internal/domain/models"
)

type ReplicateClient struct {
	baseURL    string
	apiToken   string
	httpClient *http.Client
	version    string
}

func NewReplicateClient(apiToken string) *ReplicateClient {
	return &ReplicateClient{
		baseURL:  "https://api.replicate.com/v1",
		apiToken: apiToken,
		version:  "8cf61ea6c56afd61d8f5b9ffd14d7c216c0a93844ce2d82ac1c9ecc9c7f24e05",
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c *ReplicateClient) Create(ctx context.Context, input models.Input) (*models.Prediction, error) {
	payload := models.ReplicateRequest{
		Version: c.version,
		Input: models.ReplicateInput{
			Alpha:          0.5,
			PromptA:        input.PromptA,
			PromptB:        input.PromptB,
			Denoising:      0.75,
			SeedImageID:    "vibes",
			InferenceSteps: 50,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/predictions", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var prediction models.Prediction
	if err := c.doRequest(req, &prediction); err != nil {
		return nil, err
	}

	return &prediction, nil
}

func (c *ReplicateClient) Get(ctx context.Context, id string) (*models.Prediction, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/predictions/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var prediction models.Prediction
	if err := c.doRequest(req, &prediction); err != nil {
		return nil, err
	}

	return &prediction, nil
}

func (c *ReplicateClient) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *ReplicateClient) doRequest(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var apiErr models.ReplicateError
		if err := json.Unmarshal(body, &apiErr); err == nil && apiErr.Detail != "" {
			return &models.APIError{
				Operation:  req.URL.Path,
				StatusCode: resp.StatusCode,
				Message:    apiErr.Detail,
			}
		}
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	if v != nil {
		if err := json.Unmarshal(body, v); err != nil {
			return fmt.Errorf("unmarshal response: %w", err)
		}
	}

	return nil
}
