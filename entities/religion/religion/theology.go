package religion

import "fmt"

type Theology struct {
	religion *Religion

	// traits []trait
	Rules   *Rules
	Taboos  *Taboos
	Rituals *Rituals
}

func NewTheology(r *Religion) *Theology {
	t := &Theology{religion: r}
	t.Rules = t.generateRules()
	t.Taboos = t.generateTaboos()
	t.Rituals = t.generateRituals()

	return t
}

func (t *Theology) Print() {
	fmt.Printf("Theology (religion_name=%s):\n", t.religion.Name)
	t.Rules.Print()
	t.Taboos.Print()
	t.Rituals.Print()
}
