package handlers

import (
	"encoding/json"
	"net/http"

	"Melodify/internal/domain/models"
	"Melodify/internal/usecase"
)

type PredictionHandler struct {
	useCase *usecase.PredictionUseCase
}

func NewPredictionHandler(useCase *usecase.PredictionUseCase) *PredictionHandler {
	return &PredictionHandler{
		useCase: useCase,
	}
}

func (h *PredictionHandler) CreatePrediction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	prediction, err := h.useCase.CreatePrediction(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prediction)
}

func (h *PredictionHandler) GetPrediction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing prediction ID", http.StatusBadRequest)
		return
	}

	prediction, err := h.useCase.GetPrediction(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prediction)
}
