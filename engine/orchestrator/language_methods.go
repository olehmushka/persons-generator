package orchestrator

import (
	"context"
	"fmt"
	"persons_generator/core/storage"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
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

func (o *Orchestrator) FindDefaultLanguage(id, name string) *language.Language {
	if id == "" && name == "" {
		return nil
	}

	return language.FindLanguage(id, name)
}

func (o *Orchestrator) GetDefaultLanguageSubfamilies(opts storage.PaginationSortingOpts) []*language.Subfamily {
	return tools.Paginate(language.AllSubfamilies, opts.Pagination.Offset, opts.Pagination.Limit)
}

func (o *Orchestrator) CountDefaultLanguageSubfamilies() (int, error) {
	return len(language.AllLanguages), nil
}

func (o *Orchestrator) GetRandomWordByLanguageID(ctx context.Context, id string) (string, error) {
	lang, err := o.ReadLanguageByID(ctx, id)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not find language by id (id=%s) for random word", id))
	}
	if lang == nil {
		return "", wrapped_error.NewNotFoundError(err, fmt.Sprintf("not found language by id (id=%s) for random word", id))
	}

	word, err := lang.GetWord()
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get random word by language (lang_name=%s)", lang.Name))
	}

	return word, nil
}

func (o *Orchestrator) GetRandomOwnNameByLanguageID(ctx context.Context, id string, sex gender.Sex) (string, error) {
	lang, err := o.ReadLanguageByID(ctx, id)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not find language by id (id=%s) for random own_name", id))
	}
	if lang == nil {
		return "", wrapped_error.NewNotFoundError(err, fmt.Sprintf("not found language by id (id=%s) for random own_name", id))
	}

	ownName, err := lang.GetOwnName(sex)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get random own_name by language (lang_name=%s)", lang.Name))
	}

	return ownName, nil
}

func (o *Orchestrator) GetRandomCultureNameByLanguageID(ctx context.Context, id string) (string, error) {
	lang, err := o.ReadLanguageByID(ctx, id)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not find language by id (id=%s) for random culture_name", id))
	}
	if lang == nil {
		return "", wrapped_error.NewNotFoundError(err, fmt.Sprintf("not found language by id (id=%s) for random culture_name", id))
	}

	name, err := lang.GetCultureName()
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get random culture_name by language (lang_name=%s)", lang.Name))
	}

	return name, nil
}

func (o *Orchestrator) GetRandomReligionNameByLanguageID(ctx context.Context, id string) (string, error) {
	lang, err := o.ReadLanguageByID(ctx, id)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not find language by id (id=%s) for random religion_name", id))
	}
	if lang == nil {
		return "", wrapped_error.NewNotFoundError(err, fmt.Sprintf("not found language by id (id=%s) for random religion_name", id))
	}

	name, err := lang.GetReligionName()
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get random religion_name by language (lang_name=%s)", lang.Name))
	}

	return name, nil
}

func (o *Orchestrator) CreateLanguage(ctx context.Context, in *language.Language) (string, error) {
	id := uuid.New().String()
	in.ID = id
	if _, err := o.mongodb.InsertOne(ctx, o.dbName, LanguagesCollName, in); err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not insert sevaral persons to db")
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

func (o *Orchestrator) ReadLanguageByID(ctx context.Context, id string) (*language.Language, error) {
	return o.ReadLanguage(ctx, id, "")
}

func (o *Orchestrator) ReadLanguage(ctx context.Context, id, name string) (*language.Language, error) {
	if id == "" && name == "" {
		return nil, nil
	}

	if lang := o.FindDefaultLanguage(id, name); lang != nil {
		return lang, nil
	}

	filter := bson.M{}
	if id != "" {
		filter["id"] = id
	}
	if name != "" {
		filter["name"] = name
	}
	result, err := o.mongodb.FindOne(ctx, o.dbName, LanguagesCollName, filter)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not read language by id")
	}
	if result == nil {
		return nil, wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not find language by id (id=%s)", id))
	}

	var out language.Language
	if err := result.Decode(&out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not decode language")
	}

	return &out, nil
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

func (o *Orchestrator) DeleteLanguageByID(ctx context.Context, id string) error {
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
