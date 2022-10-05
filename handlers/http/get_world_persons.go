package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
	"persons_generator/core/wrapped_error"

	"github.com/google/uuid"
)

func (h *handlers) GetWorldPersons(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		qWorldID = r.URL.Query().Get("world_id")
		limit    = ExtractLimitFromReq(r)
		offset   = ExtractOffsetFromReq(r)
	)
	worldID, err := uuid.Parse(qWorldID)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, wrapped_error.NewBadRequestError(err, fmt.Sprintf("can not parse world_id (world_id=%v)", qWorldID)))
		return
	}

	persons, total, err := h.personsSrv.GetPersonsByWorldID(ctx, worldID.String(), offset, limit)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	respJSON, err := json.Marshal(GetWorldPersonsResponse{
		Data:  serializePersons(persons),
		Total: total,
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
