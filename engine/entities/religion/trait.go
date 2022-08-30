package religion

import (
	"fmt"
	"net/http"
	"persons_generator/core/wrapped_error"
)

type trait struct {
	_religionMetadata *religionMetadata
	baseCoef          float64

	Name string
	Calc func(r *Religion, self *trait, selectedTraits []*trait) (bool, error)
}

type generateTraitsOpts struct {
	Label string
	Min   int
	Max   int
}

func generateTraits(r *Religion, traitsToSelect []*trait, opts generateTraitsOpts) ([]*trait, error) {
	if opts.Min < 0 {
		return nil, wrapped_error.New(http.StatusInternalServerError, nil, fmt.Sprintf("[%s] min can not be less than 0", opts.Label))
	}
	if opts.Max > len(traitsToSelect) {
		return nil, wrapped_error.New(http.StatusInternalServerError, nil, fmt.Sprintf("[%s] max can not be greater than traitsToSelect length", opts.Label))
	}

	traits := make([]*trait, 0, len(traitsToSelect))
	for count := 0; count < 20; count++ {
		for _, t := range traitsToSelect {
			isTrue, err := t.Calc(r, t, traits)
			if err != nil {
				return nil, err
			}
			if isTrue {
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
		md, err := UpdateReligionMetadata(r, *r.metadata, *t._religionMetadata)
		if err != nil {
			return nil, err
		}
		r.UpdateMetadata(md)
	}

	return traits, nil
}
