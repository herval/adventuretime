package engine

import "github.com/herval/adventuretime/util"

type GameState struct {
	Player  *Player
	Dungeon *Dungeon
	CurrentTurn uint32
}

func NewGame() *GameState {
	util.Debug("Generating dungeon...")
	dungeon := NewDungeon()

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
