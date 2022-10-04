package orchestrator

import (
	"context"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/person"
)

func (o *Orchestrator) SavePersons(ctx context.Context, ps []*person.Person) error {
	chunks := tools.Chunk(100, ps)
	for i := 0; i < len(chunks); i++ {
		if _, err := o.mongodb.InsertMany(ctx, o.dbName, PersonsCollName, PreparePopulationBeforeSaving(chunks[i])); err != nil {
			return wrapped_error.NewInternalServerError(err, "can not insert sevaral persons to db")
		}
	}
	return nil
}

func PreparePopulationBeforeSaving(people []*person.Person) []any {
	return tools.SliceToAnySlice(person.SerializePeople(people))
}

func (o *Orchestrator) DeleteAllPersons(ctx context.Context) error {
	if err := o.mongodb.Truncate(ctx, o.dbName, PersonsCollName); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not delete all persons")
	}

	return nil
}
