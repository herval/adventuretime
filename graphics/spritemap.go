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

	Stair1 = "67.png"

	LargeStairs1 = "68.png"

	Door1   = "69.png"
	Door2   = "70.png"
	Door3   = "71.png"
	Door4   = "72.png"
	Passage = "73.png"

	BannerRed1 = "290.png"

	HeroUnarmed2 = "302.png"
	HeroArmed2   = "303.png"

	GoblinArmed = "327.png"

	GorgonArmed = "326.png"

	TableHorizontal = "75.png"

	TheUnknown = "23.png"

	SmallShadow = "804.png"
	LargeShadow = "805.png"
)

var Walls = []string{
	"60.png",
	"61.png",
	"62.png",
	"63.png",
}

var FloorTopLefts = []string{
	"211.png",
	"270.png",
}

var FloorTopRights = []string{
	"212.png",
	"269.png",
}

var FloorBottomRights = []string{
	"214.png",
	"267.png",
}

var FloorLeftRights = []string{
	"233.png",
	"234.png",
}

var FloorTopBottoms = []string{
	"231.png",
	"232.png",
}

var FloorBottomLefts = []string{
	"268.png",
	"213.png",
}

var FloorLeftRightTops = []string{
	"228.png",
}

var FloorLeftRightBottoms = []string{
	"266.png",
}

var FloorTopBottomRights = []string{
	"229.png",
}

var FloorTopBottomLefts = []string{
	"227.png",
}

var FloorBottoms = []string{
	"221.png",
	"222.png",
	"223.png",
}

var Shadow = "804.png"

var FloorTops = []string{
	"220.png",
	"219.png",
	"218.png",
}

var FloorLefts = []string{
	"215.png",
	"216.png",
	"217.png",
}

var FloorRights = []string{
	"225.png",
	"226.png",
	"224.png",
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

type Spritemap struct {
	SmallMonsters []*Frame
	Frames        map[string]Frame
	Spritesheet   *image.RGBA
}

func (s *Spritemap) BlipInto(dst *image.RGBA, x int, y int, sprite string) {
	sheetPosition := s.frameFor(sprite)
	if sheetPosition != nil {
		// offsets "project" the sprites up (for longer or wider sprites)
		offsetY := sheetPosition.H - SquareSize
		offsetX := sheetPosition.W - SquareSize

		point := image.Point{
			X: x,
			Y: y,
		}

		if offsetX != 0 || offsetY != 0 {
			util.Debug(fmt.Sprintf("BLIPPING: %+v - %+v - %+v\n", sprite, point, sheetPosition))
		}

		basePosX := x - offsetX
		basePosY := y - offsetY

		// draw shadow, if sprite has one
		shadow := s.shadowFor(sprite)

		// blip a piece of the sprite sheet into a position on the dst image
		pointOnSpritesheet := image.Point{sheetPosition.X, sheetPosition.Y}
		position := image.Rect(basePosX, basePosY, basePosX+sheetPosition.W, basePosY+sheetPosition.H)
		draw.Draw(dst, position, s.Spritesheet, pointOnSpritesheet, draw.Over)

		if shadow != nil {
			// mask := image.NewUniform(color.Alpha{128})
			shadowOffset := offsetY + (shadow.H + SquareSize)
			pointOnSpritesheet := image.Point{shadow.X, shadow.Y}
			position := image.Rect(basePosX, basePosY+shadowOffset, basePosX+shadow.W, basePosY+shadow.H)
			draw.Draw(dst, position, s.Spritesheet, pointOnSpritesheet, draw.Over)
		}
	}
}

func (s *Spritemap) shadowFor(name string) *Dimensions {
	// if name == HeroArmed2 || name == GoblinArmed {
	// 	return s.frameFor(SmallShadow)
	// }

	return nil
}

func (s *Spritemap) frameFor(name string) *Dimensions {
	frame, found := s.Frames[name]
	if found {
		return &frame.Dimensions
	}
	util.Debug(fmt.Sprintf("Sprite not found: %s", name))
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
