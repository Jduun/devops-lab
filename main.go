package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()

	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello World",
		})
	}).Methods(http.MethodGet)

	app.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":32777", app))
}
