package person

import (
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/person/traits"
)

func (p *Person) HaveSex(other *Person) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "can not <nil> have sex with someone else")
	}
	if other == nil {
		return wrapped_error.NewInternalServerError(nil, "can not person have sex with <nil>")
	}
	if p.Human.Sex == other.Human.Sex {
		p.AppendTrait(traits.HomosexualistTrait)
		other.AppendTrait(traits.HomosexualistTrait)
	}
	if err := p.Human.HaveSex(other.Human); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not have sex")
	}

	return nil
}

func (p *Person) AppendTrait(t *traits.Trait) {
	p.Traits = append(p.Traits, t)
	p.Traits = traits.Unique(p.Traits)
}
