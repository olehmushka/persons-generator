package religion

import (
	"persons_generator/core/tools"
	"persons_generator/engine/entities/culture"
)

type CultureReference struct {
	Religion *Religion
	Culture  *culture.Culture
}

func ExtractCulturesFromCultureReferences(crs []*CultureReference) []*culture.Culture {
	if len(crs) == 0 {
		return []*culture.Culture{}
	}
	cultures := make([]*culture.Culture, 0, len(crs))
	for _, cr := range crs {
		if cr == nil {
			continue
		}
		cultures = append(cultures, cr.Culture)
	}

	return culture.UniqueCultures(cultures)
}

func ExtractReligionsFromCultureReferences(crs []*CultureReference) []*Religion {
	if len(crs) == 0 {
		return []*Religion{}
	}
	religions := make([]*Religion, 0, len(crs))
	for _, cr := range crs {
		if cr == nil {
			continue
		}
		religions = append(religions, cr.Religion)
	}

	return UniqueReligions(religions)
}

func GetReligionByCultureNameFromCultureReferences(crs []*CultureReference, name string) *Religion {
	for _, cr := range crs {
		if cr == nil || cr.Religion == nil || cr.Culture == nil {
			continue
		}
		if tools.ContainString(cr.Culture.Name, name) {
			return cr.Religion
		}
	}

	return nil
}
