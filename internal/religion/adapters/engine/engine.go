package engine

import (
	"context"

	"persons_generator/config"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/orchestrator"
	"persons_generator/internal/religion/entities"

	"github.com/google/uuid"
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

		MongoDBURL:              cfg.MongoDB.URL,
		MongoDBUsername:         cfg.MongoDB.Username,
		MongoDBPassword:         cfg.MongoDB.Password,
		MongoDBMaxBulkItemsSize: cfg.MongoDB.MaxBulkItemsSize,
		MongoDBDBName:           cfg.MongoDB.DBName,
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
			c, err := a.engine.ReadCultureByID(ctx, cultureID)
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

	religions, err := a.engine.CreateReligions(amount, preferences)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create religions")
	}
	for _, r := range religions {
		if err := a.engine.SaveReligion(ctx, r); err != nil {
			return nil, err
		}
	}

	return religions, nil
}

func (a *adapter) GetReligionByID(ctx context.Context, id uuid.UUID) (*religion.Religion, error) {
	// return a.engine.GetCultureByID(id)
	return a.engine.ReadReligionByID(ctx, id)
}

func (a *adapter) DeleteReligionByID(ctx context.Context, id uuid.UUID) error {
	return a.engine.DeleteReligionByID(ctx, id)
}

func (a *adapter) DeleteAllReligions(ctx context.Context) error {
	return a.engine.DeleteAllReligions(ctx)
}
