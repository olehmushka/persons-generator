package world

import (
	"persons_generator/engine/entities/religion"
)

type Preference struct {
	Human            HumanPreference
	ReligionCultures []*religion.CultureReference
}

type HumanPreference struct {
	Amount       int `default:"100"`
	MaleAmount   int `default:"50"`
	FemaleAmount int `default:"50"`
}
