package location

import "fmt"

func GetLocationsInFrame(frame [][]*Location) int {
	var sum int
	for y := 0; y < len(frame); y++ {
		sum += len(frame[y])
	}

	return sum
}

func SplitLocationsIntoFrames(framesAmount int, locations [][]*Location) [][]*Location {
	locNumber := GetLocationsInFrame(locations)
	if framesAmount > locNumber {
		panic(fmt.Sprintf("[SplitLocationsIntoFrames] location entities (%d) more than frame_amount (%d)", locNumber, framesAmount))
	}
	locPerFrame := locNumber / framesAmount
	reminder := locNumber - (locPerFrame * framesAmount)
	out := make([][]*Location, framesAmount)
	for i := range out {
		l := locPerFrame
		if (i + 1) == framesAmount {
			l += reminder
		}
		out[i] = make([]*Location, l)
	}

	return out
}
