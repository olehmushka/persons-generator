package world

import (
	"persons_generator/engine/entities/religion"
)

type SerializedPreWorld struct {
	ID                        string                                 `json:"id" bson:"id"`
	Size                      int                                    `json:"size" bson:"size"`
	MaxPersonsNumberPerLoc    int                                    `json:"max_persons_number_per_loc" bson:"max_persons_number_per_loc"`
	MaxDistanceValue          float64                                `json:"max_distance_value" bson:"max_distance_value"`
	Year                      int                                    `json:"year" bson:"year"`
	CulturesIDs               []string                               `json:"cultures_ids" bson:"cultures_ids"`
	ReligionsIDs              []string                               `json:"religions_ids" bson:"religions_ids"`
	CultureReligionReferences []*religion.SerializedCultureReference `json:"culture_religion_references" bson:"culture_religion_references"`
}

func (w *World) SerializePreworld() *SerializedPreWorld {
	culturesIDs := make([]string, len(w.Cultures))
	for i := range culturesIDs {
		culturesIDs[i] = w.Cultures[i].ID
	}
	religionsIDs := make([]string, len(w.Religions))
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

type SerializedWorld struct {
	ID                        string                                 `json:"id" bson:"id"`
	Size                      int                                    `json:"size" bson:"size"`
	MaxPersonsNumberPerLoc    int                                    `json:"max_persons_number_per_loc" bson:"max_persons_number_per_loc"`
	MaxDistanceValue          float64                                `json:"max_distance_value" bson:"max_distance_value"`
	Year                      int                                    `json:"year" bson:"year"`
	CulturesIDs               []string                               `json:"cultures_ids" bson:"cultures_ids"`
	ReligionsIDs              []string                               `json:"religions_ids" bson:"religions_ids"`
	CultureReligionReferences []*religion.SerializedCultureReference `json:"culture_religion_references" bson:"culture_religion_references"`
	PopulationNumber          int                                    `json:"population_number" bson:"population_number"`
	DeadPopulationNumber      int                                    `json:"dead_population_number" bson:"dead_population_number"`
	Duration                  string                                 `json:"duration" bson:"duration"`
}
