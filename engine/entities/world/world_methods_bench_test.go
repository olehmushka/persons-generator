package world

import (
	"fmt"
	"os"
	"testing"

	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/language"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

// TestMain loads word bases once (needed for culture/religion name
// generation, via an embedded FS -- see language.SetWordBases) before any
// test or benchmark in this package runs.
func TestMain(m *testing.M) {
	if err := language.SetWordBases(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

// newBenchWorld builds a real, fully-seeded World of the given size --
// real cultures/religions/persons via world.New's normal seeding path, not
// hand-rolled zero-value Persons -- since RunYear's seekPartners /
// DoesWantBeMarried require non-empty religion/culture similarity maps and
// real Culture/Religion references to run without erroring out.
func newBenchWorld(b *testing.B, size int) *World {
	b.Helper()

	cultures, err := culture.NewMany(culture.Config{}, 1, nil)
	if err != nil {
		b.Fatalf("can not build bench cultures: %+v", err)
	}
	refs, err := religion.NewReferences(religion.Config{}, 1, cultures)
	if err != nil {
		b.Fatalf("can not build bench culture-religion references: %+v", err)
	}
	religions := religion.ExtractReligionsFromCultureReferences(refs)

	w, err := New(Config{}, uuid.NewString(), size, cultures, religions, refs)
	if err != nil {
		b.Fatalf("can not build bench world: %+v", err)
	}

	return w
}

// BenchmarkRunYear measures the cost of simulating a single year at a few
// world sizes -- documented (README "Known limitations") as re-scanning the
// whole location grid per unmarried person per year in seekPartners, so
// this is expected to scale worse than linearly with world size/population.
func BenchmarkRunYear(b *testing.B) {
	for _, size := range []int{2, 4, 8} {
		b.Run(fmt.Sprintf("size=%dx%d", size, size), func(bb *testing.B) {
			w := newBenchWorld(bb, size)
			bb.ResetTimer()
			for i := 0; i < bb.N; i++ {
				if err := w.RunYear(); err != nil {
					bb.Fatalf("RunYear error: %+v", err)
				}
			}
		})
	}
}
