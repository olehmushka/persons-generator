package coordinate

func InContain(list []*Coordinate, trg *Coordinate) bool {
	for _, c := range list {
		if c.X == trg.X && c.Y == trg.Y {
			return true
		}
	}

	return false
}
