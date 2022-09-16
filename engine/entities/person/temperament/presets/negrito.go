package presets

import "persons_generator/engine/entities/person/temperament"

var NegritoTemperamentPreset = temperament.NewTemperamentGene(
	map[string]float64{
		temperament.Sanguine:    0.25,
		temperament.Choleric:    0.35,
		temperament.Melancholic: 0.15,
		temperament.Phlegmatic:  0.25,
	},
)
