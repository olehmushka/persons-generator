package location

import (
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person"
	"persons_generator/engine/entities/religion"
)

type Location struct {
	Coordinate *coordinate.Coordinate
	Population []*person.Person

	InitCulture  *culture.Culture
	InitReligion *religion.Religion
}
