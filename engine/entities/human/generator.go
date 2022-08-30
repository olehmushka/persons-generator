package human

import (
	"persons_generator/engine/entities/culture"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/human/human"
	"persons_generator/engine/entities/human/presets"
	"persons_generator/engine/entities/religion"
)

func GenerateHuman(cfg Config) (*human.Human, error) {
	c, err := culture.New(culture.Config{StorageFolderName: cfg.StorageFolderName}, &culture.Preference{Names: []string{"ruthenian"}})
	if err != nil {
		return nil, err
	}
	r, err := religion.New(religion.Config{StorageFolderName: cfg.StorageFolderName}, c)
	if err != nil {
		return nil, err
	}
	sex, err := g.GetRandomSex()
	if err != nil {
		return nil, err
	}
	h, err := human.NewHuman(c, r, sex, presets.SlavicHumanPreset)
	if err != nil {
		return nil, err
	}

	return h, nil
}
