package religion

import (
	"fmt"

	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Afterlife struct {
	religion *Religion `json:"-"`
	doctrine *Doctrine `json:"-"`

	IsExists     bool                   `json:"is_exists"`
	Participants *AfterlifeParticipants `json:"participants"`
	Traits       []*trait               `json:"traits"`
}

func (al *Afterlife) IsZero() bool {
	return al == nil
}

func (al *Afterlife) Print() {
	if !al.IsExists {
		fmt.Printf("Afterlife does not exists (religion=%s)\n", al.religion.Name)
		return
	}
	fmt.Printf("Afterlife exists (religion=%s)\n", al.religion.Name)
	al.Participants.Print()
	if len(al.Traits) != 0 {
		for _, trait := range al.Traits {
			fmt.Printf(" - %s\n", trait.Name)
		}
	}
}

func (al *Afterlife) SerializeParticipant() map[string]string {
	if al == nil {
		return map[string]string{}
	}

	return al.Participants.Serialize()
}

func (al *Afterlife) SerializeTraits() []string {
	if al == nil {
		return []string{}
	}

	return extractTraitNames(al.Traits)
}

func (d *Doctrine) generateAfterlife() (*Afterlife, error) {
	al := &Afterlife{religion: d.religion, doctrine: d}
	isExists, err := al.generateIsExistsAfterlife()
	if err != nil {
		return nil, err
	}
	al.IsExists = isExists
	if al.IsExists {
		if err := al.generateAfterlifeContent(); err != nil {
			return nil, err
		}
	}

	return al, nil
}

func (al *Afterlife) generateAfterlifeContent() error {
	ps, err := al.generateAfterlifeParticipants()
	if err != nil {
		return err
	}
	al.Participants = ps
	ts, err := al.generateTraits(0, len(al.getAllAfterlifeTraits()))
	if err != nil {
		return err
	}
	al.Traits = ts

	return nil
}

func (al *Afterlife) clearAfterlifeContent() {
	al.Participants = nil
	al.Traits = nil
}

func (al *Afterlife) generateIsExistsAfterlife() (bool, error) {
	if al.religion.Type.IsAtheism() {
		return false, nil
	}
	if al.religion.HasReincarnation() {
		return false, nil
	}

	var probability float64
	switch {
	case al.religion.Type.IsMonotheism():
		p, err := pm.RandFloat64InRange(0.75, 1)
		if err != nil {
			return false, err
		}
		probability += p
	case al.religion.Type.IsPolytheism():
		p, err := pm.RandFloat64InRange(0.3, 1)
		if err != nil {
			return false, err
		}
		probability += p
	case al.religion.Type.IsDeityDualism():
		p, err := pm.RandFloat64InRange(0.5, 1)
		if err != nil {
			return false, err
		}
		probability += p
	case al.religion.Type.IsDeism():
		p, err := pm.RandFloat64InRange(0.05, 1)
		if err != nil {
			return false, err
		}
		probability += p
	}

	for _, trait := range al.doctrine.Human.Nature.Traits {
		if trait.Name == HasSoulHumanTrait {
			p, err := pm.RandFloat64InRange(0.25, 0.75)
			if err != nil {
				return false, err
			}
			probability += p
		}
	}

	return pm.GetRandomBool(probability)
}

type AfterlifeParticipants struct {
	religion  *Religion  `json:"-"`
	afterlife *Afterlife `json:"-"`

	ForTopBelievers    AfterlifeOption `json:"for_top_belivers"`
	ForBelievers       AfterlifeOption `json:"for_belivers"`
	ForUntrueBelievers AfterlifeOption `json:"for_untrue_belivers"`
	ForAtheists        AfterlifeOption `json:"for_atheists"`
}

func (al *Afterlife) generateAfterlifeParticipants() (*AfterlifeParticipants, error) {
	alp := &AfterlifeParticipants{religion: al.religion, afterlife: al}
	topBelivers, err := al.generateForTopBelieversAfterlife(alp)
	if err != nil {
		return nil, err
	}
	alp.ForTopBelievers = topBelivers
	believers, err := al.generateForBelieversAfterlife(alp)
	if err != nil {
		return nil, err
	}
	alp.ForBelievers = believers
	untrueBelievers, err := al.generateForUntrueBelieversAfterlife(alp)
	if err != nil {
		return nil, err
	}
	alp.ForUntrueBelievers = untrueBelievers
	atheists, err := al.generateForAtheistsAfterlife(alp)
	if err != nil {
		return nil, err
	}
	alp.ForAtheists = atheists

	return alp, nil
}

func (alp *AfterlifeParticipants) Print() {
	fmt.Printf("Afterlife for top believers is %s\n", alp.ForTopBelievers)
	fmt.Printf("Afterlife for believers is %s\n", alp.ForBelievers)
	fmt.Printf("Afterlife for untrue believers is %s\n", alp.ForUntrueBelievers)
	fmt.Printf("Afterlife for atheists is %s\n", alp.ForAtheists)
}

func (alp *AfterlifeParticipants) Serialize() map[string]string {
	if alp == nil {
		return map[string]string{}
	}

	return map[string]string{
		"for_top_belivers":    alp.ForTopBelievers.String(),
		"for_belivers":        alp.ForBelievers.String(),
		"for_untrue_belivers": alp.ForUntrueBelievers.String(),
		"for_atheists":        alp.ForAtheists.String(),
	}
}

func (al *Afterlife) generateForTopBelieversAfterlife(alp *AfterlifeParticipants) (AfterlifeOption, error) {
	good, err := pm.RandFloat64InRange(0.5, 1)
	if err != nil {
		return "", err
	}
	depends, err := pm.RandFloat64InRange(0.3, 0.8)
	if err != nil {
		return "", err
	}
	bad, err := pm.RandFloat64InRange(0, 0.4)
	if err != nil {
		return "", err
	}

	return getAfterlifeOptionByProbability(good, depends, bad)
}

func (al *Afterlife) generateForBelieversAfterlife(alp *AfterlifeParticipants) (AfterlifeOption, error) {
	good, err := pm.RandFloat64InRange(0.5, 0.9)
	if err != nil {
		return "", err
	}
	depends, err := pm.RandFloat64InRange(0.3, 0.75)
	if err != nil {
		return "", err
	}
	bad, err := pm.RandFloat64InRange(0, 0.45)
	if err != nil {
		return "", err
	}

	switch {
	case alp.ForTopBelievers.GetScore() < GoodAfterlife.GetScore():
		good = 0
	case alp.ForTopBelievers.GetScore() < DependsAfterlife.GetScore():
		good = 0
		depends = 0
	}

	return getAfterlifeOptionByProbability(good, depends, bad)
}

func (al *Afterlife) generateForUntrueBelieversAfterlife(alp *AfterlifeParticipants) (AfterlifeOption, error) {
	good, err := pm.RandFloat64InRange(0.1, 0.5)
	if err != nil {
		return "", err
	}
	depends, err := pm.RandFloat64InRange(0.1, 0.6)
	if err != nil {
		return "", err
	}
	bad, err := pm.RandFloat64InRange(0.1, 0.7)
	if err != nil {
		return "", err
	}
	if al.religion.Type.IsPolytheism() {
		if al.religion.Type.IsMonolatryPolytheism() {
			p, err := pm.RandFloat64InRange(0.01, 0.1)
			if err != nil {
				return "", err
			}
			good += p
		}
		if al.religion.Type.IsOmnismPolytheism() {
			p, err := pm.RandFloat64InRange(0.05, 0.2)
			if err != nil {
				return "", err
			}
			good += p
		}
	}
	switch {
	case alp.ForBelievers.GetScore() < GoodAfterlife.GetScore():
		good = 0
	case alp.ForBelievers.GetScore() < DependsAfterlife.GetScore():
		good = 0
		depends = 0
	}

	return getAfterlifeOptionByProbability(good, depends, bad)
}

func (al *Afterlife) generateForAtheistsAfterlife(alp *AfterlifeParticipants) (AfterlifeOption, error) {
	good, err := pm.RandFloat64InRange(0.05, 0.35)
	if err != nil {
		return "", err
	}
	depends, err := pm.RandFloat64InRange(0.05, 0.45)
	if err != nil {
		return "", err
	}
	bad, err := pm.RandFloat64InRange(0.1, 0.8)
	if err != nil {
		return "", err
	}

	if al.religion.Type.IsDeism() {
		goodP, err := pm.RandFloat64InRange(0.01, 0.1)
		if err != nil {
			return "", err
		}
		good += goodP
		dependsP, err := pm.RandFloat64InRange(0.01, 0.1)
		if err != nil {
			return "", err
		}
		depends += dependsP
	}
	switch {
	case alp.ForBelievers.GetScore() < GoodAfterlife.GetScore():
		good = 0
	case alp.ForBelievers.GetScore() < DependsAfterlife.GetScore():
		good = 0
		depends = 0
	}

	return getAfterlifeOptionByProbability(good, depends, bad)
}

func (al *Afterlife) updateAfterlife(isExists bool) {
	if !al.IsExists && isExists {
		al.generateAfterlifeContent()
		al.IsExists = isExists
		return
	}
	if al.IsExists && !isExists {
		al.clearAfterlifeContent()
		al.IsExists = isExists
		return
	}
}

type AfterlifeOption string

func (alo AfterlifeOption) String() string {
	return string(alo)
}

const (
	GoodAfterlife    AfterlifeOption = "good"
	DependsAfterlife AfterlifeOption = "depends"
	BadAfterlife     AfterlifeOption = "bad"
)

func getAfterlifeOptionByProbability(good, depends, bad float64) (AfterlifeOption, error) {
	ao, err := pm.GetRandomFromSeveral(map[string]float64{
		string(GoodAfterlife):    pm.PrepareProbability(good),
		string(DependsAfterlife): pm.PrepareProbability(depends),
		string(BadAfterlife):     pm.PrepareProbability(bad),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate afterlife option")
	}
	return AfterlifeOption(ao), nil
}

func (o AfterlifeOption) GetScore() int {
	switch o {
	case GoodAfterlife:
		return 1
	case DependsAfterlife:
		return 0
	case BadAfterlife:
		return -1
	default:
		return 0
	}
}

func (al *Afterlife) generateTraits(min, max int) ([]*trait, error) {
	return generateTraits(al.religion, al.getAllAfterlifeTraits(), generateTraitsOpts{
		Label: "Afterlife.generateTraits",
		Min:   min,
		Max:   max,
	})
}

func (al *Afterlife) getAllAfterlifeTraits() []*trait {
	return []*trait{
		{
			Name: HeavenlyPalaceAfterlife,
			ReligionMetadata: &religionMetadata{
				Simple: 0.5,
			},
			BaseCoef: al.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Type.IsPolytheism() {
					c, err := pm.RandFloat64InRange(0.01, 0.1)
					if err != nil {
						return false, err
					}
					addCoef += c
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PsychopompAfterlife,
			ReligionMetadata: &religionMetadata{
				Chthonic: 0.5,
			},
			BaseCoef: al.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Type.IsPolytheism() {
					c, err := pm.RandFloat64InRange(0.01, 0.1)
					if err != nil {
						return false, err
					}
					addCoef += c
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetAfterlifeParticipantsSimilarityCoef(p1, p2 *AfterlifeParticipants) float64 {
	if p1 == nil && p2 == nil {
		return 1
	}
	if p1 == nil || p2 == nil {
		return 0
	}

	similarityTraits := []struct {
		enable bool
		coef   float64
	}{
		{
			enable: p1.ForTopBelievers == p2.ForTopBelievers,
			coef:   0.25,
		},
		{
			enable: p1.ForBelievers == p2.ForBelievers,
			coef:   0.25,
		},
		{
			enable: p1.ForUntrueBelievers == p2.ForUntrueBelievers,
			coef:   0.25,
		},
		{
			enable: p1.ForAtheists == p2.ForAtheists,
			coef:   0.25,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		if t.enable {
			out += t.coef
		}
	}

	return out
}

func GetAfterlifeSimilarityCoef(al1, al2 *Afterlife) float64 {
	if al1.IsZero() && al2.IsZero() {
		return 1
	}
	if al1.IsZero() || al2.IsZero() {
		return 0
	}

	if al1.IsExists != al2.IsExists {
		return 0
	}
	if !al1.IsExists {
		return 1
	}

	similarityTraits := []struct {
		value float64
		coef  float64
	}{
		{
			value: GetAfterlifeParticipantsSimilarityCoef(al1.Participants, al2.Participants),
			coef:  0.6,
		},
		{
			value: GetTraitsSimilarityCoef(al1.Traits, al2.Traits),
			coef:  0.4,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		out += t.value * t.coef
	}

	return out
}
