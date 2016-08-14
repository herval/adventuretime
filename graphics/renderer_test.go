package graphics

import (
	"fmt"
	"testing"
)

func TestRenderer(t *testing.T) {
	println("Rendering...")

	scene := NewScene(`                    
.........................................
.........................................
.........................................
....|||||||||||||........................
....|___________|........................
....|____m__M___||||||||||^||||..........
....|___________D_____________|..........
....|____C_TX___|||||||_____Ü_|..........
....|______C____|.....|||||||||..........
....|__H________|........................
....|||||||||||||........................
.........................................
.........................................
.........................................`)

	fmt.Println("Rendering:\n", scene.Blipmap)

	renderer := NewRenderer("../resources", 700, 700)

	img := renderer.DrawScene(scene)

	SaveImage(img, "../new.png")

	// m := image.NewRGBA(image.Rect(0, 0, 400, 400))
	// for i := 5; i < 20; i++ {
	// 	spritemap.BlipInto(m, i*SquareSize, 4*SquareSize, FloorTop1)
	// 	for j := 5; j < 20; j++ {
	// 		spritemap.BlipInto(m, i*SquareSize, j*SquareSize, FloorCenter1)
	// 	}
	// 	spritemap.BlipInto(m, i*SquareSize, 20*SquareSize, FloorBottom1)
	// }

	// for j := 5; j < 20; j++ {
	// 	spritemap.BlipInto(m, 4*SquareSize, j*SquareSize, FloorLeft1)
	// 	spritemap.BlipInto(m, 20*SquareSize, j*SquareSize, FloorRight1)
	// }

	// spritemap.BlipInto(m, 20*SquareSize, 4*SquareSize, FloorTopRight1)
	// spritemap.BlipInto(m, 4*SquareSize, 4*SquareSize, FloorTopLeft1)
	// spritemap.BlipInto(m, 4*SquareSize, 20*SquareSize, FloorBottomLeft1)
	// spritemap.BlipInto(m, 20*SquareSize, 20*SquareSize, FloorBottomRight1)

	// spritemap.BlipInto(m, 10*SquareSize, 18*SquareSize, HeroArmed2)
	// spritemap.BlipInto(m, 15*SquareSize, 5*SquareSize, GorgonArmed)
	// spritemap.BlipInto(m, 14*SquareSize, 9*SquareSize, TableHorizontal)

	// toimg, _ := os.Create("new.png")
	// defer toimg.Close()

	// png.Encode(toimg, m)
}
