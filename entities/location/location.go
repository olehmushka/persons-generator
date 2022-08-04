package location

import (
	"persons_generator/entities/coordinate"
	"persons_generator/entities/culture"
	"persons_generator/entities/human"
	"persons_generator/entities/religion"
)

type Location struct {
	Coordinate *coordinate.Coordinate
	Population []*human.Human

	InitCulture  *culture.Culture
	InitReligion *religion.Religion
}
