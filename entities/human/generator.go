package human

import (
	"persons_generator/entities/culture"
	g "persons_generator/entities/gender"
	"persons_generator/entities/human/human"
	"persons_generator/entities/human/presets"
	"persons_generator/entities/religion"
)

func GenerateHuman() (*human.Human, error) {
	c, err := culture.New([]string{"ruthenian"})
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
