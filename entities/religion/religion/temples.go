package religion

import "fmt"

type Temples struct {
	religion *Religion
	attrs    *Attributes
}

func (as *Attributes) generateTemples() *Temples {
	ts := &Temples{religion: as.religion, attrs: as}

	return ts
}

func (ts *Temples) Print() {
	fmt.Printf("Temples (religion_name=%s):\n", ts.religion.Name)
}
