package http

import (
	"encoding/json"
	"net/http"

	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
	"persons_generator/core/wrapped_error"
)

func (h *handlers) CreateReligions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := PostCreateReligionsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http_server_tools.SendErrorResp(ctx, w, wrapped_error.New(http.StatusBadRequest, err, "can not decode request"))
		return
	}

	prefs, err := deserializeReligionPreferences(req.Preferred)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	rs, err := h.religionSrv.CreateReligions(ctx, req.Amount, prefs)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	respJSON, err := json.Marshal(PostCreateReligionsResponse{
		Data:  serializeReligions(rs),
		Total: len(rs),
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
