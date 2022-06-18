package religion

import "fmt"

type Conversion struct {
	religion *Religion
}

func (t *Theology) generateConversion() *Conversion {
	c := &Conversion{religion: t.religion}

	return c
}

func (c *Conversion) Print() {
	fmt.Printf("Conversion (religion_name=%s):\n", c.religion.Name)
}
