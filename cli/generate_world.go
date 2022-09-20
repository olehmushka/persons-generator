package cli

import (
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
	w, err := o.CreateWorld(100, c, []*religion.Religion{r}, []*religion.CultureReference{{
		Religion: r,
		Culture:  c[0],
	}})
	if err != nil {
		return err
	}
	ps, err := w.GetPersons(world.GetHumansParams{})
	if err != nil {
		return err
	}
	for _, p := range ps {
		p.Print()
	}

	return nil
}
