package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var BushmanFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.15,
		string(face.Circle):    0.15,
		string(face.Oblong):    0.15,
		string(face.VTriangle): 0.2,
		string(face.Diamond):   0.25,
		string(face.Heart):     0.25,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.DroopyEyesType):       0.2,
			string(face.DroopyHoodedEyesType): 0.3,
			string(face.HoodedEyesType):       0.4,
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
		string(face.FleshyNoseType):   0.2,
		string(face.TurnedUpNoseType): 0.25,
		string(face.BumpyNoseType):    0.05,
		string(face.BulbousNoseType):  0.05,
		string(face.ComboNoseType):    0.075,
		string(face.SnubNoseType):     0.05,
		string(face.NubianNoseType):   0.3,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.3,
		string(face.WideLipsType):           0.25,
		string(face.HeardShapedLipsType):    0.05,
		string(face.DownwardTurnedLipsType): 0.25,
	}),
)
