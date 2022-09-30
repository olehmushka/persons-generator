package person

import (
	"persons_generator/engine/entities/person/human"
	"persons_generator/engine/entities/person/traits"

	"github.com/google/uuid"
)

type SerializedPerson struct {
	ID         uuid.UUID              `json:"id" bson:"id"`
	OwnName    string                 `json:"own_name" bson:"own_name"`
	CultureID  uuid.UUID              `json:"culture_id,omitempty" bson:"culture_id,omitempty"`
	ReligionID uuid.UUID              `json:"religion_id,omitempty" bson:"religion_id,omitempty"`
	Human      *human.SerializedHuman `json:"human" bson:"human"`
	Traits     []string               `json:"traits" bson:"traits"`
	Spouces    []uuid.UUID            `json:"spouces" bson:"spouces"`
	DeathYear  int                    `json:"death_year" bson:"death_year"`
}

func (p *Person) Serialize() *SerializedPerson {
	if p == nil {
		return nil
	}
	spouces := make([]uuid.UUID, len(p.Spouces))
	for i := range spouces {
		spouces[i] = p.Spouces[i].ID
	}

	return &SerializedPerson{
		ID:         p.ID,
		OwnName:    p.OwnName,
		CultureID:  p.Culture.ID,
		ReligionID: p.Religion.ID,
		Human:      p.Human.Serialize(),
		Traits:     traits.ExtractTraitNames(p.Traits),
		Spouces:    spouces,
		DeathYear:  p.Chronology.DeathYear,
	}
}
