package orchestrator

import (
	"encoding/json"
	"fmt"
	js "persons_generator/core/storage/json_storage"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/person"
	"persons_generator/engine/entities/world"
	"sync"

	"github.com/google/uuid"
)

func (o *Orchestrator) GetHumans(world *world.World, offset, limit int) ([]*person.Person, error) {
	return nil, nil
}

type DirectPersonsQuery struct {
	MinAge *int `json:"min_age"`
	MaxAge *int `json:"max_age"`
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
}

type PersonsQuery struct {
	IsAlive       *bool                    `json:"is_alive"`
	Coordinates   []*coordinate.Coordinate `json:"coordinates"`
	DirectPersons DirectPersonsQuery       `json:"direct_persons"`
}

func (o *Orchestrator) QueryPersons(worldID uuid.UUID, q PersonsQuery) ([]*person.SerializedPerson, error) {
	if worldID == uuid.Nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not query for <nil> world_id")
	}

	filenames, err := world.GetWorldDirFilenames(o.storageFolderName, worldID)
	if err != nil {
		return nil, wrapped_error.NewNotFoundError(err, "can not get world dir filenames")
	}
	// loc_4871eaa7-9de9-408f-b7e3-5295871cdd61_y7_x0.json
	filenamesToSearch := make([]*world.Filename, 0, len(filenames))
	for _, f := range filenames {
		ok := true
		if q.IsAlive != nil {
			ok = f.IsAlive == *q.IsAlive
		}
		if !ok {
			continue
		}
		if len(q.Coordinates) > 0 {
			ok = coordinate.InContain(q.Coordinates, f.Coordinate)
		}

		if ok {
			filenamesToSearch = append(filenamesToSearch, f)
		}
	}

	storage := js.New(js.Config{StorageFolderName: o.storageFolderName})
	m := tools.NewConcurrencyMap[*person.SerializedPerson](len(filenamesToSearch))
	var wg sync.WaitGroup
	wg.Add(len(filenamesToSearch))
	errCh := make(chan error)
	for _, f := range filenamesToSearch {
		go func() {
			filename := fmt.Sprintf("%s/%s", world.GetDirname(worldID), f.Filename)
			b, err := storage.Get(filename)
			if err != nil {
				errCh <- err
				return
			}
			var loc location.SerializedLocation
			if err := json.Unmarshal(b, &loc); err != nil {
				errCh <- err
				return
			}
			m.Append(loc.Population...)
			wg.Done()
		}()
		select {
		case err := <-errCh:
			return nil, err
		default:
		}
	}
	select {
	case err := <-errCh:
		return nil, err
	default:
	}
	wg.Wait()

	return queryPersons(m.Get(), q.DirectPersons)
}

func queryPersons(list []*person.SerializedPerson, q DirectPersonsQuery) ([]*person.SerializedPerson, error) {
	var (
		minAge = 0
		maxAge = 200
	)
	if q.MinAge != nil {
		minAge = *q.MinAge
	}
	if q.MaxAge != nil {
		maxAge = *q.MaxAge
	}
	if minAge < maxAge {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("min age can not be greater than max age (min_age=%d, max_age=%d)", minAge, maxAge))
	}
	filtered := make([]*person.SerializedPerson, 0, len(list))

	for _, p := range list {
		if minAge <= p.Human.Age && p.Human.Age <= maxAge {
			filtered = append(filtered, p)
		}
	}

	return tools.Paginate(filtered, q.Offset, q.Limit), nil
}
