package person

import (
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/person/human"
	"persons_generator/engine/entities/person/traits"
)

type SerializedPerson struct {
	ID         string                 `json:"id" bson:"id"`
	OwnName    string                 `json:"own_name" bson:"own_name"`
	CultureID  string                 `json:"culture_id,omitempty" bson:"culture_id,omitempty"`
	ReligionID string                 `json:"religion_id,omitempty" bson:"religion_id,omitempty"`
	Human      *human.SerializedHuman `json:"human" bson:"human"`
	Traits     []string               `json:"traits" bson:"traits"`
	Spouces    []string               `json:"spouces" bson:"spouces"`
	DeathYear  int                    `json:"death_year" bson:"death_year"`
	FatherID   string                 `json:"father_id" bson:"father_id"`
	MotherID   string                 `json:"mother_id" bson:"mother_id"`
	Coordinate coordinate.Coordinate  `json:"coordinate" bson:"coordinate"`
	WorldID    string                 `json:"world_id" bson:"world_id"`
}

func (p *Person) Serialize(wID string) *SerializedPerson {
	if p == nil {
		return nil
	}
	spouces := make([]string, len(p.Spouces))
	for i := range spouces {
		spouces[i] = p.Spouces[i].ID
	}

	var cultureID string
	if p.Culture != nil {
		cultureID = p.Culture.ID
	}
	var religionID string
	if p.Religion != nil {
		religionID = p.Religion.ID
	}
	var fatherID string
	if p.Father != nil {
		fatherID = p.Father.ID
	}
	var motherID string
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
		WorldID:    wID,
	}
}

func SerializePeople(wID string, in []*Person) []*SerializedPerson {
	out := make([]*SerializedPerson, len(in))
	for i := range out {
		out[i] = in[i].Serialize(wID)
	}

	return out
}
