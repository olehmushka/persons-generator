package human

import "persons_generator/entities"

func New(w *entities.World, father, mother *entities.Human) *entities.Human {
	if father == nil && mother == nil {
		return newWithoutParents()
	}

	return &entities.Human{
		Parents: &entities.Parents{
			Father: father,
			Mother: mother,
		},
	}
}

func newWithoutParents() *entities.Human {
	return &entities.Human{}
}

func MakeOlder(h *entities.Human) *entities.Human {
	h.Body.Age++

	return h
}
