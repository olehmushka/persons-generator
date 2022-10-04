package cli

import (
	"context"
	"persons_generator/engine/entities/religion"
	"persons_generator/engine/orchestrator"
)

const RunGenerateWorldCommand = "generate_world"

func runGenerateWorldCommand() error {
	o, err := orchestrator.New(orchestrator.Config{
		StorageFolderName: "tmp",
		RedisURL:          "redis://localhost:6379",

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
	w, err := o.CreateWorld(5, c, []*religion.Religion{r}, []*religion.CultureReference{{
		Religion: r,
		Culture:  c[0],
	}})
	if err != nil {
		return err
	}
	if err := o.RunAndSaveWorld(context.Background(), w, 150); err != nil {
		return err
	}

	return nil
}
