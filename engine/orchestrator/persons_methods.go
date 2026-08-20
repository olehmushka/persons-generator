package orchestrator

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/person"

	"go.mongodb.org/mongo-driver/bson"
)

func (o *Orchestrator) SavePersons(ctx context.Context, wID string, ps []*person.Person) error {
	chunks := tools.Chunk(100, ps)
	for i := 0; i < len(chunks); i++ {
		if _, err := o.mongodb.InsertMany(ctx, o.dbName, PersonsCollName, PreparePopulationBeforeSaving(wID, chunks[i])); err != nil {
			return wrapped_error.NewInternalServerError(err, "can not insert sevaral persons to db")
		}
	}
	return nil
}

func PreparePopulationBeforeSaving(wID string, people []*person.Person) []any {
	return tools.SliceToAnySlice(person.SerializePeople(wID, people))
}

func (o *Orchestrator) DeletePersonByID(ctx context.Context, id string) error {
	filter := bson.M{
		"id": id,
	}
	if _, err := o.mongodb.DeleteOne(ctx, o.dbName, PersonsCollName, filter); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete person by id")
	}

	return nil
}

func (o *Orchestrator) DeleteAllPersons(ctx context.Context) error {
	if err := o.mongodb.Truncate(ctx, o.dbName, PersonsCollName); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete all persons")
	}

	return nil
}

func (o *Orchestrator) ReadPersonsByWorldID(ctx context.Context, worldID string, opts storage.PaginationSortingOpts) ([]*person.SerializedPerson, error) {
	findOpt := parseFindOpts(opts)
	filter := bson.M{
		"world_id": worldID,
	}
	cursor, err := o.mongodb.Find(ctx, o.dbName, PersonsCollName, filter, findOpt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*person.SerializedPerson
	for cursor.Next(ctx) {
		var elem person.SerializedPerson

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

func (o *Orchestrator) CountPersonsByWorldID(ctx context.Context, worldID string) (int, error) {
	filter := bson.M{
		"world_id": worldID,
	}
	count, err := o.mongodb.CountDocuments(ctx, o.dbName, PersonsCollName, filter)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, "can not count persons by world id")
	}

	return count, nil
}
