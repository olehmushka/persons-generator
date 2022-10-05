package orchestrator

import (
	"context"
	"encoding/json"
	"fmt"
	"persons_generator/core/storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/entities/world"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (o *Orchestrator) CreateWorld(
	size int,
	cultures []*culture.Culture,
	religions []*religion.Religion,
	refs []*religion.CultureReference,
) (*world.World, error) {
	return world.New(world.Config{}, size, cultures, religions, refs)
}

func getWorldRunningProgressChannelName(worldID uuid.UUID) string {
	return fmt.Sprintf("w_%s", worldID.String())
}

func (o *Orchestrator) RunAndSaveWorld(ctx context.Context, w *world.World, stopYear int) error {
	if w == nil {
		return nil
	}

	progressCh := make(chan world.ProgressRunWorld)
	startDate := time.Now()
	errCh := make(chan error)
	var isFinished bool
	go w.RunWorld(stopYear, progressCh, errCh)
	for {
		select {
		case err := <-errCh:
			if err != nil {
				close(progressCh)
				return err
			}
			isFinished = true
			break
		case progress := <-progressCh:
			bProgres, err := json.Marshal(progress)
			if err != nil {
				close(progressCh)
				return wrapped_error.NewInternalServerError(err, "can not marshal progress before setting to redis")
			}
			if err := o.storage.Set(
				context.Background(),
				getWorldRunningProgressChannelName(w.ID),
				string(bProgres),
				time.Hour,
			); err != nil {
				close(progressCh)
				return wrapped_error.NewInternalServerError(err, "can not set world running progress")
			}

			if progress.Progress == 100 {
				isFinished = true

				break
			}
		}
		if isFinished {
			break
		}
	}
	close(progressCh)

	// return w.Save(time.Since(startDate))
	return o.SaveWorld(ctx, w, time.Since(startDate))
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

func (o *Orchestrator) PresaveWorld(ctx context.Context, w *world.World) error {
	if _, err := o.mongodb.InsertOne(ctx, o.dbName, WorldsCollName, w.SerializePreworld()); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not save preworld")
	}

	return nil
}

func (o *Orchestrator) SaveWorld(ctx context.Context, w *world.World, duration time.Duration) error {
	filter := bson.M{
		"id": w.ID,
	}
	update := bson.M{
		"$set": bson.M{
			"duration":               duration.String(),
			"population_number":      w.PopulationNumber,
			"dead_population_number": w.DeadPopulationNumber,
		},
	}
	if _, err := o.mongodb.UpdateOne(ctx, o.dbName, WorldsCollName, filter, update); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not update world")
	}

	if err := o.SaveWorldPopulation(ctx, w); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not save world population")
	}

	return nil
}

func (o *Orchestrator) SaveWorldPopulation(ctx context.Context, w *world.World) error {
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not run year for <nil> location (x: %d, y: %d)", x, y))
			}

			if err := o.SavePersons(ctx, w.ID, w.Locations[y][x].Population); err != nil {
				return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not save population (x: %d, y: %d)", x, y))
			}
			if err := o.SavePersons(ctx, w.ID, w.DeathWorldLocations[y][x].Population); err != nil {
				return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not save dead population (x: %d, y: %d)", x, y))
			}
		}
	}

	return nil
}

func (o *Orchestrator) DeleteWorldByID(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{
		"id": id,
	}
	if _, err := o.mongodb.DeleteOne(ctx, o.dbName, WorldsCollName, filter); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete world by id")
	}

	populationFilter := bson.M{
		"world_id": id,
	}
	if _, err := o.mongodb.DeleteOne(ctx, o.dbName, PersonsCollName, populationFilter); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete persons by world_id")
	}

	return nil
}

func (o *Orchestrator) DeleteAllWorlds(ctx context.Context) error {
	if err := o.mongodb.Truncate(ctx, o.dbName, WorldsCollName); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete all worlds")
	}

	return nil
}

func (o *Orchestrator) ReadWorlds(ctx context.Context, opts storage.PaginationSortingOpts) ([]*world.SerializedWorld, error) {
	findOpt := parseFindOpts(opts)
	cursor, err := o.mongodb.Find(ctx, o.dbName, WorldsCollName, bson.M{}, findOpt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*world.SerializedWorld
	for cursor.Next(ctx) {
		var elem world.SerializedWorld

		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (o *Orchestrator) CountWorlds(ctx context.Context, opts storage.PaginationSortingOpts) (int, error) {
	count, err := o.mongodb.CountDocuments(ctx, o.dbName, WorldsCollName, bson.M{})
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, "can not count worlds")
	}

	return count, nil
}
