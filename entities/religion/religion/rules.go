package religion

import "fmt"

type Rules struct {
	religion *Religion
}

func (t *Theology) generateRules() *Rules {
	rs := &Rules{religion: t.religion}

	return rs
}

func (rs *Rules) Print() {
	fmt.Printf("Rules (religion_name=%s):\n", rs.religion.Name)
}

/*
tax to support the poor, tithe
		DressCode           bool
		BeHumble           bool
		LearnKeyScriptures           bool
		LiveUnderGround            bool
		PrayWithFrequency bool
		TaxNonbelievers            bool
*/
