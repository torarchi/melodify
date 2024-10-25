package repository

import (
	"Melodify/internal/domain/models"
	"context"
)

type PredictionRepository interface {
	Create(ctx context.Context, input models.Input) (*models.Prediction, error)
	Get(ctx context.Context, id string) (*models.Prediction, error)
}
