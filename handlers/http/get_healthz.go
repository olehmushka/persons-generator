package http

import (
	"encoding/json"
	"net/http"
)

type GetHealthzResponse struct {
	Status string `json:"status"`
}

// GetHealthz is a shallow liveness check: no DB/Redis dependency, just
// confirms the process is up and serving requests.
func (h *handlers) GetHealthz(w http.ResponseWriter, r *http.Request) {
	respJSON, err := json.Marshal(GetHealthzResponse{Status: "ok"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respJSON)
}
