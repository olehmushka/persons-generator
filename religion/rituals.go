package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
)

func generateRituals(r *entities.Religion) {
	r.Theology.Precepts.Rituals = &religion.Rituals{}
}
