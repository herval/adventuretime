package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"
)

const SquareSize = 16

func blip(m *image.RGBA, x int, y int, spriteName string, spritemap *Spritemap) {
	sprite := spritemap.Sprite(spriteName)

	point := image.Point{
		X: (x * SquareSize),
		Y: (y * SquareSize),
	}

	fmt.Printf("%+v - %+v\n", sprite.Dimensions, point)

	// rect := image.Rect(point.X, point.Y, point.X + sprite.Dimensions.W, point.Y + sprite.Dimensions)

	// zero := image.Point{0, 0}
	draw.Draw(m, m.Bounds(), sprite.Image, point, draw.Over)
}

func TestRenderer(t *testing.T) {

	println("Rendering...")

	spritemap := LoadSpritemap()

	// var floorTiles = []image.Rectangle{}
	// for i := 0; i < 20; i++ {
	// 	floorTiles = append(
	// 		floorTiles,
	// 		image.Rectangle{
	// 			Min: image.Point{X: (i * 16), Y: 112},
	// 			Max: image.Point{X: (i * 16) + 16, Y: 128},
	// 		})
	// 	floorTiles = append(
	// 		floorTiles,
	// 		image.Rectangle{
	// 			Min: image.Point{X: (i * 16), Y: 129},
	// 			Max: image.Point{X: (i * 16) + 16, Y: 145},
	// 		})
	// }
	// fmt.Println("%v", floorTiles)

	m := image.NewRGBA(image.Rect(0, 0, 400, 400))
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			blip(m, i, j, "220.png", &spritemap)
		}
	}

	blip(m, 3, 5, "290.png", &spritemap)
	blip(m, 3, 5, "303.png", &spritemap)
	blip(m, 4, 5, "308.png", &spritemap)

	toimg, _ := os.Create("new.png")
	defer toimg.Close()

	png.Encode(toimg, m)
}
