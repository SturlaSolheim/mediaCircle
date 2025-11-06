package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/models"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Message: "API is healthy",
		Status:  200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}