package religion

import (
	"fmt"

	"persons_generator/entities"
	"persons_generator/entities/religion"
)

func generateTheology(r *entities.Religion) {
	fmt.Println("[generateTheology] started")
	defer fmt.Println("[generateTheology] finished")

	r.Theology = &religion.Theology{}
	generatePrecepts(r)
	generateHolyScripture(r)
}
