package religion

import "fmt"

type HolyScripture struct {
	religion *Religion
	attrs    *Attributes
}

func (as *Attributes) generateHolyScripture() *HolyScripture {
	hs := &HolyScripture{religion: as.religion, attrs: as}

	return hs
}

func (hs *HolyScripture) Print() {
	fmt.Printf("HolyScripture (religion_name=%s):\n", hs.religion.Name)
}
