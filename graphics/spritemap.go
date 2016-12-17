package graphics

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"

	"github.com/herval/adventuretime/util"
)

var SquareSize = 16

var TheUnknown = Sprite{"23.png", nil}

var SmallShadow = Sprite{"804.png", nil}

var LargeShadow = Sprite{"805.png", nil}

var Stair1 = Sprite{"67.png", nil}

var LargeStairs1 = Sprite{"68.png", nil}

var Passage = Sprite{"73.png", nil}

var BannerRed1 = Sprite{"290.png", nil}

var HeroUnarmed2 = Sprite{"302.png", &SmallShadow}
var HeroArmed2 = Sprite{"303.png", &SmallShadow}

var BigMonsters = []Sprite{
	Sprite{"326.png", &SmallShadow},
	Sprite{"325.png", &SmallShadow},
}

var SmallMonsters = []Sprite{
	Sprite{"327.png", &SmallShadow},
	Sprite{"328.png", &SmallShadow},
	Sprite{"329.png", &SmallShadow},
	Sprite{"330.png", &SmallShadow},
	Sprite{"331.png", &SmallShadow},
	Sprite{"332.png", &SmallShadow},
	Sprite{"333.png", &SmallShadow},
	Sprite{"334.png", &SmallShadow},
	Sprite{"335.png", &SmallShadow},
}

var TableHorizontal = Sprite{"75.png", &LargeShadow}
var TableVertical = Sprite{"74.png", &LargeShadow}

var Fire = Sprite{"802.png", nil}

var DoorsFrontFacing = []Sprite{
	Sprite{"69.png", nil},
	Sprite{"70.png", nil},
	Sprite{"71.png", nil},
	Sprite{"72.png", nil},
}

var Banners = []Sprite{
	Sprite{"286.png", nil},
	Sprite{"287.png", nil},
	Sprite{"288.png", nil},
	Sprite{"289.png", nil},
	Sprite{"290.png", nil},
	Sprite{"291.png", nil},
}

var Walls = []Sprite{
	Sprite{"60.png", nil},
	Sprite{"61.png", nil},
	Sprite{"62.png", nil},
	Sprite{"63.png", nil},
}

var FloorTopLefts = []Sprite{
	Sprite{"211.png", nil},
	Sprite{"270.png", nil},
}

var FloorTopRights = []Sprite{
	Sprite{"212.png", nil},
	Sprite{"269.png", nil},
}

var FloorBottomRights = []Sprite{
	Sprite{"214.png", nil},
	Sprite{"267.png", nil},
}

var FloorLeftRights = []Sprite{
	Sprite{"233.png", nil},
	Sprite{"234.png", nil},
}

var FloorTopBottoms = []Sprite{
	Sprite{"231.png", nil},
	Sprite{"232.png", nil},
}

var FloorBottomLefts = []Sprite{
	Sprite{"268.png", nil},
	Sprite{"213.png", nil},
}

var FloorLeftRightTops = []Sprite{
	Sprite{"228.png", nil},
}

var FloorLeftRightBottoms = []Sprite{
	Sprite{"266.png", nil},
}

var FloorTopBottomRights = []Sprite{
	Sprite{"229.png", nil},
}

var FloorTopBottomLefts = []Sprite{
	Sprite{"227.png", nil},
}

var FloorBottoms = []Sprite{
	Sprite{"221.png", nil},
	Sprite{"222.png", nil},
	Sprite{"223.png", nil},
}

var FloorTops = []Sprite{
	Sprite{"220.png", nil},
	Sprite{"219.png", nil},
	Sprite{"218.png", nil},
}

var FloorLefts = []Sprite{
	Sprite{"215.png", nil},
	Sprite{"216.png", nil},
	Sprite{"217.png", nil},
}

var FloorRights = []Sprite{
	Sprite{"225.png", nil},
	Sprite{"226.png", nil},
	Sprite{"224.png", nil},
}

var Floors = []Sprite{
	Sprite{"271.png", nil},
	Sprite{"271.png", nil},
	Sprite{"271.png", nil},
	Sprite{"271.png", nil},
	Sprite{"272.png", nil},
	Sprite{"272.png", nil},
	Sprite{"272.png", nil},
	Sprite{"272.png", nil},
	Sprite{"273.png", nil},
	Sprite{"274.png", nil},
	Sprite{"275.png", nil},
	Sprite{"276.png", nil},
	Sprite{"277.png", nil},
	Sprite{"278.png", nil},
}

type Sprite struct {
	id     string
	shadow *Sprite
}

type Spritemap struct {
	Frames      map[string]Frame
	Spritesheet *image.RGBA
}

func (s *Spritemap) BlipInto(dst *image.RGBA, x int, y int, sprite *Sprite) {
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
			util.DebugFmt("BLIPPING: %+v - %+v - %+v\n", sprite, point, sheetPosition)
		}

		basePosX := x - offsetX
		basePosY := y - offsetY

		// draw shadow, if sprite has one
		shadow := s.shadowFor(sprite)
		if shadow != nil {
			// TODO darken shadow (tint)
			// mask := image.NewUniform(color.RGBA{100, 200, 100, 200})
			shadowOffset := (shadow.H / 4) + SquareSize/5
			pointOnSpritesheet := image.Point{shadow.X, shadow.Y}
			position := image.Rect(x, y+shadowOffset, x+shadow.W, y+shadow.H+shadowOffset)
			// draw.DrawMask(dst, position, s.Spritesheet, pointOnSpritesheet, mask, image.Point{}, draw.Over)
			draw.Draw(dst, position, s.Spritesheet, pointOnSpritesheet, draw.Over)
		}

		// blip a piece of the sprite sheet into a position on the dst image
		pointOnSpritesheet := image.Point{sheetPosition.X, sheetPosition.Y}
		position := image.Rect(basePosX, basePosY, basePosX+sheetPosition.W, basePosY+sheetPosition.H)
		draw.Draw(dst, position, s.Spritesheet, pointOnSpritesheet, draw.Over)
	}
}

func (s *Spritemap) shadowFor(sprite *Sprite) *Dimensions {
	return s.frameFor(sprite.shadow)
}

func (s *Spritemap) frameFor(sprite *Sprite) *Dimensions {
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
		Frames:      frames,
		Spritesheet: spritesheet,
	}
}
