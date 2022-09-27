package mq_run_and_save_world

import (
	"context"

	"go.uber.org/fx"
)

type handler struct{}

func New() RunAndSaveWorld {
	return &handler{}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (h *handler) ProcessRunAndSaveWorld(context.Context, []byte) error {
	return nil
}
