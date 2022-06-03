package entities

import "persons_generator/entities/religion"

type Religion struct {
	Theology  *religion.Theology
	Clerics   *religion.Clerics
	Taboos    *religion.Taboos
	Doctrines *religion.Doctrines
}

type ReligionConversion struct{}

type Rituals struct{}

type RelicHunts struct{}

type ReligionBuilding struct{}

type ReligionEsthetic struct{}

type Anthropology struct{}

type Epistemology struct{}

type Ethics struct{}

type CulticPractices struct{}

type Temporal struct{}

type Cosmology struct{}

type ReligiousBehavior struct {
	Music     *ReligiousMusic
	Poetry    *ReligiousPoetry
	Symbolism *ReligiousSymbolism
	Sacrifice *ReligiousSacrifice
}

type ReligiousMusic struct{}

type ReligiousPoetry struct{}

type ReligiousSymbolism struct{}

type ReligiousSacrifice struct{}

/*
1) addressing the sacred
2) music
3) poetry
4) physiological alteration through drugs, deprivation, selfmortification, and isolation
5) exhortation — addressing others as a representative of a
divine being
6) recitation of the code — the use of sacred, written and
oral statements to state doctrine
7) sympathetic ritual — imitation of sacred beings and
events
8) wielding sacred power (mana) — touching objects
containing sacred power, including laying on of hands
9) taboos — rules specifying avoidance of contact and
action to prevent activating unwanted manifestations of
sacred power
10) feasts — sacred meals
11) sacrifice — ritual killing, often as part of a feast
12) congregation — meetings, processions, and other forms
of coming together
13) symbolism — using objects symbolizing the sacred
14) inspiration — the pursuit of visions and revelation
15) extension and modification of the code
16) extended consequences — the results of applying
religious values outside of the context of specifically
ritual occasions
*/
