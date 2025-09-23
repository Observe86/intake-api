package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Observe86/intake-api/internal/api"
	"github.com/Observe86/intake-api/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	collectorURL := os.Getenv("COLLECTOR_URL")
	if collectorURL == "" {
		collectorURL = "http://collector-gateway:9090"
	}

	// Metrics
	metricsService := &service.MetricsService{CollectorURL: collectorURL}
	metricsHandler := &api.MetricsHandler{Service: metricsService}
	http.HandleFunc("/api/v1/intake/metrics", metricsHandler.HandleMetrics)

	// Logs
	logsService := &service.LogsService{CollectorURL: collectorURL}
	logsHandler := &api.LogsHandler{Service: logsService}
	http.HandleFunc("/api/v1/intake/logs", logsHandler.HandleLogs)

	// Traces
	tracesService := &service.TracesService{CollectorURL: collectorURL}
	tracesHandler := &api.TracesHandler{Service: tracesService}
	http.HandleFunc("/api/v1/intake/traces", tracesHandler.HandleTraces)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting intake-api on port %s, forwarding to Collector-Gateway %s", port, collectorURL)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
