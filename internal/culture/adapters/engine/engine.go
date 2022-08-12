package engine

import (
	"context"

	"persons_generator/engine/orchestrator"
	"persons_generator/internal/culture/entities"

	"go.uber.org/fx"
)

type adapter struct {
	engine *orchestrator.Orchestrator
}

func New() (Adapter, error) {
	e, err := orchestrator.New()
	if err != nil {
		return nil, err
	}

	return &adapter{engine: e}, nil
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) CreateCultures(ctx context.Context, amount int, preferred []string) ([]*entities.Culture, error) {
	c, err := a.engine.CreateCultures(amount, preferred)
	if err != nil {
		return nil, err
	}

	return serializeCultures(c), nil
}
