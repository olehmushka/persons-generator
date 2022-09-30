package orchestrator

import (
	"context"
	"errors"

	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"

	"github.com/google/uuid"
)

func (o *Orchestrator) CreateCultures(amount int, preferred []*culture.Preference) ([]*culture.Culture, error) {
	return culture.NewMany(culture.Config{StorageFolderName: o.storageFolderName}, amount, preferred)
}

func (o *Orchestrator) GetCultureByID(id uuid.UUID) (*culture.Culture, error) {
	return culture.ReadByID(o.storageFolderName, id)
}

func (o *Orchestrator) SearchCultures(search string) ([]*culture.Culture, error) {
	if search == "" {
		return culture.AllCultures, nil
	}

	return culture.GetCulturesByName(search, culture.AllCultures), nil
}

func (o *Orchestrator) HybridCultures(cultures []*culture.Culture) (*culture.Culture, error) {
	if len(cultures) == 0 {
		return nil, errors.New("base cultures can not be empty")
	}

	return culture.NewWithProto(culture.Config{StorageFolderName: o.storageFolderName}, cultures)
}

func (o *Orchestrator) SaveCulture(ctx context.Context, c *culture.Culture) error {
	if _, err := o.mongodb.InsertOne(ctx, o.dbName, CulturesCollName, c); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not save culture")
	}

	return nil
}
