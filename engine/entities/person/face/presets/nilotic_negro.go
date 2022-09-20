package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var NiloticNegroFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.45,
		string(face.Oblong):    0.3,
		string(face.VTriangle): 0.15,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.AlmondEyesType):       0.15,
			string(face.DroopyEyesType):       0.35,
			string(face.DroopyHoodedEyesType): 0.25,
			string(face.HoodedEyesType):       0.15,
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
		string(face.BulbousNoseType): 0.1,
		string(face.ComboNoseType):   0.05,
		string(face.NubianNoseType):  0.45,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.35,
		string(face.WideLipsType):           0.3,
		string(face.DownwardTurnedLipsType): 0.3,
	}),
)
