package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/SturlaSolheim/mediaCircleBackend/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := database.InitInMemoryDB(); err != nil {
		log.Fatal("Det skjedde en feil under oppstart av database", err)
	}
	defer database.CloseDB()

	// Create tables
	if err := database.CreateTables(); err != nil {
		log.Fatal("Det skjedde en feil under tabelldannelse", err)
	}

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

	container := routes.NewContainer()
	container.SetupRoutes(r)

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
