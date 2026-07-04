package health

import (
	"encoding/json"
	"net/http"
	"time"
)

const DefaultReadHeaderTimeout = 5 * time.Second

type response struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, response{
		Status:  "ok",
		Service: "example-ci-go-backend",
	})
}

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, response{
		Status:  "ready",
		Service: "example-ci-go-backend",
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}
