package engine

import (
	"context"

	"persons_generator/config"
	"persons_generator/engine/orchestrator"
	cultureEntities "persons_generator/internal/culture/entities"
	"persons_generator/internal/religion/entities"

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

func (a *adapter) CreateReligions(ctx context.Context, cultures []*cultureEntities.Culture) ([]*entities.Religion, error) {
	c, err := a.engine.CreateReligions(nil)
	if err != nil {
		return nil, err
	}

	return serializeReligions(c), nil
}
