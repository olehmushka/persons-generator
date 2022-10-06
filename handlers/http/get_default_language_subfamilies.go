package http

import (
	"encoding/json"
	"net/http"
	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
	"persons_generator/core/storage"
)

func (h *handlers) GetDefaultLanguageSubfamilies(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		limit  = ExtractLimitFromReq(r)
		offset = ExtractOffsetFromReq(r)
	)

	sfs, total, err := h.languageSrv.ReadDefaultLanguageSubfamilies(ctx, storage.PaginationSortingOpts{
		Pagination: &storage.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	})
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	data, err := serializeLanguageSubfamilies(sfs)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	respJSON, err := json.Marshal(GetLanguageSubfamiliesResponse{
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
