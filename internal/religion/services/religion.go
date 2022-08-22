package services

import (
	"context"
	"fmt"

	cultureServices "persons_generator/internal/culture/services"
	"persons_generator/internal/religion/entities"
)

type religion struct {
	cultureSrv cultureServices.Culture
}

func New(cultureSrv cultureServices.Culture) Religion {
	return &religion{
		cultureSrv: cultureSrv,
	}
}

func (s *religion) CreateReligions(ctx context.Context, amount int, preferences []*entities.Preference) ([]*entities.Religion, error) {
	if amount < len(preferences) {
		return nil, fmt.Errorf("amount (%d) can not be less than preferences number (%d)", amount, len(preferences))
	}
	rs := make([]*entities.Religion, 0, amount)
	// for _, pref := range preferences {
	// 	//
	// }

	return rs, nil
}

func (s *religion) CreateReligionByOriginalCulture(ctx context.Context, oc []byte) (*entities.Religion, error) {
	return nil, nil
}
