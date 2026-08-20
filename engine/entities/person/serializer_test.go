package person

import (
	"encoding/json"
	"testing"

	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/gender"
	bodyPresets "persons_generator/engine/entities/person/body/presets"
	"persons_generator/engine/entities/person/human"
	psychoPresets "persons_generator/engine/entities/person/psycho/presets"
	"persons_generator/engine/entities/person/traits"
	"persons_generator/engine/entities/religion"
)

func newTestPerson(t *testing.T) *Person {
	t.Helper()

	h, err := human.New(gender.FemaleSex, human.NewGene(
		bodyPresets.DeltaBodyPreset,
		psychoPresets.DeltaPsychoPreset,
	), nil, nil)
	if err != nil {
		t.Fatalf("can not build test human: %+v", err)
	}

	return &Person{
		ID:         "person-1",
		OwnName:    "Test Name",
		Culture:    &culture.Culture{ID: "culture-1"},
		Religion:   &religion.Religion{ID: "religion-1"},
		Human:      h,
		Coordinate: &coordinate.Coordinate{X: 3, Y: 4},
		Traits:     []*traits.Trait{{Name: "brave"}},
		Spouces:    []*Person{{ID: "spouce-1"}},
		Chronology: Chronology{BirthYear: 0, DeathYear: 12, Events: []Event{}},
	}
}

func TestPersonSerialize(t *testing.T) {
	p := newTestPerson(t)

	sp := p.Serialize("world-1")
	if sp == nil {
		t.Fatal("Serialize returned <nil> for a non-nil person")
	}
	if sp.ID != p.ID {
		t.Errorf("unexpected ID (expected = %s, actual = %s)", p.ID, sp.ID)
	}
	if sp.CultureID != p.Culture.ID {
		t.Errorf("unexpected CultureID (expected = %s, actual = %s)", p.Culture.ID, sp.CultureID)
	}
	if sp.ReligionID != p.Religion.ID {
		t.Errorf("unexpected ReligionID (expected = %s, actual = %s)", p.Religion.ID, sp.ReligionID)
	}
	if sp.WorldID != "world-1" {
		t.Errorf("unexpected WorldID (expected = world-1, actual = %s)", sp.WorldID)
	}
	if sp.DeathYear != p.Chronology.DeathYear {
		t.Errorf("unexpected DeathYear (expected = %d, actual = %d)", p.Chronology.DeathYear, sp.DeathYear)
	}
	if len(sp.Spouces) != 1 || sp.Spouces[0] != "spouce-1" {
		t.Errorf("unexpected Spouces (expected = [spouce-1], actual = %+v)", sp.Spouces)
	}
	if sp.Human == nil {
		t.Fatal("unexpected <nil> Human in serialized person")
	}

	// Regression check for the "mose_type"->"nose_type" json tag typo: the
	// field must survive a real marshal/unmarshal round trip, since that's
	// exactly how the HTTP handler layer forwards this data.
	b, err := json.Marshal(sp)
	if err != nil {
		t.Fatalf("can not marshal serialized person: %+v", err)
	}
	var roundTripped map[string]any
	if err := json.Unmarshal(b, &roundTripped); err != nil {
		t.Fatalf("can not unmarshal serialized person: %+v", err)
	}
	humanMap, isOk := roundTripped["human"].(map[string]any)
	if !isOk {
		t.Fatal("expected \"human\" to unmarshal into an object")
	}
	if _, hasKey := humanMap["nose_type"]; !hasKey {
		t.Error("expected serialized human to have a \"nose_type\" key")
	}
	if _, hasKey := humanMap["mose_type"]; hasKey {
		t.Error("did not expect serialized human to still have the typo'd \"mose_type\" key")
	}
}

func TestPersonSerializeNil(t *testing.T) {
	var p *Person
	if sp := p.Serialize("world-1"); sp != nil {
		t.Errorf("expected <nil> Person to serialize to <nil>, got %+v", sp)
	}
}

func TestSerializePeople(t *testing.T) {
	people := []*Person{newTestPerson(t), newTestPerson(t)}

	out := SerializePeople("world-2", people)
	if len(out) != len(people) {
		t.Fatalf("unexpected length (expected = %d, actual = %d)", len(people), len(out))
	}
	for _, sp := range out {
		if sp.WorldID != "world-2" {
			t.Errorf("unexpected WorldID (expected = world-2, actual = %s)", sp.WorldID)
		}
	}
}
