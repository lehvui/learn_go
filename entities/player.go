package entities

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	Img *ebiten.Image
	X   float64
	Y   float64
}

func (p *Player) ListenEvents() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X += 1
	}
}
