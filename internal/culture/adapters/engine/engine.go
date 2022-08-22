package engine

import (
	"context"
	"encoding/json"
	"errors"

	"persons_generator/core/tools"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/orchestrator"
	"persons_generator/internal/culture/entities"

	"go.uber.org/fx"
)

type adapter struct {
	engine *orchestrator.Orchestrator
}

func New() (Adapter, error) {
	e, err := orchestrator.New()
	if err != nil {
		return nil, err
	}

	return &adapter{engine: e}, nil
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*entities.Culture, error) {
	c, err := a.engine.CreateCultures(amount, deserializeCulturePreferences(preferred))
	if err != nil {
		return nil, err
	}

	return serializeCultures(c), nil
}

func (a *adapter) GetProtoCultures(ctx context.Context, q string, limit, offset int) ([]*entities.Culture, int, error) {
	c, err := a.engine.SearchCultures(q)
	if err != nil {
		return nil, 0, err
	}
	if len(c) == 0 {
		return []*entities.Culture{}, 0, nil
	}

	return serializeCultures(tools.Paginate(c, offset, limit)), len(c), nil
}

func (a *adapter) HybridCulture(ctx context.Context, cultures []*entities.Culture) (*entities.Culture, error) {
	originalCultures := make([]*culture.Culture, 0, len(cultures))
	for _, c := range cultures {
		if c == nil {
			return nil, errors.New("base culture for hybrid can not be nil")
		}
		b, err := a.GetOriginalCulture(ctx, c)
		if err != nil {
			return nil, err
		}
		var oc culture.Culture
		if err := json.Unmarshal(b, &oc); err != nil {
			return nil, err
		}
		originalCultures = append(originalCultures, &oc)
	}

	ohc, err := a.engine.HybridCultures(originalCultures)
	if err != nil {
		return nil, err
	}

	return serializeCulture(ohc), nil
}

func (a *adapter) GetOriginalCulture(ctx context.Context, c *entities.Culture) ([]byte, error) {
	if c == nil {
		return nil, errors.New("can not get original culture from nil")
	}
	oc := &culture.Culture{}

	return json.Marshal(oc)
}
