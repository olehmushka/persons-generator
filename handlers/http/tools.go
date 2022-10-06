package http

import (
	"net/http"
	"persons_generator/engine/entities/gender"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
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

func ExtractIDFromPath(r *http.Request, key string) string {
	val := chi.URLParam(r, key)
	if val == "" {
		return ""
	}
	if id, err := uuid.Parse(val); err == nil {
		return id.String()
	}

	return ""
}

func ExtractSexFromReq(r *http.Request) gender.Sex {
	val := r.URL.Query().Get("sex")
	if val == "" {
		return ""
	}
	switch val {
	case "male":
		return gender.MaleSex
	case "female":
		return gender.FemaleSex
	}

	return ""
}
