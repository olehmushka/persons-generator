package face

type EarsType string

func (et EarsType) String() string {
	return string(et)
}

const (
	SquareEarType        EarsType = "square_ear"
	PointedEarType       EarsType = "pointed_ear"
	NarrowEarType        EarsType = "narrow_ear"
	StickingEarType      EarsType = "sticking_ear"
	RoundEarFreeLobeType EarsType = "round_ear_free_lobe"
	AttachedLobeType     EarsType = "attached_lobe"
	BroadLobeType        EarsType = "broad_lobe"
)
