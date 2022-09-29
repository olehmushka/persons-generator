package json_storage

import (
	"fmt"
	"io/ioutil"
	"os"
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
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get file by filename (filename=%s)", s.getFullFilename(filename)))
	}

	return b, nil
}

func (s *storage) Store(filename string, file []byte) error {
	if err := ioutil.WriteFile(s.getFullFilename(filename), file, 0o644); err != nil {
		return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not save file (filename=%s)", s.getFullFilename(filename)))
	}

	return nil
}

func (s *storage) getFullFilename(filename string) string {
	return strings.Join([]string{s.storageFolderName, filename}, "/")
}

func (s *storage) MkDir(dirname string) error {
	if err := os.Mkdir(s.getFullFilename(dirname), os.ModePerm); err != nil {
		return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not create dir (dirname=%s)", s.getFullFilename(dirname)))
	}
	return nil
}

func (s *storage) IsDirExists(dirname string) bool {
	_, err := os.Stat(s.getFullFilename(dirname))

	return err == nil
}

func (s *storage) GetDirInnerFilenames(dirname string) ([]string, error) {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/", s.getFullFilename(dirname)))
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read from dir (dirname=%s)", dirname))
	}

	out := make([]string, len(files))
	for i, file := range files {
		out[i] = file.Name()
	}

	return out, nil
}
