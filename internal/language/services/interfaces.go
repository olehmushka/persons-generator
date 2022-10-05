package services

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/internal/language/entities"
)

type Language interface {
	QueryDefaultLanguages(q string, opts storage.PaginationSortingOpts) ([]*entities.Language, int, error)
	CreateLanguage(ctx context.Context, in *entities.Language) (string, error)
	ReadLanguagesByName(ctx context.Context, name string, opts storage.PaginationSortingOpts) ([]*entities.Language, int, error)
	DeleteLanguageByID(ctx context.Context, id string) error
	DeleteAllLanguages(ctx context.Context) error
}
