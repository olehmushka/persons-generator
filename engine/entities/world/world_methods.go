package world

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person"
	"persons_generator/engine/entities/religion"
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

func (w *World) GetPersons(params GetHumansParams) ([]*person.Person, error) {
	out := make([]*person.Person, 0, w.Size*w.Size)
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] != nil || len(w.Locations[y][x].Population) > 0 {
				out = append(out, w.Locations[y][x].Population...)
			}
		}
	}

	return out, nil
}
