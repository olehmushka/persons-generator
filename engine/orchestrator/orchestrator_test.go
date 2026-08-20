package orchestrator

import (
	"os"
	"testing"

	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/language"
	"persons_generator/engine/entities/religion"
)

// TestMain loads word bases once before running tests. language.SetWordBases
// (needed by culture/religion name generation) used to read files via a
// path relative to the repo root, which broke under `go test`'s per-package
// working directory -- word_base_ref.go now reads from an embedded FS
// instead, so this is a plain call with no chdir workaround needed.
func TestMain(m *testing.M) {
	if err := language.SetWordBases(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

// CreateCultures, CreateReligion, and CreateWorld are thin, DB-free wrappers
// around culture.NewMany/religion.New/world.New -- none of them touch
// o.storage or o.mongodb, so they're safe to call on a zero-value
// Orchestrator without a live Mongo/Redis connection.

func TestOrchestratorCreateCultures(t *testing.T) {
	o := &Orchestrator{}

	t.Run("happy path: amount with no preferred", func(tt *testing.T) {
		cultures, err := o.CreateCultures(2, nil)
		if err != nil {
			tt.Fatalf("unexpected error: %+v", err)
		}
		if len(cultures) != 2 {
			tt.Fatalf("unexpected length (expected = 2, actual = %d)", len(cultures))
		}
		for _, c := range cultures {
			if c.ID == "" {
				tt.Error("expected a non-empty culture ID")
			}
			if c.Name == "" {
				tt.Error("expected a non-empty culture name")
			}
		}
	})

	t.Run("invalid: amount less than len(preferred)", func(tt *testing.T) {
		_, err := o.CreateCultures(1, []*culture.Preference{{}, {}})
		if err == nil {
			tt.Fatal("expected an error when amount < len(preferred), got <nil>")
		}
	})
}

func TestOrchestratorCreateReligion(t *testing.T) {
	o := &Orchestrator{}

	cultures, err := o.CreateCultures(1, nil)
	if err != nil {
		t.Fatalf("can not set up test culture: %+v", err)
	}

	r, err := o.CreateReligion(cultures[0])
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	if r.ID == "" {
		t.Error("expected a non-empty religion ID")
	}
	if r.Name == "" {
		t.Error("expected a non-empty religion name")
	}
}

func TestOrchestratorCreateWorld(t *testing.T) {
	o := &Orchestrator{}

	cultures, err := o.CreateCultures(1, nil)
	if err != nil {
		t.Fatalf("can not set up test cultures: %+v", err)
	}
	refs, err := religion.NewReferences(religion.Config{}, 1, cultures)
	if err != nil {
		t.Fatalf("can not set up test culture-religion references: %+v", err)
	}
	religions := religion.ExtractReligionsFromCultureReferences(refs)

	w, err := o.CreateWorld("world-1", 1, cultures, religions, refs)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	if w.ID != "world-1" {
		t.Errorf("unexpected world ID (expected = world-1, actual = %s)", w.ID)
	}
	if w.Size != 1 {
		t.Errorf("unexpected world size (expected = 1, actual = %d)", w.Size)
	}
	if w.PopulationNumber == 0 {
		t.Error("expected a freshly seeded world to have a non-zero population")
	}
}
