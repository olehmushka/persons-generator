package http_server

import (
	"net/http"

	cont "persons_generator/core/context"

	"github.com/google/uuid"
)

func AppendTraceID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceIDToAppend := uuid.New().String()
		if traceID := r.Header.Get(TraceIDHeader); traceID != "" {
			traceIDToAppend = traceID
		}
		ctx = cont.AppendTraceID(ctx, traceIDToAppend)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
