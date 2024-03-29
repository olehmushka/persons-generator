package services

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	engineWorld "persons_generator/engine/entities/world"
	"persons_generator/internal/world/adapters/engine"
	"persons_generator/internal/world/adapters/mq"
	"persons_generator/internal/world/entities"

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
	religionCultureRels map[string]string,
) (string, error) {
	id := uuid.New().String()
	if err := s.mqAdp.RunAndSaveWorld(ctx, id, stopYear, amount, maleAmount, femaleAmount, religionCultureRels); err != nil {
		return "", err
	}

	return id, nil
}

func (s *world) RunAndSaveWorld(ctx context.Context, in []byte) error {
	payload, err := s.mqAdp.ParseRunAndSaveWorldMsg(ctx, in)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not parse message")
	}
	w, err := s.engineAdp.CreateWorld(ctx, payload.WorldID, payload.Amount, payload.MaleAmount, payload.FemaleAmount, payload.ReligionCultureRels)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not create world")
	}
	w.ID = payload.WorldID
	if err := s.engineAdp.RunAndSaveWorld(ctx, w, payload.StopYear); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not run and save world from the world")
	}

	return nil
}

func (s *world) GetWorldRunningProgress(worldID string) (engineWorld.ProgressRunWorld, error) {
	return s.engineAdp.GetWorldRunningProgress(worldID)
}

func (s *world) ParseRunAndSaveWorldMsg(ctx context.Context, in []byte) (*mq.RunAndSaveWorldPayload, error) {
	return s.mqAdp.ParseRunAndSaveWorldMsg(ctx, in)
}

func (s *world) DeleteWorldByID(ctx context.Context, id string) error {
	return s.engineAdp.DeleteWorldByID(ctx, id)
}

func (s *world) DeleteAllWorlds(ctx context.Context) error {
	return s.engineAdp.DeleteAllWorlds(ctx)
}

func (s *world) ReadWorlds(ctx context.Context, opts storage.PaginationSortingOpts) ([]*entities.World, int, error) {
	worlds, err := s.engineAdp.ReadWorlds(ctx, opts)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not read worlds")
	}
	serializedWorlds, err := serializeSerializedWorlds(worlds)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not serialize worlds")
	}

	count, err := s.engineAdp.CountWorlds(ctx, opts)
	if err != nil {
		return nil, 0, wrapped_error.NewInternalServerError(err, "can not count worlds")
	}

	return serializedWorlds, count, nil
}
