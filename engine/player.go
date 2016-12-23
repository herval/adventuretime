package engine

import "github.com/herval/adventuretime/util"

const MAX_HP = 100

type Item struct {
}

// Hp
// Attack power
// Energy level
// Fear
// Inventory - max 2 items
type Player struct {
	Hp              int
	Attack          int
	Energy          int
	Fear            int
	CurrentLocation *Room
	Inventory       *[]Item
}

func NewPlayer(room *Room) *Player {
	player := Player{
		Hp:              100,
		Attack:          0,
		Energy:          100,
		Fear:            0,
		CurrentLocation: room,
		Inventory:       nil,
	}

	return &player
}

func (p *Player) Heal(amount int) {
	if p.Hp < MAX_HP {
		p.Hp += util.Min(amount, MAX_HP-p.Hp)
	}
}
