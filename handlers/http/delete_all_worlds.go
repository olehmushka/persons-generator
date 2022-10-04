package http

import (
	"net/http"

	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
)

func (h *handlers) DeleteAllWorlds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.worldSrv.DeleteAllWorlds(ctx); err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(hs.TraceIDHeader, cont.GetTraceID(ctx))
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}
