package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
)

func generateHolyScripture(r *entities.Religion) {
	r.Theology.HolyScripture = &religion.HolyScripture{}
}
