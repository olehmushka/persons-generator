package http

import "persons_generator/internal/culture/entities"

func deserializeCultures(cultures []*entities.Culture) []*Culture {
	out := make([]*Culture, len(cultures))
	for i := range out {
		out[i] = deserializeCulture(cultures[i])
	}

	return out
}

func deserializeCulture(in *entities.Culture) *Culture {
	return &Culture{
		Name:            in.Name,
		Proto:           deserializeCultures(in.Proto),
		WideCulture:     deserializeWideCulture(in.WideCulture),
		RootCultureName: in.RootCultureName,
		Language: &Language{
			Name:      in.Language.Name,
			Subfamily: deserializeLanguageSubfamily(in.Language.Subfamily),
		},
		EthosName:     in.EthosName,
		Traditions:    in.Traditions,
		DominantSex:   in.DominantSex,
		MartialCustom: in.MartialCustom,
	}
}

func deserializeWideCulture(in *entities.WideCulture) *WideCulture {
	if in == nil {
		return nil
	}

	return &WideCulture{
		Name:            in.Name,
		RootCultureName: in.RootCultureName,
	}
}

func deserializeLanguageSubfamily(in *entities.LanguageSubfamily) *LanguageSubfamily {
	if in == nil {
		return nil
	}
	return &LanguageSubfamily{
		Name:              in.Name,
		FamilyName:        in.FamilyName,
		ExtendedSubfamily: deserializeLanguageSubfamily(in.ExtendedSubfamily),
	}
}
