package world

import (
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/religion"
)

func FillLocationsWithReligions(cfg Config, locations []*location.Location) ([]*location.Location, error) {
	out := make([]*location.Location, len(locations))
	for i := range out {
		r, err := religion.New(religion.Config{StorageFolderName: cfg.StorageFolderName}, locations[i].InitCulture)
		if err != nil {
			return nil, err
		}
		out[i] = &location.Location{
			Coordinate: locations[i].Coordinate,
			Population: locations[i].Population,

			InitCulture:  locations[i].InitCulture,
			InitReligion: r,
		}
	}

	return out, nil
}

func ExtractReligionFromLocations(locations []*location.Location) []*religion.Religion {
	out := make([]*religion.Religion, len(locations))
	for i := range out {
		out[i] = locations[i].InitReligion
	}

	return out
}
