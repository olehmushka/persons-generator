package engine

import (
	"context"

	"persons_generator/config"
	"persons_generator/core/tools"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/orchestrator"
	"persons_generator/internal/culture/entities"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type adapter struct {
	engine *orchestrator.Orchestrator
}

func New(cfg *config.Config) (Adapter, error) {
	e, err := orchestrator.New(orchestrator.Config{
		StorageFolderName: cfg.JSONStorage.StorageFolder,
	})
	if err != nil {
		return nil, err
	}

	return &adapter{engine: e}, nil
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*culture.Culture, error) {
	return a.engine.CreateCultures(amount, deserializeCulturePreferences(preferred))
}

func (a *adapter) GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*culture.Culture, int, error) {
	c, err := a.engine.SearchCultures(q)
	if err != nil {
		return nil, 0, err
	}
	if len(c) == 0 {
		return []*culture.Culture{}, 0, nil
	}

	return tools.Paginate(c, offset, limit), len(c), nil
}

func (a *adapter) GetCultureByID(ctx context.Context, id uuid.UUID) (*culture.Culture, error) {
	return a.engine.GetCultureByID(id)
}
