package religion

import (
	"fmt"

	"persons_generator/entities"
)

func New() *entities.Religion {
	fmt.Println("[religion.New] started")
	defer fmt.Println("[religion.New] finished")

	r := &entities.Religion{}
	generateDoctrine(r)
	generateTheology(r)
	generateTaboos(r)
	generateClerics(r)

	return r
}

func NewMany(n int) []*entities.Religion {
	religions := make([]*entities.Religion, 0, n)
	for i := 0; i < n; i++ {
		religions = append(religions, New())
	}

	return religions
}
