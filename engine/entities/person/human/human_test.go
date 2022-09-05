package human

import (
	"encoding/json"
	"fmt"
	"persons_generator/engine/entities/gender"
	bodyPresets "persons_generator/engine/entities/person/body/presets"
	psychoPresets "persons_generator/engine/entities/person/psycho/presets"
	"testing"
)

func TestNew(t *testing.T) {
	sex, err := gender.GetRandomSex()
	if err != nil {
		t.Fatal(err)
		return
	}
	h, err := New(sex, NewGene(
		bodyPresets.SlavicBodyPreset,
		psychoPresets.SlavicPsychoPreset,
	), nil, nil)
	if err != nil {
		t.Fatal(err)
		return
	}
	b, err := json.Marshal(h)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(string(b))
}
