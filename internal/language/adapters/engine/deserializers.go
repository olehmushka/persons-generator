package engine

import (
	"persons_generator/engine/entities/language"
	"persons_generator/internal/language/entities"
)

func deserializeLanguage(in *entities.Language) *language.Language {
	if in == nil {
		return nil
	}

	wb := &language.WordBase{}
	if in.WordBase != nil {
		wb = &language.WordBase{
			FemaleOwnNames: in.WordBase.FemaleOwnNames,
			MaleOwnNames:   in.WordBase.MaleOwnNames,
			Words:          in.WordBase.Words,
			Name:           in.WordBase.Name,
			Min:            in.WordBase.Min,
			Max:            in.WordBase.Max,
			Dupl:           in.WordBase.Dupl,
			M:              in.WordBase.M,
		}
	}

	return &language.Language{
		ID:          in.ID,
		Name:        in.Name,
		Subfamily:   deserializeSubfamily(in.Subfamily),
		WordBaseRef: nil,
		WordBase:    wb,
		IsLiving:    in.IsLiving,
	}
}

func deserializeSubfamily(in *entities.Subfamily) *language.Subfamily {
	if in == nil {
		return nil
	}

	return &language.Subfamily{
		Name:              in.Name,
		Family:            language.Family(in.FamilyName),
		ExtendedSubfamily: deserializeSubfamily(in.ExtendedSubfamily),
		IsLiving:          in.IsLiving,
	}
}
