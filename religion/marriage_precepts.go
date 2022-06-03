package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
)

func generateMarriagePercepts(r *entities.Religion) {
	r.Theology.Precepts.Marriage = &religion.Marriage{}
}
