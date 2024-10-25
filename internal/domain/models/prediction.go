package models

import (
	"fmt"
)

type PredictionStatus string

const (
	StatusStarting   PredictionStatus = "starting"
	StatusProcessing PredictionStatus = "processing"
	StatusSucceeded  PredictionStatus = "succeeded"
	StatusFailed     PredictionStatus = "failed"
	StatusCanceled   PredictionStatus = "canceled"
)

type Input struct {
	PromptA string `json:"prompt_a"`
	PromptB string `json:"prompt_b"`
}

type Output struct {
	Audio       string `json:"audio"`
	Spectrogram string `json:"spectrogram"`
}

type Prediction struct {
	ID        string           `json:"id"`
	Version   string           `json:"version"`
	Status    PredictionStatus `json:"status"`
	Input     Input            `json:"input"`
	Output    *Output          `json:"output,omitempty"`
	Error     string           `json:"error,omitempty"`
	CreatedAt string           `json:"created_at"`
}

type ReplicateRequest struct {
	Version string         `json:"version"`
	Input   ReplicateInput `json:"input"`
}

type ReplicateInput struct {
	Alpha          float64 `json:"alpha"`
	PromptA        string  `json:"prompt_a"`
	PromptB        string  `json:"prompt_b"`
	Denoising      float64 `json:"denoising"`
	SeedImageID    string  `json:"seed_image_id"`
	InferenceSteps int     `json:"num_inference_steps"`
}

type ReplicateError struct {
	Detail string `json:"detail"`
	Type   string `json:"type"`
}

type APIError struct {
	Operation  string
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s failed with status code %d: %s", e.Operation, e.StatusCode, e.Message)
}
