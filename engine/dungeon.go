package engine

import (
	"fmt"

	"github.com/herval/adventuretime/util"
)

type Dungeon struct {
	Entrance *Room
}

// Generate dungeon... lazily
func NewDungeon() *Dungeon {
	// the main dungeon door in all it's glory
	youShallNotPass := &Door{
		facing: NORTH,
		locked: true,
	}

	mainHall := RandomRoom(nil, youShallNotPass)
	mainHall.details = "This is the entrance of the Dungeon."

	util.Debug(fmt.Sprintf("Init dungeon: %s", mainHall))

	return &Dungeon{
		Entrance: mainHall,
	}
}
