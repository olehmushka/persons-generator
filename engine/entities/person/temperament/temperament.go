package temperament

import "encoding/json"

type Temperament struct {
	Name string
}

func (t Temperament) Zero() bool {
	return t == Temperament{}
}

func (t Temperament) Serialize() string {
	return t.Name
}

func (t Temperament) Bytes() []byte {
	if t.Zero() {
		return nil
	}

	if out, err := json.Marshal(t); err == nil {
		return out
	}

	return nil
}

func (n Temperament) Print() {
}

func (t Temperament) IsSanguine() bool {
	return t.Name == Sanguine
}

func (t Temperament) IsCholeric() bool {
	return t.Name == Choleric
}

func (t Temperament) IsMelancholic() bool {
	return t.Name == Melancholic
}

func (t Temperament) IsPhlegmatic() bool {
	return t.Name == Phlegmatic
}
