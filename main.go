package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/config"
	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"github.com/SturlaSolheim/mediaCircleBackend/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Det skjedde en feil under lasting av konfigurasjon", err)
	}

	if err := database.InitDB(); err != nil {
		log.Fatal("Det skjedde en feil under oppstart av database", err)
	}
	defer database.CloseDB()

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

	openAPIContainer := routes.NewOpenAPIContainer()
	openAPIContainer.SetupOpenAPIRoutes(r)

	address := config.AppConfig.GetServerAddress()
	log.Printf("Server started on %s", address)
	http.ListenAndServe(address, r)
}
