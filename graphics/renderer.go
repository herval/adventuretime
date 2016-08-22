package graphics

import (
	"image"
	"image/png"
	"math/rand"
	"os"
)

// TODO position sprite
// TODO render stats

// maps a scene to a spritemap and renders it as an image
type Renderer struct {
	Sprites    Spritemap
	CanvasSize image.Rectangle
}

func (r *Renderer) DrawScene(scene *Scene) *image.RGBA {
	m := image.NewRGBA(r.CanvasSize)

	// intentionally ineficient :)

	// convert scene to an array of floor/background sprites
	r.renderFloorAndWalls(scene, m)

	// convert scene to an array of props & characters

	// r.drawBackgrounds(m, stuff)
	// r.drawFloors(m, stuff)
	// r.drawStuff(m, stuff)

	return m
}

func (r *Renderer) renderFloorAndWalls(scene *Scene, m *image.RGBA) {
	floorMap := scene.Tiles
	for row := 0; row < len(floorMap); row++ {
		for col := 0; col < len(floorMap[row]); col++ {
			if scene.IsSprite(row, col) {
				floorMap[row][col] = RoomFloor
			} else if scene.IsTile(row, col, WallWithDecoration) {
				floorMap[row][col] = Wall
			} else if scene.IsTile(row, col, Door) {
				floorMap[row][col] = RoomFloor
			}
		}
	}

	for row := 0; row < len(floorMap); row++ {
		for col := 0; col < len(floorMap[row]); col++ {
			var sprite string

			if scene.IsTile(row, col, Nothing) {
				sprite = TheUnknown
			} else {
				surrounds := scene.Surroundings(row, col)

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

	}
}

// is a given floor/wall/ceiling/etc tile surrounded by more of the same kind?
// const (
// 	center = iota
// 	topLeft
// 	bottomLeft
// 	topOnly
// 	topRight
// 	bottomRight
// 	centerAllSides
// )

// func cornerType(stuff [][]string, row int, col int, kind string) int {
// 	var top, left, bottom, right bool

// 	if row > 0 && stuff[row-1][col] == kind {
// 		top = true
// 	}
// 	if row < len(stuff)-1 && stuff[row+1][col] == kind {
// 		bottom = true
// 	}
// 	if col > 0 && stuff[row][col-1] == kind {
// 		left = true
// 	}
// 	if col < len(stuff[row])-1 && stuff[row][col+1] == kind {
// 		right = true
// 	}

// 	if top && left && !right && !bottom {
// 		return topLeft
// 	}
// 	if top && !left && !right && !bottom {
// 		return topOnly
// 	}
// 	if !top && left && !right && !bottom {
// 		return leftOnly
// 	}
// 	if !top && !left && right && !bottom {
// 		return topOnly
// 	}
// 	if top && !left && !right && !bottom {
// 		return topOnly
// 	}

// 	return center
// }

func (r *Renderer) drawBackgrounds(m *image.RGBA, stuff [][]string) {
	for row := 0; row < len(stuff); row++ {
		for col := 0; col < len(stuff[row]); col++ {
			if stuff[row][col] == Nothing {
				// cornerType := cornerType(stuff, row, col, Nothing)
				// switch cornerType {
				// case center:
				// r.Sprites.BlipInto(m, col*SquareSize, row*SquareSize, random(Ceilings))
				// }
			}
		}
	}
}

func (r *Renderer) drawFloors(m *image.RGBA, stuff [][]string) {
	for row := 0; row < len(stuff); row++ {
		for col := 0; col < len(stuff[row]); col++ {
			if stuff[row][col] == RoomFloor {
				r.Sprites.BlipInto(m, col*SquareSize, row*SquareSize, random(Floors))
			}
		}
	}
}

func (r *Renderer) drawStuff(m *image.RGBA, stuff [][]string) {

}

func random(source []string) string {
	// if len(source) > 1 {
	return source[rand.Intn(len(source))]
	// } else {
	// return source[0]
	// }
}

func SaveImage(img *image.RGBA, destination string) {
	dest, _ := os.Create(destination)
	defer dest.Close()

	png.Encode(dest, img)
}

func NewRenderer(spritesPath string, width int, height int) Renderer {
	//var floorTiles = []image.Rectangle{}
	//for i := 0; i < 480; i++ {
	//	floorTiles = append(
	//		floorTiles,
	//		image.Rectangle{
	//			Min: image.Point{X: (i * 16), Y: 112},
	//			Max: image.Point{X: (i * 16) + 16, Y: 128},
	//		})
	//	floorTiles = append(
	//		floorTiles,
	//		image.Rectangle{
	//			Min: image.Point{X: (i * 16), Y: 129},
	//			Max: image.Point{X: (i * 16) + 16, Y: 145},
	//		})
	//}
	//println(floorTiles)
	//
	//// TODO handle errors
	//// TODO build a more flexible tilemap format
	//path, _ := filepath.Abs("./sprites.png")
	//
	//sheet, err := os.Open(path)
	//if err != nil {
	//	println(err.Error())
	//}
	//defer sheet.Close()
	//
	//spritesheet, err := png.Decode(sheet)
	//if err != nil {
	//	println(err.Error())
	//}
	//
	//// copy spritesheet to memory so we can subimage pieces of it
	//sprites := image.NewRGBA(image.Rect(0, 0, spritesheet.Bounds().Size().X, spritesheet.Bounds().Size().Y))
	//draw.Draw(sprites, sprites.Bounds(), spritesheet, image.Point{0, 0}, draw.Src)
	//
	//m := image.NewRGBA(image.Rect(0, 0, 600, 600))
	//for i := 0; i < 5; i++ {
	//	for j := 0; j < 10; j++ {
	//		draw.Draw(m, m.Bounds(), sprites.SubImage(floorTiles[i+1]), image.Point{10 + (i * 16), 10 + (j * 16)}, draw.Src)
	//	}
	//}
	//
	//toimg, _ := os.Create("new.jpg")
	//defer toimg.Close()
	//
	//jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})

	return Renderer{
		Sprites:    LoadSpritemap(spritesPath),
		CanvasSize: image.Rect(0, 0, width, height),
	}
}
