package engine

import (
	"persons_generator/engine/entities/language"
	"persons_generator/internal/language/entities"
)

func serializeLanguages(in []*language.Language) []*entities.Language {
	out := make([]*entities.Language, len(in))
	for i := range out {
		out[i] = serializeLanguage(in[i])
	}

	return out
}

func serializeLanguage(in *language.Language) *entities.Language {
	if in == nil {
		return nil
	}

	wb := &entities.WordBase{}
	if in.WordBase != nil {
		wb = &entities.WordBase{
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

	return &entities.Language{
		ID:        in.ID,
		Name:      in.Name,
		Subfamily: serializeSubfamily(in.Subfamily),
		WordBase:  wb,
		IsLiving:  in.IsLiving,
	}
}

func serializeSubfamilies(in []*language.Subfamily) []*entities.Subfamily {
	out := make([]*entities.Subfamily, len(in))
	for i := range out {
		out[i] = serializeSubfamily(in[i])
	}

	return out
}

func serializeSubfamily(in *language.Subfamily) *entities.Subfamily {
	if in == nil {
		return nil
	}

	return &entities.Subfamily{
		Name:              in.Name,
		FamilyName:        in.Family.String(),
		ExtendedSubfamily: serializeSubfamily(in.ExtendedSubfamily),
		IsLiving:          in.IsLiving,
	}
}
