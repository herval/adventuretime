package main

import "fmt"

type Dungeon struct {
	entrance *Room
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

	Debug(fmt.Sprintf("Init dungeon: %s", mainHall))

	return &Dungeon{
		entrance: mainHall,
	}
}
