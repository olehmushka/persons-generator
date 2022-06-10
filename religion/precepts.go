package religion

import (
	"fmt"

	"persons_generator/entities"
	"persons_generator/entities/religion"
)

func generatePrecepts(r *entities.Religion) {
	fmt.Println("[generatePrecepts] started")
	defer fmt.Println("[generatePrecepts] finished")

	r.Theology.Precepts = &religion.Precepts{}
	generateSinsAndVirtues(r)
	generateMarriagePercepts(r)
	generateRituals(r)
}
