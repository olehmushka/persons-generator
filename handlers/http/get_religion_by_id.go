package http

import (
	"encoding/json"
	"net/http"

	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
)

func (h *handlers) GetReligionByID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = ExtractIDFromPath(r, "id")
	)

	religion, err := h.religionSrv.GetReligionByID(ctx, id)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	data, err := serializeReligion(religion)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	respJSON, err := json.Marshal(GetReligionByIDResponse{
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
