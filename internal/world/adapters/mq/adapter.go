package mq

import (
	"context"
	"encoding/json"
	"persons_generator/config"
	"persons_generator/core/redis"
	"persons_generator/core/wrapped_error"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type adapter struct {
	driver redis.Publisher
	chName string
}

func New(cfg *config.Config, driver redis.Publisher) (Adapter, error) {
	return &adapter{driver: driver, chName: cfg.Redis.Channels.RunAndSaveWorld.Name}, nil
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) RunAndSaveWorld(
	ctx context.Context,
	worldID uuid.UUID,
	amount int,
	maleAmount int,
	femaleAmount int,
	religionCultureRels map[uuid.UUID]uuid.UUID,
) error {
	b, err := json.Marshal(RunAndSaveWorldPayload{
		WorldID:             worldID,
		Amount:              amount,
		MaleAmount:          maleAmount,
		FemaleAmount:        femaleAmount,
		ReligionCultureRels: religionCultureRels,
	})
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not marshal payload for RunAndSaveWorld")
	}

	if err := a.driver.Publish(ctx, a.chName, b); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not pubish RunAndSaveWorld msg")
	}

	return nil
}

func (a *adapter) ParseRunAndSaveWorldMsg(ctx context.Context, in []byte) (RunAndSaveWorldPayload, error) {
	var out RunAndSaveWorldPayload
	if err := json.Unmarshal(in, &out); err != nil {
		return RunAndSaveWorldPayload{}, wrapped_error.NewInternalServerError(err, "can not unmarshal ParseRunAndSaveWorld msg")
	}

	return out, nil
}
