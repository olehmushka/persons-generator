package cli

import (
	"fmt"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/entities/world"
	"persons_generator/engine/orchestrator"
)

const RunGenerateWorldCommand = "generate_world"

func runGenerateWorldCommand() error {
	o, err := orchestrator.New(orchestrator.Config{StorageFolderName: "tmp"})
	if err != nil {
		return err
	}
	c, err := o.CreateCultures(1, nil)
	if err != nil {
		return err
	}
	r, err := o.CreateReligion(c[0])
	if err != nil {
		return err
	}
	w, err := o.CreateWorld(10, c, []*religion.Religion{r}, []*religion.CultureReference{{
		Religion: r,
		Culture:  c[0],
	}})
	if err != nil {
		return err
	}
	if err := w.RunWorld(200); err != nil {
		return err
	}
	ps, err := w.QueryPerson(world.QueryPersonsOpts{PersonsCount: 10})
	if err != nil {
		return err
	}
	fmt.Printf("count = %d\n\n", len(ps))
	for _, p := range ps {
		p.Print()
	}

	return nil
}
