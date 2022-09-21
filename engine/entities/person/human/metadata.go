package human

type Metadata struct {
	ChildrenGivenBirth int `json:"children_given_birth"`
}

func (m Metadata) GetFertilityCoef() float64 {
	switch {
	case m.ChildrenGivenBirth == 0:
		return 0.98
	case 1 <= m.ChildrenGivenBirth && m.ChildrenGivenBirth < 3:
		return 0.99
	case 3 <= m.ChildrenGivenBirth && m.ChildrenGivenBirth < 5:
		return 0.97
	case 5 <= m.ChildrenGivenBirth && m.ChildrenGivenBirth < 8:
		return 0.74
	case 8 <= m.ChildrenGivenBirth && m.ChildrenGivenBirth < 12:
		return 0.49
	case 12 <= m.ChildrenGivenBirth && m.ChildrenGivenBirth < 16:
		return 0.25
	case 16 <= m.ChildrenGivenBirth:
		return 0
	}
	return 0
}
