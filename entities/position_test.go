package entities

import (
	"testing"
)

func TestPath(t *testing.T) {
	post1 := Position{X: 2, Y: 3}
	post2 := Position{X: 6, Y: 7}

	t1 := 0.5

	reslt := Path(post1, post2, t1)

	if reslt.X != 4 || reslt.Y != 5 {
		t.Errorf("Expected (4, 5), got (%f, %f)", reslt.X, reslt.Y)
	}
}
