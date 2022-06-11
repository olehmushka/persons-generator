package religion

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
	Hedonism               bool
}

func (d *Doctrine) generateSocialDoctrine() *SocialDoctrine {
	sd := &SocialDoctrine{religion: d.religion}

	return sd
}
