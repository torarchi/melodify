package usecase

import (
	"context"
	"fmt"
	"time"

	"Melodify/internal/domain/models"
	"Melodify/internal/domain/repository"
)

type PredictionUseCase struct {
	repo    repository.PredictionRepository
	timeout time.Duration
}

func NewPredictionUseCase(repo repository.PredictionRepository) *PredictionUseCase {
	return &PredictionUseCase{
		repo:    repo,
		timeout: 60 * time.Second,
	}
}

func (uc *PredictionUseCase) CreatePrediction(ctx context.Context, input models.Input) (*models.Prediction, error) {
	if err := uc.validateInput(input); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.Create(ctx, input)
}

func (uc *PredictionUseCase) GetPrediction(ctx context.Context, id string) (*models.Prediction, error) {
	if id == "" {
		return nil, fmt.Errorf("prediction id is required")
	}

	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.Get(ctx, id)
}

func (uc *PredictionUseCase) validateInput(input models.Input) error {
	if input.PromptA == "" {
		return fmt.Errorf("prompt_a is required")
	}
	if input.PromptB == "" {
		return fmt.Errorf("prompt_b is required")
	}
	if len(input.PromptA) > 20 {
		return fmt.Errorf("prompt_a is too long (max 20 characters)")
	}
	if len(input.PromptB) > 20 {
		return fmt.Errorf("prompt_b is too long (max 20 characters)")
	}
	return nil
}
