package orchestrator

import (
	"context"
	"persons_generator/engine/entities/person"
)

func (o *Orchestrator) SavePersons(ctx context.Context, ps []*person.SerializedPerson) error {
	return nil
}
