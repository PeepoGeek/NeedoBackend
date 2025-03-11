package healthcheck

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the response structure for the health check endpoint
type HealthResponse struct {
	Status string `json:"status"`
}

// @Summary Health Check endpoint
// @Description Returns the health status of the application
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /healthcheck [get]
func Handler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
