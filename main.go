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
}

type Court struct {
	CourtImage     *ebiten.Image
	CourtX, CourtY float64
}

type Game struct {
	Player  *Sprites
	sprites []*Sprites
	Court   *Court
}

func init() {

}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.PlayerX += 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.PlayerX -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.PlayerY -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.PlayerY += 3
	}

	for _, ball := range g.sprites {
		//Ball follows player on collision
		if g.Player.PlayerX > ball.PlayerX {
			ball.PlayerX += 2
		} else if g.Player.PlayerX < ball.PlayerX {
			ball.PlayerX -= 2
		}
		if g.Player.PlayerY > ball.PlayerY {
			ball.PlayerY += 2
		} else if g.Player.PlayerY < ball.PlayerY {
			ball.PlayerY -= 2
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	opts := ebiten.DrawImageOptions{}
	courtOpts := ebiten.DrawImageOptions{}
	courtOpts.GeoM.Translate(g.Court.CourtX, g.Court.CourtY)
	opts.GeoM.Translate(g.Player.PlayerX, g.Player.PlayerY)

	screen.DrawImage(g.Court.CourtImage.SubImage(image.Rect(0, 0, 1000, 1000)).(*ebiten.Image),
		&courtOpts,
	)

	screen.DrawImage(g.Player.PlayerImage.SubImage(image.Rect(5, 0, 35, 300)).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, sprite := range g.sprites {
		opts.GeoM.Translate(sprite.PlayerX, sprite.PlayerY)
		screen.DrawImage(sprite.PlayerImage.SubImage(image.Rect(50, 50, 65, 65)).(*ebiten.Image),
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
		Player: &Sprites{
			PlayerImage: PlayerImg,
			PlayerX:     75,
			PlayerY:     60,
		},
		sprites: []*Sprites{
			{
				//Basketball
				PlayerImage: BallImg,
				PlayerX:     115,
				PlayerY:     80,
			},
		},
		Court: &Court{
			CourtImage: CourtImg,
			CourtX:     0,
			CourtY:     0,
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
