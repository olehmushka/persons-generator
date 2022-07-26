package human

import (
	"persons_generator/entities/culture"
)

type ComplexName struct {
	FirstName *culture.Name
	LastName  *culture.Name
}

func GetRandomComplexName(sex culture.Sex, c *culture.Culture) *ComplexName {
	cn := &ComplexName{
		FirstName: culture.GetRandomFirstName(sex, c),
	}
	if c.HasLastName() {
		cn.LastName = culture.GetRandomLastName(c)
	}

	return cn
}
