package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jiaflu/example-ci-go-backend/internal/health"
	"github.com/jiaflu/example-ci-go-backend/internal/orders"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", health.HandleHealth)
	mux.HandleFunc("GET /readyz", health.HandleReadiness)
	mux.HandleFunc("GET /orders/summary", orders.HandleSummary)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: health.DefaultReadHeaderTimeout,
	}

	log.Printf("server listening on :%s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
