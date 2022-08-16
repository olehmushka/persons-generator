package http

import "persons_generator/internal/culture/entities"

func serializeCultures(cultures []*entities.Culture) []*Culture {
	out := make([]*Culture, len(cultures))
	for i := range out {
		out[i] = serializeCulture(cultures[i])
	}

	return out
}

func serializeCulture(in *entities.Culture) *Culture {
	return &Culture{
		Name:            in.Name,
		Proto:           serializeCultures(in.Proto),
		CultureGroup:    serializeWideCulture(in.CultureGroup),
		RootCultureName: in.RootCultureName,
		Language: &Language{
			Name:      in.Language.Name,
			Subfamily: serializeLanguageSubfamily(in.Language.Subfamily),
		},
		EthosName:     in.EthosName,
		Traditions:    in.Traditions,
		DominantSex:   in.DominantSex,
		MartialCustom: in.MartialCustom,
	}
}

func serializeWideCulture(in *entities.CultureGroup) *CultureGroup {
	if in == nil {
		return nil
	}

	return &CultureGroup{
		Name:            in.Name,
		RootCultureName: in.RootCultureName,
	}
}

func serializeLanguageSubfamily(in *entities.LanguageSubfamily) *LanguageSubfamily {
	if in == nil {
		return nil
	}
	return &LanguageSubfamily{
		Name:              in.Name,
		FamilyName:        in.FamilyName,
		ExtendedSubfamily: serializeLanguageSubfamily(in.ExtendedSubfamily),
	}
}
