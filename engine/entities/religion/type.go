package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type Type struct {
	religion *Religion

	Type    TypeName
	Subtype SubtypeName
}

func NewType(r *Religion) *Type {
	t := &Type{religion: r}
	t.Type = t.GenerateTypeName()
	t.Subtype = t.GenerateSubtypeName()

	return t
}

func (t *Type) Print() {
	var subType string
	if t.Subtype != "" {
		subType = fmt.Sprintf("(%s)", t.Subtype)
	}
	fmt.Printf("Type=%s%s\n", t.Type, subType)
}

func (t *Type) IsMonotheism() bool {
	return t.Type == MonotheismType
}

func (t *Type) IsPolytheism() bool {
	return t.Type == PolytheismType
}

func (t *Type) IsDeityDualism() bool {
	return t.Type == DeityDualismType
}

func (t *Type) IsDeism() bool {
	return t.Type == DeismType
}

func (t *Type) IsAtheism() bool {
	return t.Type == AtheismType
}

type TypeName string

const (
	MonotheismType   TypeName = "monotheism"
	PolytheismType   TypeName = "polytheism"
	DeityDualismType TypeName = "deity_dualism"
	DeismType        TypeName = "deism" // Deism is the belief in the existence of God solely based on rational thought without any reliance on revealed religions or religious authority
	AtheismType      TypeName = "atheism"
)

func (t *Type) GenerateTypeName() TypeName {
	var (
		monotheism   = 0.55
		polytheism   = 0.61
		deityDualism = 0.4
		deism        = 0.35
		atheism      = 0.05
	)

	return TypeName(pm.GetRandomFromSeveral(map[string]float64{
		string(MonotheismType):   pm.PrepareProbability(monotheism),
		string(PolytheismType):   pm.PrepareProbability(polytheism),
		string(DeityDualismType): pm.PrepareProbability(deityDualism),
		string(DeismType):        pm.PrepareProbability(deism),
		string(AtheismType):      pm.PrepareProbability(atheism),
	}))
}

type SubtypeName string

const (
	ClassicPolytheismSubtype   SubtypeName = "classic"
	HenothismPolytheismSubtype SubtypeName = "henothism" //  is the worship of a single, supreme god while not denying the existence or possible existence of other lower deities.[
	MonolatryPolytheismSubtype SubtypeName = "monolatry" // is the belief in the existence of many gods, but with the consistent worship of only one deity
	OmnismPolytheismSubtype    SubtypeName = "omnism"
)

func (t *Type) GenerateSubtypeName() SubtypeName {
	if t.Type == "" {
		t.GenerateTypeName()
	}

	for _, tn := range []TypeName{MonotheismType, DeityDualismType, DeismType, AtheismType} {
		if tn == t.Type {
			return ""
		}
	}

	var (
		classic   = 0.65
		henothism = 0.25
		monolatry = 0.45
		omnism    = 0.3
	)

	m := map[string]float64{
		string(ClassicPolytheismSubtype):   pm.PrepareProbability(classic),
		string(HenothismPolytheismSubtype): pm.PrepareProbability(henothism),
		string(MonolatryPolytheismSubtype): pm.PrepareProbability(monolatry),
		string(OmnismPolytheismSubtype):    pm.PrepareProbability(omnism),
	}
	return SubtypeName(pm.GetRandomFromSeveral(m))
}

func (t *Type) IsClassicPolytheism() bool {
	return t.Subtype == ClassicPolytheismSubtype
}

func (t *Type) IsHenothismPolytheism() bool {
	return t.Subtype == HenothismPolytheismSubtype
}

func (t *Type) IsMonolatryPolytheism() bool {
	return t.Subtype == MonolatryPolytheismSubtype
}

func (t *Type) IsOmnismPolytheism() bool {
	return t.Subtype == OmnismPolytheismSubtype
}
