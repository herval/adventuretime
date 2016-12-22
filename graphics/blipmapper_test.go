package graphics_test

import (
	"testing"
	"github.com/herval/adventuretime/engine"
	"github.com/herval/adventuretime/graphics"
	"github.com/herval/adventuretime/util"
)

func TestBlipmapping(t *testing.T) {

	dungeon := engine.NewDungeon(50, 3)

	str := graphics.DungeonToBlipmap(dungeon)
	if len(str) == 0 {
		t.Fail()
	}

	util.Debug(str)
}
