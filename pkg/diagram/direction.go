package diagram

type Direction string

const (
	TopToBottom Direction = "TB"
	BottomToTop Direction = "BT"
	LeftToRight Direction = "LR"
	RightToLeft Direction = "RL"
)

func Directions() []Direction {
	return []Direction{TopToBottom, BottomToTop, LeftToRight, RightToLeft}
}

func validateDirection(d Direction) bool {
	dirs := Directions()

	for _, vd := range dirs {
		if vd == d {
			return true
		}
	}

	return false
}

func curveSytles() []string {
	return []string{"ortho", "curved"}
}

func outputFormats() []string {
	return []string{"png", "jpg", "svg", "pdf"}
}
