package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprites struct {
	PlayerImage      *ebiten.Image
	PlayerX, PlayerY float64
	PlayerL, PlayerH float64
	isFollowing      bool
}

type Court struct {
	CourtImage     *ebiten.Image
	CourtX, CourtY float64
}

type Game struct {
	sprite  *Sprites
	sprites []*Sprites
	Court   *Court
}

func init() {

}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.sprite.PlayerX += 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.sprite.PlayerX -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.sprite.PlayerY -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.sprite.PlayerY += 3
	}

	ball := g.sprites[1]

	//Ball follows player on collision
	if ball.PlayerX < g.sprite.PlayerX+g.sprite.PlayerL &&
		ball.PlayerX+ball.PlayerL > g.sprite.PlayerX &&
		ball.PlayerY < g.sprite.PlayerY+g.sprite.PlayerH &&
		ball.PlayerY+ball.PlayerH > g.sprite.PlayerY {
		ball.isFollowing = true
	}

	if ball.isFollowing {
		ball.PlayerX = g.sprite.PlayerX + 20
		ball.PlayerY = g.sprite.PlayerY + 50
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	opts := ebiten.DrawImageOptions{}
	courtOpts := ebiten.DrawImageOptions{}
	courtOpts.GeoM.Translate(g.Court.CourtX, g.Court.CourtY)
	opts.GeoM.Translate(g.sprite.PlayerX, g.sprite.PlayerY)

	screen.DrawImage(g.Court.CourtImage.SubImage(image.Rect(0, 0, 1000, 1000)).(*ebiten.Image),
		&courtOpts,
	)

	opts.GeoM.Reset()

	for _, sprite := range g.sprites {
		opts.GeoM.Translate(sprite.PlayerX, sprite.PlayerY)
		screen.DrawImage(sprite.PlayerImage.SubImage(image.Rect(0, 0, int(sprite.PlayerL), int(sprite.PlayerH))).(*ebiten.Image),
			&opts,
		)
		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(880, 880)
	ebiten.SetWindowTitle("Linus Torvalds")

	PlayerImg, _, err := ebitenutil.NewImageFromFile("images/jarred_mccain.png")
	if err != nil {
		log.Fatal(err)
	}

	CourtImg, _, err := ebitenutil.NewImageFromFile("images/court.png")
	if err != nil {
		log.Fatal(err)
	}

	BallImg, _, err := ebitenutil.NewImageFromFile("images/basketball.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		sprites: []*Sprites{
			{
				//Player
				PlayerImage: PlayerImg,
				PlayerX:     75,
				PlayerY:     60,
				PlayerL:     35,
				PlayerH:     300,
			},
			{
				//Basketball
				PlayerImage: BallImg,
				PlayerX:     115,
				PlayerY:     80,
				PlayerL:     10,
				PlayerH:     10,
			},
		},
		Court: &Court{
			CourtImage: CourtImg,
			CourtX:     0,
			CourtY:     0,
		},
	}

	game.sprite = game.sprites[0]

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
