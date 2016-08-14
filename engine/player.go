package engine

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
	return &Player{
		Hp:              100,
		Attack:          0,
		Energy:          100,
		Fear:            0,
		CurrentLocation: room,
		Inventory:       nil,
	}
}
