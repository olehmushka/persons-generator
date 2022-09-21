package human

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	pm "persons_generator/engine/probability_machine"
)

func (h *Human) GetAgeFertility() float64 {
	switch {
	case h.Age < 12:
		return 0
	case h.Age >= 12 && h.Age < 16:
		return 1
	case h.Age >= 16 && h.Age < 25:
		return 1
	case h.Age >= 25 && h.Age < 30:
		if h.Sex.IsMale() {
			return 1
		}
		if h.Sex.IsFemale() {
			return 0.9
		}
	case h.Age >= 30 && h.Age < 35:
		if h.Sex.IsMale() {
			return 1
		}
		if h.Sex.IsFemale() {
			return 0.7
		}
	case h.Age >= 35 && h.Age < 40:
		if h.Sex.IsMale() {
			return 0.9
		}
		if h.Sex.IsFemale() {
			return 0.5
		}
	case h.Age >= 40 && h.Age < 45:
		if h.Sex.IsMale() {
			return 0.8
		}
		if h.Sex.IsFemale() {
			return 0.33
		}
	case h.Age >= 45 && h.Age < 50:
		if h.Sex.IsMale() {
			return 0.8
		}
		if h.Sex.IsFemale() {
			return 0
		}
	case h.Age >= 50 && h.Age < 60:
		if h.Sex.IsMale() {
			return 0.7
		}
		if h.Sex.IsFemale() {
			return 0
		}
	case h.Age >= 60 && h.Age < 70:
		if h.Sex.IsMale() {
			return 0.6
		}
		if h.Sex.IsFemale() {
			return 0
		}
	case h.Age >= 70:
		if h.Sex.IsMale() {
			return 0.5
		}
		if h.Sex.IsFemale() {
			return 0
		}
	}

	return 0
}

func (h *Human) ProduceChildren(partner *Human) ([]*Human, error) {
	if h == nil || partner == nil {
		return nil, nil
	}
	if h.Sex.IsMale() {
		return nil, nil
	}

	prob := 0.25 *
		((h.GetAgeFertility() + partner.GetAgeFertility()) / 2) *
		h.Metadata.GetFertilityCoef()
	isPregnant, err := pm.GetRandomBool(prob)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get_pregnant by probability (prob=%f)", prob))
	}
	if !isPregnant {
		return nil, nil
	}
	childrenMultipleBirthName, err := pm.GetRandomFromSeveral(childrenMultipleBirth)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate children multiple birth")
	}
	count, ok := childrenMultipleBirthSet[childrenMultipleBirthName]
	if !ok {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get children count of one pregnance (name=%s)", childrenMultipleBirthName))
	}
	if count < 1 {
		return nil, wrapped_error.NewInternalServerError(nil, "can not generate less than 1 children")
	}

	children := make([]*Human, count)
	for i := range children {
		h, err := h.Pair(partner)
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair humans")
		}
		children[i] = h
	}

	return children, nil
}

func (h *Human) Pair(partner *Human) (*Human, error) {
	sex, err := gender.GetRandomSex()
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not get random sex for human pair")
	}

	g, err := h.Gene.Pair(partner.Gene)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not get random gene for human pair")
	}
	var f, m *Human
	if h.Sex.IsMale() {
		f = h
		m = partner
	} else {
		f = partner
		m = h
	}

	return New(sex, g, f, m)
}

func (h *Human) HaveSex(partner *Human) error {
	if h == nil || partner == nil {
		return nil
	}
	if h.IsPregnant() || partner.IsPregnant() {
		return nil
	}

	children, err := h.ProduceChildren(partner)
	if err != nil {
		return err
	}
	if h.Sex.IsFemale() && partner.Sex.IsMale() {
		h.PreBirthChildren = children
	}
	if h.Sex.IsMale() && partner.Sex.IsFemale() {
		partner.PreBirthChildren = children
	}

	return nil
}

func (h *Human) IsPregnant() bool {
	if h == nil {
		return false
	}

	return len(h.PreBirthChildren) > 0
}

func (h *Human) GiveBirth() []*Human {
	if h == nil {
		return nil
	}
	if len(h.PreBirthChildren) == 0 {
		return nil
	}

	children := h.PreBirthChildren
	h.PreBirthChildren = nil
	h.Metadata.ChildrenGivenBirth += len(children)

	return children
}

func (h *Human) IncreaseAge(years int) int {
	if h == nil {
		return 0
	}

	h.Age += years

	return h.Age
}

func (h *Human) IncrementAge() int {
	if h == nil {
		return 0
	}

	h.Age++

	return h.Age
}
