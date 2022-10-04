package coordinate

import "math"

type Coordinate struct {
	X int `json:"x" bson:"x"`
	Y int `json:"y" bson:"y"`
}

func CalcComplexDistance(p1, p2 *Coordinate, maxValue float64) ComplexDistance {
	return ComplexDistance{
		Value:    CalcDistanceValue(p1, p2),
		MaxValue: maxValue,
	}
}

func CalcDistanceValue(p1, p2 *Coordinate) float64 {
	var (
		xCalc    = p2.X - p1.X
		yCalc    = p2.Y - p1.Y
		sum      = math.Pow(float64(xCalc), 2) + math.Pow(float64(yCalc), 2)
		distance = math.Sqrt(sum)
	)

	return distance
}
