package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
)

const (
	SquareSize = 16

	CeilingBottomLeft1 = "1.png"

	CeilingTop1 = "2.png"

	CeilingRight1 = "4.png"
	CeilingRight2 = "6.png"

	CeilingBottom1 = "3.png"
	CeilingBottom2 = "5.png"
	CeilingBottom3 = "7.png"
	CeilingBottom4 = "25.png"

	Ceiling1  = "8.png"
	Ceiling2  = "9.png"
	Ceiling3  = "10.png"
	Ceiling4  = "11.png"
	Ceiling5  = "12.png"
	Ceiling6  = "13.png"
	Ceiling7  = "14.png"
	Ceiling8  = "15.png"
	Ceiling9  = "16.png"
	Ceiling10 = "17.png"
	Ceiling11 = "18.png"
	Ceiling12 = "19.png"
	Ceiling13 = "20.png"
	Ceiling14 = "21.png"
	Ceiling15 = "22.png"
	Ceiling16 = "23.png"

	CeilingAllSides = "24.png"

	Wall1 = "60.png"
	Wall2 = "61.png"
	Wall3 = "63.png"
	Wall4 = "64.png"
	Wall5 = "65.png"
	Wall6 = "66.png"

	Stair1 = "67.png"

	LargeStairs1 = "68.png"

	Door1   = "69.png"
	Door2   = "70.png"
	Door3   = "71.png"
	Door4   = "72.png"
	Passage = "73.png"

	FloorTopLeft1 = "211.png"

	FloorTopRight1 = "212.png"

	FloorBottomLeft1 = "213.png"

	FloorBottomRight1 = "214.png"

	FloorLeft1 = "215.png"
	FloorLeft2 = "216.png"
	FloorLeft3 = "217.png"

	FloorTop1 = "218.png"
	FloorTop2 = "219.png"
	FloorTop3 = "220.png"

	FloorBottom1 = "221.png"
	FloorBottom2 = "222.png"
	FloorBottom3 = "223.png"

	FloorRight1  = "224.png"
	FloorRight2  = "225.png"
	FloorRight13 = "226.png"

	FloorCenter1 = "259.png"

	BannerRed1 = "290.png"

	HeroUnarmed2 = "302.png"
	HeroArmed2   = "303.png"

	GorgonArmed = "326.png"

	TableHorizontal = "75.png"
)

type Spritemap struct {
	SmallMonsters []*Frame
	Frames        map[string]Frame
	Spritesheet   *image.RGBA
}

type Sprite struct {
	Dimensions Dimensions
}

func (s *Spritemap) BlipInto(dst *image.RGBA, x int, y int, spriteName string) {
	sprite := s.Sprite(spriteName)

	point := image.Point{
		X: x,
		Y: y,
	}

	Debug(fmt.Sprintf("%+v - %+v\n", sprite.Dimensions, point))

	pointOnSpritesheet := image.Point{sprite.Dimensions.X, sprite.Dimensions.Y}

	position := image.Rect(x, y, x+sprite.Dimensions.W, y+sprite.Dimensions.H)

	draw.Draw(dst, position, s.Spritesheet, pointOnSpritesheet, draw.Over)
}

func (s *Spritemap) Sprite(name string) *Sprite {
	// TODO not found?
	frame, found := s.Frames[name]
	if found {
		return &Sprite{
			Dimensions: frame.Dimensions,
		}
	}
	return nil
}

// load the default spritemap file
func LoadSpritemap() Spritemap {
	// load the map
	loader := FramesLoader{}
	data := loader.Parse("./sprites.json")

	// load the image file
	path, _ := filepath.Abs("./sprites.png")
	sheet, _ := os.Open(path)
	defer sheet.Close()
	spritesheet, _ := png.Decode(sheet)
	// copy spritesheet to memory so we can subimage pieces of it
	sprites := image.NewRGBA(image.Rect(0, 0, spritesheet.Bounds().Size().X, spritesheet.Bounds().Size().Y))
	draw.Draw(sprites, sprites.Bounds(), spritesheet, image.Point{0, 0}, draw.Src)

	// TODO handle errors

	return NewSpritemap(data, sprites)
}

func NewSpritemap(frames map[string]Frame, spritesheet *image.RGBA) Spritemap {

	return Spritemap{
		SmallMonsters: []*Frame{},
		Frames:        frames,
		Spritesheet:   spritesheet,
	}
}
