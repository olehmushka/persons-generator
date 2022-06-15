package religion

import (
	"fmt"
	"strings"
)

type SocialDoctrine struct {
	religion *Religion

	Purity                 bool
	LiveIsSacred           bool
	SanctityOfNature       bool
	SacredChildbirth       bool
	Karma                  bool
	Polyamory              bool
	BadThingForGoodPurpose bool
	Raider                 bool
}

func (d *Doctrine) generateSocialDoctrine() *SocialDoctrine {
	sd := &SocialDoctrine{religion: d.religion}

	return sd
}

func (sd *SocialDoctrine) Print() {
	socialTraits := make([]string, 0)
	if sd.Purity {
		socialTraits = append(socialTraits, "purity")
	}
	if sd.LiveIsSacred {
		socialTraits = append(socialTraits, "live is sacred")
	}
	if sd.SacredChildbirth {
		socialTraits = append(socialTraits, "sacred childbirth")
	}
	if sd.Karma {
		socialTraits = append(socialTraits, "karma")
	}
	if sd.Polyamory {
		socialTraits = append(socialTraits, "polyamory")
	}
	if sd.BadThingForGoodPurpose {
		socialTraits = append(socialTraits, "bad thing for good purpose")
	}
	if sd.Raider {
		socialTraits = append(socialTraits, "raider")
	}

	fmt.Printf("Social Doctrine (religion=%s)\n Social traits: %s\n", sd.religion.Name, strings.Join(socialTraits, ", "))
}
