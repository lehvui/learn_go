package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lehvui/learn_go/entities"
	en "github.com/lehvui/learn_go/entities"
)

type Game struct {
	Player        *en.Player
	RemotePlayers []*en.RPlayer
	Sprite        *en.Sprite
	Projectiles   []*en.Projectile
}

func (g *Game) Update() error {
	g.Player.ListenEvents()
	for _, r := range g.RemotePlayers {
		r.ListenEvents()
	}

	g.Sprite.X = g.Player.X
	g.Sprite.Y = g.Player.Y

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		curso := en.Position{X: float64(x), Y: float64(y)}
		playerP := en.Position{X: g.Player.X, Y: g.Player.Y}

		proj := entities.NewProjectile("items.png", playerP, playerP, curso)
		g.Projectiles = append(g.Projectiles, proj)
	}

	var aliveProjectiles []*en.Projectile
	for _, proj := range g.Projectiles {
		proj.Update()
		if proj.Alive {
			aliveProjectiles = append(aliveProjectiles, proj)
		}
	}
	g.Projectiles = aliveProjectiles

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Sprite != nil {
		g.Sprite.Draw(screen)
	}

	if g.Player != nil {
		pSprite := &entities.Sprite{
			Image: g.Player.Img,
			X:     g.Player.X,
			Y:     g.Player.Y,
		}
		pSprite.Draw(screen)
	}

	for _, r := range g.RemotePlayers {
		rSprite := &entities.Sprite{
			Image: r.Img,
			X:     r.X,
			Y:     r.Y,
		}
		rSprite.Draw(screen)
	}

	for _, proj := range g.Projectiles {
		proj.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	playerRect := image.Rect(96, 0, 96+64, 0+32)
	remoteRect := image.Rect(0, 0, 32, 32)

	mainSprite := en.NewSpriteFromSheet("rogues.png", playerRect, 100, 50)
	remoteSprite := en.NewSpriteFromSheet("rogues.png", remoteRect, 200, 100)

	p := &entities.Player{
		Img: mainSprite.Image,
		X:   mainSprite.X,
		Y:   mainSprite.Y,
	}

	r := &entities.RPlayer{
		Img: remoteSprite.Image,
		X:   remoteSprite.X,
		Y:   remoteSprite.Y,
	}

	game := &Game{
		Player:        p,
		RemotePlayers: []*en.RPlayer{r},
		Sprite:        mainSprite,
		Projectiles:   []*en.Projectile{},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
