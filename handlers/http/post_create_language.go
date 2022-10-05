package http

import (
	"encoding/json"
	"net/http"
	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
)

func (h *handlers) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := PostCreateLanguageRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	langIn, err := deserializeLanguage(req.Name, req.Subfamily, req.WordBase, req.IsLiving)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	langID, err := h.languageSrv.CreateLanguage(ctx, langIn)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	langIn.ID = langID

	data, err := serializeLanguage(langIn)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	respJSON, err := json.Marshal(PostCreateLanguageResponse{
		Data: data,
	})
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(hs.TraceIDHeader, cont.GetTraceID(ctx))
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(respJSON)
}
