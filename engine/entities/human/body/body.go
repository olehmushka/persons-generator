package body

import (
	"encoding/json"

	"persons_generator/engine/entities/human/face"
)

type Body struct {
	Face face.Face
}

func (b Body) Bytes() []byte {
	zero := Body{}
	if b == zero {
		return nil
	}
	if out, err := json.Marshal(b); err == nil {
		return out
	}

	return nil
}
