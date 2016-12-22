package engine_test

import (
	"testing"
	"github.com/herval/adventuretime/engine"
)

func TestGeneration(t *testing.T) {
	dungeon := engine.NewDungeon()

	dungeon.Blueprint.Print()

	if len(dungeon.Rooms) == 0 {
		t.Fail()
	}

	lockedEntrance := false
	for i := range dungeon.Entrance.Doors {
		lockedEntrance = lockedEntrance || dungeon.Entrance.Doors[i].Locked
	}
	if !lockedEntrance {
		t.Fail()
	}

	for i := range dungeon.Rooms {
		if len(dungeon.Rooms[i].Doors) == 0 {
			t.Fail()
		}
	}
}
