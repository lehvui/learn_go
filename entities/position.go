package entities

import "math"

type Position struct {
	X, Y float64
}

func Path(position1, position2 Position, t float64) Position {

	lerpX := position1.X + (position2.X-position1.X)*t
	lerpY := position1.Y + (position2.Y-position1.Y)*t

	return Position{X: lerpX, Y: lerpY}
}

func Sub(pos1, pos2 Position) Position {
	return Position{pos1.X - pos2.X, pos1.Y - pos2.Y}
}

func Add(pos1, pos2 Position) Position {
	return Position{pos1.X + pos2.X, pos1.Y + pos2.Y}
}

func Equal(pos1, pos2 Position) bool {
	if int32(pos1.X) == int32(pos2.X) && int32(pos1.Y) == int32(pos2.Y) {
		return true
	}

	return false
}

func Distance(pos1, pos2 Position) float64 {
	dir := Sub(pos1, pos2)

	return math.Sqrt(dir.X*dir.X + dir.Y*dir.Y)
}
