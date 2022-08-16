package engine

import (
	"context"

	"persons_generator/core/tools"
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

func (a *adapter) CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*entities.Culture, error) {
	c, err := a.engine.CreateCultures(amount, deserializeCulturePreferences(preferred))
	if err != nil {
		return nil, err
	}

	return serializeCultures(c), nil
}

func (a *adapter) GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*entities.Culture, int, error) {
	c, err := a.engine.SearchCultures(q)
	if err != nil {
		return nil, 0, err
	}
	if len(c) == 0 {
		return []*entities.Culture{}, 0, nil
	}

	return serializeCultures(tools.Paginate(c, offset, limit)), len(c), nil
}
