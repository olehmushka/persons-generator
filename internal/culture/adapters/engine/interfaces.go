package engine

import (
	"context"

	"persons_generator/core/storage"
	"persons_generator/engine/entities/culture"
	"persons_generator/internal/culture/entities"
)

type Adapter interface {
	CreateCultures(ctx context.Context, amount int, preferred []*entities.CulturePreference) ([]*culture.Culture, error)
	GetProtoCultures(ctx context.Context, q string, opts storage.PaginationSortingOpts) ([]*culture.Culture, int, error)
	GetCultureByID(ctx context.Context, id string) (*culture.Culture, error)
	DeleteCultureByID(ctx context.Context, id string) error
	DeleteAllCultures(context.Context) error
	UpdateCultureLanguage(ctx context.Context, cultureID, languageID string) error
}
