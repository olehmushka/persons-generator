package world

import "persons_generator/entities"

func FillWorld(w *entities.World) *entities.World {
	size := w.Size
	w.Locations = make([][]*entities.Location, 0, size)
	for y := 0; y < size; y++ {
		w.Locations = append(w.Locations, make([]*entities.Location, 0, size))
		for x := 0; x < size; x++ {
			w.Locations[y] = append(w.Locations[y], &entities.Location{
				Coordinates: &entities.Coordinate{X: x, Y: y},
				Population:  make([]*entities.Human, 0),
			})
		}
	}

	return w
}
