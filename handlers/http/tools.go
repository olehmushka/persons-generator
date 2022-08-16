package http

import (
	"net/http"
	"strconv"
)

func ExtractIntFromReq(r *http.Request, key string) *int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return nil
	}
	if n, err := strconv.Atoi(val); err == nil {
		return &n
	}

	return nil
}

func ExtractLimitFromReq(r *http.Request) int {
	if limit := ExtractIntFromReq(r, "limit"); limit != nil {
		return *limit
	}

	return DefaultLimit
}

func ExtractOffsetFromReq(r *http.Request) int {
	if limit := ExtractIntFromReq(r, "offset"); limit != nil {
		return *limit
	}

	return DefaultOffset
}
