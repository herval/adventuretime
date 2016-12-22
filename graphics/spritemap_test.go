package graphics_test

import (
	"testing"
	"github.com/herval/adventuretime/util"
	"github.com/herval/adventuretime/graphics"
)

func TestSpritemap(t *testing.T) {

	spritemap := graphics.LoadSpritemap("../resources")

	if len(spritemap.Frames) == 0 {
		t.Fail()
	}

	util.DebugFmt("%+v", spritemap)

}
