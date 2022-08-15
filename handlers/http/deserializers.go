package http

import "persons_generator/internal/culture/entities"

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
