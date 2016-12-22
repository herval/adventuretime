package graphics_test

import (
	"testing"
	"github.com/herval/adventuretime/util"
	"github.com/herval/adventuretime/graphics"
)

func TestRenderer(t *testing.T) {
	println("Rendering...")

	scene := graphics.NewScene(`.....................|||..................
.....................|_||||...............
.....................|____|...............
....|||||^|||D||||...|_|..................
....|_U__________|...|_|..................
....|____m__M____|||||_||||^||||..........
....|_m__________D_____________|..........
....|____C_TX____|||||||_____Ü_|..........
....|m_H___C_____|.....||||_||||..........
....|__X_______T_|........|__|............
....|__________X_|........|__|............
....||||||||||D||.........||||............
..........................................
..........................................
..........................................`)

	util.DebugFmt("Rendering:\n", scene)

	renderer := graphics.NewRenderer("../resources", 700, 700)

	img := renderer.DrawScene(&scene)

	graphics.SaveImage(img, "../new.png")
}
