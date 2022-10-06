package orchestrator

import (
	"context"
	"errors"
	"fmt"

	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (o *Orchestrator) CreateCultures(amount int, preferred []*culture.Preference) ([]*culture.Culture, error) {
	return culture.NewMany(culture.Config{}, amount, preferred)
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

	return culture.NewWithProto(culture.Config{}, cultures)
}

func (o *Orchestrator) SaveCulture(ctx context.Context, c *culture.Culture) error {
	if _, err := o.mongodb.InsertOne(ctx, o.dbName, CulturesCollName, c); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not save culture")
	}

	return nil
}

func (o *Orchestrator) ReadCultureByID(ctx context.Context, id string) (*culture.Culture, error) {
	filter := bson.M{
		"id": id,
	}
	result, err := o.mongodb.FindOne(ctx, o.dbName, CulturesCollName, filter)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not read culture by id")
	}
	if result == nil {
		return nil, wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not find culture by id (id=%s)", id))
	}

	var out culture.Culture
	if err := result.Decode(&out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not decode culture")
	}

	return &out, nil
}

func (o *Orchestrator) DeleteCultureByID(ctx context.Context, id string) error {
	filter := bson.M{
		"id": id,
	}
	if _, err := o.mongodb.DeleteOne(ctx, o.dbName, CulturesCollName, filter); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete culture by id")
	}

	return nil
}

func (o *Orchestrator) DeleteAllCultures(ctx context.Context) error {
	if err := o.mongodb.Truncate(ctx, o.dbName, CulturesCollName); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete all cultures")
	}

	return nil
}

func (o *Orchestrator) UpdateCultureLanguage(ctx context.Context, cultureID, languageID string) error {
	lang, err := o.ReadLanguageByID(ctx, languageID)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not get language for updating culture")
	}
	if lang == nil {
		return wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not find language by id for updating culture (lang_id=%s)", languageID))
	}

	filter := bson.M{
		"id": cultureID,
	}
	update := bson.M{
		"$set": bson.M{
			"language": lang,
		},
	}
	falseVal := false
	opts := options.Update()
	opts.Upsert = &falseVal
	if _, err := o.mongodb.UpdateOne(ctx, o.dbName, CulturesCollName, filter, update, opts); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not update culture")
	}

	return nil
}
