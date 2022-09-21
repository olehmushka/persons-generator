package world

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/person"
)

func (w *World) QueryPerson(opts QueryPersonsOpts) ([]*person.Person, error) {
	var (
		maleLimit   int
		femaleLimit int
		limit       int = 10

		maleCount   int
		femaleCount int
		count       int
	)
	if opts.MaleCount != 0 {
		maleLimit = opts.MaleCount
	}
	if opts.FemaleCount != 0 {
		femaleLimit = opts.FemaleCount
	}
	if opts.PersonsCount != 0 {
		limit = opts.PersonsCount
	}
	if maleLimit <= 0 && femaleLimit <= 0 && limit <= 0 {
		maleLimit = 5
		femaleLimit = 5
		limit = 10
	}
	if sum := maleLimit + femaleLimit; sum > limit {
		limit = sum
	}
	if maleLimit == 0 && femaleLimit == 0 {
		maleLimit = limit / 2
		femaleCount = limit - maleCount
	} else if maleLimit == 0 {
		maleCount = limit - femaleLimit
	} else if femaleLimit == 0 {
		femaleCount = limit - maleLimit
	}

	if opts.MinAge == 0 && opts.MaxAge == 0 {
		opts.MaxAge = 120
	}
	if opts.MinAge >= opts.MaxAge {
		return nil, wrapped_error.NewInternalServerError(nil, "can not query persons for min_age >= max_age")
	}
	if opts.OnlyAlive && opts.OnlyDead {
		return nil, wrapped_error.NewInternalServerError(nil, "can not query persons for only_alive is true & only_dead is true")
	}

	out := make([]*person.Person, 0, limit)
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil {
				return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not query persons for <nil> location (x: %d, y: %d)", x, y))
			}

			for i := range w.Locations[y][x].Population {
				if count == limit {
					fmt.Printf("out of limit (limit=%d)\n", limit)
					return out, nil
				}
				p := w.Locations[y][x].Population[i]
				sex, err := p.Sex()
				if err != nil {
					return nil, wrapped_error.NewInternalServerError(err, "can not get person sex")
				}
				if maleCount == maleLimit {
					if sex.IsMale() {
						fmt.Printf("out of male limit (maleLimit=%d)\n", maleLimit)
						continue
					}
				}
				if femaleCount == femaleLimit {
					if sex.IsFemale() {
						fmt.Printf("out of male limit (femaleLimit=%d)\n", femaleLimit)
						continue
					}
				}
				age, err := p.Age()
				if err != nil {
					return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not get person age for location (x: %d, y: %d)", x, y))
				}
				if opts.MinAge > age || age > opts.MaxAge {
					fmt.Printf("out of age limitation (min=%d,age=%d,max=%d)\n", opts.MinAge, age, opts.MaxAge)
					continue
				}
				isDead, err := p.IsDead()
				if err != nil {
					return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not check if person is dead for location (x: %d, y: %d)", x, y))
				}
				if opts.OnlyAlive && isDead {
					fmt.Printf("out of only_alive limitation (only_alive=%t,is_dead=%t)\n", opts.OnlyAlive, isDead)
					continue
				}
				if opts.OnlyDead && !isDead {
					fmt.Printf("out of OnlyDead limitation (OnlyDead=%t,is_dead=%t)\n", opts.OnlyDead, isDead)
					continue
				}
				if sex.IsMale() {
					maleCount++
				}
				if sex.IsFemale() {
					femaleCount++
				}
				count++
				out = append(out, p)
			}
		}
	}

	return out, nil
}
