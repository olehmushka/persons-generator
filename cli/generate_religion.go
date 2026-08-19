package cli

import (
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
)

const RunGenerateReligionCommand = "generate_religion"

// runGenerateReligionCommand generates a culture and a matching religion
// entirely in memory. Both culture.NewMany and religion.New are pure,
// DB-free functions, so this command needs no MongoDB/Redis connection.
func runGenerateReligionCommand() error {
	c, err := culture.NewMany(culture.Config{}, 1, nil)
	if err != nil {
		return err
	}
	r, err := religion.New(religion.Config{}, c[0])
	if err != nil {
		return err
	}
	r.Print()

	return nil
}
