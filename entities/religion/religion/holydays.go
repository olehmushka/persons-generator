package religion

import "fmt"

type Holydays struct {
	religion *Religion
}

func (t *Theology) generateHolydays() *Holydays {
	hs := &Holydays{religion: t.religion}

	return hs
}

func (hs *Holydays) Print() {
	fmt.Printf("Holydays (religion_name=%s):\n", hs.religion.Name)
}

/*
libations, samhain, mysteries
calendar of feasts
summer solstice
winter solstice

	DanceParty      bool
	DrumParty       bool
	SocialFestivals bool
	TreeCelebration bool
*/
