package main

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello World",
	})
}

func cpuHandler(w http.ResponseWriter, r *http.Request) {
	end := time.Now().Add(2 * time.Second)

	data := []byte("load-test")

	for time.Now().Before(end) {
		hash := sha256.Sum256(data)
		data = hash[:]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "done",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloWorldHandler).Methods(http.MethodGet)
	r.HandleFunc("/cpu", cpuHandler).Methods(http.MethodGet)
	r.HandleFunc("/ready", readyHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":32777", r))
}
