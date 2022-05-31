package nature

import (
	"math"

	probabilitymachine "persons_generator/probability-machine"
)

type World struct {
	Size                     int
	Cells                    [][]*Cell
	ContinentsNumber         int
	ExistingContinentsNumber int
	MinHeight                int
	MaxHeight                int
}

func NewWorld(size int, continentsNumber int) *World {
	w := &World{
		Size:             size,
		ContinentsNumber: continentsNumber,
		Cells:            make([][]*Cell, size),
	}
	w.worldSeedCoverage()
	w.worldContinentCoverage(size/10, size/10)

	return w
}

func (w *World) worldSeedCoverage() {
	for y := 0; y < w.Size; y++ {
		cells := make([]*Cell, w.Size)
		for x := 0; x < w.Size; x++ {
			cell := w.generateSell(x, y, 2, 0.1)
			cells = append(cells, cell)
		}
		w.Cells = append(w.Cells, cells)
	}
}

func (w *World) worldContinentCoverage(startStage, stage int) {
	if stage == 0 {
		return
	}

	for y := 0; y < w.Size; y++ {
		cells := make([]*Cell, w.Size)
		for x := 0; x < w.Size; x++ {
			groundCoef := math.Abs(1 - 0.1*float64(startStage-stage))
			cells = append(cells, w.generateSell(x, y, 1, groundCoef))
		}
		w.Cells = append(w.Cells, cells)
	}

	w.worldContinentCoverage(startStage, stage-1)
}

func (w *World) generateSell(x, y int, waterCoef, groundCoef float64) *Cell {
	if x == 0 || y == 0 || x == w.ContinentsNumber || y == w.ExistingContinentsNumber {
		return &Cell{HeightAboveWaterline: w.generateWater(x, y)}
	}
	isGround := probabilitymachine.GetRandomBool((groundCoef / (waterCoef + groundCoef)) * 100)
	if isGround {
		return &Cell{HeightAboveWaterline: w.generateGround(x, y)}
	}
	return &Cell{HeightAboveWaterline: w.generateWater(x, y)}
}

func (w *World) generateWater(x, y int) int {
	if x == 0 || y == 0 || x == w.ContinentsNumber || y == w.ExistingContinentsNumber {
		return w.MinHeight
	}
	return 0
}

func (w *World) generateGround(x, y int) int {
	return 0
}
