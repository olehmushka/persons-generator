package language

type SerializedLanguage struct {
	ID        string                       `json:"id"`
	Name      string                       `json:"name"`
	Subfamily *SerailizedLanguageSubfamily `json:"subfamily"`
}

func (l *Language) Serialize() *SerializedLanguage {
	if l == nil {
		return nil
	}

	return &SerializedLanguage{
		ID:        l.ID,
		Name:      l.Name,
		Subfamily: l.Subfamily.Serialize(),
	}
}

type SerailizedLanguageSubfamily struct {
	Name              string                       `json:"name"`
	Family            string                       `json:"family"`
	ExtendedSubfamily *SerailizedLanguageSubfamily `json:"extended_subfamily"`
}

func (l *Subfamily) Serialize() *SerailizedLanguageSubfamily {
	if l == nil {
		return nil
	}

	return &SerailizedLanguageSubfamily{
		Name:              l.Name,
		Family:            l.Family.String(),
		ExtendedSubfamily: l.ExtendedSubfamily.Serialize(),
	}
}
