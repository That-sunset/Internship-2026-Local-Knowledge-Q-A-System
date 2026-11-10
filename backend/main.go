package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("backend ok"))
	})

	fmt.Println("Backend server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", r)
}