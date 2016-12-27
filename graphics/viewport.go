package graphics

import (
	"github.com/herval/adventuretime/engine"
	"image"
)

type Viewport struct {
	renderer Renderer
	controller *engine.Controller
}

func NewViewport(controller *engine.Controller) Viewport {
	renderer := NewRenderer(
		"resources",
		(len(controller.State.Dungeon.Blueprint.Grid[0]) + 10) * SquareSize,
		(len(controller.State.Dungeon.Blueprint.Grid) + 10) * SquareSize,
	)

	return Viewport{
		renderer,
		controller,
	}
}


func (v *Viewport) DrawDungeon() *image.RGBA {
	// TODO encapsulate all that
	scene := NewScene(
		DungeonToBlipmap(
			v.controller.State.Dungeon,
			v.controller.State.Player,
		),
	)
	return v.renderer.DrawScene(&scene)
}