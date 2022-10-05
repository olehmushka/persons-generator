package orchestrator

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/language"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (o *Orchestrator) QueryDefaultLanguages(q string, opts storage.PaginationSortingOpts) ([]*language.Language, error) {
	return tools.Paginate(language.GetLanguageByName(q), opts.Pagination.Offset, opts.Pagination.Limit), nil
}

func (o *Orchestrator) CountDefaultLanguages(q string) (int, error) {
	if q == "" {
		return len(language.AllLanguages), nil
	}

	return len((language.GetLanguageByName(q))), nil
}

func (o *Orchestrator) GetDefaultLanguages(opts storage.PaginationSortingOpts) []*language.Language {
	return tools.Paginate(language.AllLanguages, opts.Pagination.Offset, opts.Pagination.Limit)
}

func (o *Orchestrator) CreateLanguage(ctx context.Context, in *language.Language) (uuid.UUID, error) {
	id := uuid.New()
	in.ID = id
	if _, err := o.mongodb.InsertOne(ctx, o.dbName, LanguagesCollName, in); err != nil {
		return uuid.Nil, wrapped_error.NewInternalServerError(err, "can not insert sevaral persons to db")
	}

	return id, nil
}

func (o *Orchestrator) ReadLanguagesByName(ctx context.Context, name string, opts storage.PaginationSortingOpts) ([]*language.Language, error) {
	findOpt := parseFindOpts(opts)
	// findOpts.Set
	filters := bson.M{"name": name}
	if name == "" {
		filters = bson.M{}
	}
	cursor, err := o.mongodb.Find(ctx, o.dbName, LanguagesCollName, filters, findOpt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*language.Language
	for cursor.Next(ctx) {
		var elem language.Language

		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (o *Orchestrator) CountLanguagesByName(ctx context.Context, name string) (int, error) {
	filters := bson.M{"name": name}
	if name == "" {
		filters = bson.M{}
	}
	count, err := o.mongodb.CountDocuments(ctx, o.dbName, LanguagesCollName, filters)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, "can not count language by name")
	}

	return count, nil
}

func (o *Orchestrator) DeleteLanguageByID(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{
		"id": id,
	}
	if _, err := o.mongodb.DeleteOne(ctx, o.dbName, LanguagesCollName, filter); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete language by id")
	}

	return nil
}

func (o *Orchestrator) DeleteAllLanguages(ctx context.Context) error {
	if err := o.mongodb.Truncate(ctx, o.dbName, LanguagesCollName); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete all languages")
	}

	return nil
}
