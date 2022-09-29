package person

import (
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/gender"
)

type Metadata struct {
	WishGetMarriedCoef float64
	DeathCoef          float64
}

func GetWishGetMarriedCoef(sex gender.Sex, age int) float64 {
	switch {
	case age < 9:
		return 0
	case 9 < age && age <= 12:
		if sex.IsMale() {
			return 0.05
		}
		return 0.1
	case 12 < age && age <= 14:
		if sex.IsMale() {
			return 0.25
		}
		return 0.75
	case 14 < age && age <= 16:
		if sex.IsMale() {
			return 0.75
		}
		return 1
	case 16 < age && age <= 18:
		return 1
	case 18 < age && age <= 21:
		if sex.IsMale() {
			return 1.05
		}
		return 1.1
	case 21 < age && age <= 25:
		if sex.IsMale() {
			return 1.1
		}
		return 1.25
	case 25 < age && age <= 30:
		if sex.IsMale() {
			return 1.25
		}
		return 1.5
	case 30 < age && age <= 40:
		if sex.IsMale() {
			return 1.5
		}
		return 2
	case 40 < age && age <= 50:
		if sex.IsMale() {
			return 1
		}
		return 1.5
	case 50 < age && age <= 60:
		if sex.IsMale() {
			return 0.75
		}
		return 1
	case 60 < age:
		return 0.01
	}

	return 0
}

func GetMarriageDistanceCoef(distance coordinate.ComplexDistance) float64 {
	if distance.Value == 0 || distance.MaxValue == 0 {
		return 1
	}

	rel := distance.Value / distance.MaxValue
	if distance.MaxValue <= 15 {
		if distance.Value < 2 {
			return 0.75
		}
		if rel <= 0.5 {
			return 0.5
		}

		return 1 - rel
	}

	return (2 - rel) / 2
}

func GetDeathCoef(age int) float64 {
	switch {
	case age < 3:
		return 0.1
	case age < 9:
		return 0.01
	case 9 < age && age <= 18:
		return 0.001
	case 18 < age && age <= 35:
		return 0.003
	case 35 < age && age <= 45:
		return 0.005
	case 45 < age && age <= 60:
		return 0.01
	case 60 < age && age <= 80:
		return 0.025
	case 80 < age && age <= 90:
		return 0.1
	case 90 < age && age <= 100:
		return 0.15
	case 100 < age && age <= 120:
		return 0.3
	case 120 < age:
		return 0.99
	}

	return 0
}
