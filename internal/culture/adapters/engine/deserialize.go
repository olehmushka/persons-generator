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
