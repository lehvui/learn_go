package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	en "github.com/lehvui/learn_go/entities"
)

type Game struct {
	Player        *en.Player
	RemotePlayers []*en.RPlayer
}

func (g *Game) Update() error {

	g.Player.ListenEvents()

	for _, r := range g.RemotePlayers {
		r.ListenEvents()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	opt := &ebiten.DrawImageOptions{}

	opt.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(g.Player.Img, opt)

	for _, r := range g.RemotePlayers {
		ropt := &ebiten.DrawImageOptions{}
		ropt.GeoM.Translate(r.X, r.Y)
		screen.DrawImage(r.Img, ropt)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	firstCharacter := image.Rect(96, 0, 64, 32)

	spriteSheet, _, err := ebitenutil.NewImageFromFile("rogues.png")

	playerImg := spriteSheet.SubImage(firstCharacter)

	remoteImg := spriteSheet.SubImage(image.Rect(0, 0, 32, 32))

	if err != nil {
		log.Fatal(err)
	}

	p1 := &en.Player{
		Img: ebiten.NewImageFromImage(playerImg),
		X:   100,
		Y:   100,
	}

	game := &Game{
		Player: p1,
		RemotePlayers: []*en.RPlayer{
			&en.RPlayer{
				Img: ebiten.NewImageFromImage(remoteImg),
				X:   150,
				Y:   150,
			},
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
