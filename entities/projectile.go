package entities

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	*Sprite
	Pos       Position
	Dir       Position
	ScaledDir Position
	Life      float64
	Alive     bool
}

func NewProjectile(path string, img, start, end Position) *Projectile {
	s := NewNumberSprite(path, img.X, img.Y, 2, 23)

	dx := end.X - start.X
	dy := end.Y - start.Y
	scaledEnd := Position{
		X: start.X + dx*30,
		Y: start.Y + dy*30,
	}

	return &Projectile{
		Sprite:    s,
		Pos:       start,
		Dir:       end,
		ScaledDir: scaledEnd,
		Alive:     true,
	}
}

func (p *Projectile) Update() {
	if math.Hypot(p.Pos.X-p.ScaledDir.X, p.Pos.Y-p.ScaledDir.Y) < 1 {
		p.Alive = false
		return
	}

	if p.Life < 50 {
		p.Pos = Path(p.Pos, p.ScaledDir, p.Life)
		p.Sprite.X = p.Pos.X
		p.Sprite.Y = p.Pos.Y
		p.Life += 0.0002
	} else {
		p.Alive = false
	}
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	p.Sprite.Draw(screen)
}
