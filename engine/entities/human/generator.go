package human

import (
	"persons_generator/engine/entities/culture"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/human/human"
	"persons_generator/engine/entities/human/presets"
	"persons_generator/engine/entities/religion"
)

func GenerateHuman() (*human.Human, error) {
	c, err := culture.New(&culture.Preference{Names: []string{"ruthenian"}})
	if err != nil {
		return nil, err
	}
	r, err := religion.NewReligion(c)
	if err != nil {
		return nil, err
	}
	sex := g.GetRandomSex()
	h, err := human.NewHuman(c, r, sex, presets.SlavicHumanPreset)
	if err != nil {
		return nil, err
	}

	return h, nil
}
