package entities

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	*Sprite
	Pos  Position
	Dir  Position
	Life float64
}

func NewProjectile(path string, img, start, end Position) *Projectile {
	s := NewNumberSprite(path, img.X, img.Y, 2, 23)

	long := 50
	dir := Sub(end, start)
	distance := Distance(start, end)

	magX := dir.X / distance
	magY := dir.Y / distance

	newDir := Position{magX * float64(long), magY * float64(long)}

	return &Projectile{
		Sprite: s,
		Pos:    start,
		Dir:    newDir,
	}
}

func (p *Projectile) Update() {

	if p.Life < 1 {
		p.Pos = Path(p.Pos, p.Dir, p.Life)
		p.Sprite.X = p.Pos.X
		p.Sprite.Y = p.Pos.Y
		p.Life += 0.005
	}
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	fmt.Println(p.Life)
	p.Sprite.Draw(screen)
}

func (p *Projectile) Alive() bool {
	return p.Life < 1.0
}
