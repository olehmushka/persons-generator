package world

import (
	"math"
	"persons_generator/engine/entities/coordinate"
)

func GetMaxDistanceValue(size int) float64 {
	return coordinate.CalcDistanceValue(
		&coordinate.Coordinate{X: 0, Y: 0},
		&coordinate.Coordinate{X: size, Y: size},
	)
}

func GetSizeByHumanAmount(humanAmount int) int {
	squareSide := math.Sqrt(float64(humanAmount)) * 1.5

	return int(math.Ceil(squareSide))
}
