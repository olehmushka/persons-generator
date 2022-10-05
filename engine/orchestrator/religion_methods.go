package orchestrator

import (
	"context"
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"

	"go.mongodb.org/mongo-driver/bson"
)

func (o *Orchestrator) CreateReligion(c *culture.Culture) (*religion.Religion, error) {
	return religion.New(religion.Config{}, c)
}

func (o *Orchestrator) CreateReligions(amount int, preferred []*religion.Preference) ([]*religion.Religion, error) {
	return religion.NewManyByPreferred(religion.Config{}, amount, preferred)
}

func (o *Orchestrator) SaveReligion(ctx context.Context, r *religion.Religion) error {
	if _, err := o.mongodb.InsertOne(ctx, o.dbName, ReligionsCollName, r); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not save religion")
	}

	return nil
}

func (o *Orchestrator) ReadReligionByID(ctx context.Context, id string) (*religion.Religion, error) {
	filter := bson.M{
		"id": id,
	}
	result, err := o.mongodb.FindOne(ctx, o.dbName, ReligionsCollName, filter)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not read religion by id")
	}
	if result == nil {
		return nil, wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not find religion by id (id=%s)", id))
	}

	var out religion.Religion
	if err := result.Decode(&out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not decode culture")
	}

	return &out, nil
}

func (o *Orchestrator) DeleteReligionByID(ctx context.Context, id string) error {
	filter := bson.M{
		"id": id,
	}
	if _, err := o.mongodb.DeleteOne(ctx, o.dbName, ReligionsCollName, filter); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete religion by id")
	}

	return nil
}

func (o *Orchestrator) DeleteAllReligions(ctx context.Context) error {
	if err := o.mongodb.Truncate(ctx, o.dbName, ReligionsCollName); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete all religions")
	}

	return nil
}
