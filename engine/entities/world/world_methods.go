package world

import (
	"fmt"
	"math"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person"
	"persons_generator/engine/entities/religion"
	pm "persons_generator/engine/probability_machine"
	"time"
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
		cultures, err := culture.NewMany(culture.Config{}, culturesAmount, nil)
		if err != nil {
			return nil, err
		}
		references, err := religion.NewReferences(
			religion.Config{},
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

func (w *World) CalculatePersonsNumber() int {
	var out int
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] != nil || len(w.Locations[y][x].Population) > 0 {
				out += len(w.Locations[y][x].Population)
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
	Duration       string  `json:"duration"`
}

func (w *World) RunWorld(stopYear int, progressCh chan ProgressRunWorld, errCh chan error) {
	if stopYear <= 0 {
		errCh <- wrapped_error.NewInternalServerError(nil, "can not run world for 0 or less stop year")
		return
	}

	now := time.Now()
	for year := 1; year <= stopYear; year++ {
		if err := w.RunYear(); err != nil {
			errCh <- wrapped_error.NewInternalServerError(err, "can not run year in run world")
			return
		}

		progressCh <- ProgressRunWorld{
			Year:           year,
			Population:     w.PopulationNumber,
			DeadPopulation: w.DeadPopulationNumber,
			Progress:       calcProgress(year, stopYear),
			Duration:       time.Since(now).String(),
		}

		if w.Year == stopYear {
			break
		}
	}

	errCh <- nil
}

func calcProgress(year, stopYear int) float64 {
	out := math.Floor(100 * (float64(year) / float64(stopYear)))
	if year == stopYear {
		out = 100
	}

	return out
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
	w.PopulationNumber = w.CalculatePersonsNumber()

	return nil
}

func (w *World) AppendNewPopulation(people []*person.Person, c *coordinate.Coordinate) error {
	availableSlots, err := w.GetPersonsSlotsPerLoc(c)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not get number of available persons slots")
	}
	if availableSlots <= 0 {
		return nil
	}

	w.Locations[c.Y][c.X].Population = append(w.Locations[c.Y][c.X].Population, people...)
	w.PopulationNumber += len(people)

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
	skip, err := pm.GetRandomBool(0.2)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not generate skip year probability")
	}
	if skip {
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
			isSPregnant, err := s.IsPregnant()
			if err != nil {
				return wrapped_error.NewInternalServerError(err, "can not check if spounce is pregnant")
			}
			if isSPregnant {
				continue
			}

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

				doesWantBeMarried, err := p.DoesWantBeMarried(
					el,
					coordinate.CalcComplexDistance(c, &coordinate.Coordinate{X: x, Y: y}, w.MaxDistanceValue),
					w.religionsSimilarity,
					w.culturesSimilarity,
				)
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

func (w *World) collectDeadPopulation() error {
	dead := make(map[string]*coordinate.Coordinate, 100)
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not run year for <nil> location (x: %d, y: %d)", x, y))
			}

			for i := range w.Locations[y][x].Population {
				isDead, err := w.Locations[y][x].Population[i].IsDead()
				if err != nil {
					return wrapped_error.NewInternalServerError(err, "can not check if person is dead")
				}
				if isDead {
					dead[w.Locations[y][x].Population[i].ID] = w.Locations[y][x].Coordinate
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

func (w *World) appendPersonToDeathWorld(pID string, c *coordinate.Coordinate) error {
	var p *person.Person
	var index int
	for i := range w.Locations[c.Y][c.X].Population {
		if w.Locations[c.Y][c.X].Population[i].ID == pID {
			p = w.Locations[c.Y][c.X].Population[i]
			index = i
		}
	}
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not find person by id (id=%s)", pID))
	}

	w.DeathWorldLocations[c.Y][c.X].Population = append(w.DeathWorldLocations[c.Y][c.X].Population, p)
	w.Locations[c.Y][c.X].Population = tools.Pop(w.Locations[c.Y][c.X].Population, index)

	w.PopulationNumber += -1
	w.DeadPopulationNumber += 1

	return nil
}

func (w *World) GetPersonsSlotsPerLoc(c *coordinate.Coordinate) (int, error) {
	if c == nil {
		return 0, wrapped_error.NewInternalServerError(nil, "can not get persons slots per loc for <nil> coordinate")
	}
	if w.Locations[c.Y][c.X] == nil {
		return 0, wrapped_error.NewInternalServerError(nil, "can not get persons slots per loc for <nil> location")
	}

	return w.MaxPersonsNumberPerLoc - len(w.Locations[c.Y][c.X].Population), nil
}
