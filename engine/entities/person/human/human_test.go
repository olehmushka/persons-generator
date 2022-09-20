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
		bodyPresets.NordicBodyPreset,
		psychoPresets.NordicPsychoPreset,
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

func TestProduceChildren(t *testing.T) {
	father, err := New(gender.MaleSex, NewGene(
		bodyPresets.NordicBodyPreset,
		psychoPresets.NordicPsychoPreset,
	), nil, nil)
	if err != nil {
		t.Errorf("can not create father (err=%+v)", err)
		return
	}
	father.IncreaseAge(18)
	mother, err := New(gender.FemaleSex, NewGene(
		bodyPresets.NordicBodyPreset,
		psychoPresets.NordicPsychoPreset,
	), nil, nil)
	if err != nil {
		t.Errorf("can not create mother (err=%+v)", err)
		return
	}
	mother.IncreaseAge(18)
	var (
		i      int
		result []*Human
	)
	for ; i < 100; i++ {
		children, err := mother.ProduceChildren(father)
		if err != nil {
			t.Errorf("can not create mother (err=%+v)", err)
			return
		}
		if len(children) != 0 {
			result = children
			break
		}
	}
	b, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println("Count::", i, "\n\n\n", string(b))
}
