package religion

import "fmt"

type trait struct {
	_religionMetadata *religionMetadata
	baseCoef          float64

	Name string
	Calc func(r *Religion, self *trait, selectedTraits []*trait) bool
}

type generateTraitsOpts struct {
	Label string
	Min   int
	Max   int
}

func generateTraits(r *Religion, traitsToSelect []*trait, opts generateTraitsOpts) []*trait {
	if opts.Min < 0 {
		panic(fmt.Sprintf("[%s] min can not be less than 0", opts.Label))
	}
	if opts.Max > len(traitsToSelect) {
		panic(fmt.Sprintf("[%s] max can not be greater than traitsToSelect length", opts.Label))
	}

	traits := make([]*trait, 0, len(traitsToSelect))
	for count := 0; count < 20; count++ {
		for _, t := range traitsToSelect {
			if t.Calc(r, t, traits) {
				traits = append(traits, &trait{
					_religionMetadata: t._religionMetadata,
					Name:              t.Name,
					Calc:              t.Calc,
				})
			}
			if len(traits) == opts.Max {
				break
			}
		}
		if len(traits) == opts.Max {
			break
		}
		if len(traits) >= opts.Min {
			break
		}
	}

	for _, t := range traits {
		r.UpdateMetadata(UpdateReligionMetadata(r, *r.metadata, *t._religionMetadata))
	}

	return traits
}
