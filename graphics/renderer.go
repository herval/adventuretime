package graphics

import (
	"image"
	"image/png"
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
	// TODO draw vertically to fix occlusion?

	for row := 0; row < len(scene.FloorMap); row++ {
		for col := 0; col < len(scene.FloorMap[row]); col++ {
			r.drawSprite(m, row, col, scene.FloorMap[row][col])
		}
	}
	for row := 0; row < len(scene.WallsMap); row++ {
		for col := 0; col < len(scene.WallsMap[row]); col++ {
			r.drawSprite(m, row, col, scene.WallsMap[row][col])
		}
	}
	for row := 0; row < len(scene.SpriteMap); row++ {
		for col := 0; col < len(scene.SpriteMap[row]); col++ {
			r.drawSprite(m, row, col, scene.SpriteMap[row][col])
		}
	}
	// TODO redraw bottom walls to occlude sprites
}

func (r *Renderer) drawSprite(m *image.RGBA, row int, col int, sprite *Sprite) {
	if sprite != nil {
		r.Sprites.BlipInto(m, col*SquareSize, row*SquareSize, sprite)
	}
}
