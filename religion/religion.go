package religion

import (
	"persons_generator/entities"
)

func New() *entities.Religion {
	r := &entities.Religion{}
	generateDoctrine(r)

	return r
}

func NewMany(n int) []*entities.Religion {
	religions := make([]*entities.Religion, 0, n)
	for i := 0; i < n; i++ {
		religions = append(religions, New())
	}

	return religions
}
