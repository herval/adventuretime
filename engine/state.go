package engine

import (
	"github.com/herval/adventuretime/util"
	"math/rand"
)

type GameState struct {
	Player  *Player
	Dungeon *Dungeon
	CurrentTurn uint32
}

func NewGame() *GameState {
	util.Debug("Generating dungeon...")

	totalRooms := rand.Intn(30) + 10
	size := 100

	dungeon := NewDungeon(size, totalRooms)

	util.Debug("Configuring character...")
	player := NewPlayer(dungeon.Entrance)

	return &GameState{
		Player:  player,
		Dungeon: dungeon,
		CurrentTurn: 0,
	}
}

func (self *GameState) Describe() string {
	return self.Player.CurrentLocation.Describe()
}
