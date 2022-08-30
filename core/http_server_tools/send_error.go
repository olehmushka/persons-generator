package http_server_tools

import (
	"context"
	"encoding/json"
	"net/http"

	cont "persons_generator/core/context"
	"persons_generator/core/wrapped_error"
)

func SendErrorResp(ctx context.Context, w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	wErr := wrapped_error.Cast(err)
	resp := ErrorResp{
		Error:        wErr.Err.Error(),
		ErrorMessage: wErr.Msg,
		TraceID:      cont.GetTraceID(ctx),
	}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Trace-ID", cont.GetTraceID(ctx))
	w.WriteHeader(wErr.StatusCode)
	_, _ = w.Write(respJSON)
}
