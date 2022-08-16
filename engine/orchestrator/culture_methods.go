package orchestrator

import (
	"fmt"

	"persons_generator/engine/entities/culture"
)

func (o *Orchestrator) ShowCultures() {
	fmt.Println()
	for _, c := range o.cultures {
		c.Print()
	}
	o.w.PrintLocationCultures()
	fmt.Println()
}

func (o *Orchestrator) CreateCultures(amount int, preferred []*culture.Preference) ([]*culture.Culture, error) {
	return culture.NewCultures(amount, preferred)
}

func (o *Orchestrator) SearchCultures(search string) ([]*culture.Culture, error) {
	if search == "" {
		return culture.AllCultures, nil
	}

	return culture.GetCulturesByName(search, culture.AllCultures), nil
}
