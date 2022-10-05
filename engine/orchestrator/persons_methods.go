package orchestrator

import (
	"context"
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
