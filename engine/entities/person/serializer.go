package person

import (
	"persons_generator/engine/entities/coordinate"
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
	FatherID   uuid.UUID              `json:"father_id" bson:"father_id"`
	MotherID   uuid.UUID              `json:"mother_id" bson:"mother_id"`
	Coordinate coordinate.Coordinate  `json:"coordinate" bson:"coordinate"`
}

func (p *Person) Serialize() *SerializedPerson {
	if p == nil {
		return nil
	}
	spouces := make([]uuid.UUID, len(p.Spouces))
	for i := range spouces {
		spouces[i] = p.Spouces[i].ID
	}

	var cultureID uuid.UUID
	if p.Culture != nil {
		cultureID = p.Culture.ID
	}
	var religionID uuid.UUID
	if p.Religion != nil {
		religionID = p.Religion.ID
	}
	var fatherID uuid.UUID
	if p.Father != nil {
		fatherID = p.Father.ID
	}
	var motherID uuid.UUID
	if p.Mother != nil {
		motherID = p.Mother.ID
	}

	return &SerializedPerson{
		ID:         p.ID,
		OwnName:    p.OwnName,
		CultureID:  cultureID,
		ReligionID: religionID,
		Human:      p.Human.Serialize(),
		Traits:     traits.ExtractTraitNames(p.Traits),
		Spouces:    spouces,
		DeathYear:  p.Chronology.DeathYear,
		FatherID:   fatherID,
		MotherID:   motherID,
		Coordinate: *p.Coordinate,
	}
}

func SerializePeople(in []*Person) []*SerializedPerson {
	out := make([]*SerializedPerson, len(in))
	for i := range out {
		out[i] = in[i].Serialize()
	}

	return out
}
