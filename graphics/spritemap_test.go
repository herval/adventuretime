package graphics

import (
	"testing"
	"github.com/herval/adventuretime/util"
)

func TestSpritemap(t *testing.T) {

	spritemap := LoadSpritemap("../resources")

	if len(spritemap.Frames) == 0 {
		t.Fail()
	}

	util.DebugFmt("%+v", spritemap)

}
