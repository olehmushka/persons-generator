package cli

import (
	"persons_generator/engine/orchestrator"
)

const RunGenerateReligionCommand = "generate_religion"

func runGenerateReligionCommand() error {
	o, err := orchestrator.New(orchestrator.Config{
		RedisURL: "redis://localhost:6379",

		MongoDBURL:              "mongodb://localhost:27017",
		MongoDBUsername:         "root",
		MongoDBPassword:         "rootPassword",
		MongoDBMaxBulkItemsSize: 500,
		MongoDBDBName:           "persons_generator_db",
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
	r.Print()

	return nil
}
