package entities

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Image *ebiten.Image
	X, Y  float64
}

func NewSprite(path string, x, y float64) *Sprite {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("failed to load sprite %s: %v", path, err)
	}
	return &Sprite{Image: img, X: x, Y: y}
}

func NewSpriteFromSheet(path string, rect image.Rectangle, x, y float64) *Sprite {
	sheet, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("failed to load sprite sheet %s: %v", path, err)
	}
	subImg := sheet.SubImage(rect).(*ebiten.Image)
	return &Sprite{Image: subImg, X: x, Y: y}
}

func NewNumberSprite(path string, x, y float64, a, d int) *Sprite {
	startX := a * 32
	startY := d * 32
	endX := startX + 32
	endY := startY + 32

	rect := image.Rect(startX, startY, endX, endY)

	return NewSpriteFromSheet(path, rect, x, y)
}
func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.Image, op)
}
