package services

import (
	"context"
	"persons_generator/core/wrapped_error"
	engineWorld "persons_generator/engine/entities/world"
	"persons_generator/internal/world/adapters/engine"
	"persons_generator/internal/world/adapters/mq"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type world struct {
	engineAdp engine.Adapter
	mqAdp     mq.Adapter
}

func New(engineAdp engine.Adapter, mqAdp mq.Adapter) World {
	return &world{
		engineAdp: engineAdp,
		mqAdp:     mqAdp,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *world) CreateWorld(
	ctx context.Context,
	stopYear,
	amount,
	maleAmount,
	femaleAmount int,
	religionCultureRels map[uuid.UUID]uuid.UUID,
) (uuid.UUID, error) {
	id := uuid.New()
	if err := s.mqAdp.RunAndSaveWorld(ctx, id, stopYear, amount, maleAmount, femaleAmount, religionCultureRels); err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s *world) RunAndSaveWorld(ctx context.Context, in []byte) error {
	payload, err := s.mqAdp.ParseRunAndSaveWorldMsg(ctx, in)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not parse message")
	}
	w, err := s.engineAdp.CreateWorld(ctx, payload.Amount, payload.MaleAmount, payload.FemaleAmount, payload.ReligionCultureRels)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not create world")
	}
	w.ID = payload.WorldID
	if err := s.engineAdp.RunAndSaveWorld(ctx, w, payload.StopYear); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not run and save world from the world")
	}

	return nil
}

func (s *world) GetWorldRunningProgress(worldID uuid.UUID) (engineWorld.ProgressRunWorld, error) {
	return s.engineAdp.GetWorldRunningProgress(worldID)
}

func (s *world) ParseRunAndSaveWorldMsg(ctx context.Context, in []byte) (mq.RunAndSaveWorldPayload, error) {
	return s.mqAdp.ParseRunAndSaveWorldMsg(ctx, in)
}

func (s *world) DeleteAllWorlds(ctx context.Context) error {
	return s.engineAdp.DeleteAllWorlds(ctx)
}
