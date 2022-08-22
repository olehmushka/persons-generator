package engine

import (
	"persons_generator/engine/entities/culture"
	"persons_generator/internal/culture/entities"
)

func deserializeCulturePreferred(in *entities.CulturePreference) *culture.Preference {
	if in == nil {
		return nil
	}

	return &culture.Preference{
		Names:  in.Names,
		Amount: in.Amount,
		Kind:   culture.PreferenceKind(in.Kind),
	}
}

func deserializeCulturePreferences(in []*entities.CulturePreference) []*culture.Preference {
	out := make([]*culture.Preference, len(in))
	for i := range out {
		out[i] = deserializeCulturePreferred(in[i])
	}

	return out
}

func deserializeCultures(in []*entities.Culture) []*culture.Culture {
	out := make([]*culture.Culture, len(in))
	for i := range out {
		out[i] = deserializeCulture(in[i])
	}

	return out
}

func deserializeCulture(in *entities.Culture) *culture.Culture {
	if in == nil {
		return nil
	}

	return &culture.Culture{
		Name:    in.Name,
		Proto:   deserializeCultures(in.Proto),
		Abstuct: deserializeAbstractCulture(in.CultureGroup),
		Root:    deserializeRoot(in.RootCultureName),
	}
}

func deserializeAbstractCulture(in *entities.CultureGroup) *culture.AbstructCulture {
	if in == nil {
		return nil
	}
	return &culture.AbstructCulture{
		Name: in.Name,
		Root: deserializeRoot(in.RootCultureName),
	}
}

func deserializeRoot(in string) *culture.Root {
	if in == "" {
		return nil
	}
	return &culture.Root{Name: in}
}
