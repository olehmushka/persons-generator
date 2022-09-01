package religion

import "persons_generator/engine/entities/culture"

type Preference struct {
	Cultures []*culture.Culture `json:"cultures"`
	Amount   int                `json:"amount"`
}
