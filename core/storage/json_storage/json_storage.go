package json_storage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"persons_generator/core/wrapped_error"
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
	b, err := ioutil.ReadFile(s.getFullFilename(filename))
	if err != nil {
		return nil, wrapped_error.New(http.StatusInternalServerError, err, fmt.Sprintf("can not get file by filename (filename=%s)", filename))
	}

	return b, nil
}

func (s *storage) Store(filename string, file []byte) error {
	if err := ioutil.WriteFile(s.getFullFilename(filename), file, 0o644); err != nil {
		return wrapped_error.New(http.StatusInternalServerError, err, fmt.Sprintf("can not save file (filename=%s)", filename))
	}

	return nil
}

func (s *storage) getFullFilename(filename string) string {
	return strings.Join([]string{s.storageFolderName, filename}, "/")
}
