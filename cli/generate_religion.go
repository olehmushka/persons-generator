package cli

import (
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/orchestrator"
)

const RunGenerateReligionCommand = "generate_religion"

func runGenerateReligionCommand() error {
	o, err := orchestrator.New()
	if err != nil {
		return err
	}
	c, err := o.CreateCultures(1, []*culture.Preference{{Names: []string{"ruthenian"}}})
	if err != nil {
		return err
	}
	r, err := o.CreateReligion(c[0])
	if err != nil {
		return err
	}
	r.Print()

	return nil
}
