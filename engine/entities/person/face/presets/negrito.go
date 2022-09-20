package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var NegritoFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.25,
		string(face.Circle):    0.2,
		string(face.Oblong):    0.25,
		string(face.VTriangle): 0.25,
		string(face.Diamond):   0.2,
		string(face.Heart):     0.2,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.HoodedEyesType):             0.2,
			string(face.AsianEyesType):              0.35,
			string(face.ChildishRoundAsianEyesType): 0.3,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette: 1,
		}),
	),
	face.NewEarsGene(map[string]float64{
		string(face.SquareEarType):        0.25,
		string(face.PointedEarType):       0.25,
		string(face.NarrowEarType):        0.25,
		string(face.StickingEarType):      0.25,
		string(face.RoundEarFreeLobeType): 0.25,
		string(face.AttachedLobeType):     0.25,
		string(face.BroadLobeType):        0.25,
	}),
	face.NewNoseGene(map[string]float64{
		string(face.FleshyNoseType):  0.2,
		string(face.NixonNoseType):   0.05,
		string(face.BulbousNoseType): 0.05,
		string(face.ComboNoseType):   0.1,
		string(face.NubianNoseType):  0.35,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):       0.2,
		string(face.HeavyUpperLipsType): 0.15,
		string(face.WideLipsType):       0.35,
		string(face.HeavyLowerLipsType): 0.15,
	}),
)
