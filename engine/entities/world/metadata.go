package world

import (
	"encoding/json"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"time"
)

type metadata struct {
	Timestamp        string                         `json:"timestamp"`
	Size             int                            `json:"size"`
	Year             int                            `json:"year"`
	Cultures         []*culture.SerializedCulture   `json:"cultures"`
	Religions        []*religion.SerializedReligion `json:"religions"`
	PopulationNumber int                            `json:"population_number"`
}

func newMetadata(w *World) *metadata {
	if w == nil {
		return nil
	}

	cultures := make([]*culture.SerializedCulture, len(w.Cultures))
	for i := range cultures {
		cultures[i] = w.Cultures[i].Serialize()
	}
	religions := make([]*religion.SerializedReligion, len(w.Religions))
	for i := range cultures {
		religions[i] = w.Religions[i].Serialize()
	}

	return &metadata{
		Timestamp:        tools.SerializeTime(time.Now()),
		Size:             w.Size,
		Year:             w.Year,
		Cultures:         cultures,
		Religions:        religions,
		PopulationNumber: w.populationNumber,
	}
}

func (m *metadata) Marshal() ([]byte, error) {
	if m == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not marshal <nil> world.metadata")
	}

	b, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not marshal world.metadata")
	}

	return b, nil
}
