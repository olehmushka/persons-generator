package religion

import (
	"fmt"
)

type SocialDoctrine struct {
	religion *Religion

	Traits []*socialTrait
}

func (d *Doctrine) generateSocialDoctrine() *SocialDoctrine {
	sd := &SocialDoctrine{religion: d.religion}
	sd.Traits = sd.generateTraits()

	return sd
}

func (sd *SocialDoctrine) Print() {
	fmt.Printf("Social Traits (religion_name=%s):\n", sd.religion.Name)
	for _, trait := range sd.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
}

func (sd *SocialDoctrine) generateTraits() []*socialTrait {
	allTraits := sd.getAllSocialTraits()
	traits := make([]*socialTrait, 0, len(allTraits))
	for i, trait := range allTraits {
		if trait.Calc(sd.religion, trait, traits) {
			traits = append(traits, &socialTrait{
				_religionMetadata: trait._religionMetadata,
				Index:             i,
				Name:              trait.Name,
				Calc:              trait.Calc,
			})
		}
	}

	return traits
}

type socialTrait struct {
	_religionMetadata *updateReligionMetadata
	Index             int
	Name              string
	Calc              func(r *Religion, self *socialTrait, selectedTraits []*socialTrait) bool
}

func (sd *SocialDoctrine) getAllSocialTraits() []*socialTrait {
	return []*socialTrait{ // @TODO use HighGoals for generation social doctrines
		{
			Name: "Purity",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				InsideDirected:   Float64(0.02),
				Strictness:       Float64(0.05),
				Ascetic:          Float64(0.5),
				Elitaric:         Float64(0.09),
				Organized:        Float64(0.08),
			},
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "LiveIsSacred",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Strictness:       Float64(0.01),
			},
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "SanctityOfNature",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Strictness:       Float64(0.01),
				Primitive:        Float64(0.03),
			},
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "SacredChildbirth",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.08),
				OutsideDirected:  Float64(0.08),
				Primitive:        Float64(0.03),
			},
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "Karma",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.05),
				OutsideDirected:  Float64(0.07),
				Strictness:       Float64(0.04),
			},
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "Polyamory",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.05),
				OutsideDirected:  Float64(0.05),
				Hedonism:         Float64(0.1),
				Chthonic:         Float64(0.01),
				Primitive:        Float64(0.01),
			},
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				if CalculateProbabilityFromReligionMetadata(0.9, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "BadThingForGoodPurpose",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.06),
				OutsideDirected:  Float64(0.06),
			},
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				if CalculateProbabilityFromReligionMetadata(0.9, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "Raider",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Chthonic:         Float64(0.1),
				Primitive:        Float64(0.1),
			},
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				if CalculateProbabilityFromReligionMetadata(0.9, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
	}
}
