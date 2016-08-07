package main

// Hp
// Attack power
// Energy level
// Fear
// Inventory - max 2 items
type Player struct {
	hp              int
	attack          int
	energy          int
	fear            int
	currentLocation *Room
	inventory       *[]Item
}

func NewPlayer(room *Room) *Player {
	return &Player{
		hp:              100,
		attack:          0,
		energy:          100,
		fear:            0,
		currentLocation: room,
		inventory:       nil,
	}
}
