package graphics

import (
	"fmt"
	"testing"
)

func TestSpritemap(t *testing.T) {

	spritemap := LoadSpritemap("../resources")

	fmt.Printf("%+v", spritemap.SmallMonsters)

}
