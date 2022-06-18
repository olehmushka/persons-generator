package religion

import "fmt"

type MarriageTradition struct {
	religion *Religion

	Kind          MarriageKind
	Bastardry     Bastardry
	Consanguinity Consanguinity
}

func (t *Theology) generateMarriageTradition() *MarriageTradition {
	mt := &MarriageTradition{religion: t.religion}

	return mt
}

func (mt *MarriageTradition) Print() {
	fmt.Printf("MarriageTradition (religion_name=%s):\n", mt.religion.Name)
	fmt.Printf("MarriageKind=%s\n", mt.Kind)
	fmt.Printf("Bastardry=%s\n", mt.Bastardry)
	fmt.Printf("Consanguinity=%s\n", mt.Consanguinity)
}

type MarriageKind string

const (
	Monogamy              MarriageKind = "monogamy"
	Polygamy              MarriageKind = "polygamy"
	ConsortsAndConcubines MarriageKind = "consorts_and_concubines"
)

type Bastardry string

const (
	NoBastards       Bastardry = "no_bastards"
	Legitimization   Bastardry = "legitimization"
	NoLegitimization Bastardry = "no_legitimization"
)

type Consanguinity string

const (
	CloseKinTaboo        Consanguinity = "close_kin_taboo"
	CousinMarriage       Consanguinity = "cousin_marriage"
	AvunculateMarriage   Consanguinity = "avunculate_marriage"
	UnrestrictedMarriage Consanguinity = "unrestricted_marriage"
)
