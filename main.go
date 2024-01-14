package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Write hello world as json
    w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			map[string]string{"hello": "world"},
		)
	})

  log.Println("Server running on port 8080")

  if err := http.ListenAndServe(":8080", r); err != nil {
    log.Fatal(err)
  }
}
