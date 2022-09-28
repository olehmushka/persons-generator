package world

import (
	"fmt"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person"
	"persons_generator/engine/entities/religion"
	pm "persons_generator/engine/probability_machine"

	"github.com/google/uuid"
)

func (w *World) preparePreference(in *Preference) (*Preference, error) {
	if in == nil {
		hp := HumanPreference{
			Amount:       w.defaultHumanAmount,
			MaleAmount:   int(float64(w.defaultHumanAmount) * w.defaultMalePercentage),
			FemaleAmount: int(float64(w.defaultHumanAmount) * w.defaultFemalePercentage),
		}
		var culturesAmount, religionsAmount int
		switch {
		case hp.Amount < 10:
			culturesAmount = 1
			religionsAmount = 1
		case hp.Amount >= 10 && hp.Amount < 25:
			culturesAmount = 2
			religionsAmount = 2
		case hp.Amount >= 25 && hp.Amount < 50:
			culturesAmount = 3
			religionsAmount = 3
		case hp.Amount >= 50 && hp.Amount < 75:
			culturesAmount = 5
			religionsAmount = 4
		case hp.Amount >= 75 && hp.Amount < 100:
			culturesAmount = 8
			religionsAmount = 5
		case hp.Amount >= 100 && hp.Amount < 250:
			culturesAmount = 13
			religionsAmount = 6
		case hp.Amount >= 250 && hp.Amount < 500:
			culturesAmount = 21
			religionsAmount = 7
		case hp.Amount >= 500 && hp.Amount < 750:
			culturesAmount = 34
			religionsAmount = 8
		case hp.Amount >= 750 && hp.Amount < 1000:
			culturesAmount = 55
			religionsAmount = 9
		case hp.Amount > 1000:
			culturesAmount = 89
			religionsAmount = 10
		}
		cultures, err := culture.NewMany(culture.Config{StorageFolderName: w.storageFolderName}, culturesAmount, nil)
		if err != nil {
			return nil, err
		}
		references, err := religion.NewReferences(
			religion.Config{StorageFolderName: w.storageFolderName},
			religionsAmount,
			cultures,
		)
		if err != nil {
			return nil, err
		}

		return &Preference{
			Human:            hp,
			ReligionCultures: references,
		}, nil
	}

	return in, nil
}

func getSizeByPreference(pref *Preference) (int, error) {
	if pref == nil {
		return 0, wrapped_error.NewInternalServerError(nil, "can not get size of world without preference")
	}

	return GetSizeByHumanAmount(pref.Human.Amount), nil
}

func (w *World) PrintLocationCultures() {
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if x != 0 {
				fmt.Printf("\t")
			}
			fmt.Printf("%s", w.Locations[y][x].InitCulture.Name)
		}
		fmt.Printf("\n")
	}
}

func (w *World) PrintLocationReligions() {
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if x != 0 {
				fmt.Printf("\t")
			}
			fmt.Printf("%s", w.Locations[y][x].InitReligion.Name)
		}
		fmt.Printf("\n")
	}
}

func (w *World) GetPersons() []*person.Person {
	out := make([]*person.Person, 0, w.Size*w.Size)
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] != nil || len(w.Locations[y][x].Population) > 0 {
				out = append(out, w.Locations[y][x].Population...)
			}
		}
	}

	return out
}

type ProgressRunWorld struct {
	Year           int     `json:"year"`
	Population     int     `json:"population"`
	DeadPopulation int     `json:"dead_population"`
	Progress       float64 `json:"progress"`
}

func (w *World) RunWorld(stopYear int, progressCh chan ProgressRunWorld, errCh chan error) {
	if stopYear <= 0 {
		errCh <- wrapped_error.NewInternalServerError(nil, "can not run world for 0 or less stop year")
		return
	}
	for year := 0; year < stopYear; year++ {
		if err := w.RunYear(); err != nil {
			errCh <- wrapped_error.NewInternalServerError(err, "can not run year in run world")
			return
		}
		progressCh <- ProgressRunWorld{
			Year:           year,
			Population:     w.populationNumber,
			DeadPopulation: w.deadPopulationNumber,
			Progress:       100 * (float64(year) / float64(stopYear)),
		}
		if w.Year == stopYear {
			break
		}
	}

	errCh <- nil
}

func (w *World) RunYear() error {
	w.Year++

	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not run year for <nil> location (x: %d, y: %d)", x, y))
			}

			for i := range w.Locations[y][x].Population {
				if err := w.runYearPerPerson(w.Locations[y][x].Population[i], w.Locations[y][x].Coordinate); err != nil {
					return err
				}
			}
		}
	}

	if err := w.collectDeadPopulation(); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not collect dead population")
	}

	return nil
}

func (w *World) AppendNewPopulation(people []*person.Person, c *coordinate.Coordinate) error {
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not append people into <nil> location (x: %d, y: %d)", x, y))
			}

			w.Locations[y][x].Population = append(w.Locations[y][x].Population, people...)
			w.populationNumber += len(people)
		}
	}

	return nil
}

func (w *World) runYearPerPerson(p *person.Person, c *coordinate.Coordinate) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "can not run year per <nil> person")
	}
	if c == nil {
		return wrapped_error.NewInternalServerError(nil, "can not run year per for <nil> coordinate")
	}

	isDead, err := p.IsDead()
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not check if person is dead")
	}
	if isDead {
		return nil
	}

	if err := p.IncreaseAge(w.Year); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not increase person age")
	}

	age, err := p.Age()
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not get person age")
	}
	if age < 12 {
		return p.TryDie(w.Year)
	}

	isPregnant, err := p.IsPregnant()
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not check if person is pregnant")
	}
	if isPregnant {
		children, err := p.GiveBirth(w.Year)
		if err != nil {
			return wrapped_error.NewInternalServerError(err, "can not give birth for children")
		}
		if err := w.AppendNewPopulation(children, c); err != nil {
			return wrapped_error.NewInternalServerError(err, "can not append children to location population")
		}
	}

	isMarried, err := p.IsMarried()
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not check if person is married")
	}
	if isMarried {
		for _, s := range p.Spouces {
			if err := p.HaveSex(s, w.Year); err != nil {
				return wrapped_error.NewInternalServerError(err, "can not have sex with spounce")
			}
		}

		return p.TryDie(w.Year)
	}
	availablePartners, err := w.seekPartners(p, c)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not get available partners")
	}
	if len(availablePartners) == 0 {
		return p.TryDie(w.Year)
	}
	spounce, err := tools.RandomValueOfSlice(pm.RandFloat64, availablePartners)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not get random spounce")
	}
	if err := p.GetMarried(spounce, w.Year); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not get married")
	}
	if err := p.HaveSex(spounce, w.Year); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not have sex with new spounce")
	}

	return p.TryDie(w.Year)
}

func (w *World) seekPartners(p *person.Person, c *coordinate.Coordinate) ([]*person.Person, error) {
	partners := make([]*person.Person, 0, 10)
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil {
				return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not seek partners in <nil> location (x: %d, y: %d)", x, y))
			}

			for _, el := range w.Locations[y][x].Population {
				age, err := el.Age()
				if err != nil {
					return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not get person age in <nil> location (x: %d, y: %d)", x, y))
				}
				if age < 12 {
					continue
				}

				isMarried, err := el.IsMarried()
				if err != nil {
					return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not check if person is married in <nil> location (x: %d, y: %d)", x, y))
				}
				if isMarried {
					continue
				}

				doesWantBeMarried, err := p.DoesWantBeMarried(el, coordinate.CalcComplexDistance(c, &coordinate.Coordinate{X: x, Y: y}, w.MaxDistanceValue))
				if err != nil {
					return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not check does person can married with other person in <nil> location (x: %d, y: %d)", x, y))
				}
				if doesWantBeMarried {
					partners = append(partners, el)
				}
			}
		}
	}

	return partners, nil
}

func (w *World) moveDeadPersonToDeathWorld(p *person.Person, c *coordinate.Coordinate) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "can not move <nil> person to death world")
	}
	w.DeathWorldLocations[c.Y][c.X].Population = append(w.DeathWorldLocations[c.Y][c.X].Population, p)
	w.populationNumber--

	return nil
}

func (w *World) removeDeadPersonFromWorld(p *person.Person, c *coordinate.Coordinate) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "can not move <nil> person to death world")
	}

	populations := make([]*person.Person, 0, len(w.Locations[c.Y][c.X].Population))
	for i := range w.Locations[c.Y][c.X].Population {
		if w.Locations[c.Y][c.X].Population[i].ID != p.ID {
			populations = append(populations, w.Locations[c.Y][c.X].Population[i])
		}
	}
	w.Locations[c.Y][c.X].Population = populations

	return nil
}

func (w *World) collectDeadPopulation() error {
	dead := make(map[uuid.UUID]*coordinate.Coordinate, 100)
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not run year for <nil> location (x: %d, y: %d)", x, y))
			}

			for i := range w.Locations[y][x].Population {
				p := w.Locations[y][x].Population[i]
				c := w.Locations[y][x].Coordinate
				isDead, err := p.IsDead()
				if err != nil {
					return wrapped_error.NewInternalServerError(err, "can not check if person is dead")
				}
				if isDead {
					dead[p.ID] = c
				}
			}
		}
	}

	for pID, c := range dead {
		if err := w.appendPersonToDeathWorld(pID, c); err != nil {
			return err
		}
	}

	return nil
}

func (w *World) appendPersonToDeathWorld(pID uuid.UUID, c *coordinate.Coordinate) error {
	var p *person.Person
	for i := range w.Locations[c.Y][c.X].Population {
		p = w.Locations[c.Y][c.X].Population[i]
	}

	w.DeathWorldLocations[c.Y][c.X].Population = append(w.DeathWorldLocations[c.Y][c.X].Population, p)
	w.populationNumber--
	w.deadPopulationNumber++

	populations := make([]*person.Person, 0, len(w.Locations[c.Y][c.X].Population))
	for i := range w.Locations[c.Y][c.X].Population {
		if w.Locations[c.Y][c.X].Population[i].ID != pID {
			populations = append(populations, w.Locations[c.Y][c.X].Population[i])
		}
	}
	w.Locations[c.Y][c.X].Population = populations

	return nil
}
