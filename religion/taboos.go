package religion

import (
	"persons_generator/entities"
	rel "persons_generator/entities/religion"
)

func generateTaboos(r *entities.Religion) {
	t := &rel.Taboos{}
	r.Taboos = t
}
