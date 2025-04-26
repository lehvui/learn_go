package entities

import (
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

func TestPath(t *testing.T) {
	pos1 := Position{X: 2, Y: 3}
	pos2 := Position{X: 6, Y: 7}

	t1 := 0.5

	reslt := Path(pos1, pos2, t1)

	if reslt.X != 4 || reslt.Y != 5 {
		t.Errorf("Expected (4, 5), got (%f, %f)", reslt.X, reslt.Y)
	}
}
