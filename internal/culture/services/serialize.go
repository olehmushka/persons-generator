package services

import (
	engineCulture "persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/language"
	"persons_generator/internal/culture/entities"
)

func serializeCultures(in []*engineCulture.Culture) []*entities.Culture {
	out := make([]*entities.Culture, len(in))
	for i := range out {
		out[i] = serializeCulture(in[i])
	}

	return out
}

func serializeCulture(in *engineCulture.Culture) *entities.Culture {
	if in == nil {
		return nil
	}

	return &entities.Culture{
		Name:            in.Name,
		Proto:           serializeCultures(in.Proto),
		CultureGroup:    serializeWideCulture(in.Abstuct),
		RootCultureName: serializeRoot(in.Root),
		Language:        serializeLanguage(in.Language),
		EthosName:       in.Ethos.Name,
		Traditions:      serializeTraditions(in.Traditions),
		DominantSex:     string(in.GenderDominance.Dominance),
		MartialCustom:   string(in.MartialCustom),
	}
}

func serializeWideCulture(in *engineCulture.AbstructCulture) *entities.CultureGroup {
	if in == nil {
		return nil
	}
	return &entities.CultureGroup{
		Name:            in.Name,
		RootCultureName: in.Root.Name,
	}
}

func serializeRoot(in *engineCulture.Root) string {
	if in == nil {
		return ""
	}
	return in.Name
}

func serializeLanguageSubfamily(in *language.Subfamily) *entities.LanguageSubfamily {
	if in == nil {
		return nil
	}
	return &entities.LanguageSubfamily{
		Name:              in.Name,
		FamilyName:        string(in.Family),
		ExtendedSubfamily: serializeLanguageSubfamily(in.ExtendedSubfamily),
	}
}

func serializeLanguage(in *language.Language) *entities.Language {
	if in == nil {
		return nil
	}

	return &entities.Language{
		Name:      in.Name,
		Subfamily: serializeLanguageSubfamily(in.Subfamily),
	}
}

func serializeTraditions(traditions []*engineCulture.Tradition) []string {
	if len(traditions) == 0 {
		return []string{}
	}

	out := make([]string, len(traditions))
	for i := range out {
		out[i] = traditions[i].Name
	}

	return out
}
