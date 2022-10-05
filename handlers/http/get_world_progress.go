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

func (h *handlers) GetWorldProgress(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		qWorldID = r.URL.Query().Get("world_id")
	)
	worldID, err := uuid.Parse(qWorldID)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, wrapped_error.NewBadRequestError(err, fmt.Sprintf("can not parse world_id (world_id=%v)", qWorldID)))
		return
	}

	progress, err := h.worldSrv.GetWorldRunningProgress(worldID.String())
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	respJSON, err := json.Marshal(GetWorldProgressResponse{
		Year:           progress.Year,
		Population:     progress.Population,
		DeadPopulation: progress.DeadPopulation,
		Progress:       progress.Progress,
		Duration:       progress.Duration,
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
