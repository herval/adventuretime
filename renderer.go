package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// TODO load spritemap
// TODO get sprite
// TODO position sprite
// TODO render stats

func LoadSpritemap() {
	var floorTiles = []image.Rectangle{}
	for i := 0; i < 480; i++ {
		floorTiles = append(
			floorTiles,
			image.Rectangle{
				Min: image.Point{X: (i * 16), Y: 112},
				Max: image.Point{X: (i * 16) + 16, Y: 128},
			})
		floorTiles = append(
			floorTiles,
			image.Rectangle{
				Min: image.Point{X: (i * 16), Y: 129},
				Max: image.Point{X: (i * 16) + 16, Y: 145},
			})
	}
	println(floorTiles)

	// TODO handle errors
	// TODO build a more flexible tilemap format
	path, _ := filepath.Abs("./sprites.png")

	sheet, err := os.Open(path)
	if err != nil {
		println(err.Error())
	}
	defer sheet.Close()

	spritesheet, err := png.Decode(sheet)
	if err != nil {
		println(err.Error())
	}

	// copy spritesheet to memory so we can subimage pieces of it
	sprites := image.NewRGBA(image.Rect(0, 0, spritesheet.Bounds().Size().X, spritesheet.Bounds().Size().Y))
	draw.Draw(sprites, sprites.Bounds(), spritesheet, image.Point{0, 0}, draw.Src)

	m := image.NewRGBA(image.Rect(0, 0, 600, 600))
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			draw.Draw(m, m.Bounds(), sprites.SubImage(floorTiles[i+1]), image.Point{10 + (i * 16), 10 + (j * 16)}, draw.Src)
		}
	}

	toimg, _ := os.Create("new.jpg")
	defer toimg.Close()

	jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})

}
