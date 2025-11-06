package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/SturlaSolheim/mediaCircleBackend/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := models.Response{
			Message: "Media Circle test",
			Status:  200,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	routes.SetupRoutes(r)

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
