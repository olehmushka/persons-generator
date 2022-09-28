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
	if c.Abstuct == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not get human gene by culture for <nil> abstract culture")
	}
	if c.Abstuct.Name == "" {
		return nil, wrapped_error.NewInternalServerError(nil, "can not get human gene by culture for abstract culture empty name")
	}

	// For ancient abstract cultures
	switch c.Abstuct.Name {
	case culture.AbsAncientBelgae.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsAncientCeltIberian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientGaelic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsAncientGallic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			DinaricHumanPreset,
			AlpineHumanPreset,
		})
	case culture.AbsAncientGermanic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			NordicAlpineHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsAncientIberian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientOccidental.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			CelticHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientPannonian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientPretani.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsAncientVeneti.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientItalic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientDacian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientHellenistic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientIllyrian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsAncientNumidian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{NiloticNegroHumanPreset})
	case culture.AbsAncientLevantine.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			ArmenoidHumanPreset,
		})
	case culture.AbsAncientAksumite.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			NiloticNegroHumanPreset,
			AfricanNegroHumanPreset,
		})
	case culture.AbsAncientArabian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			ArmenoidHumanPreset,
		})
	case culture.AbsAncientEgyptian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			NiloticNegroHumanPreset,
		})
	case culture.AbsAncientNubian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			NiloticNegroHumanPreset,
			AfricanNegroHumanPreset,
		})
	case culture.AbsAncientAnatolian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArmenoidHumanPreset})
	case culture.AbsAncientAramaic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArmenoidHumanPreset})
	case culture.AbsAncientBactrian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			ArmenoidHumanPreset,
		})
	case culture.AbsAncientCaucasian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArmenoidHumanPreset})
	case culture.AbsAncientIranian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			ArmenoidHumanPreset,
		})
	case culture.AbsAncientScythian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			ArmenoidHumanPreset,
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsAncientAryan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			ArmenoidHumanPreset,
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsAncientDravidian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MalayaMongoloidHumanPreset,
		})
	case culture.AbsAncientPracyan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AlpineHumanPreset})
	case culture.AbsAncientTibetan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	}

	switch c.Abstuct.Name {
	case culture.AbsMedievalAkan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalArabic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ArmenoidHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalBaltic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			EastBalticHumanPreset,
			NordicAlpineHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsMedievalBaltoFinnic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			EastBalticHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsMedievalBerber.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalBrythonic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			AlpineHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsMedievalBurman.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MalayaMongoloidHumanPreset,
		})
	case culture.AbsMedievalByzantine.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			NordicMedititerraneanHumanPreset,
		})
	case culture.AbsMedievalCentralAfrican.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalCentralGermanic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			CelticHumanPreset,
			DinaricHumanPreset,
			NordicAlpineHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsMedievalChinese.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalDravidian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
		})
	case culture.AbsMedievalEastAfrican.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalEastSlavic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			AlpineHumanPreset,
			EastBalticHumanPreset,
		})
	case culture.AbsMedievalFrankish.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsMedievalGoidelic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsMedievalGuineanUplander.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalHornAfrican.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalIberian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CelticHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalIndoAryan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			AlpineHumanPreset,
		})
	case culture.AbsMedievalIranian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ArmenoidHumanPreset,
		})
	case culture.AbsMedievalIsraelite.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ArmenoidHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalLatin.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
			NordicMedititerraneanHumanPreset,
		})
	case culture.AbsMedievalMagyar.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
			DinaricHumanPreset,
		})
	case culture.AbsMedievalMongolic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalNigerDelta.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalNorthGermanic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			NordicHumanPreset,
			NordicAlpineHumanPreset,
		})
	case culture.AbsMedievalQiangic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalSahelian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MedititerraneanHumanPreset,
			AfricanNegroHumanPreset,
		})
	case culture.AbsMedievalSenegambian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AfricanNegroHumanPreset})
	case culture.AbsMedievalSouthSlavic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalTibetan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalTocharian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{AlpineHumanPreset})
	case culture.AbsMedievalTurkic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ArmenoidHumanPreset,
			MedititerraneanHumanPreset,
		})
	case culture.AbsMedievalUgroPermian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			EastBalticHumanPreset,
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalVlach.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
		})
	case culture.AbsMedievalVolgaFinnic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ClassicMongoloidHumanPreset,
		})
	case culture.AbsMedievalWestGermanic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			NordicAlpineHumanPreset,
			AlpineHumanPreset,
			NordicHumanPreset,
		})
	case culture.AbsMedievalWestSlavic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			AlpineHumanPreset,
			DinaricHumanPreset,
			NordicHumanPreset,
		})
	}

	return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not get human gene preset for culture (culture abstract=%s, culture=%+v)", c.Abstuct.Name, c))
}
