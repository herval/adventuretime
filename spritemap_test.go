package main

import (
	"testing"
	"fmt"
)

func TestSpritemap(t *testing.T) {

	spritemap := LoadSpritemap()

	fmt.Printf("%+v", spritemap.SmallMonsters)

}