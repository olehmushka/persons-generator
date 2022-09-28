package orchestrator

import (
	"context"
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/entities/world"
	"time"

	"github.com/google/uuid"
)

func (o *Orchestrator) CreateWorld(
	size int,
	cultures []*culture.Culture,
	religions []*religion.Religion,
	refs []*religion.CultureReference,
) (*world.World, error) {
	return world.New(world.Config{StorageFolderName: o.storageFolderName}, size, cultures, religions, refs)
}

func getWorldRunningProgressChannelName(worldID uuid.UUID) string {
	return fmt.Sprintf("w_%s", worldID.String())
}

func (o *Orchestrator) RunAndSaveWorld(w *world.World, stopYear int) error {
	if w == nil {
		return nil
	}

	progressCh := make(chan world.ProgressRunWorld)
	errCh := make(chan error)
	var isFinished bool
	go w.RunWorld(stopYear, progressCh, errCh)
	for {
		select {
		case err := <-errCh:
			if err != nil {
				return err
			}
			isFinished = true
			close(errCh)
			break
		case progress := <-progressCh:
			bProgres, err := json.Marshal(progress)
			if err != nil {
				return wrapped_error.NewInternalServerError(err, "can not marshal progress before setting to redis")
			}
			if err := o.storage.Set(
				context.Background(),
				getWorldRunningProgressChannelName(w.ID),
				string(bProgres),
				time.Hour,
			); err != nil {
				return wrapped_error.NewInternalServerError(err, "can not set world running progress")
			}
			// fmt.Printf("year: %d ; population: %d/%d ; progress: %f\n", progress.Year, progress.Population, progress.DeadPopulation, progress.Progress)
			if progress.Progress == 100 {
				isFinished = true
				close(progressCh)
				break
			}
		}
		if isFinished {
			break
		}
	}

	return w.Save()
}

func (o *Orchestrator) GetWorldRunningProgress(worldID uuid.UUID) (world.ProgressRunWorld, error) {
	val, err := o.storage.Get(context.Background(), getWorldRunningProgressChannelName(worldID))
	if err != nil {
		return world.ProgressRunWorld{}, wrapped_error.NewInternalServerError(err, "can not get world running progress")
	}
	if val == "" {
		return world.ProgressRunWorld{}, wrapped_error.NewNotFoundError(nil, "can not find world run progress by id")
	}
	var progress world.ProgressRunWorld
	if err := json.Unmarshal([]byte(val), &progress); err != nil {
		return world.ProgressRunWorld{}, wrapped_error.NewInternalServerError(err, "can not convert world running progress")
	}

	return progress, nil
}
