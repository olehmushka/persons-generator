package json_storage

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	js "persons_generator/core/storage/json_storage"
	"persons_generator/internal/culture/entities"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

type adapter struct {
	storage js.Storage
}

func New(storage js.Storage) Adapter {
	return &adapter{storage: storage}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) SaveCulture(ctx context.Context, c *entities.Culture) error {
	if c == nil {
		return errors.New("can not save nil culture")
	}
	b, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	return a.storage.Store(ctx, getFilename(c.ID), b)
}

func (a *adapter) GetCultureByID(ctx context.Context, id uuid.UUID) (*entities.Culture, error) {
	if id == uuid.Nil {
		return nil, errors.New("can not get culture by nil uuid")
	}
	b, err := a.storage.Get(ctx, getFilename(id))
	if err != nil {
		return nil, err
	}
	var c *entities.Culture
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	return c, nil
}

func getFilename(id uuid.UUID) string {
	return strings.Join([]string{CultureFilePrefix, id.String()}, "_") + ".json"
}
