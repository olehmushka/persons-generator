package http

import (
	"net/http"

	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
)

func (h *handlers) DeleteReligionByID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = ExtractIDFromPath(r, "id")
	)

	if err := h.religionSrv.DeleteReligionByID(ctx, id); err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(hs.TraceIDHeader, cont.GetTraceID(ctx))
	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte(""))
}
