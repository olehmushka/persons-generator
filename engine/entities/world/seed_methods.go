package world

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/person"
	"persons_generator/engine/entities/religion"
)

func (w *World) seed() error {
	w.seedLocations()
	if err := w.seedCultures(); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not seed cultures")
	}
	if err := w.seedReligions(); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not seed religions")
	}
	if err := w.seedPopulation(); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not seed population")
	}

	return nil
}

func (w *World) seedLocations() {
	size := w.Size
	w.Locations = make([][]*location.Location, 0, size)
	for y := 0; y < size; y++ {
		w.Locations = append(w.Locations, make([]*location.Location, 0, size))
		for x := 0; x < size; x++ {
			w.Locations[y] = append(w.Locations[y], &location.Location{
				Coordinate: &coordinate.Coordinate{X: x, Y: y},
				Population: make([]*person.Person, 0),
			})
		}
	}
}

func (w *World) seedCultures() error {
	totalLocs := w.Size * w.Size
	var (
		locsPerCulture = totalLocs / len(w.Cultures)
		reminder       = totalLocs - (locsPerCulture * len(w.Cultures))
		toFillCultures = make(map[string]int, len(w.Cultures))
	)
	for i, c := range w.Cultures {
		toFillCultures[c.Name] = locsPerCulture
		if i == 0 {
			toFillCultures[c.Name] += reminder
		}
	}

	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			var cultureName string
			for {
				randCultureName, err := culture.GetRandomCultureName(w.Cultures)
				if err != nil {
					return wrapped_error.NewInternalServerError(err, "can not generate random culture_name")
				}
				if rem, ok := toFillCultures[randCultureName]; ok && rem > 0 {
					cultureName = randCultureName
					break
				}
			}
			w.Locations[y][x].InitCulture = culture.GetCultureByName(cultureName, w.Cultures)
			toFillCultures[cultureName]--
		}
	}

	return nil
}

func (w *World) seedReligions() error {
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil || w.Locations[y][x].InitCulture == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("location is nil or its culture is nil (x: %d, y: %d)", x, y))
			}

			cultureName := w.Locations[y][x].InitCulture.Name
			w.Locations[y][x].InitReligion = religion.GetReligionByCultureNameFromCultureReferences(w.CultureReligionReferences, cultureName)
		}
	}

	return nil
}

func (w *World) seedPopulation() error {
	return nil
}
