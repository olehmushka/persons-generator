package json_storage

import (
	"io/ioutil"
	"strings"

	"go.uber.org/fx"
)

type storage struct {
	storageFolderName string
}

func New(cfg Config) Storage {
	return &storage{
		storageFolderName: cfg.StorageFolderName,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (s *storage) Get(filename string) ([]byte, error) {
	return ioutil.ReadFile(s.getFullFilename(filename))
}

func (s *storage) Store(filename string, file []byte) error {
	return ioutil.WriteFile(s.getFullFilename(filename), file, 0o644)
}

func (s *storage) getFullFilename(filename string) string {
	return strings.Join([]string{s.storageFolderName, filename}, "/")
}
