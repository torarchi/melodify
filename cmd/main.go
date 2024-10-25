package main

import (
	"log"
	"net/http"

	"Melodify/internal/delivery/http/handlers"
	"Melodify/internal/infrastructure/api"
	"Melodify/internal/usecase"
	"Melodify/pkg/config"
)

func main() {
	log.Println("Starting application...")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	replicateClient := api.NewReplicateClient(cfg.ReplicateAPIToken)
	predictionUseCase := usecase.NewPredictionUseCase(replicateClient)
	predictionHandler := handlers.NewPredictionHandler(predictionUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("/predictions", predictionHandler.CreatePrediction)
	mux.HandleFunc("/predictions/get", predictionHandler.GetPrediction)

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mux,
	}

	log.Printf("Server starting on :%s", cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
