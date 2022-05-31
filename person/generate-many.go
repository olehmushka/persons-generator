package person

import "persons_generator/person/entities"

type GenerateManyConfig struct {
	Mode string
}

func GenerateMany(cfg *GenerateManyConfig) []*entities.Person {
	return []*entities.Person{}
}
