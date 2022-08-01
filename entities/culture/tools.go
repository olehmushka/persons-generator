package culture

import (
	pm "persons_generator/probability_machine"
	"persons_generator/tools"
)

func GetReligionCultures(amount int, cultures []*Culture) []*Culture {
	if len(cultures) == 0 {
		return []*Culture{}
	}
	out := make([]*Culture, amount)
	if amount >= len(cultures) {
		for i := range out {
			if len(cultures) > i {
				out[i] = cultures[i]
				continue
			}
			out[i] = tools.RandomValueOfSlice(pm.RandFloat64, cultures)
		}
		return out
	}

	shuffledCultures := tools.Shuffle(cultures)
	for i := range out {
		out[i] = shuffledCultures[i]
	}

	return out
}
