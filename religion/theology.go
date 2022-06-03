package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
)

func generateTheology(r *entities.Religion) {
	r.Theology = &religion.Theology{}
	generatePrecepts(r)
	generateHolyScripture(r)
}
