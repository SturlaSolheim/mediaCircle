package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SturlaSolheim/mediaCircleBackend/generated"
	"github.com/SturlaSolheim/mediaCircleBackend/generated/health"
)

// HealthHandlerImpl implements the generated.HealthHandler interface
type HealthHandlerImpl struct{}

// NewOpenAPIHealthHandler creates a new HealthHandlerImpl instance
func NewOpenAPIHealthHandler() health.ServerInterface {
	return &HealthHandlerImpl{}
}

// HealthCheck implements the health check endpoint
func (h *HealthHandlerImpl) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := generated.Response{
		Message: "Media Circle API is healthy",
		Status:  200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}