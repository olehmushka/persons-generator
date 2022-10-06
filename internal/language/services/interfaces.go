package services

import (
	"context"
	"persons_generator/core/storage"
	"persons_generator/engine/entities/gender"
	"persons_generator/internal/language/entities"
)

type Language interface {
	QueryDefaultLanguages(q string, opts storage.PaginationSortingOpts) ([]*entities.Language, int, error)
	ReadDefaultLanguageSubfamilies(ctx context.Context, opts storage.PaginationSortingOpts) ([]*entities.Subfamily, int, error)
	CreateLanguage(ctx context.Context, in *entities.Language) (string, error)
	ReadLanguagesByName(ctx context.Context, name string, opts storage.PaginationSortingOpts) ([]*entities.Language, int, error)
	ReadLanguage(ctx context.Context, id, name string) (*entities.Language, error)
	DeleteLanguageByID(ctx context.Context, id string) error
	DeleteAllLanguages(ctx context.Context) error

	ReadLanguageByID(ctx context.Context, id string) (*entities.Language, error)
	GetRandomWordByLanguageID(ctx context.Context, id string) (string, error)
	GetRandomOwnNameByLanguageID(ctx context.Context, id string, sex gender.Sex) (string, error)
	GetRandomCultureNameByLanguageID(ctx context.Context, id string) (string, error)
	GetRandomReligionNameByLanguageID(ctx context.Context, id string) (string, error)
}
