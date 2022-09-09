package person

import (
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/person/traits"
)

func (p *Person) HaveSex(other *Person, year int) error {
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
	p.AppendHaveSexEvent(year)
	if err := other.Human.HaveSex(p.Human); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not have sex")
	}
	other.AppendHaveSexEvent(year)

	return nil
}

func (p *Person) AppendTrait(t *traits.Trait) {
	if p == nil {
		return
	}

	p.Traits = append(p.Traits, t)
	p.Traits = traits.Unique(p.Traits)
}

func (p *Person) appendEvent(e Event) {
	if p == nil {
		return
	}

	p.Chronology.Events = append(p.Chronology.Events, e)
}

func (p *Person) AppendHaveSexEvent(year int) {
	if p == nil {
		return
	}

	p.appendEvent(Event{Year: year, Name: HaveSexEventName})
}

func (p *Person) IsDead() bool {
	if p == nil {
		return false
	}

	return p.Chronology.DeathYear == -1
}

func (p *Person) Die(year int) {
	if p == nil {
		return
	}

	p.Chronology.DeathYear = year
}

func (p *Person) HadSex(year int) bool {
	if p == nil {
		return false
	}

	for _, r := range p.Chronology.Events {
		if r.Year == year && r.Name == HaveSexEventName {
			return true
		}
	}

	return false
}
