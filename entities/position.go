package entities

type Position struct {
	X, Y float64
}

func Path(position1, position2 Position, t float64) Position {

	lerpX := position1.X + (position2.X-position1.X)*t
	lerpY := position1.Y + (position2.Y-position1.Y)*t

	return Position{X: lerpX, Y: lerpY}
}
