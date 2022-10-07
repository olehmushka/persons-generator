package cli

import (
	"context"
	"persons_generator/engine/orchestrator"
)

const RunRefreshDataCommand = "refresh_data"

func runRefreshDataCommand() error {
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
	ctx := context.Background()
	if err := o.RefreshNativeLanguageSubfamilies(ctx); err != nil {
		return err
	}
	if err := o.RefreshNativeLanguages(ctx); err != nil {
		return err
	}
	if err := o.RefreshNativeCultures(ctx); err != nil {
		return err
	}

	return nil
}
