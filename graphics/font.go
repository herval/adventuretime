package graphics

import (
	"image"
	"image/draw"
	"github.com/herval/adventuretime/util"
)

type Font struct {
	Frames      map[string]Frame
	Spritesheet *image.RGBA
}

func (s *Font) Write(dst *image.RGBA, x int, y int, text string) {
	// TODO
	//sheetPosition := s.frameFor(sprite)
	//if sheetPosition != nil {
	//	// offsets "project" the sprites up (for longer or wider sprites)
	//	offsetY := sheetPosition.H - SquareSize
	//	offsetX := sheetPosition.W - SquareSize
	//
	//	point := image.Point{
	//		X: x,
	//		Y: y,
	//	}
	//
	//	if offsetX != 0 || offsetY != 0 {
	//		util.DebugFmt("BLIPPING: %+v - %+v - %+v\n", sprite, point, sheetPosition)
	//	}
	//
	//	basePosX := x - offsetX
	//	basePosY := y - offsetY
	//
	//	// blip a piece of the sprite sheet into a position on the dst image
	//	pointOnSpritesheet := image.Point{sheetPosition.X, sheetPosition.Y}
	//	position := image.Rect(basePosX, basePosY, basePosX+sheetPosition.W, basePosY+sheetPosition.H)
	//	draw.Draw(dst, position, s.Spritesheet, pointOnSpritesheet, draw.Over)
	//}
}

func (s *Font) frameFor(sprite *Sprite) *Dimensions {
	if sprite == nil {
		return nil
	}

	frame, found := s.Frames[sprite.id]
	if found {
		return &frame.Dimensions
	}
	util.DebugFmt("Sprite not found: %s", sprite)
	return nil
}

// load the default spritemap file
func LoadFont(spritesPath string) Font {
	data, sprites := loadSheet(spritesPath, "boxy_font")

	return Font{
		data,
		sprites,
	}
}