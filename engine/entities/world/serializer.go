package world

import (
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

type SerializedPreWorld struct {
	ID                        uuid.UUID
	Size                      int
	MaxPersonsNumberPerLoc    int
	MaxDistanceValue          float64
	Year                      int
	CulturesIDs               []uuid.UUID
	ReligionsIDs              []uuid.UUID
	CultureReligionReferences []*religion.SerializedCultureReference
}

func (w *World) SerializePreworld() *SerializedPreWorld {
	culturesIDs := make([]uuid.UUID, len(w.Cultures))
	for i := range culturesIDs {
		culturesIDs[i] = w.Cultures[i].ID
	}
	religionsIDs := make([]uuid.UUID, len(w.Religions))
	for i := range religionsIDs {
		religionsIDs[i] = w.Religions[i].ID
	}
	cultureReligionReferences := make([]*religion.SerializedCultureReference, len(w.CultureReligionReferences))
	for i := range cultureReligionReferences {
		cultureReligionReferences[i] = w.CultureReligionReferences[i].Serialize()
	}

	return &SerializedPreWorld{
		ID:                        w.ID,
		Size:                      w.Size,
		MaxPersonsNumberPerLoc:    w.MaxPersonsNumberPerLoc,
		MaxDistanceValue:          w.MaxDistanceValue,
		Year:                      w.Year,
		CulturesIDs:               culturesIDs,
		ReligionsIDs:              religionsIDs,
		CultureReligionReferences: cultureReligionReferences,
	}
}
