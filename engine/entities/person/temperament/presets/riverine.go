package presets

import "persons_generator/engine/entities/person/temperament"

var RiverineTemperamentPreset = temperament.NewTemperamentGene(
	map[string]float64{
		temperament.Sanguine:    0.25,
		temperament.Choleric:    0.3,
		temperament.Melancholic: 0.25,
		temperament.Phlegmatic:  0.2,
	},
)
