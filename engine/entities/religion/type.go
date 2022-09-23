package religion

import (
	"fmt"

	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Type struct {
	religion *Religion `json:"-"`

	Type    TypeName    `json:"type"`
	Subtype SubtypeName `json:"subtype"`
}

func NewType(r *Religion) (*Type, error) {
	t := &Type{religion: r}
	tp, err := t.GenerateTypeName()
	if err != nil {
		return nil, err
	}
	t.Type = tp
	st, err := t.GenerateSubtypeName()
	if err != nil {
		return nil, err
	}
	t.Subtype = st

	return t, nil
}

func (t *Type) String() string {
	if t == nil {
		return ""
	}
	out := t.Type.String()
	if st := t.Subtype.String(); st != "" {
		out = fmt.Sprintf("%s (%s)", out, st)
	}

	return out
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

func (t TypeName) String() string {
	return string(t)
}

const (
	MonotheismType   TypeName = "monotheism"
	PolytheismType   TypeName = "polytheism"
	DeityDualismType TypeName = "deity_dualism"
	DeismType        TypeName = "deism" // Deism is the belief in the existence of God solely based on rational thought without any reliance on revealed religions or religious authority
	AtheismType      TypeName = "atheism"
)

func (t *Type) GenerateTypeName() (TypeName, error) {
	var (
		monotheism   = 0.55
		polytheism   = 0.61
		deityDualism = 0.4
		deism        = 0.35
		atheism      = 0.05
	)
	tp, err := pm.GetRandomFromSeveral(map[string]float64{
		string(MonotheismType):   pm.PrepareProbability(monotheism),
		string(PolytheismType):   pm.PrepareProbability(polytheism),
		string(DeityDualismType): pm.PrepareProbability(deityDualism),
		string(DeismType):        pm.PrepareProbability(deism),
		string(AtheismType):      pm.PrepareProbability(atheism),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate religion type")
	}

	return TypeName(tp), nil
}

type SubtypeName string

func (st SubtypeName) String() string {
	return string(st)
}

const (
	ClassicPolytheismSubtype   SubtypeName = "classic"
	HenothismPolytheismSubtype SubtypeName = "henothism" //  is the worship of a single, supreme god while not denying the existence or possible existence of other lower deities.[
	MonolatryPolytheismSubtype SubtypeName = "monolatry" // is the belief in the existence of many gods, but with the consistent worship of only one deity
	OmnismPolytheismSubtype    SubtypeName = "omnism"
)

func (t *Type) GenerateSubtypeName() (SubtypeName, error) {
	if t.Type == "" {
		t.GenerateTypeName()
	}

	for _, tn := range []TypeName{MonotheismType, DeityDualismType, DeismType, AtheismType} {
		if tn == t.Type {
			return "", nil
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
	st, err := pm.GetRandomFromSeveral(m)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate religion subtype")
	}

	return SubtypeName(st), nil
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

func IsTypesEqual(t1, t2 *Type) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Type != t2.Type {
		return false
	}
	if t1.Subtype != t2.Subtype {
		return false
	}

	return true
}
