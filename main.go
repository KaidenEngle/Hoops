package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/KaidenEngle/8-Bit_Hoops/spritesheet"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprites struct {
	PlayerImage      *ebiten.Image
	PlayerX, PlayerY float64
	PlayerL, PlayerH float64
	isFollowing      bool
}

type Meter struct {
	MeterImage                     *ebiten.Image
	MeterX, MeterY, MeterL, MeterH float64
	Active                         bool
}

type Court struct {
	CourtImage     *ebiten.Image
	CourtX, CourtY float64
}

type Game struct {
	jarred            *Sprites
	ball              *Sprites
	sprites           []*Sprites
	Court             *Court
	Meter             *Meter
	playerSpriteSheet *spritesheet.SpriteSheet
}

func init() {

}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.jarred.PlayerX += 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.jarred.PlayerX -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.jarred.PlayerY -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.jarred.PlayerY += 3
	}

	ball := g.sprites[1]

	//Ball follows player on collision
	if ball.PlayerX < g.jarred.PlayerX+g.jarred.PlayerL &&
		ball.PlayerX+ball.PlayerL > g.jarred.PlayerX &&
		ball.PlayerY < g.jarred.PlayerY+g.jarred.PlayerH &&
		ball.PlayerY+ball.PlayerH > g.jarred.PlayerY {
		ball.isFollowing = true
	}

	if ball.isFollowing {
		ball.PlayerX = g.jarred.PlayerX + 20
		ball.PlayerY = g.jarred.PlayerY + 50
	}

	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.Meter.Active = true
	}

	if !g.Meter.Active {
		g.Meter.MeterL = 0
		g.Meter.MeterH = 0
	}

	if g.Meter.Active && ball.isFollowing {
		g.Meter.MeterL = 5
		g.Meter.MeterH = 500
		g.Meter.MeterX = g.jarred.PlayerX + 40
		g.Meter.MeterY = g.jarred.PlayerY
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	opts := ebiten.DrawImageOptions{}
	courtOpts := ebiten.DrawImageOptions{}
	meterOpts := ebiten.DrawImageOptions{}
	courtOpts.GeoM.Translate(g.Court.CourtX, g.Court.CourtY)
	opts.GeoM.Translate(g.jarred.PlayerX, g.jarred.PlayerY)
	meterOpts.GeoM.Translate(g.Meter.MeterX, g.Meter.MeterY)

	screen.DrawImage(g.Court.CourtImage.SubImage(image.Rect(0, 0, 1000, 1000)).(*ebiten.Image),
		&courtOpts,
	)

	opts.GeoM.Reset()

	screen.DrawImage(g.Meter.MeterImage.SubImage(image.Rect(0, 0, int(g.Meter.MeterL), int(g.Meter.MeterH))).(*ebiten.Image),
		&meterOpts)

	for _, spriteV := range g.sprites {
		opts.GeoM.Translate(spriteV.PlayerX, spriteV.PlayerY)
		screen.DrawImage(spriteV.PlayerImage.SubImage(image.Rect(0, 0, int(spriteV.PlayerL), int(spriteV.PlayerH))).(*ebiten.Image),
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
	ebiten.SetWindowTitle("Never had a doubt inside me")

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

	MeterImg, _, err := ebitenutil.NewImageFromFile("images/greenmeter.png")
	if err != nil {
		log.Fatal(err)
	}

	playerSpriteSheet := spritesheet.NewSpriteSheet(4, 7, 16)

	game := Game{
		sprites: []*Sprites{
			{
				//Player
				PlayerImage: PlayerImg,
				PlayerX:     75,
				PlayerY:     60,
				PlayerL:     35,
				PlayerH:     300,
				isFollowing: false,
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

		playerSpriteSheet: playerSpriteSheet,

		Court: &Court{
			CourtImage: CourtImg,
			CourtX:     0,
			CourtY:     0,
		},
		Meter: &Meter{
			MeterImage: MeterImg,
			MeterX:     0,
			MeterY:     0,
			MeterL:     5,
			MeterH:     500,
			Active:     false,
		},
	}

	game.jarred = game.sprites[0]
	game.ball = game.sprites[1]

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
