package world

type GetHumansParams struct{}

type QueryPersonsOpts struct {
	MaleCount    int
	FemaleCount  int
	PersonsCount int
	MinAge       int `default:"0"`
	MaxAge       int `default:"120"`
	OnlyAlive    bool
	OnlyDead     bool
}
