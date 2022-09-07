package color

func getColorsByPalette(palette string, colors []Color) []Color {
	out := make([]Color, 0, len(colors))
	for _, c := range colors {
		if c.Palette == palette {
			out = append(out, c)
		}
	}

	return out
}
