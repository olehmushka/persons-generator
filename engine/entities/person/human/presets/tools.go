package presets

import (
	"fmt"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

func GetPresetByCulture(c *culture.Culture) (gene.Gene, error) {
	if c == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not get human gene by culture for <nil>")
	}

	// For ancient abstract cultures
	switch c.Abstuct {
	case culture.AbsAncientBelgae:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsAncientCeltIberian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientGaelic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsAncientGallic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			DinaricHumanPreset,
			AlpineHumanPreset,
		})
	case culture.AbsAncientGermanic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			NordicAlpineHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsAncientIberian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientOccidental:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			CelticHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientPannonian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientPretani:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsAncientVeneti:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientItalic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientDacian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientHellenistic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientIllyrian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientNumidian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{NiloticNegroHumanPreset})
	case culture.AbsAncientLevantine:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			ArmenoidHumanPreset,
		})
	case culture.AbsAncientAksumite:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			NiloticNegroHumanPreset,
			AfricanNegroHumanPreset,
		})
	case culture.AbsAncientArabian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			ArmenoidHumanPreset,
		})
	case culture.AbsAncientEgyptian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			NiloticNegroHumanPreset,
		})
	case culture.AbsAncientNubian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			NiloticNegroHumanPreset,
			AfricanNegroHumanPreset,
		})
	case culture.AbsAncientAnatolian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArmenoidHumanPreset})
	case culture.AbsAncientAramaic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArmenoidHumanPreset})
	case culture.AbsAncientBactrian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			ArmenoidHumanPreset,
		})
	case culture.AbsAncientCaucasian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArmenoidHumanPreset})
	case culture.AbsAncientIranian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			ArmenoidHumanPreset,
		})
	case culture.AbsAncientScythian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			ArmenoidHumanPreset,
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsAncientAryan:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			ArmenoidHumanPreset,
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsAncientDravidian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MalayaMongoloidHumanPreset,
		})
	case culture.AbsAncientPracyan:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AlpineHumanPreset})
	case culture.AbsAncientTibetan:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	}

	switch c.Abstuct {
	case culture.AbsMedievalAkan:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalArabic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ArmenoidHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalBaltic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			EastBalticHumanPreset,
			NordicAlpineHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsMedievalBaltoFinnic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			EastBalticHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsMedievalBerber:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalBrythonic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			AlpineHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsMedievalBurman:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MalayaMongoloidHumanPreset,
		})
	case culture.AbsMedievalByzantine:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			NordicMedititerraneanHumanPreset,
		})
	case culture.AbsMedievalCentralAfrican:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalCentralGermanic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			CelticHumanPreset,
			DinaricHumanPreset,
			NordicAlpineHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsMedievalChinese:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalDravidian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
		})
	case culture.AbsMedievalEastAfrican:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalEastSlavic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			AlpineHumanPreset,
			EastBalticHumanPreset,
		})
	case culture.AbsMedievalFrankish:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsMedievalGoidelic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsMedievalGuineanUplander:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalHornAfrican:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalIberian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalIndoAryan:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			AlpineHumanPreset,
		})
	case culture.AbsMedievalIranian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ArmenoidHumanPreset,
		})
	case culture.AbsMedievalIsraelite:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ArmenoidHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalLatin:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
			NordicMedititerraneanHumanPreset,
		})
	case culture.AbsMedievalMagyar:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
			DinaricHumanPreset,
		})
	case culture.AbsMedievalMongolic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalNigerDelta:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalNorthGermanic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			NordicHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsMedievalQiangic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalSahelian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			AfricanNegroHumanPreset,
		})
	case culture.AbsMedievalSenegambian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalSouthSlavic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalTibetan:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalTocharian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AlpineHumanPreset})
	case culture.AbsMedievalTurkic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ArmenoidHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalUgroPermian:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			EastBalticHumanPreset,
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalVlach:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
		})
	case culture.AbsMedievalVolgaFinnic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalWestGermanic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			NordicAlpineHumanPreset,
			AlpineHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsMedievalWestSlavic:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			NordicHumanPreset,
		})
	}

	return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not get human gene preset for culture (culture abstract=%s, culture=%+v)", c.Abstuct.Name, c))
}
