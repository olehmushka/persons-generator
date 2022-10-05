package culture

import (
	"persons_generator/engine/entities/language"
)

type SerializedCulture struct {
	ID              string                       `json:"id"`
	Name            string                       `json:"name"`
	ProtoCultures   []*SerializedCulture         `json:"proto_cultures"`
	AbstructCulture string                       `json:"abstruct_culture"`
	Root            string                       `json:"root"`
	Language        *language.SerializedLanguage `json:"language"`
	Ethos           string                       `json:"ethos"`
	Traditions      []string                     `json:"traditions"`
	GenderDominance string                       `json:"gender_dominance"`
	MartialCustom   string                       `json:"martial_custom"`
}

func (c *Culture) Serialize() *SerializedCulture {
	if c == nil {
		return nil
	}

	protoCultures := make([]*SerializedCulture, len(c.Proto))
	for i := range protoCultures {
		protoCultures[i] = c.Proto[i].Serialize()
	}

	traditions := make([]string, len(c.Traditions))
	for i := range traditions {
		traditions[i] = c.Traditions[i].String()
	}

	return &SerializedCulture{
		ID:              c.ID,
		Name:            c.Name,
		AbstructCulture: c.Abstuct.SerializeName(),
		Root:            c.Root.SerializeName(),
		ProtoCultures:   protoCultures,
		Language:        c.Language.Serialize(),
		Ethos:           c.Ethos.Serialize(),
		Traditions:      traditions,
		GenderDominance: c.GenderDominance.String(),
		MartialCustom:   c.MartialCustom.String(),
	}
}
