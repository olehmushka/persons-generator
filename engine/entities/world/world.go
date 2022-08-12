package world

import (
	"errors"
	"fmt"

	"persons_generator/core/tools"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/human/human"
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/religion"
	pm "persons_generator/engine/probability_machine"
)

type World struct {
	Size      int
	Locations [][]*location.Location
	Cultures  []*culture.Culture
	Religions []*religion.Religion
}

func New(s int) *World {
	return &World{
		Size: s,
	}
}

func (w *World) Fill() *World {
	size := w.Size
	w.Locations = make([][]*location.Location, 0, size)
	for y := 0; y < size; y++ {
		w.Locations = append(w.Locations, make([]*location.Location, 0, size))
		for x := 0; x < size; x++ {
			w.Locations[y] = append(w.Locations[y], &location.Location{
				Coordinate: &coordinate.Coordinate{X: x, Y: y},
				Population: make([]*human.Human, 0),
			})
		}
	}

	return w
}

func (w *World) CulturesPropagate(amount int, preferred []string) (*World, error) {
	if amount == 0 {
		return nil, errors.New("[World.CulturesPropagate] culture amount can not be zero")
	}

	totalLocs := w.Size * w.Size
	if totalLocs < amount {
		return nil, fmt.Errorf("[World.CulturesPropagate] incorrect cultures number (max culture_number=%d, actual culture_number=%d)", totalLocs, amount)
	}
	cultures, err := culture.NewCultures(amount, preferred)
	if err != nil {
		return nil, err
	}
	w.Cultures = cultures

	var (
		locsPerCulture = totalLocs / amount
		reminder       = totalLocs - (locsPerCulture * amount)
		toFillCultures = make(map[string]int, amount)
	)
	for i, c := range w.Cultures {
		toFillCultures[c.Name] = locsPerCulture
		if i == 0 {
			toFillCultures[c.Name] += reminder
		}
	}

	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			var cultureName string
			for {
				randCultureName, err := culture.GetRandomCultureName(w.Cultures)
				if err != nil {
					return nil, fmt.Errorf("[World.CulturesPropagate] can not generate random culture_name(error=%s)", err.Error())
				}
				if rem, ok := toFillCultures[randCultureName]; ok && rem > 0 {
					cultureName = randCultureName
					break
				}
			}
			w.Locations[y][x].InitCulture = culture.GetCultureByName(cultureName, w.Cultures)
			toFillCultures[cultureName]--
		}
	}

	return w, nil
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

func (w *World) ReligionsPropagate(amount int) (*World, error) {
	if amount == 0 {
		return nil, errors.New("[World.ReligionsPropagate] religion amount can not be zero")
	}
	totalLocs := w.Size * w.Size
	if totalLocs < amount {
		return nil, fmt.Errorf("[World.ReligionsPropagate] incorrect religion number (max religion_number=%d, actual religion_number=%d)", totalLocs, amount)
	}

	if amount > len(w.Cultures) {
		return w.religionsPropagateForCultureNumberLess(amount)
	}
	if amount == len(w.Cultures) {
		return w.religionsPropagateForCultureNumberEqual(amount)
	}
	if amount < len(w.Cultures) {
		return w.religionsPropagateForCultureNumberGreater(amount)
	}

	return w, nil
}

func (w *World) religionsPropagateForCultureNumberLess(amount int) (*World, error) {
	var (
		sampleReligions      = make([]*religion.Religion, amount)
		religions            = make([]*religion.Religion, 0, amount)
		chunkSampleReligions = tools.ChunkFor(sampleReligions, len(w.Cultures))
		cultureReligionMap   = make(map[string][]string, amount)
	)

	for i := range chunkSampleReligions {
		for range chunkSampleReligions[i] {
			c := w.Cultures[i]
			r, err := religion.NewReligion(c)
			if err != nil {
				return nil, err
			}
			religions = append(religions, r)
			val, ok := cultureReligionMap[c.Name]
			if !ok || len(val) == 0 {
				cultureReligionMap[c.Name] = []string{r.Name}
			} else {
				cultureReligionMap[c.Name] = append(cultureReligionMap[c.Name], r.Name)
			}
		}
	}
	w.Religions = religions

	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil || w.Locations[y][x].InitCulture == nil {
				return nil, fmt.Errorf("[World.religionsPropagateForCultureNumberLess] location is nil or its culture is nil (x: %d, y: %d)", x, y)
			}

			cultureName := w.Locations[y][x].InitCulture.Name
			religionNames, ok := cultureReligionMap[cultureName]
			if !ok {
				return nil, fmt.Errorf("[World.religionsPropagateForCultureNumberLess] can not get culture by name (list=%+v, name=%s)", culture.MapCultureNames(w.Cultures), cultureName)
			}
			if len(religionNames) == 0 {
				return nil, fmt.Errorf("[World.religionsPropagateForCultureNumberLess] can not get religion names from empty list (culture_religion_map=%+v, name=%s)", cultureReligionMap, cultureName)
			}
			religionName := tools.RandomValueOfSlice(pm.RandFloat64, religionNames)
			w.Locations[y][x].InitReligion = religion.GetReligionByName(religionName, religions)
		}
	}

	return w, nil
}

func (w *World) religionsPropagateForCultureNumberEqual(amount int) (*World, error) {
	var (
		religions          = make([]*religion.Religion, 0, amount)
		cultureReligionMap = make(map[string]string, amount)
	)

	for _, c := range w.Cultures {
		r, err := religion.NewReligion(c)
		if err != nil {
			return nil, err
		}
		religions = append(religions, r)
		cultureReligionMap[c.Name] = r.Name
	}
	w.Religions = religions

	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil || w.Locations[y][x].InitCulture == nil {
				return nil, fmt.Errorf("[World.religionsPropagateForCultureNumberEqual] location is nil or its culture is nil (x: %d, y: %d)", x, y)
			}

			cultureName := w.Locations[y][x].InitCulture.Name
			religionName, ok := cultureReligionMap[cultureName]
			if !ok {
				return nil, fmt.Errorf("[World.religionsPropagateForCultureNumberEqual] can not get culture by name (list=%+v, name=%s)", culture.MapCultureNames(w.Cultures), cultureName)
			}
			w.Locations[y][x].InitReligion = religion.GetReligionByName(religionName, religions)
		}
	}

	return w, nil
}

func (w *World) religionsPropagateForCultureNumberGreater(amount int) (*World, error) {
	var (
		religions          = make([]*religion.Religion, 0, amount)
		culturesChunks     = tools.ChunkFor(w.Cultures, amount)
		cultureReligionMap = make(map[string]string, amount)
	)

	for _, chunk := range culturesChunks {
		hybridCulture, err := culture.NewWithProto(chunk)
		if err != nil {
			return nil, err
		}
		r, err := religion.NewReligion(hybridCulture)
		if err != nil {
			return nil, err
		}
		religions = append(religions, r)
		for _, c := range chunk {
			cultureReligionMap[c.Name] = r.Name
		}
	}
	w.Religions = religions

	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil || w.Locations[y][x].InitCulture == nil {
				return nil, fmt.Errorf("[World.religionsPropagateForCultureNumberGreater] location is nil or its culture is nil (x: %d, y: %d)", x, y)
			}

			cultureName := w.Locations[y][x].InitCulture.Name
			religionName, ok := cultureReligionMap[cultureName]
			if !ok {
				return nil, fmt.Errorf("[World.religionsPropagateForCultureNumberGreater] can not get culture by name (list=%+v, name=%s)", culture.MapCultureNames(w.Cultures), cultureName)
			}
			w.Locations[y][x].InitReligion = religion.GetReligionByName(religionName, religions)
		}
	}

	return w, nil
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
