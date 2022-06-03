package coordinate

import (
	"math"

	"persons_generator/entities"
)

func CalcDistance(p1, p2 *entities.Coordinate) int {
	var (
		xCalc    = p2.X - p1.X
		yCalc    = p2.Y - p1.Y
		sum      = math.Pow(float64(xCalc), 2) + math.Pow(float64(yCalc), 2)
		distance = math.Sqrt(sum)
	)

	return int(math.Round(distance))
}
