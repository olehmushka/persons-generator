package json_storage

import (
	"context"
	"io/ioutil"
	"strings"

	"persons_generator/config"

	"go.uber.org/fx"
)

type storage struct {
	storageFolder string
}

func New(cfg *config.Config) Storage {
	return &storage{
		storageFolder: cfg.JSONStorage.StorageFolder,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *storage) Get(ctx context.Context, filename string) ([]byte, error) {
	return ioutil.ReadFile(s.getFullFilename(filename))
}

func (s *storage) Store(ctx context.Context, filename string, file []byte) error {
	return ioutil.WriteFile(s.getFullFilename(filename), file, 0o644)
}

func (s *storage) getFullFilename(filename string) string {
	return strings.Join([]string{s.storageFolder, filename}, "/")
}
