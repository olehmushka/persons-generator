package json_storage

import "context"

type Storage interface {
	Get(ctx context.Context, filename string) ([]byte, error)
	Store(ctx context.Context, filename string, file []byte) error
}
