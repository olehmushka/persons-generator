package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var NordicAlpineFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.05,
		string(face.Circle):    0.3,
		string(face.Oblong):    0.2,
		string(face.VTriangle): 0.05,
		string(face.Square):    0.4,
		string(face.Diamond):   0.05,
		string(face.Rectangle): 0.6,
		string(face.Heart):     0.05,
		string(face.ATriangle): 0.05,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.RoundEyesType):          0.55,
			string(face.RoundishAlmondEyesType): 0.55,
			string(face.AlmondEyesType):         0.4,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette:    0.15,
			color.HazelEyesColorPalette:    0.2,
			color.AmberEyesColorPalette:    0.2,
			color.GreenEyesColorPalette:    0.1,
			color.GreyEyesColorPalette:     0.25,
			color.TealBlueEyesColorPalette: 0.235,
			color.CyanEyesColorPalette:     0.05,
			color.IndigoEyesColorPalette:   0.025,
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
		string(face.TurnedUpNoseType): 0.2,
		string(face.HawkNoseType):     0.2,
		string(face.RomanNoseType):    0.2,
		string(face.BumpyNoseType):    0.2,
		string(face.NixonNoseType):    0.2,
		string(face.BulbousNoseType):  0.2,
		string(face.ComboNoseType):    0.2,
		string(face.SnubNoseType):     0.05,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.22,
		string(face.HeavyUpperLipsType):     0.1,
		string(face.WideLipsType):           0.1,
		string(face.HeavyLowerLipsType):     0.1,
		string(face.ThinLipsType):           0.05,
		string(face.HeardShapedLipsType):    0.05,
		string(face.DownwardTurnedLipsType): 0.05,
	}),
)
