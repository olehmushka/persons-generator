package orchestrator

import (
	"context"
	"errors"
	"fmt"

	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (o *Orchestrator) ReadCultureByID(ctx context.Context, id uuid.UUID) (*culture.Culture, error) {
	filter := bson.M{}
	result, err := o.mongodb.FindOne(ctx, o.dbName, CulturesCollName, filter)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not read culture by id")
	}
	if result == nil {
		return nil, wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not find culture by id (id=%s)", id.String()))
	}

	var out culture.Culture
	if err := result.Decode(&out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not decode culture")
	}

	return &out, nil
}
