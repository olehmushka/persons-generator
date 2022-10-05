package engine

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/internal/language/entities"

	"github.com/google/uuid"
)

type Adapter interface {
	QueryDefaultLanguages(q string, opts storage.PaginationSortingOpts) ([]*entities.Language, error)
	CountDefaultLanguages(q string) (int, error)
	CreateLanguage(ctx context.Context, in *entities.Language) (uuid.UUID, error)
	ReadLanguagesByName(ctx context.Context, name string, opts storage.PaginationSortingOpts) ([]*entities.Language, error)
	CountLanguagesByName(ctx context.Context, name string) (int, error)
	DeleteLanguageByID(ctx context.Context, id uuid.UUID) error
	DeleteAllLanguages(ctx context.Context) error
}
