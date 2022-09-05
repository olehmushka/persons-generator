package body

import (
	"encoding/json"

	"persons_generator/engine/entities/person/face"
	"persons_generator/engine/entities/person/hair"
)

type Body struct {
	Face face.Face `json:"face"`
	Hair hair.Hair `json:"hair"`
}

func (b Body) Zero() bool {
	return b == Body{}
}

func (b Body) Bytes() []byte {
	if b.Zero() {
		return nil
	}
	if out, err := json.Marshal(b); err == nil {
		return out
	}

	return nil
}
