package engine_test

import (
	"testing"
	"github.com/herval/adventuretime/engine"
)

func TestGeneration(t *testing.T) {
	dungeon := engine.NewDungeon(20, 5)

	dungeon.Blueprint.Print()

	// there must be rooms
	if len(dungeon.Rooms) == 0 {
		t.Fail()
	}

	// there must be an entrance door
	lockedEntrance := false
	for i := range dungeon.Entrance.Doors {
		lockedEntrance = lockedEntrance || dungeon.Entrance.Doors[i].Locked
	}
	if !lockedEntrance {
		t.Fail()
	}

	// each room must have at least a door
	for i := range dungeon.Rooms {
		if len(dungeon.Rooms[i].Doors) == 0 {
			t.Fail()
		}
	}
}
