package http

import (
	"encoding/json"
	"net/http"

	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
	"persons_generator/core/storage"
)

func (h *handlers) GetProtoCultures(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		q      = r.URL.Query().Get("q")
		limit  = ExtractLimitFromReq(r)
		offset = ExtractOffsetFromReq(r)
	)

	cultures, total, err := h.cultureSrv.GetProtoCultures(ctx, q, storage.PaginationSortingOpts{
		Pagination: &storage.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	})
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	data, err := serializeCultures(cultures)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	respJSON, err := json.Marshal(GetProtoCulturesResponse{
		Data:  data,
		Total: total,
	})
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(hs.TraceIDHeader, cont.GetTraceID(ctx))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respJSON)
}
