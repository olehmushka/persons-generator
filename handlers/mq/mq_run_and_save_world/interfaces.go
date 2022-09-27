package mq_run_and_save_world

import "context"

type RunAndSaveWorld interface {
	ProcessRunAndSaveWorld(context.Context, []byte) error
}
