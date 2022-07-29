package human

import (
	// "strings"

	"persons_generator/entities/culture"
	"persons_generator/entities/religion"
)

type Human struct {
	Name     *ComplexName
	Culture  *culture.Culture
	Religion *religion.Religion

	Body *Body
}

func NewPrimaryHuman(
	name *ComplexName,
	c *culture.Culture,
	r *religion.Religion,
	sex Sex,
) *Human {
	return &Human{
		Name:     name,
		Culture:  c,
		Religion: r,

		Body: &Body{Sex: sex},
	}
}

func (h *Human) GetFullName() string {
	if h == nil || h.Name == nil {
		return ""
	}
	// ns := make([]string, 0, 2)
	// if h.Name.FirstName != nil {
	// 	ns = append(ns, h.Name.FirstName.Name)
	// }
	// if h.Name.LastName != nil {
	// 	ns = append(ns, h.Name.LastName.Name)
	// }

	// return strings.Join(ns, " ")
	return ""
}

func PrimaryGeneration(m Metadata) []*Human {
	var (
		humans          = make([]*Human, m.PrimaryGenerationHumansCount)
		cultureCounter  int
		religionCounter int
	)
	for i := range humans {
		c := m.PrimaryGenerationCultures[cultureCounter]
		sex := GetRandomSex()
		cn := GetRandomComplexName(sex.ToCultureSex(), c)
		r := m.PrimaryGenerationReligions[religionCounter]
		humans[i] = NewPrimaryHuman(cn, c, r, sex)
		cultureCounter++
		if len(m.PrimaryGenerationCultures) == cultureCounter {
			cultureCounter = 0
		}
		religionCounter++
		if len(m.PrimaryGenerationReligions) == religionCounter {
			religionCounter = 0
		}
	}

	return humans
}
