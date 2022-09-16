package presets

import "persons_generator/engine/entities/person/temperament"

var ArcticMongoloidTemperamentPreset = temperament.NewTemperamentGene(
	map[string]float64{
		temperament.Sanguine:    0.2,
		temperament.Choleric:    0.2,
		temperament.Melancholic: 0.25,
		temperament.Phlegmatic:  0.35,
	},
)
