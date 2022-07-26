package world

import (
	"persons_generator/entities/coordinate"
	"persons_generator/entities/human"
	"persons_generator/entities/location"
)

type World struct {
	Size      int
	Locations [][]*location.Location
}

func New(s int) *World {
	return &World{
		Size: s,
	}
}

func (w *World) Fill() *World {
	size := w.Size
	w.Locations = make([][]*location.Location, 0, size)
	for y := 0; y < size; y++ {
		w.Locations = append(w.Locations, make([]*location.Location, 0, size))
		for x := 0; x < size; x++ {
			w.Locations[y] = append(w.Locations[y], &location.Location{
				Coordinate: &coordinate.Coordinate{X: x, Y: y},
				Population: make([]*human.Human, 0),
			})
		}
	}

	return w
}
