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
			HighlandHumanPreset,
			ValleyHumanPreset,
		})
	case culture.AbsAncientCeltIberian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			HighlandHumanPreset,
			MontaneHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsAncientGaelic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			HighlandHumanPreset,
			ValleyHumanPreset,
		})
	case culture.AbsAncientGallic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			HighlandHumanPreset,
			MontaneHumanPreset,
			CoastalHumanPreset,
		})
	case culture.AbsAncientGermanic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			MontaneHumanPreset,
			ValleyHumanPreset,
			DeltaHumanPreset,
		})
	case culture.AbsAncientIberian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			HighlandHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsAncientOccidental.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			HighlandHumanPreset,
			MontaneHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsAncientPannonian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			MontaneHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsAncientPretani.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			HighlandHumanPreset,
			ValleyHumanPreset,
		})
	case culture.AbsAncientVeneti.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			MontaneHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsAncientItalic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MontaneHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsAncientDacian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			MontaneHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsAncientHellenistic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RiverineHumanPreset,
		})
	case culture.AbsAncientIllyrian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MontaneHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsAncientNumidian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{TimberlandHumanPreset})
	case culture.AbsAncientLevantine.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RiverineHumanPreset,
			FoothillHumanPreset,
		})
	case culture.AbsAncientAksumite.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			TimberlandHumanPreset,
			ArchipelagoHumanPreset,
		})
	case culture.AbsAncientArabian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RiverineHumanPreset,
			FoothillHumanPreset,
		})
	case culture.AbsAncientEgyptian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RiverineHumanPreset,
			TimberlandHumanPreset,
		})
	case culture.AbsAncientNubian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			TimberlandHumanPreset,
			ArchipelagoHumanPreset,
		})
	case culture.AbsAncientAnatolian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{FoothillHumanPreset})
	case culture.AbsAncientAramaic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{FoothillHumanPreset})
	case culture.AbsAncientBactrian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			FoothillHumanPreset,
		})
	case culture.AbsAncientCaucasian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{FoothillHumanPreset})
	case culture.AbsAncientIranian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			FoothillHumanPreset,
		})
	case culture.AbsAncientScythian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			FoothillHumanPreset,
			LowlandHumanPreset,
		})
	case culture.AbsAncientAryan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			FoothillHumanPreset,
			LowlandHumanPreset,
		})
	case culture.AbsAncientDravidian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RidgeHumanPreset,
		})
	case culture.AbsAncientPracyan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{CoastalHumanPreset})
	case culture.AbsAncientTibetan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			LowlandHumanPreset,
		})
	}

	switch c.Abstuct.Name {
	case culture.AbsMedievalAkan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArchipelagoHumanPreset})
	case culture.AbsMedievalArabic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			FoothillHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsMedievalBaltic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			PeninsulaHumanPreset,
			ValleyHumanPreset,
			DeltaHumanPreset,
		})
	case culture.AbsMedievalBaltoFinnic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			PeninsulaHumanPreset,
			DeltaHumanPreset,
		})
	case culture.AbsMedievalBerber.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RiverineHumanPreset,
		})
	case culture.AbsMedievalBrythonic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			HighlandHumanPreset,
			CoastalHumanPreset,
			ValleyHumanPreset,
		})
	case culture.AbsMedievalBurman.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RidgeHumanPreset,
		})
	case culture.AbsMedievalByzantine.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RiverineHumanPreset,
			WetlandHumanPreset,
		})
	case culture.AbsMedievalCentralAfrican.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArchipelagoHumanPreset})
	case culture.AbsMedievalCentralGermanic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			HighlandHumanPreset,
			MontaneHumanPreset,
			ValleyHumanPreset,
			DeltaHumanPreset,
		})
	case culture.AbsMedievalChinese.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			LowlandHumanPreset,
		})
	case culture.AbsMedievalDravidian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
		})
	case culture.AbsMedievalEastAfrican.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArchipelagoHumanPreset})
	case culture.AbsMedievalEastSlavic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MontaneHumanPreset,
			CoastalHumanPreset,
			PeninsulaHumanPreset,
		})
	case culture.AbsMedievalFrankish.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			HighlandHumanPreset,
			ValleyHumanPreset,
		})
	case culture.AbsMedievalGoidelic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			HighlandHumanPreset,
			ValleyHumanPreset,
		})
	case culture.AbsMedievalGuineanUplander.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArchipelagoHumanPreset})
	case culture.AbsMedievalHornAfrican.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArchipelagoHumanPreset})
	case culture.AbsMedievalIberian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			HighlandHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsMedievalIndoAryan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MontaneHumanPreset,
			CoastalHumanPreset,
		})
	case culture.AbsMedievalIranian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			FoothillHumanPreset,
		})
	case culture.AbsMedievalIsraelite.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			FoothillHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsMedievalLatin.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			MontaneHumanPreset,
			RiverineHumanPreset,
			WetlandHumanPreset,
		})
	case culture.AbsMedievalMagyar.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			LowlandHumanPreset,
			MontaneHumanPreset,
		})
	case culture.AbsMedievalMongolic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			LowlandHumanPreset,
		})
	case culture.AbsMedievalNigerDelta.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArchipelagoHumanPreset})
	case culture.AbsMedievalNorthGermanic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			DeltaHumanPreset,
			ValleyHumanPreset,
		})
	case culture.AbsMedievalQiangic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			LowlandHumanPreset,
		})
	case culture.AbsMedievalSahelian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			RiverineHumanPreset,
			ArchipelagoHumanPreset,
		})
	case culture.AbsMedievalSenegambian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{ArchipelagoHumanPreset})
	case culture.AbsMedievalSouthSlavic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			MontaneHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsMedievalTibetan.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			LowlandHumanPreset,
		})
	case culture.AbsMedievalTocharian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{CoastalHumanPreset})
	case culture.AbsMedievalTurkic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			FoothillHumanPreset,
			RiverineHumanPreset,
		})
	case culture.AbsMedievalUgroPermian.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			PeninsulaHumanPreset,
			LowlandHumanPreset,
		})
	case culture.AbsMedievalVlach.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			MontaneHumanPreset,
		})
	case culture.AbsMedievalVolgaFinnic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			LowlandHumanPreset,
		})
	case culture.AbsMedievalWestGermanic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			ValleyHumanPreset,
			CoastalHumanPreset,
			DeltaHumanPreset,
		})
	case culture.AbsMedievalWestSlavic.Name:
		return tools.RandomValueOfSlice(pm.RandFloat64, []gene.Gene{
			CoastalHumanPreset,
			MontaneHumanPreset,
			DeltaHumanPreset,
		})
	}

	return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not get human gene preset for culture (culture abstract=%s, culture=%+v)", c.Abstuct.Name, c))
}
