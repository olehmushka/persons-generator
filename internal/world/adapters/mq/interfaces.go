package mq

import (
	"context"
)

type Adapter interface {
	RunAndSaveWorld(context.Context, string, int, int, int, int, map[string]string) error
	ParseRunAndSaveWorldMsg(context.Context, []byte) (*RunAndSaveWorldPayload, error)
}
