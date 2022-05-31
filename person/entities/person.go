package entities

type Person struct {
	FirstName         string
	LastName          string
	Birddate          int
	Sex               *Sex
	SexualOrientation *SexualOrientation
	Race              *Race
	Health            *Health
	IQ                *IQ
	Temperament       *Temperament
	SocionicType      *Temperament // LII, ILE, ESE, SEI  LSI, SLE, EIE, IEI...
	Culture           *Culture
	Religion          *Religion
	Mother            *Person
	Father            *Person
}
