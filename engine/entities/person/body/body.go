package body

import (
	"encoding/json"

	"persons_generator/engine/entities/person/face"
	"persons_generator/engine/entities/person/hair"
	"persons_generator/engine/entities/person/size"
)

type Body struct {
	Face face.Face `json:"face"`
	Hair hair.Hair `json:"hair"`
	Size size.Size `json:"size"`
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
