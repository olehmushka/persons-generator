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
	w.seedDeathWorldLocations()
	if err := w.seedSimilarities(); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not seed similarities")
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

func (w *World) seedDeathWorldLocations() {
	size := w.Size
	w.DeathWorldLocations = make([][]*location.Location, 0, size)
	for y := 0; y < size; y++ {
		w.DeathWorldLocations = append(w.Locations, make([]*location.Location, 0, size))
		for x := 0; x < size; x++ {
			w.DeathWorldLocations[y] = append(w.DeathWorldLocations[y], &location.Location{
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
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil || w.Locations[y][x].InitCulture == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("location is nil or its culture is nil (x: %d, y: %d)", x, y))
			}
			if w.Locations[y][x].InitReligion == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("location is nil or its religion is nil (x: %d, y: %d)", x, y))
			}

			w.Locations[y][x].Population = make([]*person.Person, 0, 10)
			for i := 0; i < 10; i++ {
				p, err := person.NewBase(w.Locations[y][x].InitCulture, w.Locations[y][x].InitReligion, w.Locations[y][x].Coordinate)
				if err != nil {
					return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not generate person for location (x: %d, y: %d)", x, y))
				}
				w.Locations[y][x].Population = append(w.Locations[y][x].Population, p)
			}
		}
	}

	return nil
}

func (w *World) seedSimilarities() error {
	if len(w.Cultures) == 0 {
		return wrapped_error.NewInternalServerError(nil, "can not seed similarities for empty cultures")
	}
	if len(w.Religions) == 0 {
		return wrapped_error.NewInternalServerError(nil, "can not seed similarities for empty religions")
	}

	religionSimilarities := make(map[string]float64, len(w.Religions)*len(w.Religions))
	for _, r1 := range w.Religions {
		for _, r2 := range w.Religions {
			religionSimilarityCoef, err := religion.GetReligionSimilarityCoef(r1, r2)
			if err != nil {
				return wrapped_error.NewInternalServerError(err, "can not get religion similarity coef")
			}
			religionSimilarities[fmt.Sprintf("%s:%s", r1.ID.String(), r2.ID.String())] = religionSimilarityCoef
		}
	}
	w.religionsSimilarity = religionSimilarities

	cultureSimilarities := make(map[string]float64, len(w.Cultures)*len(w.Cultures))
	for _, c1 := range w.Cultures {
		for _, c2 := range w.Cultures {
			cultureSimilarityCoef, err := culture.GetCultureSimilarityCoef(c1, c2)
			if err != nil {
				return wrapped_error.NewInternalServerError(err, "can not get culture similarity coef")
			}
			cultureSimilarities[fmt.Sprintf("%s:%s", c1.ID.String(), c2.ID.String())] = cultureSimilarityCoef
		}
	}
	w.culturesSimilarity = cultureSimilarities

	return nil
}
