package cli

import (
	"encoding/json"
	"fmt"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/orchestrator"
)

const RunGenerateWorldCommand = "generate_world"

func runGenerateWorldCommand() error {
	o, err := orchestrator.New(orchestrator.Config{
		StorageFolderName: "tmp",
		RedisURL:          "redis://localhost:6379",
	})
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
	w, err := o.CreateWorld(5, c, []*religion.Religion{r}, []*religion.CultureReference{{
		Religion: r,
		Culture:  c[0],
	}})
	if err != nil {
		return err
	}
	if err := o.RunAndSaveWorld(w, 150); err != nil {
		return err
	}
	persons, err := o.QueryPersons(w.ID, orchestrator.PersonsQuery{})
	if err != nil {
		return err
	}
	for i, p := range persons {
		b, err := json.Marshal(p)
		if err != nil {
			return err
		}
		fmt.Printf("#%d\n%s\n\n", i, string(b))
	}
	fmt.Printf("count = %d\n\n", len(persons))
	// ps, err := w.QueryPerson(world.QueryPersonsOpts{PersonsCount: 10})
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("count = %d\n\n", len(ps))
	// for _, p := range ps {
	// 	p.Print()
	// }

	return nil
}
