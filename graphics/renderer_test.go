package graphics_test

import (
	"testing"
	"github.com/herval/adventuretime/util"
	"github.com/herval/adventuretime/graphics"
	"github.com/herval/adventuretime/engine"
	"fmt"
)

func TestRenderer(t *testing.T) {
	util.Debug("Rendering...")

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

func TestRandomRendering(t *testing.T) {
	dungeon := engine.NewDungeon(50, 5)

	dungeon.Blueprint.Print()
	fmt.Println(graphics.DungeonToBlipmap(dungeon))

	scene := graphics.NewScene(
		graphics.DungeonToBlipmap(
			dungeon,
		),
	)

	renderer := graphics.NewRenderer(
		"../resources",
		(len(dungeon.Blueprint.Grid[0])+10) * graphics.SquareSize,
		(len(dungeon.Blueprint.Grid)+10) * graphics.SquareSize,
	)

	img := renderer.DrawScene(&scene)

	graphics.SaveImage(img, "../random.png")
}
