package mq_run_and_save_world

import (
	"context"
	worldServices "persons_generator/internal/world/services"

	"go.uber.org/fx"
)

type handler struct {
	worldSrv worldServices.World
}

func New(
	worldSrv worldServices.World,
) RunAndSaveWorld {
	return &handler{
		worldSrv: worldSrv,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (h *handler) ProcessRunAndSaveWorld(ctx context.Context, in []byte) error {
	return h.worldSrv.RunAndSaveWorld(ctx, in)
}
