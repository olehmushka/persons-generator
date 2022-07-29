package culture

import (
	pm "persons_generator/probability_machine"
	"persons_generator/tools"
)

func GetReligionNames(amount int, cultures []*Culture) []string {
	if len(cultures) == 0 {
		return []string{}
	}
	out := make([]string, amount)
	if amount >= len(cultures) {
		for i := range out {
			if len(cultures) > i {
				out[i] = cultures[i].Language.GetReligionName()
				continue
			}
			out[i] = tools.RandomValueOfSlice(pm.RandFloat64, cultures).Language.GetReligionName()
		}
		return out
	}

	shuffledCultures := tools.Shuffle(cultures)
	for i := range out {
		out[i] = shuffledCultures[i].Language.GetReligionName()
	}

	return out
}
