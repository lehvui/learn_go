package entities

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	pos1 := Position{X: 6, Y: 7}
	pos2 := Position{X: 2, Y: 3}

	dir := Sub(pos1, pos2)

	assert.Equal(t, dir.X, 4.0)
	assert.Equal(t, dir.Y, 4.0)
}

func TestDistance(t *testing.T) {
	pos1 := Position{X: 3, Y: 4}
	pos2 := Position{X: 3, Y: 5}

	assert.Equal(t, 1.0, Distance(pos1, pos2))
}

func TestScaled(t *testing.T) {

	start := Position{100, 50}
	end := Position{307, 223}

	long := 50
	dir := Sub(end, start)
	distance := Distance(start, end)

	magX := dir.X / distance
	magY := dir.Y / distance

	newDir := Position{magX * float64(long), magY * float64(long)}

	lastDir := Add(start, newDir)

	fmt.Println(lastDir)
}

func TestPath(t *testing.T) {
	pos1 := Position{X: 2, Y: 3}
	pos2 := Position{X: 6, Y: 7}

	t1 := 0.5

	reslt := Path(pos1, pos2, t1)

	if reslt.X != 4 || reslt.Y != 5 {
		t.Errorf("Expected (4, 5), got (%f, %f)", reslt.X, reslt.Y)
	}
}
