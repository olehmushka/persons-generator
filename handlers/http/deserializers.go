package http

import (
	"persons_generator/internal/culture/entities"
	languageEntities "persons_generator/internal/language/entities"
	religionEntities "persons_generator/internal/religion/entities"
)

func deserializeCulturePreferred(in *CulturePreferred) *entities.CulturePreference {
	if in == nil {
		return nil
	}

	return &entities.CulturePreference{
		Names:  in.Names,
		Amount: in.Amount,
		Kind:   in.Kind,
	}
}

func deserializeCulturePreferences(in []*CulturePreferred) []*entities.CulturePreference {
	out := make([]*entities.CulturePreference, len(in))
	for i := range out {
		out[i] = deserializeCulturePreferred(in[i])
	}

	return out
}

func deserializeReligionPreferred(in *ReligionPreferred) (*religionEntities.Preference, error) {
	if in == nil {
		return nil, nil
	}

	return &religionEntities.Preference{
		CultureIDs: in.CultureIDs,
		Amount:     in.Amount,
	}, nil
}

func deserializeReligionPreferences(in []*ReligionPreferred) ([]*religionEntities.Preference, error) {
	out := make([]*religionEntities.Preference, 0, len(in))
	for _, i := range in {
		o, err := deserializeReligionPreferred(i)
		if err != nil {
			return nil, err
		}
		out = append(out, o)
	}

	return out, nil
}

func deserializeLanguage(name string, sf *LanguageSubfamily, wb *WordBase, isLiving bool) (*languageEntities.Language, error) {
	var wbOut *languageEntities.WordBase
	if wb != nil {
		wbOut = &languageEntities.WordBase{
			FemaleOwnNames: wb.FemaleOwnNames,
			MaleOwnNames:   wb.MaleOwnNames,
			Words:          wb.Words,
			Name:           wb.Name,
			Min:            wb.Min,
			Max:            wb.Max,
			Dupl:           wb.Dupl,
			M:              wb.M,
		}
	}

	return &languageEntities.Language{
		Name:      name,
		Subfamily: deserializeSubfamily(sf),
		WordBase:  wbOut,
		IsLiving:  isLiving,
	}, nil
}

func deserializeSubfamily(sf *LanguageSubfamily) *languageEntities.Subfamily {
	if sf == nil {
		return nil
	}

	return &languageEntities.Subfamily{
		Name:              sf.Name,
		FamilyName:        sf.FamilyName,
		ExtendedSubfamily: deserializeSubfamily(sf.ExtendedSubfamily),
		IsLiving:          sf.IsLiving,
	}
}
