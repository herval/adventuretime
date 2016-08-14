package graphics

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"

	"github.com/herval/adventuretime/util"
)

const (
	SquareSize = 16

	CeilingBottomLeft1 = "1.png"

	CeilingTop1 = "2.png"

	CeilingAllSides = "24.png"

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

	BannerRed1 = "290.png"

	HeroUnarmed2 = "302.png"
	HeroArmed2   = "303.png"

	GorgonArmed = "326.png"

	TableHorizontal = "75.png"
)

const (
	TheUnknown = "23.png"
)

var Walls = []string{
	"60.png",
	"61.png",
	"62.png",
	"63.png",
}

var Ceilings = []string{
	"8.png",
	"9.png",
	"10.png",
	"11.png",
	"12.png",
	"13.png",
	"14.png",
	"15.png",
	"16.png",
	"17.png",
	"18.png",
	"19.png",
	"20.png",
	"21.png",
	"22.png",
	"23.png",
}

var CeilingLefts = []string{
	"33.png",
}

var CeilingTops = []string{
	"46.png",
}

var CeilingRights = []string{
	"4.png",
	"6.png",
}

var CeilingBottoms = []string{
	"3.png",
	"22.png",
	"25.png",
}

var Floors = []string{
	"271.png",
	"271.png",
	"271.png",
	"271.png",
	"272.png",
	"272.png",
	"272.png",
	"272.png",
	"273.png",
	"274.png",
	"275.png",
	"276.png",
	"277.png",
	"278.png",
}

var CeilingBottomLefts = []string{CeilingBottomLeft1}

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

	util.Debug(fmt.Sprintf("%+v - %+v\n", sprite.Dimensions, point))

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
func LoadSpritemap(spritesPath string) Spritemap {
	// load the map
	loader := FramesLoader{}
	data := loader.Parse(spritesPath + "/sprites.json")

	// load the image file
	path, _ := filepath.Abs(spritesPath + "/sprites.png")
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
