package http_server_tools

import (
	"context"
	"encoding/json"
	"net/http"

	cont "persons_generator/core/context"
)

func SendErrorResp(ctx context.Context, w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	resp := ErrorResp{
		Error:   err.Error(),
		TraceID: cont.GetTraceID(ctx),
	}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Trace-ID", cont.GetTraceID(ctx))
	_, _ = w.Write(respJSON)
}
