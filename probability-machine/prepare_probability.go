package probabilitymachine

func PrepareProbability(pp float64) float64 {
	if pp < 0 {
		return 0
	}
	if pp > 1 {
		return 1
	}

	return pp
}
