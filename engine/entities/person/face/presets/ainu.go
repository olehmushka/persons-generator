package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var AinuFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.6,
		string(face.Circle):    0.6,
		string(face.Oblong):    0.5,
		string(face.Square):    0.5,
		string(face.Rectangle): 0.3,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.DroopyHoodedEyesType):       0.3,
			string(face.HoodedEyesType):             0.4,
			string(face.AsianEyesType):              0.1,
			string(face.ChildishRoundAsianEyesType): 0.25,
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
		string(face.FleshyNoseType):    0.4,
		string(face.BumpyNoseType):     0.4,
		string(face.BulbousNoseType):   0.1,
		string(face.ComboNoseType):     0.1,
		string(face.SnubNoseType):      0.05,
		string(face.EastAsianNoseType): 0.25,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.1,
		string(face.HeavyLowerLipsType):     0.2,
		string(face.ThinLipsType):           0.4,
		string(face.HeardShapedLipsType):    0.2,
		string(face.DownwardTurnedLipsType): 0.1,
	}),
)
