package human

import (
	"persons_generator/entities/culture"
	"persons_generator/entities/religion/religion"
)

type Metadata struct {
	PrimaryGenerationFirstNameDuplicatesNumber int
	PrimaryGenerationCultures                  []*culture.Culture
	PrimaryGenerationHumansCount               int
	PrimaryGenerationReligions                 []*religion.Religion
}
