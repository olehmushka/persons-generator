package location

import (
	"fmt"
	"persons_generator/core/wrapped_error"
)

func GetLocationsInFrame(frame [][]*Location) int {
	var sum int
	for y := 0; y < len(frame); y++ {
		sum += len(frame[y])
	}

	return sum
}

func SplitLocationsIntoFrames(framesAmount int, locations [][]*Location) ([][]*Location, error) {
	locNumber := GetLocationsInFrame(locations)
	if framesAmount > locNumber {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("[SplitLocationsIntoFrames] location entities (%d) more than frame_amount (%d)", locNumber, framesAmount))
	}
	locPerFrame := locNumber / framesAmount
	reminder := locNumber - (locPerFrame * framesAmount)
	out := make([][]*Location, framesAmount)
	occupied := map[string]bool{}
	for i := range out {
		l := locPerFrame
		if (i + 1) == framesAmount {
			l += reminder
		}
		frame := make([]*Location, 0, l)
		var cur *Location

		for {
			for y := 0; y < len(locations); y++ {
				for x := 0; x < len(locations[y]); x++ {
					if cur == nil {
						cur = locations[y][x]
					}
					if len(occupied) == 0 {
						occupied[fmt.Sprintf("%d.%d", y, x)] = true
					}
				}
			}
			if len(frame) == l {
				break
			}
		}

		out[i] = frame
	}

	return out, nil
}

/*
Amount 4

total = 6 x 6 = 36
per_frame = 36 / 4 = 9

  1 2 3 4 5 6
A y y y x x x
B y y y x x x
C y y y x x x
D x x x x x x
E x x x x x x
F x x x x x x

*/
