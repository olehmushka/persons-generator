package http

import (
	"encoding/json"
	"net/http"

	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
)

func (h *handlers) GenerateWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	respJSON, err := json.Marshal(PostCreateWorldResponse{})
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(hs.TraceIDHeader, cont.GetTraceID(ctx))
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(respJSON)
}
