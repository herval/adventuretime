package main

import (
	"testing"
	"fmt"
)

func TestSpritemap(t *testing.T) {

	loader := FramesLoader{}
	data := loader.Parse("./sprites.json")

	spritemap := NewSpritemap(data)

	fmt.Printf("%+v", spritemap.SmallMonsters)

}