package http_server

import "time"

const (
	readTimeout       = 300 * time.Second
	readHeaderTimeout = 300 * time.Second
	writeTimeout      = 300 * time.Second
	idleTimeout       = 300 * time.Second

	TraceIDHeader = "X-Trace-ID"
)
