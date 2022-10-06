package http

import (
	"encoding/json"
	"net/http"

	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
)

func (h *handlers) GetWordLanguageByID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		id   = ExtractIDFromPath(r, "id")
		kind = r.URL.Query().Get("kind")
		word string
	)

	switch kind {
	case RandomWordKind:
		result, err := h.languageSrv.GetRandomWordByLanguageID(ctx, id)
		if err != nil {
			http_server_tools.SendErrorResp(ctx, w, err)
			return
		}
		word = result
	case OwnNameWordKind:
		result, err := h.languageSrv.GetRandomOwnNameByLanguageID(ctx, id, ExtractSexFromReq(r))
		if err != nil {
			http_server_tools.SendErrorResp(ctx, w, err)
			return
		}
		word = result
	case CultureNameWordKind:
		result, err := h.languageSrv.GetRandomCultureNameByLanguageID(ctx, id)
		if err != nil {
			http_server_tools.SendErrorResp(ctx, w, err)
			return
		}
		word = result
	case ReligionNameWordKind:
		result, err := h.languageSrv.GetRandomReligionNameByLanguageID(ctx, id)
		if err != nil {
			http_server_tools.SendErrorResp(ctx, w, err)
			return
		}
		word = result
	default:
		result, err := h.languageSrv.GetRandomWordByLanguageID(ctx, id)
		if err != nil {
			http_server_tools.SendErrorResp(ctx, w, err)
			return
		}
		word = result
	}

	respJSON, err := json.Marshal(GetWordResponse{
		Word: word,
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
