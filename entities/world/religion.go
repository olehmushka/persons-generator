package world

import (
	"persons_generator/entities/location"
	"persons_generator/entities/religion"
)

func FillLocationsWithReligions(locations []*location.Location) []*location.Location {
	out := make([]*location.Location, len(locations))
	for i := range out {
		out[i] = &location.Location{
			Coordinate: locations[i].Coordinate,
			Population: locations[i].Population,

			InitCulture:  locations[i].InitCulture,
			InitReligion: religion.NewReligion(locations[i].InitCulture),
		}
	}

	return out
}

func ExtractReligionFromLocations(locations []*location.Location) []*religion.Religion {
	out := make([]*religion.Religion, len(locations))
	for i := range out {
		out[i] = locations[i].InitReligion
	}

	return out
}
