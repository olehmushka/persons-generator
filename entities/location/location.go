package location

import (
	"persons_generator/entities/coordinate"
	"persons_generator/entities/human"
)

type Location struct {
	Coordinate *coordinate.Coordinate
	Population []*human.Human
}
