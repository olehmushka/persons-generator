package culture

type Origin string

func (o Origin) String() string {
	return string(o)
}

const (
	NativeOrigin Origin = "native"
	CustomOrigin Origin = "custom"
)
