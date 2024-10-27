package spritesheet

import "image"

type SpriteSheet struct {
	width, height, tileSize int
}

func (s *SpriteSheet) Rect(index int) image.Rectangle {
	x := (index % s.width) * s.tileSize
	y := (index / s.width) * s.tileSize

	return image.Rect(
		x, y, x+s.tileSize, y+s.tileSize,
	)
}

func NewSpriteSheet(w, h, t int) *SpriteSheet {
	return &SpriteSheet{w, h, t}
}
