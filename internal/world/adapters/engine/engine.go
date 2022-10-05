package engine

import (
	"context"
	"persons_generator/config"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/entities/world"
	"persons_generator/engine/orchestrator"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type adapter struct {
	engine *orchestrator.Orchestrator
}

func New(cfg *config.Config) (Adapter, error) {
	e, err := orchestrator.New(orchestrator.Config{
		RedisURL:      cfg.Redis.URL,
		RedisUsername: cfg.Redis.Username,
		RedisPassword: cfg.Redis.Password,

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

func (a *adapter) CreateWorld(ctx context.Context, amount, maleAmount, femaleAmount int, religionCultureRels map[uuid.UUID]uuid.UUID) (*world.World, error) {
	var (
		religions = make([]*religion.Religion, 0, len(religionCultureRels))
		cultures  = make([]*culture.Culture, 0, len(religionCultureRels))
		refs      = make([]*religion.CultureReference, 0, len(religionCultureRels))
	)
	for religionID, cultureID := range religionCultureRels {
		culture, err := a.engine.ReadCultureByID(ctx, cultureID)
		if err != nil {
			return nil, err
		}
		cultures = append(cultures, culture)
		relig, err := a.engine.ReadReligionByID(ctx, religionID)
		if err != nil {
			return nil, err
		}
		religions = append(religions, relig)
		refs = append(refs, &religion.CultureReference{
			Culture:  culture,
			Religion: relig,
		})
	}
	cultures = culture.UniqueCultures(cultures)
	religions = religion.UniqueReligions(religions)
	size := world.GetSizeByHumanAmount(amount)
	w, err := a.engine.CreateWorld(
		size,
		cultures,
		religions,
		refs,
	)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create world in adapter")
	}
	if err := a.engine.PresaveWorld(ctx, w); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not presave world")
	}

	return w, nil
}

func (a *adapter) RunAndSaveWorld(ctx context.Context, w *world.World, stopYear int) error {
	if err := a.engine.RunAndSaveWorld(ctx, w, stopYear); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not run and save world")
	}

	return nil
}

func (a *adapter) GetWorldRunningProgress(worldID uuid.UUID) (world.ProgressRunWorld, error) {
	return a.engine.GetWorldRunningProgress(worldID)
}

func (a *adapter) DeleteWorldByID(ctx context.Context, id uuid.UUID) error {
	return a.engine.DeleteWorldByID(ctx, id)
}

func (a *adapter) DeleteAllWorlds(ctx context.Context) error {
	return a.engine.DeleteAllWorlds(ctx)
}

func (a *adapter) ReadWorlds(ctx context.Context, opts storage.PaginationSortingOpts) ([]*world.SerializedWorld, error) {
	return a.engine.ReadWorlds(ctx, opts)
}

func (a *adapter) CountWorlds(ctx context.Context, opts storage.PaginationSortingOpts) (int, error) {
	return a.engine.CountWorlds(ctx, opts)
}
