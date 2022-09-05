package http

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/internal/culture/entities"
	religionEntities "persons_generator/internal/religion/entities"

	"github.com/google/uuid"
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
	cultureIDs := make([]uuid.UUID, 0, len(in.CultureIDs))
	for _, id := range in.CultureIDs {
		cultureID, err := uuid.Parse(id)
		if err != nil {
			return nil, wrapped_error.NewBadRequestError(err, fmt.Sprintf("can not parse culture_id (input=%s)", id))
		}
		cultureIDs = append(cultureIDs, cultureID)
	}

	return &religionEntities.Preference{
		CultureIDs: cultureIDs,
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
