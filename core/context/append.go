package context

import "context"

func AppendTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, ContextTraceID, traceID)
}

func GetTraceID(ctx context.Context) string {
	if traceID, ok := ctx.Value(ContextTraceID).(string); ok {
		return traceID
	}

	return ""
}
