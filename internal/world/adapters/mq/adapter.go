package mq

import (
	"context"
	"encoding/json"
	"persons_generator/config"
	"persons_generator/core/redis"
	"persons_generator/core/wrapped_error"

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
	worldID string,
	stopYear,
	amount,
	maleAmount,
	femaleAmount int,
	religionCultureRels map[string]string,
) error {
	b, err := json.Marshal(RunAndSaveWorldPayload{
		WorldID:             worldID,
		StopYear:            stopYear,
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

func (a *adapter) ParseRunAndSaveWorldMsg(ctx context.Context, in []byte) (*RunAndSaveWorldPayload, error) {
	var out RunAndSaveWorldPayload
	if err := json.Unmarshal(in, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal ParseRunAndSaveWorld msg")
	}

	return &out, nil
}
