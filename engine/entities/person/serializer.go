package person

import (
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person/human"
	"persons_generator/engine/entities/person/traits"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

type SerializedPerson struct {
	ID       uuid.UUID                    `json:"id"`
	OwnName  string                       `json:"own_name"`
	Culture  *culture.SerializedCulture   `json:"culture,omitempty"`
	Religion *religion.SerializedReligion `json:"religion,omitempty"`
	Human    *human.SerializedHuman       `json:"human"`
	Traits   []string                     `json:"traits"`
	Spouces  []*SerializedPerson          `json:"spouces"`
	IsAlive  bool                         `json:"is_alive"`
}

func (p *Person) Serialize() *SerializedPerson {
	if p == nil {
		return nil
	}
	spouces := make([]*SerializedPerson, len(p.Spouces))
	for i := range spouces {
		spouces[i] = p.Spouces[i].Serialize()
	}

	return &SerializedPerson{
		ID:       p.ID,
		OwnName:  p.OwnName,
		Culture:  p.Culture.Serialize(),
		Religion: p.Religion.Serialize(),
		Human:    p.Human.Serialize(),
		Traits:   traits.ExtractTraitNames(p.Traits),
		Spouces:  spouces,
		IsAlive:  p.Chronology.DeathYear == -1,
	}
}
