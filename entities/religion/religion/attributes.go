package religion

import "fmt"

type Attributes struct {
	religion *Religion

	Clerics       *Clerics
	HolyScripture *HolyScripture
	Temples       *Temples
}

func NewAttributes(r *Religion) *Attributes {
	attrs := &Attributes{religion: r}
	attrs.Clerics = attrs.generateClerics()
	attrs.HolyScripture = attrs.generateHolyScripture()
	attrs.Temples = attrs.generateTemples()

	return attrs
}

func (a *Attributes) Print() {
	fmt.Printf("Attributes (religion_name=%s):\n", a.religion.Name)
	a.Clerics.Print()
	a.HolyScripture.Print()
	a.Temples.Print()
}
