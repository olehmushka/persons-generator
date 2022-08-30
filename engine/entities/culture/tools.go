package culture

import (
	"persons_generator/core/tools"
	pm "persons_generator/engine/probability_machine"
)

func GetReligionCultures(amount int, cultures []*Culture) ([]*Culture, error) {
	if len(cultures) == 0 {
		return []*Culture{}, nil
	}
	out := make([]*Culture, amount)
	if amount >= len(cultures) {
		for i := range out {
			if len(cultures) > i {
				out[i] = cultures[i]
				continue
			}
			c, err := tools.RandomValueOfSlice(pm.RandFloat64, cultures)
			if err != nil {
				return nil, err
			}
			out[i] = c
		}
		return out, nil
	}

	shuffledCultures := tools.Shuffle(cultures)
	for i := range out {
		out[i] = shuffledCultures[i]
	}

	return out, nil
}
