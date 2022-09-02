package orchestrator

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

func (o *Orchestrator) ShowReligions() {
	fmt.Println()
	for _, r := range o.religions {
		r.Print()
	}
	o.w.PrintLocationReligions()
	fmt.Println()
}

func (o *Orchestrator) CreateReligion(c *culture.Culture) (*religion.Religion, error) {
	return religion.New(religion.Config{StorageFolderName: o.storageFolderName}, c)
}

func (o *Orchestrator) CreateReligions(amount int, preferred []*religion.Preference) ([]*religion.Religion, error) {
	return religion.NewManyByPreferred(religion.Config{StorageFolderName: o.storageFolderName}, amount, preferred)
}

func (o *Orchestrator) GetReligionByID(id uuid.UUID) (*religion.Religion, error) {
	return religion.ReadByID(o.storageFolderName, id)
}
