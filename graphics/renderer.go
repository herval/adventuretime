package graphics

import (
	"image"
	"image/png"
	"math/rand"
	"os"
)

// TODO render stats

// maps a scene to a spritemap and renders it as an image
type Renderer struct {
	Sprites    Spritemap
	CanvasSize image.Rectangle
}

func NewRenderer(spritesPath string, width int, height int) Renderer {
	return Renderer{
		Sprites:    LoadSpritemap(spritesPath),
		CanvasSize: image.Rect(0, 0, width, height),
	}
}

func (r *Renderer) DrawScene(scene *Scene) *image.RGBA {
	m := image.NewRGBA(r.CanvasSize)

	// intentionally ineficient :)
	r.render(scene, m)

	return m
}

// save scene to a file
func SaveImage(img *image.RGBA, destination string) {
	dest, _ := os.Create(destination)
	defer dest.Close()

	png.Encode(dest, img)
}

func (r *Renderer) render(scene *Scene, m *image.RGBA) {
	for row := 0; row < len(scene.Tiles); row++ {
		for col := 0; col < len(scene.Tiles[row]); col++ {
			r.drawScenario(scene, m, row, col)
			r.drawCharacter(scene, m, row, col)
		}
	}
}

func (r *Renderer) drawScenario(scene *Scene, m *image.RGBA, row int, col int) {
	var sprite string

	if scene.IsTile(row, col, Nothing) {
		sprite = TheUnknown
	} else {
		surrounds := scene.SurroundingScenario(row, col)

		if scene.IsTile(row, col, RoomFloor) {
			// TODO corners on intersections
			if surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top != RoomFloor && surrounds.bottom == RoomFloor { // cap up
				sprite = random(FloorLeftRightTops)
			} else if surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top == RoomFloor && surrounds.bottom != RoomFloor { // cap up
				sprite = random(FloorLeftRightBottoms)
			} else if surrounds.right != RoomFloor && surrounds.left == RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor { // cap up
				sprite = random(FloorTopBottomRights)
			} else if surrounds.right == RoomFloor && surrounds.left != RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor { // cap up
				sprite = random(FloorTopBottomLefts)
			} else if surrounds.right == RoomFloor && surrounds.left == RoomFloor && surrounds.top != RoomFloor && surrounds.bottom != RoomFloor { // corridor sideways
				sprite = random(FloorTopBottoms)
			} else if surrounds.right != RoomFloor && surrounds.left != RoomFloor && surrounds.top == RoomFloor && surrounds.bottom == RoomFloor { // corridor up
				sprite = random(FloorLeftRights)
			} else if surrounds.top != RoomFloor && surrounds.left != RoomFloor && surrounds.right == RoomFloor && surrounds.bottom == RoomFloor { // top-left
				sprite = random(FloorTopLefts)
			} else if surrounds.top != RoomFloor && surrounds.right != RoomFloor && surrounds.left == RoomFloor && surrounds.bottom == RoomFloor { // top-right
				sprite = random(FloorTopRights)
			} else if surrounds.top != RoomFloor && surrounds.bottom == RoomFloor { // top
				sprite = random(FloorTops)
			} else if surrounds.left != RoomFloor && surrounds.right == RoomFloor { // left walls
				sprite = random(FloorLefts)
			} else if surrounds.right != RoomFloor && surrounds.left == RoomFloor { // right walls
				sprite = random(FloorRights)
			} else { // everything else
				sprite = random(Floors)
			}
		} else if scene.IsTile(row, col, Wall) { // left/right/bottom "walls" are just empty space w/ shadows
			// if surrounds.right == RoomFloor && surrounds.top == Nothing {
			if surrounds.right == RoomFloor && surrounds.left == Nothing {
				sprite = TheUnknown
			} else if surrounds.left == Nothing && surrounds.right == Wall {
				sprite = TheUnknown
			} else if surrounds.left == RoomFloor && surrounds.right == Nothing {
				sprite = TheUnknown
			} else if surrounds.right == Nothing && surrounds.left == Wall {
				sprite = TheUnknown
			} else if surrounds.bottom == Nothing && surrounds.top == RoomFloor {
				sprite = TheUnknown
			} else if surrounds.top == RoomFloor && surrounds.bottom == Wall {
				sprite = TheUnknown
			} else {
				sprite = random(Walls)
			}
		}
	}
	if sprite != "" {
		r.Sprites.BlipInto(m, col*SquareSize, row*SquareSize, sprite)
	}
}

func (r *Renderer) drawCharacter(scene *Scene, m *image.RGBA, row int, col int) {
	sprite := ""

	// draw stuff on top of the floor
	if scene.IsTile(row, col, Hero) {
		sprite = HeroArmed2
	} else if scene.IsTile(row, col, WallWithDecoration) {
		sprite = BannerRed1
	} else if scene.IsTile(row, col, SmallMonster) {
		sprite = GoblinArmed
	}

	if sprite != "" {
		r.Sprites.BlipInto(m, col*SquareSize, row*SquareSize, sprite)
	}
}

func random(source []string) string {
	// if len(source) > 1 {
	return source[rand.Intn(len(source))]
	// } else {
	// return source[0]
	// }
}
