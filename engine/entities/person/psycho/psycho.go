package psycho

import "encoding/json"

type Psycho struct{}

func (b Psycho) Zero() bool {
	return b == Psycho{}
}

func (b Psycho) Bytes() []byte {
	if b.Zero() {
		return nil
	}
	if out, err := json.Marshal(b); err == nil {
		return out
	}

	return nil
}
