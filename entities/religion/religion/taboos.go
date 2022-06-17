package religion

import "fmt"

type Taboos struct {
	religion *Religion
}

func (t *Theology) generateTaboos() *Taboos {
	ts := &Taboos{religion: t.religion}

	return ts
}

func (rs *Taboos) Print() {
	fmt.Printf("Taboos (religion_name=%s):\n", rs.religion.Name)
}
