package engine

import (
	"context"

	"persons_generator/config"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/orchestrator"
	"persons_generator/internal/religion/entities"

	"go.uber.org/fx"
)

type adapter struct {
	engine *orchestrator.Orchestrator
}

func New(cfg *config.Config) (Adapter, error) {
	e, err := orchestrator.New(orchestrator.Config{
		StorageFolderName: cfg.JSONStorage.StorageFolder,
		RedisURL:          cfg.Redis.URL,
		RedisUsername:     cfg.Redis.Username,
		RedisPassword:     cfg.Redis.Password,
	})
	if err != nil {
		return nil, err
	}

	return &adapter{engine: e}, nil
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) CreateReligions(ctx context.Context, amount int, preferred []*entities.Preference) ([]*religion.Religion, error) {
	preferences := make([]*religion.Preference, len(preferred))
	for i := range preferences {
		cultures := make([]*culture.Culture, 0, len(preferred[i].CultureIDs))
		for _, cultureID := range preferred[i].CultureIDs {
			c, err := a.engine.GetCultureByID(cultureID)
			if err != nil {
				return nil, err
			}
			if c != nil {
				cultures = append(cultures, c)
			}
		}
		preferences[i] = &religion.Preference{
			Cultures: cultures,
			Amount:   preferred[i].Amount,
		}
	}

	return a.engine.CreateReligions(amount, preferences)
}
