package http

import (
	"encoding/json"
	"net/http"

	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
)

func (h *handlers) CreateCultures(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := PostCreatCulturesRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	cultures, err := h.cultureSrv.CreateCultures(ctx, req.Amount, deserializeCulturePreferences(req.Preferred))
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	data, err := serializeCultures(cultures)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	respJSON, err := json.Marshal(PostCreateCulturesResponse{
		Data: data,
	})
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(hs.TraceIDHeader, cont.GetTraceID(ctx))
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(respJSON)
}
