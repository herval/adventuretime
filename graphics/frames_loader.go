package graphics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"image/png"
	"image"
	"image/draw"
)

type FramesLoader struct {
}

func (l *FramesLoader) Parse(fileName string) map[string]Frame {
	path, e := filepath.Abs(fileName)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))

	var data jsonobject
	json.Unmarshal(file, &data)

	// util.Debug(fmt.Sprintf("Results: %+v\n", data.Frames))

	return data.Frames
}

type Dimensions struct {
	X int
	Y int
	W int
	H int
}

type Frame struct {
	Dimensions Dimensions `json:"frame"`
}

type jsonobject struct {
	Frames map[string]Frame
}

func loadSheet(spritesPath string, filename string) (map[string]Frame, *image.RGBA){
	// load the map
	loader := FramesLoader{}
	data := loader.Parse(fmt.Sprintf("%s/%s.json", spritesPath, filename))

	// load the image file
	path, _ := filepath.Abs(fmt.Sprintf("%s/%s.png", spritesPath, filename))
	sheet, _ := os.Open(path)
	defer sheet.Close()
	spritesheet, _ := png.Decode(sheet)
	// copy spritesheet to memory so we can subimage pieces of it
	sprites := image.NewRGBA(image.Rect(0, 0, spritesheet.Bounds().Size().X, spritesheet.Bounds().Size().Y))
	draw.Draw(sprites, sprites.Bounds(), spritesheet, image.Point{0, 0}, draw.Src)

	// TODO handle errors

	return data, sprites
}
