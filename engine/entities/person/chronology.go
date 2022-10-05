package person

import (
	"fmt"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
)

type Chronology struct {
	BirthYear int     `json:"birth_year"`
	DeathYear int     `json:"death_year"`
	Events    []Event `json:"events"`
}

func (c Chronology) IsLustful() (bool, error) {
	if len(c.Events) == 0 {
		return false, nil
	}

	sexes := make(map[int][]string, len(c.Events))
	for _, e := range c.Events {
		if e.Name != HaveSexEventName {
			continue
		}

		p, isOk := e.Value.(*Person)
		if !isOk {
			return false, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not cast event.value to *Person (value=%+v)", e.Value))
		}
		if _, ok := sexes[e.Year]; ok {
			sexes[e.Year] = append(sexes[e.Year], p.ID)
		} else {
			sexes[e.Year] = []string{p.ID}
		}
	}
	if len(sexes) == 0 {
		return false, nil
	}
	for year, personIDs := range sexes {
		sexes[year] = tools.Unique(personIDs)
		if len(sexes[year]) > 2 {
			return true, nil
		}
	}
	allSexPartners := make([]string, 0, len(sexes))
	for _, partnerIDs := range sexes {
		allSexPartners = append(allSexPartners, partnerIDs...)
	}

	return len(tools.Unique(allSexPartners)) > 5, nil
}
