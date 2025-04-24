package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *Player
}

type Player struct {
	Img *ebiten.Image
	X   float64
	Y   float64
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.Player.Y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Player.X -= 1
	}

	fmt.Println(g.Player.Y)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	opt := &ebiten.DrawImageOptions{}

	opt.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(g.Player.Img, opt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	i, _ := os.Open("gopher.png")
	playerImg, _ := png.Decode(i)

	p1 := &Player{
		Img: ebiten.NewImageFromImage(playerImg),
		X:   100,
		Y:   100,
	}

	if err := ebiten.RunGame(&Game{
		Player: p1,
	}); err != nil {
		log.Fatal(err)
	}
}
