package engine

import (
	"context"
	"persons_generator/config"
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

func (a *adapter) CreateWorld(ctx context.Context, amount, maleAmount, femaleAmount int, religionCultureRels map[uuid.UUID]uuid.UUID) (*world.World, error) {
	refs := make([]*religion.CultureReference, 0, len(religionCultureRels))
	for religionID, cultureID := range religionCultureRels {
		culture, err := a.engine.GetCultureByID(cultureID)
		if err != nil {
			return nil, err
		}
		relig, err := a.engine.GetReligionByID(religionID)
		if err != nil {
			return nil, err
		}
		refs = append(refs, &religion.CultureReference{
			Culture:  culture,
			Religion: relig,
		})
	}

	// (*orchestrator.Orchestrator).CreateWorld(size int, cultures []*culture.Culture, religions []*religion.Religion) (*world.World, error)
	// return a.engine.CreateWorld(amount, []*culture.Culture{culture}, []*religion.Religion{relig})
	return nil, nil
}
