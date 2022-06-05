package probabilitymachine

func PrepareProbability(pp int) int {
	if pp < 0 {
		return 0
	}
	if pp > 100 {
		return 100
	}

	return pp
}
