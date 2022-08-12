package http_server_tools

type ErrorResp struct {
	Error        string `json:"error"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
	TraceID      string `json:"trace_id"`
}
