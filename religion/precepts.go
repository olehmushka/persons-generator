package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
)

func generatePrecepts(r *entities.Religion) {
	r.Theology.Precepts = &religion.Precepts{}
	generateSinsAndVirtues(r)
	generateMarriagePercepts(r)
	generateRituals(r)
}
