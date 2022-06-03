package entities

type Psycho struct {
	IQ                float64
	Sociotype         *Sociotype
	Temperament       *Temperament
	SexualOrientation *SexualOrientation
}

type Sociotype struct{}

type Temperament struct{}

type SexualOrientation struct{}
