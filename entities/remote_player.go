package entities

import "github.com/hajimehoshi/ebiten/v2"

type RPlayer struct {
	Img *ebiten.Image
	X   float64
	Y   float64
}

func (p *RPlayer) ListenEvents() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= 3
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.X -= 6
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.X += 4
	}
}
